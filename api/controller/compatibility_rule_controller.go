package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompatibilityRuleController struct {
	CompatibilityRuleRepository domain.CompatibilityRuleRepository
}

func (te *CompatibilityRuleController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var CompatibilityRule domain.CompatibilityRule

	err := c.ShouldBind(&CompatibilityRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	CompatibilityRule.CompatibilityRuleID = uuid.New()

	err = te.CompatibilityRuleRepository.Create(c, CompatibilityRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "CompatibilityRule created successfully",
	})
}

func (te *CompatibilityRuleController) Fetch(c *gin.Context) {
	CompatibilityRules, err := te.CompatibilityRuleRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CompatibilityRules)
}

func (te *CompatibilityRuleController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	CompatibilityRules, err := te.CompatibilityRuleRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CompatibilityRules)
}

func (te *CompatibilityRuleController) Update(c *gin.Context) {
	updatedCompatibilityRule := &domain.CompatibilityRule{}

	err := c.ShouldBind(updatedCompatibilityRule)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedCompatibilityRule.CompatibilityRuleID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "CompatibilityRuleID is requiered to update"})
		return
	}

	err = te.CompatibilityRuleRepository.Update(c, *updatedCompatibilityRule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "CompatibilityRule updated succesfully"})
}

func (te *CompatibilityRuleController) Delete(c *gin.Context) {
	CompatibilityRuleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.CompatibilityRuleRepository.Delete(c, CompatibilityRuleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "CompatibilityRule delete succesfully"})
}
