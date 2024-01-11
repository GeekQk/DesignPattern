package main

import (
	"database/sql"
	"fmt"
	"log"
)

/*
适配器模式来解决不同数据库兼容性的问题。
适配器模式是一种设计模式，它允许将一个类的接口转换成客户端所期望的另一种接口，
从而使得原本由于接口不兼容而无法协同工作的类能够一起工作。
*/

// 定义数据库操作的通用接口
type Database interface {
	Connect() error
	Query(query string) (*sql.Rows, error)
	Close() error
}

// MySQL数据库适配器
type MySQLAdapter struct {
	conn *sql.DB
}

func (a *MySQLAdapter) Connect() error {
	// 连接MySQL数据库的代码
	conn, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		return err
	}
	a.conn = conn
	return nil
}

func (a *MySQLAdapter) Query(query string) (*sql.Rows, error) {
	return a.conn.Query(query)
}

func (a *MySQLAdapter) Close() error {
	return a.conn.Close()
}

// PostgreSQL数据库适配器
type PostgreSQLAdapter struct {
	conn *sql.DB
}

func (a *PostgreSQLAdapter) Connect() error {
	// 连接PostgreSQL数据库的代码
	conn, err := sql.Open("postgres", "user=username password=password host=localhost port=5432 dbname=database")
	if err != nil {
		return err
	}
	a.conn = conn
	return nil
}

func (a *PostgreSQLAdapter) Query(query string) (*sql.Rows, error) {
	return a.conn.Query(query)
}

func (a *PostgreSQLAdapter) Close() error {
	return a.conn.Close()
}

func main() {
	// 使用MySQL数据库适配器进行查询操作
	mysqlAdapter := &MySQLAdapter{}
	err := mysqlAdapter.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlAdapter.Close()
	rows, err := mysqlAdapter.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("MySQL:", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 使用PostgreSQL数据库适配器进行查询操作
	postgresAdapter := &PostgreSQLAdapter{}
	err = postgresAdapter.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresAdapter.Close()
	rows, err = postgresAdapter.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PostgreSQL:", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
