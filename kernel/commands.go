package idiom_kernel

var Commands Registrar

func init() {
	Commands = &internalRegistrar{}
}
