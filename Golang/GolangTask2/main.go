package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.length
}

func (rec Rectangle) Perimeter() float64 {
	return (rec.width + rec.length) * 2
}

type Circle struct {
	radius float64
}

func (cir Circle) Area() float64 {
	return math.Pi * math.Pow(cir.radius, 2)
}

func (cir Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID string
	Person
}

func (ep Employee) PrintInfo() {
	str, _ := json.Marshal(ep)
	fmt.Println("员工的信息为：", string(str))
}

var wg sync.WaitGroup
var mutx sync.Mutex

func main() {
	/************************** 指针 **************************/
	// 1. 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
	// x := 10
	// modifyValue(&x)
	// fmt.Println(x)
	// 2. 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
	// arr := []int{2, 4, 8, 16}
	// modifySlice(&arr)
	// fmt.Println(arr)

	/************************** Gorutine **************************/
	// 1. 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
	// wg.Add(1)
	// go printOddNumber()
	// wg.Add(1)
	// go printEvenNumber()
	// wg.Wait()
	// 2. 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		timeStart := time.Now().UnixMilli()
	// 		random := rand.IntN(20)
	// 		for i := 0; i < random; i++ {
	// 			time.Sleep(10 * time.Millisecond)
	// 		}
	// 		timeEnd := time.Now().UnixMilli()
	// 		fmt.Printf("任务%v的执行时间为%vms\n", i+1, timeEnd-timeStart)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	/************************** 面向对象 **************************/
	// 1. 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
	// var shape1, shape2 Shape
	// rec := Rectangle{
	// 	width: 3,
	// 	length: 5,
	// }
	// shape1 = rec
	// fmt.Printf("矩形面积为：%v，周长为：%v\n", shape1.Area(), shape1.Perimeter())
	// cir := Circle{
	// 	radius: 3,
	// }
	// shape2 = cir
	// fmt.Printf("圆形面积为：%v，周长为：%v\n", shape2.Area(), shape2.Perimeter())
	// 2. 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
	// employee := Employee{
	// 	EmployeeID: "20250919",
	// 	Person: Person{
	// 		Name: "张三",
	// 		Age:  18,
	// 	},
	// }
	// employee.PrintInfo()

	/************************** Channel **************************/
	// 1. 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
	// ch := make(chan int)
	// wg.Add(1)
	// go func(ch chan int) {
	// 	for i := 1; i <= 10; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Add(1)
	// go func(ch chan int) {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}
	// 	wg.Done()
	// }(ch)
	// wg.Wait()
	// 2. 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
	// ch := make(chan int, 100)
	// wg.Add(1)
	// go func(ch chan int) {
	// 	for i := 1; i <= 100; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// 	wg.Done()
	// }(ch)
	// wg.Add(1)
	// go func(ch chan int) {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}
	// 	wg.Done()
	// }(ch)
	// wg.Wait()

	/************************** 锁机制 **************************/
	// 1. 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	// count := 0
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(count *int) {
	// 		mutx.Lock()
	// 		defer mutx.Unlock()
	// 		for i := 0; i < 1000; i++ {
	// 			*count++
	// 		}
	// 		wg.Done()
	// 	}(&count)
	// }
	// wg.Wait()
	// fmt.Println("计数器的值为：", count)
	// 2. 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	var count int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&count, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("计数器的值为：", count)
}

func modifyValue(x *int) {
	*x += 10
}

func modifySlice(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}

func printOddNumber() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Printf("奇数%v\n", i)
		}
	}
	wg.Done()
}

func printEvenNumber() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("偶数%v\n", i)
		}
	}
	wg.Done()
}
