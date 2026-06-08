package main

import "fmt"

func main() {
	var limit int
	fmt.Println("Enter the Slice Limit: ")
	fmt.Scan(&limit)
	nums := []int{}
	for i := 1; i <= limit; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
}
