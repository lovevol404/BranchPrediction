package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const length  = 10000000
func main() {
	var nums []int
	for i := 0; i < length; i++ {
		nums =append(nums,rand.Intn(length * 10))
	}

	indexNum := length * 5

	start := time.Now().Nanosecond()
	count := 0
	for _,i := range nums{
		if i > indexNum{
			count ++
		}
	}

	sort.Ints(nums)
	end := time.Now().Nanosecond()
	fmt.Printf("count: %d,耗时：%d\n", count, (end-start)/1000)

	start = time.Now().Nanosecond()
	count = 0
	for _,i := range nums{
		if i > indexNum{
			count ++
		}
	}
	end = time.Now().Nanosecond()
	fmt.Printf("count: %d,耗时：%d\n", count, (end-start)/1000)

}