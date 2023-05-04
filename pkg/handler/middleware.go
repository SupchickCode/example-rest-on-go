package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationToken = "Authorization"
)

func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationToken)

	if header == "" {
		newErrorResponse(c, 403, "Authorization header is empty")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, 403, "Authorization format is invalid")
		return
	}
}
