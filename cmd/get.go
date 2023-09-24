package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetches and decrypts the specified password from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		passwordName := args[0]
		getPassword(passwordName)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getPassword(passwordName string) {
	password := decrypt(passwordName)
	clipboard.WriteAll(password)

	fmt.Printf("%s password copied to clipboard\n", passwordName)
}

func decrypt(fileName string) string {
	homeDir, _ := os.UserHomeDir()
	path := homeDir + "/.store/" + fileName + ".asc"
	cmd := exec.Command("gpg", "-d", path)

	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	result := string(data)
	return result
}
