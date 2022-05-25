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

