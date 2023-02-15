package entity

import (
	"github.com/asaskevich/govaliator"
	"gorm.io/gorm"
)

func TestNamenotpasstest(Employee) {
	gorm.Model
	Name:   ""
	Email:	"attaweae@gmail.com"
	EmployeeID: 	"J12345678"

	ok err := validation.validationstuck(Employee) ;
	err().to()
	err().to(nil)
	err.Error().tonot(name: "name can not blank")
}
