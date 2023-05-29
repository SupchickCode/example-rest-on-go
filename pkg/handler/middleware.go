package handler

import (
	"errors"
	"fmt"
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

func (h Handler) getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s is not found", userCtx))
		return 0, errors.New(fmt.Sprintf("%s is not found", userCtx))
	}

	idINT, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("%s is of invalid type", userCtx))
		return 0, errors.New(fmt.Sprintf("%s is of invalid type", userCtx))
	}

	return idINT, nil
}
