//go:build darwin

package selection

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

type darwinSelectionReader struct{}

func DefaultSelectionReader() SelectionReader {
	return darwinSelectionReader{}
}

func (darwinSelectionReader) Read() (string, error) {
	before, err := exec.Command("/usr/bin/pbpaste").Output()
	if err != nil {
		before = nil
	}
	beforeNorm := strings.TrimSpace(string(before))

	script := `tell application "System Events" to keystroke "c" using command down`
	if err := exec.Command("/usr/bin/osascript", "-e", script).Run(); err != nil {
		return "", err
	}
	time.Sleep(120 * time.Millisecond)

	after, err := exec.Command("/usr/bin/pbpaste").Output()
	if err != nil {
		_ = restorePasteboardDarwin(before)
		return "", err
	}
	afterNorm := strings.TrimSpace(string(after))

	_ = restorePasteboardDarwin(before)

	if afterNorm == beforeNorm {
		return "", ErrNoSelection
	}
	return afterNorm, nil
}

func restorePasteboardDarwin(data []byte) error {
	cmd := exec.Command("/usr/bin/pbcopy")
	cmd.Stdin = bytes.NewReader(data)
	return cmd.Run()
}
