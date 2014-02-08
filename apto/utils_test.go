package apto

func ExampleDebug() {
	a, b := 5, 18
	Debug("Value of a is %v - Value of b is %v", a, b)
	// Output: --- DEBUG: Value of a is 5 - Value of b is 18
}
