package persist

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {}()
	return out

}
