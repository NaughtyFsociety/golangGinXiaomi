package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "Index")
}
