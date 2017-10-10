package simple

func fn1_1() bool { return false }
func fn1_2() bool { return false }

func fn1() {
	type T bool
	var x T
	const t T = false
	if x == t {
	}
	if fn1_1() == true { // MATCH "simplified to fn1_1()"
	}
	if fn1_1() != true { // MATCH "simplified to !fn1_1()"
	}
	if fn1_1() == false { // MATCH "simplified to !fn1_1()"
	}
	if fn1_1() != false { // MATCH "simplified to fn1_1()"
	}
	if fn1_1() && (fn1_1() || fn1_1()) || (fn1_1() && fn1_1()) == true { // MATCH "simplified to (fn1_1() && fn1_1())"
	}

	if (fn1_1() && fn1_2()) == false { // MATCH "simplified to !(fn1_1() && fn1_2())"
	}

	var y bool
	for y != true { // MATCH /simplified to !y/
	}
	if !y == true { // MATCH /simplified to !y/
	}
	if !y == false { // MATCH /simplified to y/
	}
	if !y != true { // MATCH /simplified to y/
	}
	if !y != false { // MATCH /simplified to !y/
	}
	if !!y == false { // MATCH /simplified to !y/
	}
	if !!!y == false { // MATCH /simplified to y/
	}
	if !!y == true { // MATCH /simplified to y/
	}
	if !!!y == true { // MATCH /simplified to !y/
	}
	if !!y != true { // MATCH /simplified to !y/
	}
	if !!!y != true { // MATCH /simplified to y/
	}
	if !y == !false { // not matched because we expect true/false on one side, not !false
	}

	var z interface{}
	if z == true {
	}
}
