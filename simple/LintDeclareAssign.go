package simple

func fn7() {
	// MATCH "should merge variable declaration with assignment on next line"
	var x = 1
	_ = x

	// MATCH "should merge variable declaration with assignment on next line"
	var y interface{} = 1
	_ = y

	if true {
		var x string // MATCH "should merge variable declaration with assignment on next line"

		_ = x
	}

	var z []string
	z = append(z, "")
	_ = z

	var f func()
	f = func() { f() }
	_ = f
}
