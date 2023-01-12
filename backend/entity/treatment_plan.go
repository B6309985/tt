package entity

import (
	"time"

	"gorm.io/gorm"
)

type Treatment_plan struct {
	gorm.Model

	PatientID *uint
	Patient   Patient

	DentistID *uint
	Dentist   Dentist

	Order_of_treatment int

	Type_Of_Treatment_ID *uint
	Type_Of_Treatment    Type_Of_Treatment

	Treatment string

	Treatment_Time time.Time
}
