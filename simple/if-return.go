package simple

func fn4() bool { return true }
func fn4_1() bool {
	x := true
	if x { // MATCH /should use 'return <expr>'/
		return true
	}
	return false
}

func fn4_2() bool {
	x := true
	if !x {
		return true
	}
	if x {
		return true
	}
	return false
}

func fn4_3() int {
	var x bool
	if x {
		return 1
	}
	return 2
}

func fn4_4() bool { return true }

func fn4_5() bool {
	if fn4() { // MATCH /should use 'return <expr>'/
		return false
	}
	return true
}

func fn4_6() bool {
	if fn4_3() != fn4_3() { // MATCH /should use 'return <expr>'/
		return true
	}
	return false
}

func fn4_7() bool {
	if 1 > 2 { // MATCH /should use 'return <expr>'/
		return true
	}
	return false
}

func fn4_8() bool {
	if fn4() || fn4() {
		return true
	}
	return false
}
