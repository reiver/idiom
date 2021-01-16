package idioms

//  HandlerFunc is an adaptor that turns a func with signature:…
//
//	func(...string)(string, error)
//
// … into a idioms.Handler
type HandlerFunc func(...string)(string, error)

func (fn HandlerFunc) Run(parameters ...string) (string, error) {
	return fn(parameters...)
}
