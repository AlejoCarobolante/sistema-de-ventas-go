package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EstadoPedidoController struct {
	EstadoPedidoRepository domain.EstadoPedidoRepository
}

func (te *EstadoPedidoController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var EstadoPedido domain.EstadoPedido

	err := c.ShouldBind(&EstadoPedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if EstadoPedido.NombreEP == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	EstadoPedido.ID = uuid.New()

	err = te.EstadoPedidoRepository.Create(c, EstadoPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "EstadoPedido created successfully",
	})
}

func (te *EstadoPedidoController) Fetch(c *gin.Context) {
	EstadoPedidos, err := te.EstadoPedidoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, EstadoPedidos)
}

func (te *EstadoPedidoController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	EstadoPedidos, err := te.EstadoPedidoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, EstadoPedidos)
}

func (te *EstadoPedidoController) Update(c *gin.Context) {
	updatedEstadoPedido := &domain.EstadoPedido{}

	err := c.ShouldBind(updatedEstadoPedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEstadoPedido.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID EstadoPedido is requiered to update"})
		return
	}

	err = te.EstadoPedidoRepository.Update(c, *updatedEstadoPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "EstadoPedido updated succesfully"})
}

func (te *EstadoPedidoController) Delete(c *gin.Context) {
	EstadoPedidoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.EstadoPedidoRepository.Delete(c, EstadoPedidoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "EstadoPedido delete succesfully"})
}
