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

func (te *VehicleTypeController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
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

	err = te.VehicleTypeRepository.Create(c, VehicleType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "VehicleType created successfully",
	})
}

func (te *VehicleTypeController) Fetch(c *gin.Context) {
	VehicleTypes, err := te.VehicleTypeRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, VehicleTypes)
}

func (te *VehicleTypeController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	VehicleTypes, err := te.VehicleTypeRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, VehicleTypes)
}

func (te *VehicleTypeController) Update(c *gin.Context) {
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

	err = te.VehicleTypeRepository.Update(c, *updatedVehicleType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "VehicleType updated succesfully"})
}

func (te *VehicleTypeController) Delete(c *gin.Context) {
	VehicleTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.VehicleTypeRepository.Delete(c, VehicleTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "VehicleType delete succesfully"})
}
