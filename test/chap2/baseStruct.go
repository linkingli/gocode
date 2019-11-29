package main

import "fmt"

func main() {
	//array,slice,map的方法传递使用
	// 将切片或者map传递给函数成本很小，并且不会复制底层的数据结,只是传递的是
	useArray()
	useSlice()
	useMap()


}




func useArray() {
	//数组声明的三种方式
	var array[10] string
	array[0]="sss"

	strings := [5]string{"1"}
	strings[0]="pp"

	arrayList1 := [...]string{""}
	arrayList1[0]="kk"

	arrayList2 := [5] string{1: "ss", 2: "ddf"}


	arrayList2[3]="wsa"

	//使用指针的数组
	arrayList3 := [5]*int{0: new(int), 1: new(int)}
	*arrayList3[0] = 10
	*arrayList3[1] = 20
	//数组赋值，只有数组的类型和数组的长度一样的情况下，2个数组才认为是同一类型的数组，才能互相赋值
	var arrayCopy[10]string
	arrayCopy=array

	arrayCopy[1]="es"

	//复制指针数组，只会复制指正，不会复制指针对应的值，
	//换一句话说，复制数组指针，就是2个数组指向了同一个数组
	var arrayIndex1 [3]*string
	arrayIndex2 := [3]*string{new(string), new(string), new(string)}
	*arrayIndex2[0] = "Red"
	*arrayIndex2[1] = "Blue"
	*arrayIndex2[2] = "Green"
	arrayIndex1 = arrayIndex2
	//使用指针去修改数组的值，会修改到所有指向这个数组的指针的引用
	*arrayIndex1[0]="YELLOW"


	//多维指针的值
	var arrayTensor [4][2]int
	arrayTensor[1][1]=10


	//方法传递时候，使用指针比使用数组更加高效
	var arrayLarge[1100000] int
	foo1(arrayLarge)
	foo2(&arrayLarge)
}

//每次调用需要传递需要创建一个8M的数组，把传入的参数拷贝到方法中
func foo1(ints [1100000]int) {

}
//每次调用需要传递只需要传递8bit的内存指针
func foo2(ints *[1100000]int) {

}

func useSlice() {

	//创建一个切片，如果不指定capcityn，那么默认认为和length的长度一样
	slice1:= make([]int, 3, 5)
	slice1[2]=1

	//声明创建一个切片,和数组不同的是，数组必须加上长度的声明，切片的声明是不需要长度的声明
	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	slice2[1]="dfds"

	//使用索引引声明切片，这种方式比较巧妙
	slice3 := []string{99: ""}
	slice3[98]=""
	//声明一个nil的切片，空切片的长度为0,值为nil
/*	var slice4[] int
	slice5 := make([]int, 0)
	slice6 := []int{}*/


	//slice7和slice8共享一个底层数组,调用append的时候在slice8看来是由2,3->2,3,60，实际也修改底层4的值为60
	//如果底层数组的长度不够append增加数组使用，会创建一个新的数组，并且把原来的值copy进去
	slice7:=[]int{1,2,3,4,5,6,7,8}
	slice8 := slice7[1:3]
	slice8 = append(slice8, 60)

	//三索引声明一个切片,分别是index,end,capcity-end
	//长度: j – i 或3 - 2 = 1
	//容量: k – i 或4 - 2 = 2
	slice9 := slice7[2:3:4]
	slice9[0]=33


	// 数组分层:append参数会创建一个新的底层数组，供切片使用，可以使用这种方式使得切片分层(即不使用同一个底层数组)


	// 创建字符串切片
	// 其长度和容量都是5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是1 个元素
	slice := source[2:3:3]
	// 向slice 追加新字符串
	slice = append(slice, "Kiwi")

	//_ = append(slice, source)

	for index,value:=range slice{
		fmt.Printf("Index: %d Value: %d\n", index, &value,&slice[index])
	}

		//切片的函数传递
	/*	在 64 位架构的机器上，使用切片传递参数需要24 字节的内存：指针字段需要8 字节，长度和容量字段分别需要8 字节*/
	// 分配包含100 万个整型值的切片
	slicefunction := make([]int, 1e6)
	// slicefunction 传递到函数foo
	slicefunction = fooslice(slicefunction)


}

func fooslice(ints []int) []int {

	return  ints
}



func useMap() {
	//使用make和字面量创建map的两种方式
	//切片、函数以及包含切片的结构类型这些类型由于具有引用语义， 不能作为映射的键
	map1 := make(map[string]int)
	map1["green"]=12
	map2 := map[string]string{}
	map2["red"]="yamap"

	//获取map中的元素会返回value和exsits两个值
	value,exsits := map2["red"]
	if exsits {
		print(value)
	}

	//range遍历map
	for key,valu:=range map2{
		fmt.Printf("Key: %s Value: %s\n", key, valu)
	}

	//map删除不存在的元素也不回报错
	delete(map2,"pop")
	delete(map2,"red")

}




