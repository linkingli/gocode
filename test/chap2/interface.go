package main

import "fmt"

func main()  {
	//初始化interface
	var skill Skills
	//初始化 struct=
	var stu1 Student
	stu1.Name = "wd"
	stu1.Age = 22
	//interface的多态
	var teacher Teacher
	teacher.Name="ttt"
	skill = stu1
	skill.Running()  //调用接口
	skill=teacher
	skill.Running()
	//判断interface的类型
	i,ok := skill.(Teacher)
	fmt.Printf("",i,ok)
}

//定义接口
type Skills interface {
	Running()
	Getname() string

}
type Test interface {
	sleeping()
	Skills   //继承Skills
}
type Student struct {
	Name string
	Age int
}
type Teacher struct {
	Name string
	Salary int
}

func (teacher Teacher) Running() {
	fmt.Printf("%s running",teacher.Name)
}

func (teacher Teacher) Getname() string {
	fmt.Println(teacher.Name )
	return teacher.Name
}



// 实现接口
func (student Student) Getname() string{   //实现Getname方法
	fmt.Println(student.Name )
	return student.Name
}

func (student Student) Running()  {   // 实现 Running方法
	fmt.Printf("%s running",student.Name)
}
