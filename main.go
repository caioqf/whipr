package main

import (
	"fmt"
	"net"
	"os"

	"github.com/caioqf/whipr/cmd"
)

const sockPath = "/tmp/whipr.sock"

func main() {
	cmd.LoadSettings()

	if len(os.Args) > 1 && os.Args[1] == "shortcut" {
		conn, err := net.Dial("unix", sockPath)
		if err == nil {
			_, err = conn.Write([]byte("translate"))
			if err != nil {
				fmt.Printf("Write error: %v\n", err)
			}
			conn.Close()
			return
		}
		fmt.Println("whipr not running, using default settings.")
	}

	cmd.Execute()
}
