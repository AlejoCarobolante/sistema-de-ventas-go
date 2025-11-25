package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SpotController struct {
	SpotRepository domain.SpotRepository
}

func (te *SpotController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Spot domain.Spot

	err := c.ShouldBind(&Spot)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	Spot.SpotID = uuid.New()

	err = te.SpotRepository.Create(c, Spot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Spot created successfully",
	})
}

func (te *SpotController) Fetch(c *gin.Context) {
	Spots, err := te.SpotRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Spots)
}

func (te *SpotController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Spots, err := te.SpotRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Spots)
}

func (te *SpotController) Update(c *gin.Context) {
	updatedSpot := &domain.Spot{}

	err := c.ShouldBind(updatedSpot)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedSpot.SpotID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "SpotID is requiered to update"})
		return
	}

	err = te.SpotRepository.Update(c, *updatedSpot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Spot updated succesfully"})
}

func (te *SpotController) Delete(c *gin.Context) {
	SpotID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.SpotRepository.Delete(c, SpotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Spot delete succesfully"})
}
