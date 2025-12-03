package main

import (
	"fmt"
	"sort"
)

func main() {

	//onlyNumber()
	//palindrome()
	//validStr()
	//longestCommonPrefix()
	//plusOne()
	//removeDuplicates()
	//merge()
	twoSum()
}

func twoSum() {
	/**
	两数之和
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
	*/

	num := []int{2, 7, 11, 15}
	target := 18
	fmt.Println("原始数据", num, target)

	needMap := make(map[int]int)
	for i := 0; i < len(num); i++ {
		needNum := target - num[i]
		if ni, ok := needMap[needNum]; !ok {
			needMap[num[i]] = i
		} else {
			fmt.Printf("两个数之和为%d的索引为[%d,%d], 两个数分别是%d和%d\n", target, ni, i, num[i], needNum)
			return
		}
	}
}

func merge() {
	/*
		56. 合并区间
		以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
	*/
	intervals := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	fmt.Println("原始数据", intervals)

	//按区间的起始值升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("排序后的数据", intervals)

	// 将第一个区间先存入结果
	result := [][]int{intervals[0]}

	// 从第二个区间开始遍历
	for _, v := range intervals[1:] {
		// 如果当前区间的开始位置大于结果数组最后一个区间的结束位置，则将当前区间更新到数组中
		if v[0] < result[len(result)-1][1] {
			result[len(result)-1][1] = v[1]
		} else {
			result = append(result, v)
		}
	}
	fmt.Println(result)
}

func removeDuplicates() {
	/*
		26. 删除有序数组中的重复项
		给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
		nums 已按 非递减 顺序排列。
	*/

	numbers := []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8, 9, 9, 10}
	fmt.Println("原始数据", numbers)

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] == numbers[i+1] {
			numbers = append(numbers[:i], numbers[i+1:]...)
		}
	}
	fmt.Println("去重后的数据", numbers)
}

func plusOne() {
	digits := []int{9, 9, 9, 9, 9}
	fmt.Println("原始数据", digits)
	// 初始化进位为1（因为要加1）
	carry := 1
	// 从数组末尾（最低位）向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位 = 原数字 + 进位
		sum := digits[i] + carry
		// 更新当前位的值（取余，处理10的情况）
		digits[i] = sum % 10
		// 更新进位（整除，只有sum>=10时进位为1）
		carry = sum / 10
		// 进位为0，无需继续遍历
		if carry == 0 {
			break
		}
	}

	// 如果遍历完所有位后仍有进位（说明所有位都是9），在开头插入1
	if carry == 1 {
		// 新建数组，首元素为1，拼接原数组
		digits = append([]int{1}, digits...)
	}

	fmt.Println("+1后的数据", digits)
}

func longestCommonPrefix() {
	// 最长公共前缀
	// 编写一个函数来查找字符串数组中的最长公共前缀。
	// 如果不存在公共前缀，返回空字符串 ""。
	strArr := []string{"flower", "flow", "floight"}
	prefix := make([]rune, 0)

forStr:
	for i, r := range strArr[0] {
		for _, str := range strArr {
			if r != rune(str[i]) {
				break forStr
			}
		}
		prefix = append(prefix, r)
	}
	fmt.Println(string(prefix))

}

func validStr() {
	// 有效的括号
	// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	str := "{(){}[[]]}"
	bracketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	valid := func(str string) bool {
		// 切片是有序的，可能根据切片的后进先出的原则，去判断
		bracket := []rune{}
		// range 字符串迭代的是rune
		for _, val := range str {
			if v, ok := bracketMap[val]; ok {
				// 压入栈
				bracket = append(bracket, v)
			} else {
				if len(bracket) == 0 {
					return false
				}
				// 获取栈顶元素
				s := bracket[len(bracket)-1]
				// 弹出栈
				bracket = bracket[:len(bracket)-1]

				if s != val {
					return false
				}
			}

		}
		return true
	}

	fmt.Println(str, "进行字符串验证", valid(str))

}

func palindrome() {
	// 回文数
	num := 12221

	isPalindrome := func(num int) bool {
		// 复数不可能是回文数
		if num < 0 {
			return false
		}
		// 方式一.将数字转为字符串，然后将字符串转为rune数组，再倒序
		//var numRune []rune = []rune(strconv.Itoa(num))
		//var reversalNumRune []rune = []rune{}
		//for i := len(numRune) - 1; i >= 0; i-- {
		//	reversalNumRune = append(reversalNumRune, numRune[i])
		//}
		//reversalNum, _ := strconv.Atoi(string(reversalNumRune))
		//fmt.Println("reversalNum", reversalNum)
		//return reversalNum == num

		// 方式二，将数字转为字符串，然后将字符串转为byte数组，再倒序
		s := fmt.Sprintf("%d", num)
		var numBytes []byte = []byte(s)
		reversalNumBytes := make([]byte, 0)
		for i := len(numBytes) - 1; i >= 0; i-- {
			reversalNumBytes = append(reversalNumBytes, numBytes[i])
		}
		fmt.Println(string(reversalNumBytes))
		return s == string(reversalNumBytes)
	}
	fmt.Printf("数字%d,回文检查为%v\n", num, isPalindrome(num))
}

func onlyNumber() {

	// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

	numbers := []int{
		4, 1, 2, 1, 2,
	}

	numMap := make(map[int]int)
	for _, val := range numbers {
		numMap[val]++
	}

	for num, count := range numMap {
		if count == 1 {
			fmt.Println(num)
		}
	}
}
