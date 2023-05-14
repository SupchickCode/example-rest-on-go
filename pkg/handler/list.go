package handler

import (
	"fmt"
	"net/http"

	"github.com/SupchickCode/simpleRestAPI"
	"github.com/gin-gonic/gin"
)

func (h *Handler) indexList(c *gin.Context) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s is not found", userCtx))
		return
	}

	var input simpleRestAPI.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) showList(c *gin.Context) {

}

func (h *Handler) createList(c *gin.Context) {

}

func (h *Handler) editList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
