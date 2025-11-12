package controller

import (
	"net/http"
	"strconv"

	"gorm-template/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Entity1Controller struct {
	Entity1Repository domain.Entity1Repository
}

func (te *Entity1Controller) Create(c *gin.Context) { //Hay que ingresar todos los datos necesarios para crear
	var Entity1 domain.Entity1

	err := c.ShouldBind(&Entity1)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if Entity1.Name == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Name is required"})
		return
	}

	Entity1.ID = uuid.New()

	err = te.Entity1Repository.Create(c, Entity1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Entity1 created successfully",
	})
}

func (te *Entity1Controller) Fetch(c *gin.Context) {
	Entity1s, err := te.Entity1Repository.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Entity1s)
}

func (te *Entity1Controller) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	Entity1s, err := te.Entity1Repository.FetchById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Entity1s)
}

func (te *Entity1Controller) Update(c *gin.Context) {
	updatedEntity1 := &domain.Entity1{}

	err := c.ShouldBind(updatedEntity1)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if updatedEntity1.ID == uuid.Nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "ID Entity1 is requiered to update"})
		return
	}

	err = te.Entity1Repository.Update(c, *updatedEntity1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Entity1 updated succesfully"})
}

func (te *Entity1Controller) Delete(c *gin.Context) {
	Entity1ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	err = te.Entity1Repository.Delete(c, Entity1ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Entity1 delete succesfully"})
}
