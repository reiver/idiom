package idiom

type HandlerFunc func(...string)(string, error)

func (fn HandlerFunc) Run(parameters ...string) (string, error) {
	return fn(parameters...)
}
