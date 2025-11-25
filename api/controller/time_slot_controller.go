package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TimeSlotController struct {
	TimeSlotRepository domain.TimeSlotRepository
}

func (tsc *TimeSlotController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var TimeSlot domain.TimeSlot

	err := c.ShouldBind(&TimeSlot)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if TimeSlot.DayOfWeek == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Day of Week is required"})
		return
	}

	TimeSlot.TimeSlotID = uuid.New()

	err = tsc.TimeSlotRepository.Create(c, TimeSlot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "TimeSlot created successfully",
	})
}

func (tsc *TimeSlotController) Fetch(c *gin.Context) {
	TimeSlots, err := tsc.TimeSlotRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TimeSlots)
}

func (tsc *TimeSlotController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	TimeSlots, err := tsc.TimeSlotRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, TimeSlots)
}

func (tsc *TimeSlotController) Update(c *gin.Context) {
	updatedTimeSlot := &domain.TimeSlot{}

	err := c.ShouldBind(updatedTimeSlot)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedTimeSlot.TimeSlotID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "TimeSlotID is requiered to update"})
		return
	}

	err = tsc.TimeSlotRepository.Update(c, *updatedTimeSlot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "TimeSlot updated succesfully"})
}

func (tsc *TimeSlotController) Delete(c *gin.Context) {
	TimeSlotID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = tsc.TimeSlotRepository.Delete(c, TimeSlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "TimeSlot delete succesfully"})
}
