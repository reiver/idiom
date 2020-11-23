package idiom

type Handler interface {
	Run(...Stringer) (Stringer, error)
}
