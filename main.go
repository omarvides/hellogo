//My hello world application, prints a hello world
package main

import (
	"fmt"

	"github.com/omarvides/hellogo/helpers"
)

func main() {
	x := helpers.AddPrefix("pre", "fixed")
	fmt.Println("The value is", x)
	sumFunc := helpers.GetSum(5)
	subFunc := helpers.GetSub(5)
	mulFunc := helpers.GetMul(5)
	divFunc := helpers.GetDiv(5)
	fmt.Println("The sum of 5 and 3 is", sumFunc(3))
	fmt.Println("The substraction of 5 and 3 is", subFunc(3))
	fmt.Println("5 times 3 is", mulFunc(3))
	fmt.Println("5 divided by 3 is", divFunc(3))
}
