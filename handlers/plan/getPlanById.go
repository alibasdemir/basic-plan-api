package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"

	"github.com/labstack/echo/v4"
)

func GetPlanById(c echo.Context) error {
	id := c.Param("id")
	var plan plan.Plan
	if err := database.DB.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Plan not found")
	}
	return c.JSON(http.StatusOK, plan)
}
