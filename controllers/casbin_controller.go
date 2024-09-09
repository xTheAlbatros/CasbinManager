package controllers

import (
	"CasbinManager/models"
	"CasbinManager/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CasbinController struct {
	service *services.CasbinService
}

func NewCasbinController(service *services.CasbinService) *CasbinController {
	return &CasbinController{service: service}
}

func (ctrl *CasbinController) GetRules(c *gin.Context) {
	rules, err := ctrl.service.GetAllRules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}

func (ctrl *CasbinController) AddRule(c *gin.Context) {
	var rule models.CasbinRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.AddRule(rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "rule added"})
}

func (ctrl *CasbinController) DeleteRule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.service.DeleteRule(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "rule deleted"})
}

func (ctrl *CasbinController) UpdateRule(c *gin.Context) {
	var rule models.CasbinRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateRule(rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "rule updated"})
}
