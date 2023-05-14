package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationToken = "Authorization"
	userCtx            = "userId"
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

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
