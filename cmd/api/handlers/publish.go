package handlers

import "github.com/gin-gonic/gin"

func Publish(c *gin.Context) {
	var publishParam PublishParam
	token := c.PostForm("token")
	title := c.PostForm("title")

	file, _, err := c.Request.FormFile("data")
	if err != nil {

	}
	defer file.Close()

	publishParam.Token = token
	publishParam.Title = title

}
