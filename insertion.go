package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	pl("Inserton Sort!")

	arr := []int{9, 4, 2, 7, 1}

	for index := 0; index < len(arr); index++ {
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
	}
	pl(arr)
}
