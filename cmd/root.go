package cmd

import (
	"fmt"
	"log"
	"net"
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
		systray.Run(onReady, onExit)
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

	mPopup := systray.AddMenuItemCheckbox("Popup", "Use popup", ShouldUsePopup())
	mNotify := systray.AddMenuItemCheckbox("Notify", "Use notify", ShouldUseNotify())
	systray.AddSeparator()

	mSettings := systray.AddMenuItem("Settings", "Settings")
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit Whipr", "Quit Whipr")

	doTranslate := func() {
		out, err := exec.Command("xclip", "-o", "-selection", "primary").Output()
		if err != nil {
			out = []byte("Error getting selection: " + err.Error())
		}
		DisplayTranslated(string(out))
	}

	go func() {
		if err := os.Remove("/tmp/whipr.sock"); err != nil && !os.IsNotExist(err) {
			log.Printf("Error removing socket: %v", err)
		}
		listener, err := net.Listen("unix", "/tmp/whipr.sock")
		if err != nil {
			log.Printf("Failed to listen on socket: %v", err)
			return
		}
		defer listener.Close()
		log.Println("Listening on /tmp/whipr.sock")
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Accept error: %v", err)
				continue
			}
			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, 512)
				n, err := c.Read(buf)
				if err != nil {
					log.Printf("Read error: %v", err)
					return
				}
				msg := string(buf[:n])
				if msg == "translate" {
					doTranslate()
				}
			}(conn)
		}
	}()

	go func() {
		for {
			select {
			case <-mTranslate.ClickedCh:
				doTranslate()
			case <-mPopup.ClickedCh:
				SetPopupEnabled(true)
				mPopup.Check()
				mNotify.Uncheck()
			case <-mNotify.ClickedCh:
				SetNotifyEnabled(true)
				mNotify.Check()
				mPopup.Uncheck()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			case <-mSettings.ClickedCh:
				// openSettings()
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
	// Initialize shortcut command flags
	shortcutCmd.Flags().BoolVar(&usePopup, "popup", false, "Use popup display")
	shortcutCmd.Flags().BoolVar(&useNotify, "notify", false, "Use notification display")

	rootCmd.AddCommand(shortcutCmd)
}
