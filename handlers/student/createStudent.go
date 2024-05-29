package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
)

func CreateStudent(c echo.Context) error {
	student := new(student.Student)
	if err := c.Bind(student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	database.DB.Create(&student)
	return c.JSON(http.StatusCreated, student)
}
