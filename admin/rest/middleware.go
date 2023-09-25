package rest

import (
	"frame/admin/service"
	"frame/pkg/ecode"
	"frame/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 检查并获取登录用户
func checkToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		token := ctx.GetHeader("X-Access-Token")
		if token == "" {
			app.AbortResponse(errors.New("未权限认证"))
			return
		}

		user, err := service.AuthService.CheckToken(token)
		if err != nil {
			app.AbortResponse(err)
			return
		}

		app.SetAdminUser(user)
		ctx.Next()
	}
}

// 检查游客登录信息
func checkGhost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		token := ctx.GetHeader("X-Access-Token")
		if token != "" {
			user, err := service.AuthService.CheckToken(token)
			if err == nil {
				app.SetAdminUser(user)
			}
		}

		ctx.Next()
	}
}

// 检查权限
func permit(permit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app := request.New(ctx)

		user := app.GetAdminUser()
		if user == nil || user.RoleId <= 0 {
			app.AbortResponse(ecode.ErrForbidden)
			return
		}

		pass, err := service.AdminService.CheckAdminRole(user, permit)
		if err != nil {
			app.AbortResponse(err)
			return
		}

		if !pass {
			app.AbortResponse(ecode.ErrNoPermission)
			return
		}

		ctx.Next()
	}
}
