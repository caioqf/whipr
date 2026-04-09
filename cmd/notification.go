package cmd

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

type Notification struct {
	Title   string
	Message string
	Icon    string
	Media   []string
}

func runCmd(name string, arg ...string) {
	if err := exec.Command(name, arg...).Run(); err != nil {
		log.Printf("%s: %v", name, err)
	}
}

func renderNotification(content Notification) {
	switch runtime.GOOS {
	case "linux":
		log.Println("Rendering notification on Linux")
		args := []string{content.Title, content.Message}
		if content.Icon != "" {
			args = append(args, "-i", content.Icon)
		}
		runCmd("notify-send", args...)
	case "darwin":
		log.Println("Rendering notification on Darwin")
		notifyDarwin(content)
	case "windows":
		log.Println("Rendering notification on Windows")
		runCmd("msg", content.Title, content.Message)
	}
}

func notifyDarwin(content Notification) {
	note := gosxnotifier.NewNotification(content.Message)
	note.Title = content.Title
	note.Group = "github.com.caioqf.whipr"
	if content.Icon != "" {
		if _, err := os.Stat(content.Icon); err == nil {
			note.AppIcon = content.Icon
		}
	}
	if err := note.Push(); err != nil {
		log.Printf("gosx-notifier: %v", err)
	}
}
