package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

var shortcutCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Trigger translation from clipboard selection",
	Run: func(cmd *cobra.Command, args []string) {
		// get text from selection
		output, err := exec.Command("xclip", "-o", "-selection", "primary").Output()
		if err != nil {
			output = []byte("Error: " + err.Error())
		}

		if ShouldUseNotify() {
			exec.Command("notify-send", "Translation", string(output)).Run()
		}
		if ShouldUsePopup() {
			exec.Command("zenity", "--info", "--text", string(output)).Run()
		}
	},
}
