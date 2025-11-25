package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ParkingController struct {
	ParkingRepository domain.ParkingRepository
}

func (pc *ParkingController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Parking domain.Parking

	err := c.ShouldBind(&Parking)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	Parking.ParkingID = uuid.New()

	err = pc.ParkingRepository.Create(c, Parking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Parking created successfully",
	})
}

func (pc *ParkingController) Fetch(c *gin.Context) {
	Parkings, err := pc.ParkingRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Parkings)
}

func (pc *ParkingController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Parkings, err := pc.ParkingRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Parkings)
}

func (pc *ParkingController) Update(c *gin.Context) {
	updatedParking := &domain.Parking{}

	err := c.ShouldBind(updatedParking)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedParking.ParkingID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ParkingID is requiered to update"})
		return
	}

	err = pc.ParkingRepository.Update(c, *updatedParking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Parking updated succesfully"})
}

func (pc *ParkingController) Delete(c *gin.Context) {
	ParkingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = pc.ParkingRepository.Delete(c, ParkingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Parking delete succesfully"})
}
