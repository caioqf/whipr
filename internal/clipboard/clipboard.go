package clipboard

type Reader interface {
	Read() (string, error)
}
