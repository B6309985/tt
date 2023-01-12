package controller

import (
	"net/http"

	"github.com/B6309985/team09/entity"
	"github.com/gin-gonic/gin"
)

// POST /users
func CreateTreatment(c *gin.Context) {
	var treatments entity.Treatment
	var dentists entity.Dentist
	var patients entity.Patient	
	var type_of_treatments entity.Type_Of_Treatment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร branch
	if err := c.ShouldBindJSON(&treatments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา admin ด้วย id
	if tx := entity.DB().Where("id = ?", treatments.DentistID).First(&dentists); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dentist not found"})
		return
	}

	// 10: ค้นหา academy ด้วย id
	if tx := entity.DB().Where("id = ?", treatments.PatientID).First(&patients); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 11: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", treatments.Type_Of_TreatmentID).First(&type_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	// 12: สร้าง branch
	treatment := entity.Treatment{
		Dentist:    dentists,               // โยงความสัมพันธ์กับ Entity admin
		Patient: patients,  // โยงความสัมพันธ์กับ Entity academies
		Type_Of_Treatment: type_of_treatments,     // โยงความสัมพันธ์กับ Entity course
		Number_of_cavities: treatments.Number_of_cavities,
		Number_of_swollen_gums: treatments.Number_of_swollen_gums,
		Other_teeth_problems: treatments.Other_teeth_problems,
		Treatment:  treatments.Treatment, 
		Treatment_code: treatments.Treatment_code,            
              	
	}

	// 13: บันทึก
	if err := entity.DB().Create(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(http.StatusCreated, gin.H{"data": treatment})
}

// GET /user/:id
func GetTreatment(c *gin.Context) {
	var treatments entity.Treatment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM treatments WHERE id = ?", id).Scan(&treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatments})
}

func ListTreatment(c *gin.Context) {
	var treatments []entity.Treatment
	if err := entity.DB().Raw("SELECT * FROM treatments").Scan(&treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatments})
}

////////////////////////////////////////////////////////////////////////////maybe change Listshow/////////////////////////////////////////
// GET /users
func ListTreatmentShow(c *gin.Context) {
	// var bp []entity.Behavior_Point
	result := []map[string]interface{}{}
	entity.DB().Table("branches").
		Select("branches.id, branches.brname, branches.contact, admins.aname, academies.acaname, rooms.rname").
		Joins("left join admins on admins.id = branches.admin_id").
		Joins("left join academies on academies.id = branches.academy_id").
		Joins("left join rooms on rooms.id = branches.room_id").
		Find(&result)

	c.JSON(http.StatusOK, gin.H{"data": result})

}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// DELETE /users/:id
func DeleteTreatment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatment WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateTreatment(c *gin.Context) {
	var treatments entity.Treatment
	if err := c.ShouldBindJSON(&treatments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", treatments.ID).First(&treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment not found"})
		return
	}
	if err := entity.DB().Save(&treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatments})
}
 
