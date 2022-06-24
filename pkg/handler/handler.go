package handler

import (
	"billing/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/bill")
	{

		//api.GET("/") //all bills
		api.GET("/:id", h.Balance)
		api.PUT("/", h.AddMoney)         // add money
		api.PUT("/transfer", h.Transfer) //transaction

	}
	return router
}
