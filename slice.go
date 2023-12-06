package main

import "fmt"

func main() {
	months := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := months[2:5]
	fmt.Println(slice1)
	fmt.Println(cap(slice1)) //capacity
	fmt.Println(len(slice1)) //panjang array
}
