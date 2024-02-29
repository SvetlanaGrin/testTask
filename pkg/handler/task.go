package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"testTask"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) answerTask1(c *gin.Context) {

	var input testTask.InputTask1
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	answer, err := h.tasks.AnswerTask1(input.Array)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": answer,
	})
}

func (h *Handler) answerTask2(c *gin.Context) {
	var input testTask.InputTask2
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	answer, err := h.tasks.AnswerTask2(input.InputString, input.Anagram)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": answer,
	})
}
func (h *Handler) answerTask3(c *gin.Context) {
	var input testTask.InputTask3
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}
	answer, err := h.tasks.AnswerTask3(input.InputString, input.Key)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": answer,
	})
}
