package main

import (
	"fmt"
)

/* func main() {
	os.Args
}
*/
func looper() {
	x := 0
	for i := 0; i < 10; i++ {
		fmt.Println("number:", x)
	}
}
