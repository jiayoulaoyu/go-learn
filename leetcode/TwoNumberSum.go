package main

import "fmt"

/**
两数之和
*/
func main() {
	fmt.Println(twoSum([]int{1, 3, 4}, 4))

}

func twoSum(nums []int, target int) []int {
	tab := map[int]int{}
	for i, v := range nums {
		if p, ok := tab[target-v]; ok {
			return []int{p, i}
		}
		tab[v] = i
	}

	return nil
}
