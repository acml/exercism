package erratum

// Use opens a resource, calls Frob(input) on the result resource and then
// closes that resource (in all cases)
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = o()
		if _, ok := err.(TransientError); !ok {
			break
		}
	}
	if err != nil {
		return err
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(FrobError); ok {
				res.Defrob(e.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)
	return nil
}
