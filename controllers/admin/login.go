package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(200, "admin/login/index.html", gin.H{})
}

func (con LoginController) Dologin(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
