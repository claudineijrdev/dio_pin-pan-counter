package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("PinPan")
			break
		case i%3 == 0:
			fmt.Println("Pin")
			break
		case i%5 == 0:
			fmt.Println("Pan")
			break
		default:
			fmt.Println(i)
		}
	}
}
