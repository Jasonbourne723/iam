package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/api"
)

func AddRoute(g *gin.Engine) {

	userController := &api.UserController{}
	v1 := g.Group("/api/")
	{
		v1.Use(gin.Recovery())

		v1.GET("/users", userController.List)
		v1.POST("/users", userController.Add)
		v1.DELETE("/users/:userId", userController.Delete)
		v1.PUT("/users", userController.Update)
	}

}
