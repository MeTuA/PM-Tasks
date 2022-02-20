package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		// начала изменения
		fmt.Println("any string")
		return
		// конец изменения
	}
}

// second variant
/*
func main() {
	for i := 0; i < 5; i++ {
		// начала изменения
		fmt.Println("any string")
		break
		// конец изменения
	}
}
*/
