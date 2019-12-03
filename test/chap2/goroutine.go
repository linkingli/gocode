package goroutinett

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//hello world
	go Printhellorounting()
	fmt.Println("start go Goroutine")

	//chanel只能使用一次，即读一次写一次，再写或者读都会阻塞,这里ChanelTest3就会阻塞
	ints := make(chan int)
	go ChanelTest1(ints)
	go ChanelTest2(ints)
	//go ChanelTest3(ints)
	fmt.Println("end chanel")
	
	//多个go routing使用chanel通信
	number:=5
	sqrch :=make(chan  int)
	cubech  := make(chan int)
	go MakeSqrch(number,sqrch)
	go MakeCubech(number,cubech)
	sumSqrch:=<-sqrch
	sumCubech:=<-cubech
	fmt.Println(sumCubech+sumSqrch)

	time.Sleep(4 * time.Second)
	//缓冲信道
	go UseBUfferChanel()

	//WaitGroup
	go Fatherroute()
	
	time.Sleep(400 * time.Second)
}

func Printhellorounting()   {
	fmt.Printf("hello ,go rounting")
}

func ChanelTest1(a chan int)  {
	fmt.Println("send to chanel",a)
	a<-2

}

func ChanelTest2(a chan int)  {
	b:=<-a
	fmt.Println("get from chanel",b)
}
func ChanelTest3(a chan int)  {
	fmt.Println("send to chanel3333333",a)
	a<-3

}

//使用多个协调chanel 工作
func  MakeSqrch(a int,b chan int){
	sum:=0
	if a != 0 {
		sum=a*a;

	}
	b<-sum
}
func  MakeCubech(a int,b chan int){
	sum:=0
	if a != 0 {
		sum=a*a*a;

	}
	b<-sum
}
//死锁
func DieChanl() {
	ch := make(chan int)
	ch <- 5
}

//单向信道,双向可以调用单向参数的函数，单向不可以调用双向
func SignalChanel() {
	ints := make(chan int)
	sendData1(ints)
	sendData2(ints)
}
//单向信道只写
func sendData1(sendch chan<- int) {
	sendch <- 10
}

//单向只读
func sendData2( chRecvOnly <-chan int) {
	fmt.Println(chRecvOnly)
}
//信道遍历
func RangeChanel() {
	ints := make(chan int)
	go CloseChanel(ints)

	//信道取值
	v, ok := <-ints
	if ok == false {
		fmt.Println(v)
	}
	//信道遍历
	for v :=range ints{
		fmt.Println(v)
	}

}

//关闭信道
func CloseChanel( chanel chan int) {
	for i:=0;i<10 ;i++  {
		chanel<-i
	}
	close(chanel)
}


//缓冲信道

func UseBUfferChanel(){
	ints := make(chan int, 2)
	go bufferChanel(ints)
	time.Sleep(2 * time.Second)
	for v := range ints {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)

	}

}
func bufferChanel(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}


//WaitGroup控制go rounting,wg可以控制使得
func Fatherroute()  {
	println("start wait group")
	no:=3
	var wg sync.WaitGroup
	for i:=0;i<no ; i++ {
		wg.Add(1)
		go childrounte(i,&wg)
	}
	wg.Wait()
	println("end wait group")

}

func childrounte(i int, wg *sync.WaitGroup) {
	println("start child")
	wg.Done()
	println("end child")
}
