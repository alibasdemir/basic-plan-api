package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func WeeklyPlan(c echo.Context) error {
	startOfWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var weeklyPlans []plan.Plan
	if err := database.DB.Where("day BETWEEN ? AND ?", startOfWeek, endOfWeek).Find(&weeklyPlans).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "No plans found for this week")
		}
		return c.JSON(http.StatusInternalServerError, "Error fetching weekly plans")
	}

	return c.JSON(http.StatusOK, weeklyPlans)
}
