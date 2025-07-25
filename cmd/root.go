package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whipr",
	Short: "Whipr is a tool for fast text translation",
	Long:  `Whipr is a tool that shows translations of selected text on a shortcut pressed.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.AddCommand(shortcutCmd)
	// rootCmd.AddCommand(trayCmd)
}
