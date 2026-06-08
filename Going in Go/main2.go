package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var limit int
// 	fmt.Print("Enter the Slice Limit: ")
// 	fmt.Scan(&limit)
// 	nums := []int{}
// 	for i := 1; i <= limit; i++ {
// 		nums = append(nums, i)
// 	}
// 	fmt.Println(nums)
// }

// func main() {
// 	for {
// 		go greet()
// 	}
// }
// // spawn infinite GO routines
// func greet() {
// 	fmt.Println("Hi, how are you?")
// }

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 1; i <= 3000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
		defer wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 3000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
		defer wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 3000; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println(counter)
}

/*
AI recommendations :
defer mu.Unlock() more reliable
use of atomic counter preffered
*/
