package array_example

import (
	"fmt"
)

func Array_example() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("Array_example(): recover", handler)
		}
	}()

	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}
