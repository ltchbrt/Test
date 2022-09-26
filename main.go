package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dafalo/Time_Management/views"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

const (
	BindIP = "localhost"
	Port   = ":3333"
)

// creattable(table name, database name schema name)

func main() {
	fmt.Printf("Go to port System: %v%v/\n", BindIP, Port)
	CreateDB("Time_Management")
	CreateTable("users", "Time_Management")
	CreateTable_data("data", "Time_Management")
	Handlers()
	http.ListenAndServe(Port, nil)
}

//Handlers for the design and views

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/login/", views.LogInHander)
	http.HandleFunc("/signup/", views.SignUpHandler)
	http.HandleFunc("/", views.MainHandler)
}

//Create DB(Activity)

func CreateDB(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

//Create Table(Data)

func CreateTable(table string, name string) *sql.DB {
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + table + "(user_id INT(11) PRIMARY KEY AUTO_INCREMENT, username varchar(32) UNIQUE, password varchar (1000), name varchar (1000), position varchar (32) );")
	if err != nil {
		panic(err)
	}
	db.Close()
	return db
}

func CreateTable_data(table string, name string) *sql.DB {
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + table + "(data_id INT(11) PRIMARY KEY AUTO_INCREMENT , username varchar(1000), name varchar(1000), position varchar(100), date varchar(100), time_in varchar(100), time_out varchar(100), duration varchar(100), total varchar(100) );")
	if err != nil {
		panic(err)
	}
	db.Close()
	return db
}
