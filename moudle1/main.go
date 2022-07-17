package module01

import (
	"fmt"
	"time"
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

// 基于 Channel 编写一个简单的单线程生产者消费者模型：
// 队列：
// 队列长度 10，队列元素类型为 int
// 生产者：
// 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
// 消费者：
// 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

func Exercise2() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			<-ticker.C
			ch <- i
			fmt.Printf("put %d into channel\n", i)
		}
		defer close(ch)
	}()

	for v := range ch {
		time.Sleep(time.Second)
		fmt.Printf("get %d out of channel\n", v)
	}
}
