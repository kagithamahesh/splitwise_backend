package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	// choose the appropriate gorm dialector; here we support mysql
	if Dbdriver == "mysql" {
		DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	} else {
		log.Fatalf("unsupported DB_DRIVER: %s", Dbdriver)
	}

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	// err = DB.AutoMigrate(
	// 	&Userlist{},
	// 	// &Group{},
	// 	// &GroupMember{},
	// 	// &Expense{},
	// 	// &ExpenseSplit{},
	// 	// &Payment{},
	// 	// &Balance{}, // Include the recommended balance table
	// )
	// if err != nil {
	// 	log.Fatal("AutoMigration failed:", err)
	// }
}
func (Userlist) TableName() string {
	return "userslist"
}

func (Groupslist) TableName() string {
	return "groupslist"
}

func (GroupMembers) TableName() string {
	return "groupmembers"
}

func (Expenses) TableName() string {
	return "expenses"
}

func (ExpensesList) TableName() string {
	return "expensesplits"
}
