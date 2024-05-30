package main

import (
	database "student-plan/db"
	planHandlers "student-plan/handlers/plan"
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

	e.GET("/plans", planHandlers.GetPlan)
	e.GET("/plans/:id", planHandlers.GetPlanById)
	e.POST("/plans", planHandlers.CreatePlan)
	e.PUT("/plans/:id", planHandlers.UpdatePlan)
	e.Start(":8080")
}
