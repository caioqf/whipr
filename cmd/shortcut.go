package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	usePopup  bool
	useNotify bool
)

var shortcutCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Trigger translation from clipboard selection",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Shortcut command started")

		// get text from selection
		output, err := exec.Command("xclip", "-o", "-selection", "primary").Output()
		if err != nil {
			log.Printf("Error getting clipboard: %v", err)
			output = []byte("Error: " + err.Error())
		}

		log.Printf("Clipboard content: %q", string(output))

		if usePopup {
			SetPopupEnabled(true)
		}
		if useNotify {
			SetNotifyEnabled(true)
		}

		DisplayTranslated(string(output))

		log.Println("Shortcut command completed")
	},
}
