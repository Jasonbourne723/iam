package bootstrap

import (
	"github.com/jasonbourne723/iam/docs"

	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/route"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunServe() {

	g := gin.Default()
	route.AddRoute(g)
	
	docs.SwaggerInfo.BasePath = "/api"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	g.Run(":8080")

}
