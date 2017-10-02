package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	result, err := o()

	for err != nil {
		_, ok := err.(TransientError)
		if !ok {
			return err
		}
		result, err = o()
	}

	defer result.Close()
	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(FrobError); ok {
				result.Defrob(v.defrobTag)
			}
			err = r.(error)
		}
	}()
	result.Frob(input)

	return
}
