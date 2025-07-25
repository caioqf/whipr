package main

import (
	// "fmt"
	// "net"
	// "os"
	// "github.com/getlantern/systray"

	"github.com/caioqf/whipr/cmd"
)

const sockPath = "/tmp/whipr.sock"

func main() {
	// if len(os.Args) > 1 && os.Args[1] == "--shortcut" {
	// 	conn, err := net.Dial("unix", sockPath)
	// 	if err != nil {
	// 		fmt.Println("whipr not running.")
	// 		return
	// 	}
	// 	conn.Write([]byte("translate"))
	// 	conn.Close()
	// 	return
	// }
	// systray.Run(OnReady, OnExit)

	cmd.Execute()

}
