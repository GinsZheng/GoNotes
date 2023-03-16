package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func mysqlPrint() {
	err := initDB() // 调用初始化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("数据库调用：")
	queryRowDemo()      // 查单行
	queryMultiRowDemo() // 查多行
	//insertRowDemo() // 插入
	//updateRowDemo() // 更新
	deleteRowDemo() // 删除
}

// 定义一个全局变量db
var db *sql.DB

// 初始化数据库
func initDB() (err error) {
	// DSN 格式：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:Aqjklwq5060/@tcp(127.0.0.1:3306)/test"
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil // 返回nil，说明数据库连接成功
}

type USER struct {
	id       int
	username string
	phone    string
}

// 查询单条数据 QueryRow
func queryRowDemo() {
	sqlStr := "select id, username, phone from user where id=?"
	var u USER
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 9).Scan(&u.id, &u.username, &u.phone) // QueryRow中的args = 4，4表示把sqlStr中⌜id=?⌟里问号赋值为4
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s phone:%s\n", u.id, u.username, u.phone)
}

// 查询多条数据 Query
func queryMultiRowDemo() {
	sqlStr := "select * from user where id < ?"
	rows, err := db.Query(sqlStr, 7) // 用sql语句请求，把结果赋值给⌜rows⌟，args表示query中的占位参数(即⌜?⌟)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u USER
		err := rows.Scan(&u.id, &u.username, &u.phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s phone:%s\n", u.id, u.username, u.phone)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(username, phone) values (?,?)"
	ret, err := db.Exec(sqlStr, "Gink", "13345677654") // 执行插入数据操作。插入、更新和删除操作都使用Exec方法
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 获取新插入数据的id，仅用于打印结果
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set username=? where id = ?"
	ret, err := db.Exec(sqlStr, "Gins", 5) // 执行更新数据操作
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数，仅用于打印结果
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// ❓这些操作可以封装
