package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpotTypeController struct {
	SpotTypeRepository domain.SpotTypeRepository
}

func (te *SpotTypeController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var SpotType domain.SpotType

	err := c.ShouldBind(&SpotType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	SpotType.SpotTypeID = uuid.New()

	err = te.SpotTypeRepository.Create(c, SpotType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "SpotType created successfully",
	})
}

func (te *SpotTypeController) Fetch(c *gin.Context) {
	SpotTypes, err := te.SpotTypeRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SpotTypes)
}

func (te *SpotTypeController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	SpotTypes, err := te.SpotTypeRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SpotTypes)
}

func (te *SpotTypeController) Update(c *gin.Context) {
	updatedSpotType := &domain.SpotType{}

	err := c.ShouldBind(updatedSpotType)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedSpotType.SpotTypeID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "SpotTypeID is requiered to update"})
		return
	}

	err = te.SpotTypeRepository.Update(c, *updatedSpotType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "SpotType updated succesfully"})
}

func (te *SpotTypeController) Delete(c *gin.Context) {
	SpotTypeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.SpotTypeRepository.Delete(c, SpotTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "SpotType delete succesfully"})
}
