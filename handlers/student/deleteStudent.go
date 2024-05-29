package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
)

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	if err := database.DB.Delete(&student.Student{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Student has been deleted"})
}
