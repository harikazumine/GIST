package main

import (
	"fmt"
)

func palindrom(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var check string
	fmt.Println("Mari masukkan kata: ")
	fmt.Scanf("%s", &check)

	if palindrom(check) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
