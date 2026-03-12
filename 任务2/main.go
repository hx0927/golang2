package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 引用的传递
func byPointer(x *int) {
	*x = *x + 10
}

// 值传递
func byValue(y int) {
	y = y + 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
// 初始化切片
func splitSlice(z *[]int) {
	var slice []int = *z
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * 2
	}
}

func getOddnumber() {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("打印1到10的奇数", i)
		}
	}
}

func getEvenNumber() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("打印1到10的偶数", i)
		}
	}
}

// 生产者，消费者
func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(100 * time.Microsecond)
	}
	close(ch)
}

// 消费者
func consumerTwo(ch <-chan int, id int) {
	//记录执行开始时间
	start := time.Now()
	for v := range ch {
		fmt.Printf("Comsumer %d receiver: %d\n", id, v)
		time.Sleep(50 * time.Microsecond)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Comsumer %d  执行时间 %d\n", id, elapsed)

}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 长方形结构体
type Rectangle struct {
	width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.Height)
}

// 圆形结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape, name string) {
	fmt.Printf("%s:Area=%.2f,Perimeter=%.2f\n", name, s.Area(), s.Perimeter())
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工ID: %d\n", e.Name, e.Age, e.EmployeeID)
}

// 生产者协程：生成1到10的整数并发送到通道
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("发送数字: %d\n", i)
		ch <- i // 将数字发送到通道
	}
	close(ch) // 发送完所有数字后关闭通道
}

// 消费者协程：从通道接收整数并打印
func consumer(ch <-chan int) {
	for num := range ch { // 从通道接收数字直到通道关闭
		fmt.Printf("接收到数字: %d\n", num)
	}
	fmt.Println("所有数字接收完毕")
}

// 生产者协程：向缓冲通道发送100个整数
func producer2(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("生产者发送: %d\n", i)

		// 模拟生产耗时
		if i%10 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	close(ch) // 发送完所有数据后关闭通道
	fmt.Println("生产者完成数据发送")
}

// 消费者协程：从缓冲通道接收整数并打印
func consumer2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 0
	for num := range ch {
		count++
		fmt.Printf("消费者接收: %d\n", num)

		// 模拟消费耗时
		if count%15 == 0 {
			time.Sleep(150 * time.Millisecond)
		}
	}

	fmt.Printf("消费者总共处理了 %d 个数据\n", count)
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (sc *SafeCounter) Increment() {
	for i := 0; i < 1000; i++ {
		sc.mu.Lock()
		sc.count++
		sc.mu.Unlock()
		fmt.Printf("当前的计数器值为 %d\n", sc.count)
	}

}

func (sc *SafeCounter) GetCount() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

var count int64

func atomicIncrement() {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count, 1)
	}
}

func main() {
	var x int = 10
	byPointer(&x)
	var y int = 10
	byValue(y)
	fmt.Println("值引,原变量值不会改变", y)
	fmt.Println("指针引用修改引用地址,修改后的值", x)
	var slice = []int{1, 2, 3, 4, 5}
	splitSlice(&slice)
	fmt.Println("切片元素乘以2", slice)
	// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。
	go getOddnumber()
	go getEvenNumber()
	time.Sleep(time.Millisecond)
	// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 考察点 ：协程原理、并发任务调度。
	ch := make(chan int, 5)
	go produce(ch)
	//多个消费者
	for i := 0; i < 3; i++ {
		go consumerTwo(ch, i)
	}
	time.Sleep(2 * time.Second)
	//获取不同形状对应的形状切片
	shapes := []Shape{Rectangle{width: 2, Height: 3}, Circle{Radius: 2}}
	names := []string{"矩形", "圆形"}
	for i, shape := range shapes {
		PrintShapeInfo(shape, names[i])
	}
	// 创建 Employee 实例
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: 1001,
	}
	employee.PrintInfo()
	//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
	// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

	// 创建一个无缓冲的整数通道
	ch1 := make(chan int)

	// 启动生产者协程
	go producer(ch1)

	// 启动消费者协程
	go consumer(ch1)

	// 等待足够的时间让协程执行完毕
	time.Sleep(1 * time.Second)

	// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 创建一个缓冲大小为10的通道
	bufferedChan := make(chan int, 10)

	// 使用WaitGroup等待所有协程完成
	var wg sync.WaitGroup

	// 启动生产者协程
	wg.Add(1)
	go producer2(bufferedChan, &wg)

	// 启动消费者协程
	wg.Add(1)
	go consumer2(bufferedChan, &wg)

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有数据处理完成")

	// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	counter := NewSafeCounter()
	var wg1 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			counter.Increment()
		}(i)

	}
	wg1.Wait()
	fmt.Printf("最终计数: %d (期望: 10000)\n\n", counter.GetCount())
	// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var wg2 sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomicIncrement()
		}()
	}

	// 等待所有协程完成
	wg2.Wait()
	

	fmt.Printf("最终计数器值: %d\n", atomic.LoadInt64(&count))

}
