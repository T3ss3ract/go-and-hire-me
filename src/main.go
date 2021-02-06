package main

import (
	"fmt"
	"github.com/alok87/goutils/pkg/random"
	"math/rand"
	"time"
)

func make_arr(sl []int, len int) []int {
	for i := 0; i < len; i++ {
		val := random.RangeInt(0,2,100)	//works to make random nums every time!
		sl = val
		//fmt.Println(val)
	}
	return sl
}

/**
	find longest chain of 1s in this array
	part of an interview question
 */
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

func generate_quicksort_slice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice

}


func __interview_quicksort() {

}

func main() {
	var slice []int
	fmt.Println("yeetus")
	//slice = longest_subarray_all_1s(slice)
	slice = make_arr(slice, 50)
	fmt.Println(slice)
	fmt.Println(longest_subarray_all_1s(slice))
}