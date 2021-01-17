package idioms

var Commands Registrar

func init() {
	Commands = &internalRegistrar{}
}
