package main

import "fmt"

func main() {
	nilaiUjian := 80

	if nilaiUjian >= 75 {
		fmt.Println("lulus")
	} else {
		fmt.Println("tidak lulus")
	}

	//short statement
	name := "salman"

	if length := len(name); length >= 3 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
