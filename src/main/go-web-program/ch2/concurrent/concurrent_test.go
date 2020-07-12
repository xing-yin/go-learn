package concurrent

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/// todo ========================== goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		runtime.Gosched()
		// 默认情况下，在Go 1.5将标识并发系统线程个数的
		// runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。
		fmt.Println(s)
	}
}

func Test1(t *testing.T) {
	go say("go") // 开启新的 Goroutines 执行
	say("hello")
}

/// todo ==========================channels
// 定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

func Test2(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan interface{})
	fmt.Println(c1, c2, c3)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func Test3(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

// 默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，
// 而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
// 其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。

/// todo ========================== Buffered Channels
// ch := make(chan type, value)
// - 当 value = 0 时，channel 是无缓冲阻塞读写的，
// - 当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

func Test4(t *testing.T) {
	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/// todo ========================== Range和Close
// 上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，
// 所以也可以通过range，像操作slice或者map一样操作缓存类型的channel

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Test5(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。
// 上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。
// 关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。
// 如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

//⚠️应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
//⚠️ channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的

/// todo ==========================Select
// 如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，
//通过select可以监听channel上的数据流动。

// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
// 当多个channel都准备好的时候，select是随机的选择一个执行的。

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Test6(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

// 在select里面还有default语法，select其实就是类似switch的功能，
// default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
func t(c, quit chan int) {
	select {
	case i := <-c:
		// use i
		fmt.Println(i)
	default:
		// 当 c 阻塞的时候执行这里
	}
}

/// todo ==========================超时
// 有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？
// 我们可以利用select来设置超时，通过如下的方式实现：
func Test7(t *testing.T) {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

/// todo ==========================runtime goroutine
// runtime包中有几个处理goroutine的函数：
// - Goexit
//退出当前执行的goroutine，但是defer函数还会继续调用
// - Gosched
//让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
// - NumCPU
//返回 CPU 核数量
// - NumGoroutine
//返回正在执行和排队的任务总数
// - GOMAXPROCS
//用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
