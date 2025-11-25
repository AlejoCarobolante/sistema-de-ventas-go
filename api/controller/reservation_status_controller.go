package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationStatusController struct {
	ReservationStatusRepository domain.ReservationStatusRepository
}

func (te *ReservationStatusController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var ReservationStatus domain.ReservationStatus

	err := c.ShouldBind(&ReservationStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if ReservationStatus.RSName == "" {
	c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}


	ReservationStatus.ReservationStatusID = uuid.New()

	err = te.ReservationStatusRepository.Create(c, ReservationStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "ReservationStatus created successfully",
	})
}

func (te *ReservationStatusController) Fetch(c *gin.Context) {
	ReservationStatuss, err := te.ReservationStatusRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReservationStatuss)
}

func (te *ReservationStatusController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ReservationStatuss, err := te.ReservationStatusRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ReservationStatuss)
}

func (te *ReservationStatusController) Update(c *gin.Context) {
	updatedReservationStatus := &domain.ReservationStatus{}

	err := c.ShouldBind(updatedReservationStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedReservationStatus.ReservationStatusID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ReservationStatusID is requiered to update"})
		return
	}

	err = te.ReservationStatusRepository.Update(c, *updatedReservationStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationStatus updated succesfully"})
}

func (te *ReservationStatusController) Delete(c *gin.Context) {
	ReservationStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.ReservationStatusRepository.Delete(c, ReservationStatusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "ReservationStatus delete succesfully"})
}
