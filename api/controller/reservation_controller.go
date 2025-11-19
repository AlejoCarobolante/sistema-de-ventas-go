package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationController struct {
	ReservationRepository domain.ReservationRepository
}

func (te *ReservationController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Reservation domain.Reservation

	err := c.ShouldBind(&Reservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Reservation.Start.IsZero() {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Start is required"})
		return
	}

	if Reservation.End.IsZero(){
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "End is required"})
		return
	} 
	Reservation.ReservationID = uuid.New()

	err = te.ReservationRepository.Create(c, Reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Reservation created successfully",
	})
}

func (te *ReservationController) Fetch(c *gin.Context) {
	Reservations, err := te.ReservationRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Reservations)
}

func (te *ReservationController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Reservations, err := te.ReservationRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Reservations)
}

func (te *ReservationController) Update(c *gin.Context) {
	updatedReservation := &domain.Reservation{}

	err := c.ShouldBind(updatedReservation)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedReservation.ReservationID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Reservation is requiered to update"})
		return
	}

	err = te.ReservationRepository.Update(c, *updatedReservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Reservation updated succesfully"})
}

func (te *ReservationController) Delete(c *gin.Context) {
	ReservationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.ReservationRepository.Delete(c, ReservationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Reservation delete succesfully"})
}
