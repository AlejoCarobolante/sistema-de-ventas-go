package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentController struct {
	PaymentRepository domain.PaymentRepository
}

func (pac *PaymentController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Payment domain.Payment

	err := c.ShouldBind(&Payment)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Payment.PaymentAmmount == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Ammount is required"})
		return
	}

	Payment.PaymentID = uuid.New()

	err = pac.PaymentRepository.Create(c, Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Payment created successfully",
	})
}

func (pac *PaymentController) Fetch(c *gin.Context) {
	Payments, err := pac.PaymentRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Payments)
}

func (pac *PaymentController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Payments, err := pac.PaymentRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Payments)
}

func (pac *PaymentController) Update(c *gin.Context) {
	updatedPayment := &domain.Payment{}

	err := c.ShouldBind(updatedPayment)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPayment.PaymentID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "PaymentID is requiered to update"})
		return
	}

	err = pac.PaymentRepository.Update(c, *updatedPayment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Payment updated succesfully"})
}

func (pac *PaymentController) Delete(c *gin.Context) {
	PaymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = pac.PaymentRepository.Delete(c, PaymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Payment delete succesfully"})
}
