//go:build darwin

package clipboard

import (
	"os/exec"
	"strings"
)

type darwinReader struct{}

func DefaultReader() Reader {
	return darwinReader{}
}

func (darwinReader) Read() (string, error) {
	out, err := exec.Command("/usr/bin/pbpaste").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
