package main

import (
	"fmt"
	"math/rand"
	"time"
)

const length = 1000000

func main() {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(length * 10)
	}

	indexNum := length * 5
	count := 0

	start := time.Now().UnixNano()
	for j := 0; j < 10000; j++ {
		for i := 0; i < length; i++ {
			n := nums[i]
			if n > indexNum {
				count += n
			}
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("count: %d,耗时：%d\n", count, (end-start)/1000)

	quickSort(nums, 0, len(nums)-1)
	count = 0

	start = time.Now().UnixNano()
	for j := 0; j < 10000; j++ {
		for i := 0; i < length; i++ {
			n := nums[i]
			if n > indexNum {
				count += n
			}
		}
	}
	end = time.Now().UnixNano()
	fmt.Printf("count: %d,耗时：%d\n", count, (end-start)/1000)

}

func quickSort(a []int, left int, right int) {
	if left >= right {
		return
	}
	temp := a[left]
	start := left
	stop := right
	for right != left {
		for right > left && a[right] >= temp {
			right--
		}
		for left < right && a[left] <= temp {
			left++
		}
		if right > left {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[right], a[start] = temp, a[right]
	quickSort(a, start, left)
	quickSort(a, right+1, stop)
}
