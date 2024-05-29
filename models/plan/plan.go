package plan

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	StudentID uint      `json:"student_id"`
	Title     string    `json:"title"`
	Day       time.Time `json:"day"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Status    string    `json:"status"`
}
