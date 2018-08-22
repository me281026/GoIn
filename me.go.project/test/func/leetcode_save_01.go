package main

import (
		"strings"
	"strconv"
	"fmt"
)

func main() {

	//a := []string{"200 www.me.com"}
	//fmt.Println(subdomainVisits(a))
	//fmt.Println(fizzBuzz(15))
	//c := getSum(1,2)
	//fmt.Println(c)
	//( 1 - 0000 0001 )(2 - 0000 0010 )(3 - 0000 0011)
	//( 1^2 - 0000 0011 ) ( 1&2 - 0000 0000 )
	//fmt.Println(1^2)
	nums := [][]int{{1,2},{3,4}}
	fmt.Println(matrixReshape(nums,1,4))
}


func titleToNumber(s string) int {
	arr := []rune(s)
	count := 0
	for i := range arr {
		count = (count * 26)  + ( int(arr[i]) - 64 )
	}
	return count
}

func subdomainVisits(cpdomains []string) []string {
	rt := []string{}
	hm := make(map[string]int)
	for _,v := range cpdomains {
		v1 := strings.Fields(v)
		num := v1[0]
		yu := v1[1]
		arr := strings.Split(yu,".")
		for m := len(arr) - 1; m >= 0 ; m-- {
			r1 := strings.Join(arr[m:],".")
			n1,_ := strconv.Atoi(num)
			hm[r1] += n1
		}
	}
	for a,b := range hm  {
		rt = append(rt,strings.Join([]string{strconv.Itoa(b),a}," "))
	}
	
	return rt
}



func fizzBuzz(n int) []string {
	arr := make([]string,n)
	for i := 1; i <= n ;i++  {
		if i % 3 == 0 && i % 5 == 0 {
			arr[i-1] = "FizzBuzz"
		} else if i % 3 == 0 {
			arr[i-1] = "Fizz"
		} else if i % 5 == 0 {
			arr[i-1] = "Buzz"
		} else {
			arr[i-1] = string(i)
		}

	}
	return arr
}

//两数之和,不使用+和-
func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b,(a&b)<<1)
}



func matrixReshape(nums [][]int, r int, c int) [][]int {
	a := len(nums)
	b := len(nums[0])
	if a*b != r*c || a==1 && b == 1 {
		return nums
	}
	arr := make([][]int,r)
	for i:=0;i<r;i++ {
		for j:=0;j<c;j++ {
			c := j + c*i
			arr[i][j] = nums[c/b][c%b]
		}
	}
	return arr
}





