package icon

import (
	_ "embed"
	"log"
)

//go:embed icon.png
var iconBytes []byte

func LoadIcon() []byte {
	log.Printf("[icon] Loaded icon.png with %d bytes", len(iconBytes))
	return iconBytes
}
