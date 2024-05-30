package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"time"

	"github.com/labstack/echo/v4"
)

func CreatePlan(c echo.Context) error {
	newPlan := new(plan.Plan)
	if err := c.Bind(newPlan); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if _, err := time.Parse("2006-01-02", newPlan.Day); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
	}

	var conflictingPlans []plan.Plan
	database.DB.Where("day = ? AND start_time < ? AND end_time > ?", newPlan.Day, newPlan.EndTime, newPlan.StartTime).Find(&conflictingPlans)
	if len(conflictingPlans) > 0 {
		return c.JSON(http.StatusConflict, "There is already a plan in the specified time range")
	}

	database.DB.Create(newPlan)
	return c.JSON(http.StatusCreated, newPlan)
}
