package iterators

var sumCallback int

func CallbackLoop(top int) {
	err := callbackLoopIterator(top, func(n int) error {
		sumCallback += n
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func callbackLoopIterator(top int, callback func(n int) error) error {
	for i := 0; i < top; i++ {
		err := callback(i)
		if err != nil {
			return err
		}
	}
	return nil
}
