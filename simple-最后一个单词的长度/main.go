// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/length-of-last-word
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


//此题可直接变成问返回单词？？？？？？？？？？？？？（可用此法切割返回）
package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := len(s)
	u := length - 1
	var sum int
	sum = 0
  //循环排空
	for s[u] == ' ' {
		u--

	}
  //开始计算最后一个
	for s[u] != ' ' {
		u -- 
		sum++
	}
	
	return sum

}

func main() {
	s := " trsr yo lll     "
	k := lengthOfLastWord(s)
	fmt.Println(k)
}

