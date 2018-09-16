package util

func Choose(expr bool, trueValue, falseValue string) string {
	if expr {
		return trueValue
	}

	return falseValue
}
