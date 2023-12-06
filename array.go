package main

import "fmt"


func main() {
	name := [...]string{
		"budi",
		"tono",
	}

	var names = [5]int{1, 2, 3}

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])

	fmt.Println(names[1])
}
