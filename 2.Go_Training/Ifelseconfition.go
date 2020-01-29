package main

import (
	"fmt"
)

func main() {
	var x = "suresh"
	var y = "Asutosh"
	var array = []string{"Suresh", "Asutosh", "Biswa"}

	for i := 0; i < len(array)-1; i++ {
		if array[i] == x {
			fmt.Println("suresh - Manager")
		} else if array[i] == y {
			fmt.Println("asutosh - Automation Engineer")
		} else {
			fmt.Println(array[i])
		}
	}
}
