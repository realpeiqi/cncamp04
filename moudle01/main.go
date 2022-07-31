package main

import (
	"fmt"
)

// 编写一个小程序：
// 给定一个字符串数组
// [“I”,“am”,“stupid”,“and”,“weak”]
// 用 for 循环遍历该数组并修改为
// [“I”,“am”,“smart”,“and”,“strong”]

func Exercise1() {
	inputList := []string{"I", "am", "stupid", "and", "weak"}
	for i, v := range inputList {
		if v == "stupid" {
			inputList[i] = "smart"
		}
		if v == "weak" {
			inputList[i] = "strong"
		}
	}
	fmt.Println(inputList)
}

func main(){

	Exercise1()

}
