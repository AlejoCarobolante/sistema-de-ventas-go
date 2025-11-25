package controller

import (
	"gorm-template/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClientController struct {
	ClientRepository domain.ClientRepository
}

func (cc *ClientController) Create(c *gin.Context) {
	var client domain.Client

	err := c.ShouldBind(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if client.Name == "" || client.Email == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Nombre y Mail requeridos"})
		return
	}

	client.ID = uuid.New()

	err = cc.ClientRepository.Create(c, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Client create successfully",
	})
}

func (cc *ClientController) Fetch(c *gin.Context) {
	clients, err := cc.ClientRepository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (cc *ClientController) FetchByID(c *gin.Context) {
	clientID := c.Param("id")
	client, err := cc.ClientRepository.FetchByID(c, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, client)
}

func (cc *ClientController) Update(c *gin.Context) {
	updateClient := &domain.Client{}
	err := c.ShouldBind(updateClient)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if updateClient.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID requerido"})
		return
	}

	err = cc.ClientRepository.Update(c, *updateClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Client updated successfully"})
}

func (cc *ClientController) Delete(c *gin.Context) {
	clientID := c.Param("id")
	err := cc.ClientRepository.Delete(c, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Client deleted successfully"})
}
