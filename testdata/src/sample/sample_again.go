package src

func AFunctionWithoutReturn() {
	return
}

func AFunctionWithReturn() bool {
	return true
}

func AFunctionWithEnoughReturn() (int, bool) {
	return 0, false
}

func AFunctionWithTooManyReturn() (int, int, int) {
	return 1, 2, 3
}

func TestSkipThisBro() (int, int, int) {
	return 1, 2, 3
}
