package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
)

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	var student student.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Student not found")
	}
	if err := database.DB.Delete(&student).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting student")
	}
	return c.JSON(http.StatusOK, "Student has been deleted")
}
