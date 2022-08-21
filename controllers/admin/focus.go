package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FocusController struct{}

func (con FocusController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}
