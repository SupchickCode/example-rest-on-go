package handler

import (
	"github.com/SupchickCode/simpleRestAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sing-up", h.singUp)
		auth.POST("sing-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.indexList)
			lists.GET("/:id", h.showList)
			lists.PUT("/:id", h.editList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.indexItem)
				items.GET("/:id", h.showItem)
				items.PUT("/:id", h.editItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}
	}

	return router
}
