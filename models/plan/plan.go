package plan

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	StudentID uint   `json:"student_id"`
	Title     string `json:"title"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
}
