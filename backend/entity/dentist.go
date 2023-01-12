package entity

import (
	"gorm.io/gorm"
)

type Dentist struct {
	gorm.Model
    
	Dentist_name string

	Treatment        []Treatment  `gorm:"foreignkey:DentistID"`
	
}
