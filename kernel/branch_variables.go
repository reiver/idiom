package idiom_kernel

func (receiver *Branch) Variables() Constructor {
	if nil == receiver {
		return nil
	}

	return &(receiver.variables)
}
