package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RateController struct {
	RateRepository domain.RateRepository
}

func (rc *RateController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Rate domain.Rate

	err := c.ShouldBind(&Rate)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	Rate.RateID = uuid.New()

	err = rc.RateRepository.Create(c, Rate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Rate created successfully",
	})
}

func (rc *RateController) Fetch(c *gin.Context) {
	Rates, err := rc.RateRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Rates)
}

func (rc *RateController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Rates, err := rc.RateRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Rates)
}

func (rc *RateController) Update(c *gin.Context) {
	updatedRate := &domain.Rate{}

	err := c.ShouldBind(updatedRate)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedRate.RateID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "RateID is requiered to update"})
		return
	}

	err = rc.RateRepository.Update(c, *updatedRate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Rate updated succesfully"})
}

func (rc *RateController) Delete(c *gin.Context) {
	RateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = rc.RateRepository.Delete(c, RateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Rate delete succesfully"})
}
