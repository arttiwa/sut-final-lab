package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `validation: "Name~Name can not blank"`
	Email      string `gorm : "unique" validation: "Email~Name can not blank"`
	EmployeeID string `gorm : "unique" "match{({J},{s},{m}\\8 d)}"`
}
