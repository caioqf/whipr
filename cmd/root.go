package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/caioqf/whipr/assets/icon"
	"github.com/getlantern/systray"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whipr",
	Short: "Whipr is a tool for fast text translation",
	Long:  `Whipr is a tool that shows translations of selected text on a shortcut pressed.`,
	Run: func(cmd *cobra.Command, args []string) {
		systray.Run(onReady, func() {})
	},
}

func onReady() {
	iconBytes := icon.LoadIcon()
	if len(iconBytes) > 0 {
		systray.SetTemplateIcon(iconBytes, iconBytes)
	}
	systray.SetTooltip("Click to translate selection")

	mTranslate := systray.AddMenuItem("Translate selected text", "Translate selected text")
	systray.AddSeparator()

	mPopup := systray.AddMenuItemCheckbox("Popup", "Use popup", false)
	mNotify := systray.AddMenuItemCheckbox("Notify", "Use notify", true)

	mQuit := systray.AddMenuItem("Exit", "Quit")

	go func() {
		for {
			select {
			case <-mTranslate.ClickedCh:
				out, err := exec.Command("xclip", "-o", "-selection", "primary").Output()
				if err != nil {
					out = []byte("Error getting selection: " + err.Error())
				}
				if mPopup.Checked() {
					exec.Command("zenity", "--info", "--text", string(out)).Run()
				}
				if mNotify.Checked() {
					exec.Command("notify-send", "Translation", string(out)).Run()
				}
			case <-mPopup.ClickedCh:
				SetPopupEnabled(true) // This updates settings AND UI
				mPopup.Check()
				mNotify.Uncheck()
			case <-mNotify.ClickedCh:
				SetNotifyEnabled(true) // This updates settings AND UI
				mNotify.Check()
				mPopup.Uncheck()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	systray.Quit()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(shortcutCmd)
}
