package main

import "fmt"

func main() {
	var angka int
	fmt.Println("Mari masukkan angka: ")
	fmt.Scanf("%d", &angka)
	
	if angka%3 == 0 && angka%5 == 0 {
		fmt.Println("Hello World")
	} else if angka%5 == 0 {
		fmt.Println("World")
	} else if angka%3 == 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("nothing")
	}
}
