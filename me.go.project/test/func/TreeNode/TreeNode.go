package TreeNode

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	return nil
}



func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > R {
		trimBST(root.Left,L,R)
	}
	if root.Val < L {
		trimBST(root.Right,L,R)
	}
	root.Left = trimBST(root.Left,L,R)
	root.Right = trimBST(root.Right,L,R)
	return root

}

func DistributeCandies(candies []int) int {
	l := len(candies) / 2
	m := make(map[int]int)
	for _,v := range candies {
		m[v] = m[v] + 1
	}
	l2 := len(m)
	if l2 < l {
		l = l2
	}
	return l
}


func singleNumber(nums []int) int {
	count := 0
	for i:=0;i<len(nums);i++{
		count ^= nums[i]
	}
	return count
}


//682,棒球分数
func calPoints(ops []string) int {
	arr := make([]int,len(ops))
	count := 0
	score := 0
	for _,v := range ops {
		switch v {
			case "+":
				arr[count] = arr[count-1] + arr[count-2]
				score += arr[count]
				count ++
			case "C":
				score -= arr[count-1]
				count --
			case "D":
				arr[count] = arr[count-1] * 2
				score += arr[count]
				count ++
			default:
				arr[count],_ = strconv.Atoi(v)
				score += arr[count]
				count ++
		}
	}
	return score
}

//762. 二进制表示中质数个计算置位
func countPrimeSetBits(L int, R int) func() int {
	count := 0
	for i:=L;i<=R;i++  {
		count++
	}
	return func() int {
		if count == 2 || count == 3 || count == 5 || count == 7 ||
			count == 11 || count == 13 || count == 17 || count == 19 {
			return count
		}
		return 0
	}

}


//169. 求众数
func majorityElement(nums []int) int {
	count := 0
	result := 0
	for _,v := range nums {
		if count == 0 {
			 result = v
		}
		if result == v {
			count ++
		}else {
			count --
		}
	}
	return result
}
