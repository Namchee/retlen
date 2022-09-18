package src

func FunctionWithoutReturn() {
	return
}

func FunctionWithReturn() bool {
	return true
}

func FunctionWithEnoughReturn() (int, bool) {
	return 0, false
}

func FunctionWithTooManyReturn() (int, int, int) {
	return 1, 2, 3
}

func TestSkipThis() (int, int, int) {
	return 1, 2, 3
}
