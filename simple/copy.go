package simple

func fn3() {
	var b1, b2 []byte
	// MATCH /should use copy/
	copy(b2, b1)

	// MATCH /should use copy/
	copy(b2, b1)

	type T [][16]byte
	var a T
	b := make([]interface{}, len(a))
	for i := range b {
		b[i] = a[i]
	}

	var b3, b4 []*byte
	// MATCH /should use copy/
	copy(b4, b3)

	var m map[int]byte
	for i, v := range b1 {
		m[i] = v
	}

}
