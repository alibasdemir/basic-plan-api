package database

import (
	"fmt"
	"log"
	"student-plan/models/plan"
	"student-plan/models/student"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:123456@tcp(localhost:3306)/studentPlan?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Successfully connected to the database!")
	DB.AutoMigrate(&student.Student{}, &plan.Plan{})
}
