package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PedidoController struct {
	PedidoRepository domain.PedidoRepository
}

func (te *PedidoController) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Pedido domain.Pedido

	err := c.ShouldBind(&Pedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Pedido.NumeroPedido == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Numero is required"})
		return
	}

	Pedido.ID = uuid.New()

	err = te.PedidoRepository.Create(c, Pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Pedido created successfully",
	})
}

func (te *PedidoController) Fetch(c *gin.Context) {
	Pedidos, err := te.PedidoRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Pedidos)
}

func (te *PedidoController) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Pedidos, err := te.PedidoRepository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Pedidos)
}

func (te *PedidoController) Update(c *gin.Context) {
	updatedPedido := &domain.Pedido{}

	err := c.ShouldBind(updatedPedido)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedPedido.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Pedido is requiered to update"})
		return
	}

	err = te.PedidoRepository.Update(c, *updatedPedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Pedido updated succesfully"})
}

func (te *PedidoController) Delete(c *gin.Context) {
	PedidoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.PedidoRepository.Delete(c, PedidoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Pedido delete succesfully"})
}
