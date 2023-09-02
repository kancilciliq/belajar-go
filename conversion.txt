package main

import "fmt"

func main() {
	var number32 int32 = 10000
	var number64 int64 = int64(number32)

	fmt.Println(number32)
	fmt.Println(number64)

	var name = "salman"
	var e = name[0]
	var eName = string(e)

	fmt.Println(eName)
}
