package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasonbourne723/iam/internal/dto"
	"github.com/jasonbourne723/iam/internal/service"
)

type UserController struct {
}

// @Summary	用户列表
// @Schemes
// @Description
// @Tags		User
// @Param		pageIndex	query	int64	true	"用户信息"
// @Param		pageSize	query	int64	true	"用户信息"
// @Accept		json
// @Produce	json
// @Success	200	{object}	dto.Result
// @Router		/users [Get]
func (u *UserController) List(ctx *gin.Context) {
	pageIndex := ctx.Query("pageIndex")

	fmt.Printf("pageIndex: %v\n", pageIndex)
}

// @Summary	新增用户
// @Schemes
// @Description
// @Tags		User
// @Param		default	body	dto.AddUserDto	true	"用户信息"
// @Accept		json
// @Produce	json
// @Success	200	{object}	dto.Result
// @Router		/users [post]
func (u *UserController) Add(ctx *gin.Context) {
	var addUser *dto.AddUserDto
	if err := ctx.ShouldBindJSON(addUser); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	if err := service.UserSrv.Add(addUser); err != nil {
		ctx.JSON(500, dto.Result{
			Code:    -1,
			Message: err.Error(),
		})
	}
	ctx.JSON(200, dto.Result{
		Code:    0,
		Message: "Success",
	})
}

// @Summary	删除用户
// @Schemes
// @Description
// @Tags		User
// @Param		userId	path	int64	true	"用户Id"
// @Accept		json
// @Produce	json
// @Success	200	{object}	dto.Result
// @Router		/users/{userId} [delete]
func (u *UserController) Delete(ctx *gin.Context) {

	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(400, dto.Result{
			Code:    -1,
			Message: err.Error(),
		})
	}
	service.UserSrv.Delete(userId)
	ctx.JSON(200, dto.Result{
		Code:    0,
		Message: "Success",
	})
}

// @Summary	更新用户
// @Schemes
// @Description
// @Tags		User
// @Param		default	body	dto.UpdateUserDto	true	"用户信息"
// @Accept		json
// @Produce	json
// @Success	200	{object}	dto.Result
// @Router		/users [put]
func (u *UserController) Update(ctx *gin.Context) {
	var user dto.UpdateUserDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, dto.Result{
			Code:    -1,
			Message: err.Error(),
		})
	}
	if err := service.UserSrv.Update(&user); err != nil {
		ctx.JSON(400, dto.Result{
			Code:    -1,
			Message: err.Error(),
		})
	}
	ctx.JSON(200, dto.Result{
		Code:    0,
		Message: "Success",
	})
}
