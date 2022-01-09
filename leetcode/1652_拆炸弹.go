package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//fmt.Println(sum([]int{5, 7, 1, 4}, 3, 3))
	fmt.Println(decrypt([]int{10, 5, 7, 7, 3, 2, 10, 3, 6, 9, 1, 6}, -4))
	fmt.Println(getIndex(-1, len([]int{10, 5, 7, 7, 3, 2, 10, 3, 6, 9, 1, 6})))

}

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	fmt.Println(reflect.TypeOf(res))
	for i := range code {
		res[i] = sum(code, k, i)
	}
	return res
}
func sum(code []int, k, currentIdx int) int {
	maxLen := len(code)
	if k == 0 {
		return 0
	}
	var sum int
	if k < 0 {
		for i := k; i < 0; i++ {
			idx := getIndex(i+currentIdx, maxLen)
			sum = code[idx] + sum
		}
	} else {
		for i := 0; i < k; i++ {
			idx := getIndex(i+currentIdx+1, maxLen)
			sum = code[idx] + sum
		}
	}
	return sum
}

func getIndex(idx, maxLen int) int {
	if idx < 0 {
		idx = int(math.Abs(float64((maxLen + idx) % maxLen)))
	}
	return idx % maxLen
}
