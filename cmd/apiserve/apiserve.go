package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/route"
)

func main() {

	g := gin.Default()

	route.AddRoute(g)

	g.Run(":8080")

}
