package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
)

func GetStudent(c echo.Context) error {
	var students []student.Student
	if err := database.DB.Find(&students).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, students)
}
