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

				exec.Command("zenity", "--info", "--text", string(selectedText)).Run()

			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func OnExit() {

}
