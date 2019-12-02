package main

import (
	interfacett "../interfacett"
	basestruct "../basestruct"
	json "../json"
	"fmt"
	methodfunctiontt "../methodfunctiontt"
)
func main()  {


	//json解析
	json.PrarseJson()

	//结构体与method,function
	methodfunctiontt.UseStruct()
	methodfunctiontt.UseFunction()

	//array,slice,map
	basestruct.UseArray()
	basestruct.UseSlice()
	basestruct.UseMap()

	//初始化interface
	var skill interfacett.Skills
	//初始化 struct=
	var stu1 interfacett.Student
	stu1.Name = "wd"
	stu1.Age = 22
	//interface的多态
	var teacher interfacett.Teacher
	teacher.Name="ttt"
	skill = stu1
	skill.Running()  //调用接口
	skill=teacher
	skill.Running()
	//判断interface的类型
	i,ok := skill.(interfacett.Teacher)
	fmt.Printf("",i,ok)
}
