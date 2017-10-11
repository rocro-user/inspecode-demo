package simple

import "errors"

func fn5_1() error {
	var err error

	// MATCH /simplified/
	return err

	_ = 0

	// MATCH /simplified/
	return err

}

func fn5_2() (int, string, error) {
	var x int
	var y string
	var z int
	var err error

	// MATCH /simplified/
	return x, y, err

	_ = 0

	// MATCH /simplified/
	return x, y, err

	_ = 0

	if err != nil {
		return x, y, err
	}
	return z, y, err

	_ = 0

	if err != nil {
		return 0, "", err
	}
	return x, y, err

	_ = 0

	// TODO(dominikh): currently, only returning identifiers is
	// supported
	if err != nil {
		return 42, "foo", err
	}
	return 42, "foo", err
}

func fn5_3() error {
	var err error
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func fn5_4(i int, err error) error {
	if err != nil {
		return err
	} else if i == 1 {
		return errors.New("some non-nil error")
	}
	return nil
}

func fn5_5() interface{} {
	var i *int
	if i != nil {
		return i
	}
	return nil

	var v interface{}
	// MATCH /simplified/
	return v

}

func fn5_6() {
	_ = func() error {
		var err error
		// MATCH /simplified/
		return err

	}
}
