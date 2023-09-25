package rest

import (
	"frame/admin/service"
	"frame/pkg/ecode"
	"frame/pkg/request"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthController struct {
}

type AuthLoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 登录
func (c *AuthController) Login(ctx *gin.Context) {
	app := request.New(ctx)

	var req AuthLoginRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	if req.Username == "" || req.Password == "" {
		app.Response(ecode.ErrRequest)
		return
	}

	_, count, err := service.AdminService.QueryAdminUsers(0, 0, "", 1, 20)
	if err != nil {
		app.Response(err)
		return
	}
	if count == 0 {
		if err := service.AdminService.InitAdmin(); err != nil {
			app.Response(err)
			return
		}
	}

	user, token, err := service.AuthService.Login(req.Username, req.Password)
	if err != nil {
		app.Response(err)
		return
	}

	service.AdminService.AddLog(user.ID, "登录成功", ctx.ClientIP())

	app.Response(nil, token)
}

// 退出登录
func (c *AuthController) OutLogin(ctx *gin.Context) {
	time.Sleep(time.Second)
	request.New(ctx).Response(nil)
}
