package handler

import (
	"net/http"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) singUp(c *gin.Context) {
	var input simpleRestAPI.User

	// Мы говорим что будем парсить json в &input
	if err := c.BindJSON(&input); err != nil {
		logrus.Fatalf("%s", err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	h.services.Authorization
}

func (h *Handler) singIn(c *gin.Context) {

}
