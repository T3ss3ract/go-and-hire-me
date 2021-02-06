package main

import (
	"fmt"
	"github.com/alok87/goutils/pkg/random"
	_ "math/rand"
)

func make_arr(sl []int, len int) []int {
	for i := 0; i < len; i++ {
		val := random.RangeInt(0,2,100)	//works to make random nums every time!
		sl = val
		//fmt.Println(val)
	}
	return sl
}

func longest_subarray_all_1s(arr []int) int {
	longestArray := 0
	currArraySize := 0
	fmt.Println("finding longest subarray")
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			currArraySize++
		}
		if arr[i] == 0 {
			if longestArray < currArraySize {
				longestArray = currArraySize
			}
			currArraySize = 0
		}
	}
	return longestArray
}

func main() {
	var slice []int
	fmt.Println("yeetus")
	//slice = longest_subarray_all_1s(slice)
	slice = make_arr(slice, 100)
	fmt.Println(slice)
	fmt.Println(longest_subarray_all_1s(slice))
}