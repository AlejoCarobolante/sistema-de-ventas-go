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

func (te *PaymentController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
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

	err = te.PaymentRepository.Create(c, Payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Payment created successfully",
	})
}

func (te *PaymentController) Fetch(c *gin.Context) {
	Payments, err := te.PaymentRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Payments)
}

func (te *PaymentController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Payments, err := te.PaymentRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Payments)
}

func (te *PaymentController) Update(c *gin.Context) {
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

	err = te.PaymentRepository.Update(c, *updatedPayment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Payment updated succesfully"})
}

func (te *PaymentController) Delete(c *gin.Context) {
	PaymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.PaymentRepository.Delete(c, PaymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Payment delete succesfully"})
}
