package main

import "fmt"

func main() {

	for counter := 1; counter < 10; counter++ {
		fmt.Println("perulangan ke -", counter)
	}

	slice := []string{"ahmad", "salman", "firdaus", "tado"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

}
