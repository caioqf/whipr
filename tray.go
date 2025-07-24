package main

import (
	"os/exec"

	"github.com/caioqf/whipr/assets/icon"
	"github.com/getlantern/systray"
)

func OnReady() {
	iconBytes := icon.LoadIcon()
	if len(iconBytes) > 0 {
		systray.SetTemplateIcon(iconBytes, iconBytes)
	}
	systray.SetTooltip("Click to translate selection")

	mTranslateSelected := systray.AddMenuItem("Translate selected text", "Translate clipboard")
	systray.AddSeparator()

	mOptNotify := systray.AddMenuItemCheckbox("Notification", "Show translation on a notification", true)
	mOptPopup := systray.AddMenuItemCheckbox("Pop-Up", "Show the translation on a popup window", false)
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Exit", "Close app")

	go func() {
		for {
			select {
			case <-mTranslateSelected.ClickedCh:
				cmd := exec.Command("xclip", "-o", "-selection", "primary")
				selectedText, err := cmd.Output()
				if err != nil {
					selectedText = []byte("Erro ao obter seleção: " + err.Error())
				}

				if mOptPopup.Checked() {
					exec.Command("zenity", "--info", "--text", string(selectedText)).Run()
				}
				if mOptNotify.Checked() {
					exec.Command("notify-send", "Translation", string(selectedText)).Run()
				}
			case <-mOptNotify.ClickedCh:
				mOptNotify.Check()
				mOptPopup.Uncheck()
			case <-mOptPopup.ClickedCh:
				mOptPopup.Check()
				mOptNotify.Uncheck()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func OnExit() {

}
