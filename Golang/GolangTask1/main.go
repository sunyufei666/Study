package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {

	/******************** 控制流程 ********************/
	// 1.只出现一次的数字
	// OnlyOneTimeNumber()

	// 2.回文数
	// PalindromeNumber()

	/******************** 字符串 ********************/
	// 1.括号匹配
	// BracketMatching("({[]})")

	// 2.最长公共前缀
	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"dog", "racecar", "car"}
	// fmt.Println(LongestCommonPrefix(strs))

	/******************** 基本值类型 ********************/
	// 1.加一
	// digits := []int{1, 2, 3}
	// digits := []int{9}
	// fmt.Println(DigitPlusOne(digits))

	/******************** 引用类型：切片 ********************/
	// 1.删除有序数组中的重复项
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums := []int{1, 1, 2}
	// fmt.Println(DeleteRepeatItem(nums), nums)
	// 2.合并区间
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// intervals := [][]int{{4, 7}, {1, 4}}
	// fmt.Println(mergeInterval(intervals))

	/******************** 基础 ********************/
	// nums, target := []int{2, 7, 11, 15}, 9
	// nums, target := []int{3, 2, 4}, 6
	nums, target := []int{3, 3}, 6
	fmt.Println(twoSum(nums, target))
}

func OnlyOneTimeNumber() {
	/*
		给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
		可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，
		然后再遍历 map 找到出现次数为1的元素.
	*/
	arr := [...]int{1, 1, 3, 5, 5}

	timeMap := make(map[int]int)

	for _, ele := range arr {
		timeMap[ele] += 1
	}

	for ele, times := range timeMap {
		if times == 1 {
			fmt.Printf("出现一次的元素为：%v\n", ele)
		}
	}
}

func PalindromeNumber() {
	num := 21233212
	flag := true
	strArr := strconv.FormatInt(int64(num), 10)

	for i := 0; i < len(strArr); i++ {
		if strArr[i] != strArr[len(strArr)-i-1] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Println("此数字是回文数字")
	} else {
		fmt.Println("此数字不是回文数字")
	}

}

func BracketMatching(bracketStr string) bool {
	if len(bracketStr) <= 0 || len(bracketStr)%2 != 0 {
		return false
	}

	arr := []byte{}
	mapping := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	for i := 0; i < len(bracketStr); i++ {
		condition1 := bracketStr[i] == '(' || bracketStr[i] == '[' || bracketStr[i] == '{'
		condition2 := bracketStr[i] == ')' || bracketStr[i] == ']' || bracketStr[i] == '}'
		if len(arr) == 0 && condition2 {
			return false
		} else if condition1 {
			arr = append(arr, bracketStr[i])
		} else if condition2 {
			if mapping[arr[len(arr)-1]] == bracketStr[i] {
				arr = arr[:len(arr)-1]
			}
		}
	}
	fmt.Println(len(arr) == 0)
	return len(arr) == 0
}

func LongestCommonPrefix(strs []string) (commonStr string) {
	minLen, minIndex := 0, 0
	for i, v := range strs {
		if minLen > len(v) {
			minLen = len(v)
			minIndex = i
			break
		}
	}
loop:
	for i := 0; i < len(strs[minIndex]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[minIndex][i] != strs[j][i] {
				break loop
			}
		}
		commonStr += string(strs[minIndex][i])
	}
	return
}

func DigitPlusOne(digits []int) []int {
	var num int
	for i := 0; i < len(digits); i++ {
		temp := int(math.Pow10(len(digits) - i - 1))
		num += digits[i] * temp
	}
	num++
	numStr := strconv.FormatInt(int64(num), 10)
	ret := []int{}
	for _, v := range numStr {
		temp, _ := strconv.ParseInt(fmt.Sprintf("%c", v), 10, 64)
		ret = append(ret, int(temp))
	}
	return ret
}

func DeleteRepeatItem(nums []int) int {
	a := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[a-1] {
			nums[a] = nums[i]
			a++
		}
	}
	return a
}

func mergeInterval(intervals [][]int) [][]int {
	// 先把区间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	retArr := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= retArr[len(retArr)-1][1] {
			if intervals[i][1] >= retArr[len(retArr)-1][1] {
				retArr[len(retArr)-1][1] = intervals[i][1]
			}
		} else {
			retArr = append(retArr, intervals[i])
		}
	}
	return retArr
}

func twoSum(nums []int, target int) []int {
	ret := []int{}
	tempMap := make(map[int][]int)
	for i, v := range nums {
		tempMap[v] = append(tempMap[v], i)
	}
	for num, index := range tempMap {
		if target-num == num {
			if len(tempMap[num]) > 1 {
				ret = append(ret, tempMap[num][0], tempMap[num][1])
			}
		} else {
			another, ok := tempMap[target-num]
			if ok {
				ret = append(ret, index[0], another[0])
				break
			}
		}
	}
	return ret
}
