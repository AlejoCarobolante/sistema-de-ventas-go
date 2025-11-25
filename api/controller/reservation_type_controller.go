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

func (rtc *ReservationTypeController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
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

	err = rtc.ReservationTypeRepository.Create(c, ReservationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ReservationType created successfully",
	})
}

func (rtc *ReservationTypeController) Fetch(c *gin.Context) {
	ReservationTypes, err := rtc.ReservationTypeRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReservationTypes)
}

func (rtc *ReservationTypeController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ReservationTypes, err := rtc.ReservationTypeRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ReservationTypes)
}

func (rtc *ReservationTypeController) Update(c *gin.Context) {
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

	err = rtc.ReservationTypeRepository.Update(c, *updatedReservationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationType updated succesfully"})
}

func (rtc *ReservationTypeController) Delete(c *gin.Context) {
	ReservationTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = rtc.ReservationTypeRepository.Delete(c, ReservationTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationType delete succesfully"})
}
