package cmd

import (
	"os"

	"github.com/0xMudded/mudded-password-manager/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new password store",
	Long: `
- Initializes a new password store in the user's home directory

- A valid GPG key ID must be specified as the first and only argument when running the init command`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createStore()
		keyId := args[0]
		config.AddConfig("keyId", keyId)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createStore() {
	homeDir, _ := os.UserHomeDir()
	err := os.Mkdir(homeDir+"/.store", 0755)
	if err != nil {
		panic(err)
	}
}
