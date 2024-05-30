package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"
	"time"

	"github.com/labstack/echo/v4"
)

func UpdatePlan(c echo.Context) error {
	id := c.Param("id")
	var existingPlan plan.Plan
	if err := database.DB.First(&existingPlan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Plan not found")
	}

	updatedPlan := new(plan.Plan)
	if err := c.Bind(updatedPlan); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if _, err := time.Parse("2006-01-02", updatedPlan.Day); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid date format. Use YYYY-MM-DD")
	}

	var conflictingPlans []plan.Plan
	database.DB.Where("id <> ? AND day = ? AND start_time < ? AND end_time > ?", id, updatedPlan.Day, updatedPlan.EndTime, updatedPlan.StartTime).Find(&conflictingPlans)
	if len(conflictingPlans) > 0 {
		return c.JSON(http.StatusConflict, "There is already a plan in the specified time range")
	}

	// Güncellemeleri uygula
	existingPlan.Title = updatedPlan.Title
	existingPlan.Day = updatedPlan.Day
	existingPlan.StartTime = updatedPlan.StartTime
	existingPlan.EndTime = updatedPlan.EndTime
	existingPlan.Status = updatedPlan.Status

	database.DB.Save(&existingPlan)
	return c.JSON(http.StatusOK, existingPlan)
}
