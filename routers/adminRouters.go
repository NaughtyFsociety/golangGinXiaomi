package routers

import (
	"ginXiaomi/controllers/admin"
	"ginXiaomi/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	// middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.GET("/doLogin", admin.LoginController{}.Dologin)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)

		//adminRouters.GET("/user", func(c *gin.Context) {
		//	c.String(http.StatusOK, "/user")
		//})
		//adminRouters.GET("/edit", func(c *gin.Context) {
		//	c.String(http.StatusOK, "/edit")
		//})
		//adminRouters.GET("/delete", func(c *gin.Context) {
		//	c.String(http.StatusOK, "/delete")
		//})
		//adminRouters.GET("/article", func(c *gin.Context) {
		//	c.String(http.StatusOK, "/article")
		//})
	}
}
