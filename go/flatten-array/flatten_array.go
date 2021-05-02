package flatten

// Flatten takes a nested list and returns a single flattened list with all
// values except nil/null.
func Flatten(input interface{}) []interface{} {
	if nil == input {
		return nil
	}
	r := []interface{}{}
	switch input := input.(type) {
	case []interface{}:
		for _, v := range input {
			r = append(r, Flatten(v)...)
		}
	default:
		r = append(r, input)
	}
	return r
}
