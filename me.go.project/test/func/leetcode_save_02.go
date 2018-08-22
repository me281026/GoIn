package main

import (
	"fmt"
)

const num = 50

func main() {
	f := fibfibonacci()
	var arr [num]int
	for i := 0; i < num; i++ {
		arr[i] = f()
	}
	fmt.Println(arr)
}

func fibfibonacci() func() int {
	num01, num02 := 0, 1
	return func() int {
		num01, num02 = num02, num02+num01
		return num01
	}
}

//9. 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	return (x == y || x == y/10)
}

//283. 移动零
func moveZeroes(nums []int) {
	arr := nums[:0]
	num := 0
	for _, v := range nums {
		if v != 0 {
			arr = append(arr, v)
		} else {
			num++
		}
	}
	for num > 0 {
		arr = append(arr, 0)
		num--
	}
}
