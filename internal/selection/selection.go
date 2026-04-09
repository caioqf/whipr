package selection

import "errors"

var ErrNoSelection = errors.New("no text selected or selection could not be read")

type SelectionReader interface {
	Read() (string, error)
}
