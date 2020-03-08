package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_selectUser = "select name,address from user where id = ?"
	_updateUser = "update user set name = ? where name = ?"
	_insertUser = "insert user (name,age) values (?,?)"
	_deleteUser = "delete user where id = ?"
)

func main() {
	var (
		err error
	)
	// 得到一个db对象
	// Open方法第二个参数:  用户名:密码@协议(ip:端口)/数据库?编码
	db, err := sql.Open("mysql", "root:12345@tcp(192.168.211.128:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close() //关闭数据库

	// 这里我们拿到了db这个对象，就可以对day0731这个mysql数据库进行操作了
	// 查询操作
	// Query方法，返回值具体类型自己看
	var name, address string
	// 接收返回结果， 这里的name，age必须和sql语句顺序一致，也不能多传，也不能少传
	// QueryRow方法，查询的效果和上面一致
	err = db.QueryRow(_selectUser, 1).Scan(&name, &address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	fmt.Println(address)

	// 插入、更新、删除操作
	// Exec方法，执行一段sql语句
	result, err := db.Exec(_updateUser, "Goat", "LadyGaGa")
	if err != nil {
		fmt.Println(err)
		return
	}
	// result这个对象，他有两个方法，RowsAffected(),LastInsertId(),
	// 分别返回影响的行数，还有就是最后插入的id
	// 我们这里执行的是更新语句，就看这个方法就好了RowsAffected
	rowCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if rowCount == 0 {
		fmt.Println("更新操作失败")
		return
	}
	fmt.Println("更新操作成功！")

}
