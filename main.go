package main

import (
	"fmt"
)

func main() {
	for a := 0; a <= 100; a++ {
		if (a % 3 == 0 && a % 5 == 0) {
			fmt.Println("foobar")
		} else if (a % 5 == 0) {
			fmt.Println("bar")
		} else if (a % 3 == 0) {
			fmt.Println("foo")
		} else {
			fmt.Println(a)
		}
	}
}
