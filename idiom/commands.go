package idiom

var Commands Registrar

func init() {
	Commands = &internalRegistrar{}
}
