package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VehicleController struct {
	VehicleRepository domain.VehicleRepository
}

func (vc *VehicleController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Vehicle domain.Vehicle

	err := c.ShouldBind(&Vehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Vehicle.LicensePlate == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "License Plate is required"})
		return
	}

	Vehicle.VehicleID = uuid.New()

	err = vc.VehicleRepository.Create(c, Vehicle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Vehicle created successfully",
	})
}

func (vc *VehicleController) Fetch(c *gin.Context) {
	Vehicles, err := vc.VehicleRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Vehicles)
}

func (vc *VehicleController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Vehicles, err := vc.VehicleRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Vehicles)
}

func (vc *VehicleController) Update(c *gin.Context) {
	updatedVehicle := &domain.Vehicle{}

	err := c.ShouldBind(updatedVehicle)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedVehicle.VehicleID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "VehicleID is requiered to update"})
		return
	}

	err = vc.VehicleRepository.Update(c, *updatedVehicle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Vehicle updated succesfully"})
}

func (vc *VehicleController) Delete(c *gin.Context) {
	VehicleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = vc.VehicleRepository.Delete(c, VehicleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Vehicle delete succesfully"})
}
