package idiom_kernel

type Registrar interface {
	Fetch(string) (Handler, error)
	Register(Handler, string) error
}
