package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 4, 5, 6, 8, 2}

	insertion(arr)
}

func insertion(bil []int) {
	var n = len(bil)
	charts(bil)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if bil[j-1] > bil[j] {
				time.Sleep(1 * time.Second)
				bil[j-1], bil[j] = bil[j], bil[j-1]
				charts(bil)
			}
			j = j - 1
		}
	}
}

func charts(bil []int) {
	for i := len(bil); i >= -1; i-- {
		for j := 0; j < len(bil); j++ {
			if bil[j] > i && i >= 0 {
				fmt.Printf("%3s", "|")
			} else if i == -1 {
				fmt.Printf("%3d", bil[j])
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}
}

func max(bil []int) int {
	var n int

	for i := 0; i < len(bil); i++ {
		if bil[i] > n {
			n = bil[i]
		}
	}

	return n
}
