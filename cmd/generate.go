package cmd

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/0xMudded/mudded-password-manager/config"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@£$%^&*")

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates and adds a new password to the store",
	Long: `
- Generates and adds a new password to the store

- Each password is encrypted using the the gpg key which the user specified when initializing the store, and stored in a separate .asc file inside the .store directory located in the user's home directory`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		passwordName := args[0]
		generate(passwordName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generate(passwordName string) {
	password := generatePassword(16)
	encryptedPassword := encrypt(password, config.GetViper().GetString("keyid"))

	success := createPasswordFile(passwordName, encryptedPassword)
	if success {
		clipboard.WriteAll(password)
		fmt.Println("Successfully generated new password and copied to clipboard")
	}
}

func createPasswordFile(fileName, encryptedPassword string) bool {
	homeDir, _ := os.UserHomeDir()
	path := homeDir + "/.store/" + fileName + ".asc"

	if fileExists(path) {
		fmt.Printf("Error! %s password file already exists. Delete the file or specify a different name", fileName)
		return false
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%v\n", encryptedPassword)
	w.Flush()

	return true
}

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)

	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	return string(b)
}

func encrypt(message, gpgId string) string {
	cmd := exec.Command("gpg", "-e", "-r", gpgId, "--armor")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, message)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	result := string(out)
	return result
}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
