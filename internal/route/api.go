package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/api"
)

func AddRoute(g *gin.Engine) {

	userController := &api.UserController{}
	v1 := g.Group("/api/user")
	{
		v1.GET("/user", userController.List)
		v1.POST("/user", userController.Add)
		v1.DELETE("/user", userController.Delete)
	}

}
