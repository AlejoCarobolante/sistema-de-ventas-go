package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationTypeController struct {
	ReservationTypeRepository domain.ReservationTypeRepository
}

func (te *ReservationTypeController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var ReservationType domain.ReservationType

	err := c.ShouldBind(&ReservationType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if ReservationType.RTName == "" {
	c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}


	ReservationType.ReservationTypeID = uuid.New()

	err = te.ReservationTypeRepository.Create(c, ReservationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ReservationType created successfully",
	})
}

func (te *ReservationTypeController) Fetch(c *gin.Context) {
	ReservationTypes, err := te.ReservationTypeRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReservationTypes)
}

func (te *ReservationTypeController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ReservationTypes, err := te.ReservationTypeRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ReservationTypes)
}

func (te *ReservationTypeController) Update(c *gin.Context) {
	updatedReservationType := &domain.ReservationType{}

	err := c.ShouldBind(updatedReservationType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedReservationType.ReservationTypeID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ReservationTypeID is requiered to update"})
		return
	}

	err = te.ReservationTypeRepository.Update(c, *updatedReservationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationType updated succesfully"})
}

func (te *ReservationTypeController) Delete(c *gin.Context) {
	ReservationTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.ReservationTypeRepository.Delete(c, ReservationTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationType delete succesfully"})
}
