package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	pl("Inserton Sort!")

	arr := []int{9, 4, 2, 7, 1}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; arr[j] > arr[j+1] && j >= 0; {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			if j > 0 {
				j--
			}
		}
	}
	pl(arr)
}
