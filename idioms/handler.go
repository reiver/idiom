package idiom

type Handler interface {
	Run(...string) (string, error)
}
