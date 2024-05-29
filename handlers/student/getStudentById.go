package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
)

func GetStudentById(c echo.Context) error {
	id := c.Param("id")

	var student student.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}
	return c.JSON(http.StatusOK, student)
}
