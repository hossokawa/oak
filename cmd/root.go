/*
Copyright © 2026 Leonardo Hossokawa <leonardoholiver@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oak",
	Short: "Terminal-based Pokémon encyclopedia",
	Long: `oak is a fast, lightweight, and modern CLI designed to bring the Pokémon
world directly to your terminal. 

Inspired by Kanto's professor Oak, this tool serves as your personal pocket encyclopedia.
Whether you need to quickly check type effectiveness mid-battle, inspect base stats
for competitive team building, or look up evolutionary chains, oak provides comprehensive
data without forcing you to leave your development environment.

"Welcome to the world of Pokémon! Your journey through the data begins here."`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oak.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
