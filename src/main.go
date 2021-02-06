package main

import (
	"fmt"
	"github.com/alok87/goutils/pkg/random"
	_ "math/rand"
)

func make_arr(sl []int, len int) []int {
	for i := 0; i < len; i++ {
		val := random.RangeInt(0,2,100)
		sl = val
		//fmt.Println(val)
	}
	return sl
}

func longest_subarray_all_1s(arr []int) int {
	//longestArray := 0
	fmt.Println("finding longest subarray")
	//arr = make_arr(arr, 100)
	return 420
}

func main() {
	var slice []int
	fmt.Println("yeetus")
	//slice = longest_subarray_all_1s(slice)
	slice = make_arr(slice, 100)
	fmt.Println(slice)
}