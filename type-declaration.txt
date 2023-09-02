package main

import "fmt"

func main() {
	type noKTP string
	type statusMeriage bool

	var noKTPEko noKTP = "324567674"
	var statusEko statusMeriage = true

	fmt.Println(noKTPEko)
	fmt.Println(statusEko)
}
