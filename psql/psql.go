package psql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DBB *gorm.DB = nil
var DBC *gorm.DB = nil
var err error

// ---Connect to Database---
func ConnectDBB(){
	var (
		host string 
		user string
		password string
		dbname string
		port string
	)
	fmt.Print("Enter the database host: ")
	fmt.Scanln(&host)
	fmt.Print("Enter the database user: ")
	fmt.Scanln(&user)
	fmt.Print("Enter the database password: ")
	fmt.Scanln(&password)
	fmt.Print("Enter the database dbname: ")
	fmt.Scanln(&dbname)
	fmt.Print("Enter the database port: ")
	fmt.Scanln(&port)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host ,user ,password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix: ".",
			// Schema
		},
	})
	if err != nil{
		log.Fatal("connect error")
	}
	fmt.Println("------DB Connected------")
	DBB = db
}

// ---Connect to database---
func ConnectDBC(){ 
	var (
		host string 
		user string
		password string
		dbname string
		port string
	)
	fmt.Print("Enter the database host: ")
	fmt.Scanln(&host)
	fmt.Print("Enter the database user: ")
	fmt.Scanln(&user)
	fmt.Print("Enter the database password: ")
	fmt.Scanln(&password)
	fmt.Print("Enter the database dbname: ")
	fmt.Scanln(&dbname)
	fmt.Print("Enter the database port: ")
	fmt.Scanln(&port)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host ,user ,password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix: ".",
			// Schema
		},
	})
	if err != nil{
		log.Fatal("connect error")
	}
	fmt.Println(`------ DB Connected------
********************************`)
	DBC = db
}