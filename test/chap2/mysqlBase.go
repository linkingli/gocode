package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
/**
	sql 原始连接
	参考文档:https://blog.csdn.net/pangudashu/article/details/54291558
 */

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "passwd"
	NETWORK = "tcp"
	SERVER = "ip"
	PORT = 3306
	DATABASE = "keystone"
)

//user表结构体定义
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status int   `json:"status"`      // 0 正常状态， 1删除
	Createtime int64 `json:"createtime"`
}


func main() {
	//root:cloudos@tcp(10.125.38.27:3306)/keystone
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, e := sql.Open("mysql", conn)
	if e!=nil {
		fmt.Print("connection error",e)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超时的连接就close
	DB.SetMaxOpenConns(100)                //设置最大连接数
	CreateTable(DB)
	InsertData(DB)
	QueryOne(DB)
/*	QueryMulti(DB)
	UpdateData(DB)
	DeleteData(DB)*/

}

func CreateTable(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS usersss(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	username VARCHAR(64),
	password VARCHAR(64),
	status INT(4),
	createtime INT(10)
	); `

	_, err := DB.Exec(sql);
	if err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}

//插入数据
func InsertData(DB *sql.DB) {
	result,err := DB.Exec("insert INTO usersss(username,password) values(?,?)","test","123456")
	if err != nil{
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	lastInsertID,err := result.LastInsertId()    //获取插入数据的自增ID
	if err != nil {
		fmt.Printf("Get insert id failed,err:%v", err)
		return
	}
	fmt.Println("Insert data id:", lastInsertID)

	rowsaffected,err := result.RowsAffected()  //通过RowsAffected获取受影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//查询单行
func QueryOne(DB *sql.DB) {
	user := new(User)   //用new()函数初始化一个结构体对象
	row := DB.QueryRow("select id,username,password from usersss where id=?", 1)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	//?
	if err := row.Scan(&user.Id,&user.Username,&user.Password); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("Single row data:", *user)
}

//查询多行
func QueryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select id,username,password from usersss where id = ?", 2)

	defer func() {
		if rows != nil {
			rows.Close()   //关闭掉未scan的sql连接
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)  //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Println("scan successd:", *user)
	}
}

//更新数据
func UpdateData(DB *sql.DB){
	result,err := DB.Exec("UPDATE usersss set password=? where id=?","111111",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n", err)
		return
	}
	fmt.Println("update data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}

//删除数据
func DeleteData(DB *sql.DB){
	result,err := DB.Exec("delete from usersss where id=?",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n",err)
		return
	}
	fmt.Println("delete data successd:", result)

	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("Affected rows:", rowsaffected)
}
