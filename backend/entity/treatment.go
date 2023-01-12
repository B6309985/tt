package entity

import (
	"time"

	"gorm.io/gorm"
)


type Type_Of_Treatment struct {
	gorm.Model
	Type_of_treatment_name       string
	Price 						 int
	Treatment        []Treatment  `gorm:"foreignkey:Type_Of_TreatmentID"`
	// Student       []Student `gorm:"foreignkey:AcademyID"`
} 

type Treatment struct {
	gorm.Model
	
	PatientID *uint  
	Patient   Patient

	DentistID *uint  
	Dentist   Dentist

	Type_Of_TreatmentID *uint
	Type_Of_Treatment	Type_Of_Treatment

	Number_of_cavities int

	Number_of_swollen_gums int 

	Other_teeth_problems string
		
	Treatment string

	Treatment_Time time.Time

	Treatment_code string

}
