package main
import "fmt"

func search(nums []int, target int) int {
	left, right := 0, 0
	idx := -1
	for i:=0; i<(len(nums)/2); i++ {
		left = i
		right = len(nums) - 1 - i
		if nums[left] == target {
			idx = left
		} else if nums[right] == target {
			idx = right
		} 
	}
	return idx
}

func main() {
	nums := []int{-1, 0, 3, 5, 23, 12, 20, 5, 4, 3, 2, 1, 9, -1,  0, 3, 5, 23, 12, 20, 5, 4,}
	target := 9
	fmt.Println(search(nums, target))
}