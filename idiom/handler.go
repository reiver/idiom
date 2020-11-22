package idiom

type Handler interface {
	Run(parameters ...Stringer) (Stringer, error)
}
