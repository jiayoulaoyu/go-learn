package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("mgntdygtxrvxjnwksqhxuxtrv"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	hash := make(map[uint8]int)
	var maxLength = -1
	for i := range s {
		preIdx, ok := hash[s[i]]
		if ok && i-preIdx-1 > maxLength {
			maxLength = i - preIdx - 1
		} else if !ok {
			hash[s[i]] = i
		}
	}
	return maxLength
}
