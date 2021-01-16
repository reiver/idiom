package idioms

type Registrar interface {
	Fetch(string) (Handler, error)
	Register(Handler, string) error
}
