package helpers

// GetSum function Receives a base number (int) and returns a function that receives a value (int) and returns
// the sum of the base and the value passed
func GetSum(base int) func(int) int {
	return func(value int) int {
		return base + value
	}
}

// GetSub function Receives a base number (int) and returns a function that receives a value (int) and returns
// the substraction of the base and the value passed
func GetSub(base int) func(int) int {
	return func(value int) int {
		return base - value
	}
}

// GetDiv function Receives a base number (float32) and returns a function that receives a value (float32) and returns
// the divition of the base and the value passed
func GetDiv(base float32) func(float32) float32 {
	return func(value float32) float32 {
		return value / base
	}
}

// GetMul function Receives a base number (int) and returns a function that receives a value (int) and returns
// the the base times value passed
func GetMul(base int) func(int) int {
	return func(value int) int {
		return base * value
	}
}
