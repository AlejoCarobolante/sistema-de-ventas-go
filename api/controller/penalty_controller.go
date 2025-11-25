package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PenaltyController struct {
	PenaltyRepository domain.PenaltyRepository
}

func (te *PenaltyController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Penalty domain.Penalty

	err := c.ShouldBind(&Penalty)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	
	if Penalty.Reason == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Reason is required"})
		return
	}
	if Penalty.DelayMinutes == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Delay Time is required"})
		return
	}
	
	Penalty.PenaltyID = uuid.New()

	err = te.PenaltyRepository.Create(c, Penalty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Penalty created successfully",
	})
}

func (te *PenaltyController) Fetch(c *gin.Context) {
	Penaltys, err := te.PenaltyRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Penaltys)
}

func (te *PenaltyController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Penaltys, err := te.PenaltyRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Penaltys)
}

func (te *PenaltyController) Update(c *gin.Context) {
	updatedPenalty := &domain.Penalty{}

	err := c.ShouldBind(updatedPenalty)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPenalty.PenaltyID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "PenaltyID is requiered to update"})
		return
	}

	err = te.PenaltyRepository.Update(c, *updatedPenalty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Penalty updated succesfully"})
}

func (te *PenaltyController) Delete(c *gin.Context) {
	PenaltyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.PenaltyRepository.Delete(c, PenaltyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Penalty delete succesfully"})
}
