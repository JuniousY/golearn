package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者: 生成 factor 整数倍的序列
func Producer(id string, factor int, out chan<- string) {
	for i := 0; ; i++ {
		out <- fmt.Sprintf("%s: %v", id, i*factor)
		time.Sleep(50 * time.Millisecond)
	}
}

// 消费者
func Consumer(in <-chan string) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan string, 64) // 成果队列

	go Producer("a", 3, ch) // 生成 3 的倍数的序列
	go Producer("b", 5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出
	// time.Sleep(5 * time.Second)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
