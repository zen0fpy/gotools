package tests

func Test1() bool {
	var v string
	if v == "" {
		return true
	}
	return false
}

func Test2() bool {
	var v string
	if len(v) == 0 {
		return true
	}
	return false
}
