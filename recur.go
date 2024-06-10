package main

import "fmt"

func _sum(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	return _sum(nums[1:]...) + nums[0]
}

func main() {
	fmt.Println(_sum(1, 2, 3))
}