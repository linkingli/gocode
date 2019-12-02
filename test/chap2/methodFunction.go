package main

import (
	"fmt"
	"log"
)

//go语言是静态强类型语言,所使用的值必须指定类型
//像int 64，这种第一部分，需要分配多少内存给这个值（即值的规模；在64位机器占8字节, 第二部分，这段内存表示什么,存储一个整数
func main() {

	//useStruct()
	useFunction()
}


func useStruct() {
	type user struct {
		name string
		email string
		ext string
		privileged bool
	}

	var tom user
	tom.email="tom@tom.com"

	john := user{"john", "john@john.com", "john information", false}
	john.email="john@john.cn"

	type admin struct {
		person user
		level string
	}

	boss := admin{
		user{
			"boss",
			"boss@boss.com",
			"boss information",
			false,
		},
		"1",
	}
	fmt.Printf(boss.person.email)

	type Duration int
	var inttest Duration
	inttest=25
	log.Println(inttest)

}


func useFunction() {


	// user 类型的值可以用来调用
	 // 使用值接收者声明的方法
	 bill := user{"Bill", "bill@email.com"}
	 bill.notify()
	 notifys(bill)



	 // 指向user 类型值的指针也可以用来调用
	 // 使用值接收者声明的方法
	 lisa := &user{"Lisa", "lisa@email.com"}
	 lisa.notify()

	 // user 类型的值可以用来调用
	 // 使用指针接收者声明的方法
	 bill.changeEmail("bill@newdomain.com")
	 bill.notify()
	 // 指向user 类型值的指针可以用来调用
	 // 使用指针接收者声明的方法
	 lisa.changeEmail("lisa@newdomain.com")

	 lisa.notify()
}

func notifys(u user) {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}



// notify 使用值接收者实现了一个方法
func (u user) notify()  {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

type user struct {
	name string
	email string
}



// changeEmail 使用指针接收者实现了一个方法
 func (u *user) changeEmail(email string) {
	 u.email = email
 }
// methodFor user:方法属于user类型所有
// params user:传入的对象是user类型
// user: 返回的参数类型是user类型

//对象传递: struct，数值类型、字符串类型 和布尔类:类型传入function是传入的副本，不会修改原数据，
//    slice,aray,map传入的是指针，会影响到原数据

//指针传递，无论是struct还是slice,aray,map都会修改数据

//方法对应的是struct,函数对应的是包，所以方法的引用时是lisa.notify()，而函数的引用是notifys(lisa)
func (methodFor user) describeFcuntion(params user) user {
	methodFor.email = params.email
	return  methodFor
}
