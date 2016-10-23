//My hello world application, prints a hello world
package main

import (
	"fmt"

	"github.com/omarvides/hellogo/helpers"
)

func main() {
	x := helpers.AddPrefix("-prefixed")
	fmt.Println("	The value is", x)
}
