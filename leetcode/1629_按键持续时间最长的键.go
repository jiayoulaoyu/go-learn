package main

import "fmt"

func main() {
	fmt.Printf("%c", slowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Println("")
	fmt.Printf("%c", slowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxTc := releaseTimes[0]
	var maxIdx int
	for i := 1; i < len(releaseTimes); i++ {
		temp := releaseTimes[i]
		tc := temp - releaseTimes[i-1]
		if tc > maxTc {
			maxIdx = i
			maxTc = tc
		} else if tc == maxTc && keysPressed[i] > keysPressed[maxIdx] {
			maxIdx = i
			maxTc = tc
		}
	}
	return keysPressed[maxIdx]

}

func swap(a, b int, arr []int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
