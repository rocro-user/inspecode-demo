package simple

func fn12() {
	var ch chan int
	// MATCH /should use a simple channel send/
	<-ch

outer:
	// MATCH /should use for range/

	for range ch {

		break outer

	}

	// MATCH /should use for range/

	for x := range ch {

		_ = x

	}

	for {
		// MATCH /should use a simple channel send/
		ch <- 0

	}
}
