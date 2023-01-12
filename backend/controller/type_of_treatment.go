package controller

import (
	"net/http"

	"github.com/B6309985/team09/entity"
	"github.com/gin-gonic/gin"
)

// POST /users
func Create_Type_of_treatment(c *gin.Context) {
	var type_of_treatments entity.Type_Of_Treatment
	
	if err := c.ShouldBindJSON(&type_of_treatments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&type_of_treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_of_treatments})

}

// GET /BehaviorType/:id

func Get_Type_of_treatment(c *gin.Context) {
	var type_of_treatments entity.Academy
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM type_of_treatments WHERE id = ?", id).Scan(&type_of_treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_of_treatments})

}

// GET /BehaviorTypes

func List_Type_of_treatment(c *gin.Context) {
	var type_of_treatments []entity.Type_of_treatment
	if err := entity.DB().Raw("SELECT * FROM type_of_treatments").Scan(&type_of_treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": type_of_treatments})

}

// DELETE /BehaviorTypes/:id

func Delete_Type_of_treatment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM type_of_treatments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of treatment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /BehaviorTypes

func Update_Type_of_treatment(c *gin.Context) {
	var type_of_treatments entity.Type_of_treatment
	if err := c.ShouldBindJSON(&type_of_treatments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", type_of_treatments.ID).First(&type_of_treatments); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type of treatment not found"})
		return
	}

	if err := entity.DB().Save(&type_of_treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": type_of_treatments})

}