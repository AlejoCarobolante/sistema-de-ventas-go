package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VehicleTypeController struct {
	VehicleTypeRepository domain.VehicleTypeRepository
}

func (vtc *VehicleTypeController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var VehicleType domain.VehicleType

	err := c.ShouldBind(&VehicleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if VehicleType.Name == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Ammount is required"})
		return
	}

	VehicleType.VehicleTypeID = uuid.New()

	err = vtc.VehicleTypeRepository.Create(c, VehicleType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "VehicleType created successfully",
	})
}

func (vtc *VehicleTypeController) Fetch(c *gin.Context) {
	VehicleTypes, err := vtc.VehicleTypeRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VehicleTypes)
}

func (vtc *VehicleTypeController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	VehicleTypes, err := vtc.VehicleTypeRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, VehicleTypes)
}

func (vtc *VehicleTypeController) Update(c *gin.Context) {
	updatedVehicleType := &domain.VehicleType{}

	err := c.ShouldBind(updatedVehicleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedVehicleType.VehicleTypeID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "VehicleTypeID is requiered to update"})
		return
	}

	err = vtc.VehicleTypeRepository.Update(c, *updatedVehicleType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "VehicleType updated succesfully"})
}

func (vtc *VehicleTypeController) Delete(c *gin.Context) {
	VehicleTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = vtc.VehicleTypeRepository.Delete(c, VehicleTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "VehicleType delete succesfully"})
}
