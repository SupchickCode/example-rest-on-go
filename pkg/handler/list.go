package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type indexListResponce struct {
	Data []simpleRestAPI.TodoList `json:"data"`
}

func (h *Handler) indexList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAllLists(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, indexListResponce{
		Data: lists,
	})
}

func (h *Handler) showList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetListByID(id)

	if err != nil {
		newErrorResponse(c, http.StatusAlreadyReported, err.Error())
		return
	}

	logrus.Debug(fmt.Printf("%#+v", list))

	c.JSON(http.StatusOK, list)
}

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		return
	}

	var input simpleRestAPI.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userID, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) editList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
