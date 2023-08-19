package controller

import (
	"net/http"

	"github.com/PunGrumpy/golang-mvc-simple/model"
	"github.com/PunGrumpy/golang-mvc-simple/service"
	"github.com/gin-gonic/gin"
)

type SoldierController interface {
	AddSoldier(c *gin.Context)
	UpdateSoldier(c *gin.Context)
	GetSoldierByID(c *gin.Context)
}

type soldierController struct {
	dutyService service.DutyService
}

func NewSoldierController(dutyService service.DutyService) SoldierController {
	return &soldierController{
		dutyService: dutyService,
	}
}

func (s *soldierController) AddSoldier(c *gin.Context) {
	var newSoldier model.Soldier
	if err := c.ShouldBindJSON(&newSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Soldier Input"})
		return
	}

	if err := s.dutyService.AddSoldier(&newSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"soldier": newSoldier})
}

func (s *soldierController) UpdateSoldier(c *gin.Context) {
	id := c.Param("id")
	var updateSoldier model.Soldier
	if err := c.ShouldBindJSON(&updateSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Soldier Input"})
		return
	}

	if err := s.dutyService.UpdateSoldier(id, &updateSoldier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Soldier updated successfully"})
}

func (s *soldierController) GetSoldierByID(c *gin.Context) {
	id := c.Param("id")
	soldierInfo, err := s.dutyService.GetSoldierByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"soldier": soldierInfo})
}
