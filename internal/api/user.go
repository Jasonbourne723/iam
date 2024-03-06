package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/dto"
	"github.com/jasonbourne723/iam/internal/service"
)

type UserController struct {
}

func (u *UserController) List(ctx *gin.Context) {

}

func (u *UserController) Add(ctx *gin.Context) {
	var addUser *dto.AddUserDto
	if err := ctx.ShouldBindJSON(addUser); err != nil {
		fmt.Printf("\"\": %v\n", "")
	}
	userService := service.UserService{}
	userService.Add(addUser)
}

func (u *UserController) Delete(ctx *gin.Context) {
	//ctx.ShouldBindQuery()
}
