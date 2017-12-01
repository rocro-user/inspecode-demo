package simple

func fn6(i interface{}) {
	if _, ok := i.(string); ok { // MATCH "when ok is true, i can't be nil"
	}
	if _, ok := i.(string); ok { // MATCH "when ok is true, i can't be nil"
	}
	if _, ok := i.(string); i != nil || ok {
	}
	if _, ok := i.(string); i != nil && !ok {
	}
	if _, ok := i.(string); i == nil && ok {
	}
}
