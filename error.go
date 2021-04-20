package main

func recover_from_error() {
	defer func() {
		recover()
	}()
	_ = raise_error()
}

func raise_error() float64 {
	divisor := 0.0
	dividend := 0.0
	return dividend / divisor
}
