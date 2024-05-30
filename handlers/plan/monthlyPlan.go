package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"time"

	"github.com/labstack/echo/v4"
)

func MonthlyPlan(c echo.Context) error {
	currentMonth := time.Now().Month()
	startOfMonth := time.Date(time.Now().Year(), currentMonth, 1, 0, 0, 0, 0, time.Now().Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	var monthlyPlans []plan.Plan
	result := database.DB.Where("day BETWEEN ? AND ?", startOfMonth, endOfMonth).Find(&monthlyPlans)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching monthly plans")
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "No plans found for this month")
	}

	return c.JSON(http.StatusOK, monthlyPlans)
}
