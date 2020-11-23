package idiom

type HandlerFunc func(...Stringer)(Stringer, error)

func (fn HandlerFunc) Run(parameters ...Stringer) (Stringer, error) {
	return fn(parameters...)
}
