package handler

import (
	"github.com/gin-gonic/gin"
	"testTask/pkg/service"
)

type Handler struct {
	tasks service.Task
}

func NewHandler(tasks service.Task) *Handler {
	return &Handler{tasks: tasks}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.POST("/task1", h.answerTask1)
	router.POST("/task2", h.answerTask2)
	router.POST("/task3", h.answerTask3)

	return router
}
