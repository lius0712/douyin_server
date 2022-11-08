package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

type UserRegisterParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
