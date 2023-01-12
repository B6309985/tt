package entity

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
    
	Patient_name string

	Treatment        []Treatment  `gorm:"foreignkey:PatientID"`
	
}
