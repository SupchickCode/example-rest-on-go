package handler

import (
	"net/http"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SingInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singUp(c *gin.Context) {
	var input simpleRestAPI.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Fatalf("%s", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {
	var input SingInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Fatalf("%s", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Password, input.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
