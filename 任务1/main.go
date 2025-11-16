package main

import (
	"fmt"
	"strings"
)

// 两数之和
func twoSum(nums []int, target int) []int {
	var m1 map[int]int = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var _, ok = m1[nums[i]]
		if ok {
			return []int{m1[nums[i]], i}

		}
		m1[target-nums[i]] = i
	}
	return []int{}
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	var a int = 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[a] = nums[i+1]
			a++
		}
	}
	return a
}

func xor(numArr []int) int {
	var a int = 0
	for i := 0; i < len(numArr); i++ {
		a = a ^ numArr[i]
		fmt.Println("use len(), index = ", i, "value = ", numArr[i])
	}
	return a
}

// 判断输入的数据是否为回文数,如果是则返回true
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	var revsersed int = 0

	for x > revsersed {
		revsersed = revsersed*10 + x%10
		x = x / 10
	}
	return x == revsersed || x == revsersed/10

}

// 是否回文数
func isValid(str string) bool {
	//左括号必须是在最左边
	if len(str) == 0 {
		return false
	}

	for index, v := range str {
		fmt.Printf("index: %d, char: %c\n", index, v)
		if index == 0 && v != '(' {
			return false
		}
	}
	var lastStr string
	//1.只存在()
	var str1 string = strings.Replace(str, "()", "", -1)
	if str1 == "" {
		return true
	}
	rs := []rune(str)
	if len(rs) > 0 {
		var lastStrChar rune = rs[len(rs)-1]
		fmt.Printf("len(rs): %d, lastChar: %c\n", len(rs), lastStrChar)
		lastStr = string(lastStrChar)
	}

	if strings.Contains(str, "{}") && strings.Contains(str, "{}") && (str == "()[]{}") {
		return true
	}
	if strings.Contains(str, "[]") &&
		strings.Contains(str, "(") && strings.Contains(str, ")") &&
		(!strings.Contains(str, "{") || !strings.Contains(str, "}")) && lastStr == ")" {
		return true
	}
	return false
}

// 公共前缀
func commonPrefix(arrayStr []string) string {
	//数组为空或低于两个字符串数量的数组
	if len(arrayStr) < 2 {
		return ""
	}
	//遍历循环所有字符串,以数组的第一个字符串为开头去模拟前缀匹配，匹配不上就移除不匹配的字符
	// 初始化第一个字符串的默认值
	var firstStr string = arrayStr[0]
	if firstStr == "" {
		return ""
	}

	for i := 0; i < len(arrayStr); i++ {
		fmt.Println("use len(), index = ", i, "value = ", arrayStr[i])

		if !strings.HasPrefix(arrayStr[i], firstStr) {
			//截取字符串最后一位，二次匹配
			firstStr = firstStr[:len(firstStr)-1]
			i--
		}
	}
	return firstStr
}

// 大整数加1
func addBigInteger(digits []int) []int {
	//末尾不为9就正常加，否则扩容进位
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	//如果为9 则需要重置当前位为0，并且扩容数组，go的数组不可扩容,需转化为切片
	var sliece1 []int = make([]int, len(digits)+1)
	sliece1[0] = 1
	return sliece1
}

func main() {
	// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	// 异或理论，0^任何数为自己，^自己为
	array := []int{1, 2, 1, 2, 5}
	fmt.Println("只出现一次的元素", xor(array))

	// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	// 例如，121 是回文，而 123 不是。
	fmt.Println("122是否为回文数", isPalindrome(122))

	// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	// 有效字符串需满足：
	// 左括号必须用相同类型的右括号闭合。
	// 左括号必须以正确的顺序闭合。()[]{}// 每个右括号都有一个对应的相同类型的左括号。
	fmt.Println("([])是否为回文数", isValid("([)]"))

	// 最长公共前缀
	// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。
	//位置和字符都一致，且都作为前缀
	a := []string{"flower", "flow", "flight"}
	fmt.Println("flower,flow,flight最长公共前缀", commonPrefix(a))

	// 	给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	// 将大整数加 1，并返回结果的数字数组。
	var s4 []int = []int{1, 2, 9}
	fmt.Println("结果的数字数组", addBigInteger(s4))

	// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
	// 返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
	// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
	//不参与排序，原则上移除，相同元素出现第一次以后得数字，并且优先遍历的数字应该始终位于
	var s5 []int = []int{1, 2, 3, 3, 4, 4, 5}
	fmt.Println("唯一的数字出现次数", removeDuplicates(s5))

	// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
	var s6 []int = []int{2, 7, 11, 15}
	fmt.Println("唯一的数字出现次数", twoSum(s6, 9))

}
