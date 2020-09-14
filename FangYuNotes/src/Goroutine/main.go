package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//goroutine 协程
//使用并发或者并行的方式，将统计素数的任务分配给多个Goroutine去完成
//进程就是操作系统的一次执行过程，是系统进行资源分配和调度的基本单位
//线程是进程的一个执行实例，是程序执行的最小单位，是比进程更小的能独立运行的基本单位
//一个进程可以创建和销毁多个线程，同一个进程的多个线程可以同时并发执行
//一个程序至少有一个进程，一个进程至少有一个线程
//并发和并行
//多线程程序在单核上运行，叫并发，微观的角度上来说，其实只有一个程序在运行
//多线程程序在多核上运行，叫并行，微观的角度上来说，有多个程序在同时运行
//协程是轻量级的线程【编译馆做优化】
//go协程的特点：独立的栈空间，共享程序堆空间，调度由用户控制，协程是轻量级的线程

//MFG模式的介绍
//M操作系统的主线程，P执行过程中的上下文环境，G协程
func test()  {
	for i:=0 ;i<6 ;i++{
		fmt.Println("Hello test !" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main()  {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	//go test()  //开启协程
	//go test()
	//for i:=0 ;i<10 ;i++{
	//	fmt.Println("Hello main !" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}


}
