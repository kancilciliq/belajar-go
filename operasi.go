package main

import "fmt"

func main() {
	var a = 100 + 96

	var b = 3
	var c = 7
	var d = b + c

	fmt.Println(a)
	fmt.Println(d)

	//augmented assignment

	var i = 10
	// i += 10 // i = i + 10 = 20
	i *= 20

	fmt.Println(i)

	//unary operator

	var e = 10
	e++ // e = e + 1 = 11
	fmt.Println(e)

}
