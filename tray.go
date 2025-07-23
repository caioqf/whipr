package main

import (
	"os/exec"

	"github.com/getlantern/systray"
)

func OnReady() {
	// icone comentado por enquanto
	// systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("whipr")
	systray.SetTooltip("Click to translate selection")

	mTranslate := systray.AddMenuItem("Translate selected text", "Translate clipboard")
	systray.AddSeparator()

	mOptNotify := systray.AddMenuItemCheckbox("Notification", "Show translation on a notification", true)
	mOptPopup := systray.AddMenuItemCheckbox("Popup", "Show the translation on a popup window", false)
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Exit", "Close app")

	go func() {
		for {
			select {
			case <-mTranslate.ClickedCh:
				// captra o texto selecionado usando
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
