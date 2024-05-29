package main

import (
	database "student-plan/db"
	studentHandlers "student-plan/handlers/student"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Connect()

	e.GET("/students", studentHandlers.GetStudent)
	e.GET("/students/:id", studentHandlers.GetStudentById)
	e.POST("/students", studentHandlers.CreateStudent)
	e.PUT("/students/:id", studentHandlers.UpdateStudent)
	e.DELETE("students/:id", studentHandlers.DeleteStudent)
	e.Start(":8080")
}
