package studentHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"student-plan/models/student"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetStudentPlan(c echo.Context) error {
	id := c.Param("id")
	var student student.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Student not found")
		}
	}

	var studentPlans []plan.Plan
	if err := database.DB.Where("student_id = ?", id).Find(&studentPlans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching student plans")
	}

	if len(studentPlans) == 0 {
		return c.JSON(http.StatusNotFound, "The student does not have any plans")
	}

	return c.JSON(http.StatusOK, studentPlans)
}
