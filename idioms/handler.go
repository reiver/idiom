package idioms

type Handler interface {
	Run(...string) (string, error)
}
