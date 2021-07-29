package main

import (
	"fmt"
)

func main() {
	var angka int
	a := angka%3 == 0
	b := angka%5 == 0
	fmt.Println("Mari masukkan angka: ")
	fmt.Scanf("%d", &angka)

	if a {
		fmt.Println("Hello")
	} else if b {
		fmt.Println("World")
	} else if a && b {
		fmt.Println("Hello World")
	} else {
		fmt.Println("nothing")
	}
}
