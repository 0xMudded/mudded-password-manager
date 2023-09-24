package cmd

import (
	"os"

	"github.com/0xMudded/mudded-password-manager/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mudded-password-manager",
	Aliases: []string{"mpw"},
	Short:   "A simple terminal-based password manager",
	Long: `
- A simple terminal-based password manager

- Generates and stores encrypted passwords in the .store directory

- Run 'mpw init <gpg-key-id>' to initialize your password store
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Load)
}
