package internal

type Reader interface {
	Read() (string, error)
}
