package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"time"

	"github.com/labstack/echo/v4"
)

func WeeklyPlan(c echo.Context) error {
	startOfWeek := time.Now().AddDate(0, 0, -int(time.Now().Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var weeklyPlans []plan.Plan
	result := database.DB.Where("day BETWEEN ? AND ?", startOfWeek, endOfWeek).Find(&weeklyPlans)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching weekly plans")
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "No plans found for this week")
	}

	return c.JSON(http.StatusOK, weeklyPlans)
}
