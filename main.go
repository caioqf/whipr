package main

import (
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(OnReady, OnExit)
}
