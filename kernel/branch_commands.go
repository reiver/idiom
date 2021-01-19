package idiom_kernel

func (receiver *Branch) Commands() Registrar {
	if nil == receiver {
		return nil
	}

	return &(receiver.commands)
}
