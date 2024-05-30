package planHandlers

import (
	"net/http"
	database "student-plan/db"
	"student-plan/models/plan"

	"github.com/labstack/echo/v4"
)

func GetPlan(c echo.Context) error {
	var plans []plan.Plan
	if err := database.DB.Find(&plans).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, plans)
}
