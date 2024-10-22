package main

func Multiply(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

var Double = Multiply(2)
var Triple = Multiply(3)

func main() {

	print(Double(2), Triple(3))
	print(Mutliply(2)(3))
}
