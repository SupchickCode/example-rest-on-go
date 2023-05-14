package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) indexList(c *gin.Context) {
	id, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) showList(c *gin.Context) {

}

func (h *Handler) createList(c *gin.Context) {

}

func (h *Handler) editList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
