package main

import (
	"fmt"
	"net"
	"os"
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
				runTranslation(mOptPopup, mOptNotify)
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
	_ = os.Remove(sockPath)
}

func runTranslation(mOptPopup, mOptNotify *systray.MenuItem) {
	cmd := exec.Command("xclip", "-o", "-selection", "primary")
	selectedText, err := cmd.Output()
	if err != nil {
		selectedText = []byte("error obtaining selection: " + err.Error())
	}

	if mOptPopup.Checked() {
		exec.Command("zenity", "--info", "--text", string(selectedText)).Run()
	}
	if mOptNotify.Checked() {
		exec.Command("notify-send", "Translation", string(selectedText)).Run()
	}
}

func startSocketServer(onTranslate func()) {
	_ = os.Remove(sockPath)
	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		fmt.Println("error initiating socket", err)
		return
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			buf := make([]byte, 128)
			n, _ := conn.Read(buf)
			cmd := string(buf[:n])
			if cmd == "translate" {
				onTranslate()
			}
			conn.Close()
		}
	}()
}
