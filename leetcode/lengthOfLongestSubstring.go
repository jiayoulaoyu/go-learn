package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("dvdf")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	_map := make(map[byte]int)
	arr := []byte(s)

	var res, start int
	for i, v := range arr {
		preIdx, ok := _map[v]
		if ok {
			start = max(start, preIdx+1)
		}
		_map[v] = i
		res = max(res, i-start+1)
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
