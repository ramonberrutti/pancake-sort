package main

import (
	"fmt"
)

func flip(arr []int, k int) {
	for i := 0; i < k/2; i++ {
		arr[i], arr[k-i-1] = arr[k-i-1], arr[i]
	}
}

func pancakeSort(arr []int) {
	arrLen := len(arr)

	var indexFlip int
	for i := arrLen - 1; i >= 0; i-- {
		indexFlip = i
		for j := i; j >= 0; j-- {
			if arr[indexFlip] < arr[j] {
				indexFlip = j
			}
		}

		flip(arr, indexFlip+1)
		flip(arr, i+1)
	}
}

func main() {
	arr := []int{3, 4, 1, 2, 8, 7, 9, 10, 50, 5, 500, 66, 43}

	fmt.Println(arr)
	pancakeSort(arr)
	fmt.Println(arr)
}
