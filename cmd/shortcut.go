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

		// Use flags if provided, otherwise fall back to settings
		shouldUseNotify := useNotify || (!usePopup && ShouldUseNotify())
		shouldUsePopup := usePopup || (!useNotify && ShouldUsePopup())

		if shouldUseNotify {
			log.Println("Using notification display")
			err := exec.Command("notify-send", "Translation", string(output)).Run()
			if err != nil {
				log.Printf("notify-send error: %v", err)
			}
		}
		if shouldUsePopup {
			log.Println("Using popup display")
			err := exec.Command("zenity", "--info", "--text", string(output)).Run()
			if err != nil {
				log.Printf("zenity error: %v", err)
			}
		}

		log.Println("Shortcut command completed")
	},
}
