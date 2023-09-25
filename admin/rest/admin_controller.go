package rest

import (
	"errors"
	"frame/admin/service"
	"frame/model"
	"frame/pkg/ecode"
	"frame/pkg/request"
	"frame/pkg/sql/sqltypes"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminController struct {
}

// 获取当前登录管理员
func (c *AdminController) Current(ctx *gin.Context) {
	app := request.New(ctx)

	user := app.GetAdminUser()
	time.Sleep(time.Second)
	app.Response(nil, NewAdminUserInfoVo(user))
}

type AdminUpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// 修改密码
func (c *AdminController) UpdatePassword(ctx *gin.Context) {
	app := request.New(ctx)

	var req AdminUpdatePasswordRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	admin := app.GetAdminUser()
	if !admin.Password.Check(req.OldPassword) {
		app.Response(errors.New("账号密码错误"))
		return
	}

	app.Response(service.AdminService.UpdatePassword(admin, req.NewPassword))
}

// 上传文件
func (c *AdminController) Upload(ctx *gin.Context) {
	app := request.New(ctx)

	f, err := ctx.FormFile("file")
	if err != nil {
		app.Response(err)
		return
	}

	url, err := service.AdminService.Upload(f, app.GetAdminUser())
	if err != nil {
		app.Response(err)
		return
	}

	app.Response(nil, url)
}

type QueryAdminUsersRequest struct {
	PageRequest
	ID       int64  `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	RoleId   int64  `json:"roleId" form:"roleId"`
}

// 查询管理员列表
func (c *AdminController) QueryAdminUsers(ctx *gin.Context) {
	app := request.New(ctx)

	var req QueryAdminUsersRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	data, count, err := service.AdminService.QueryAdminUsers(req.ID, req.RoleId, req.Username, req.Page, req.Limit)
	if err != nil {
		app.Response(err)
		return
	}

	vos := []*AdminUserVo{}
	for _, n := range data {
		vos = append(vos, NewAdminUserVo(n))
	}

	app.Response(err, NewListResult(count, vos))
}

type QueryAdminRolesRequest struct {
	PageRequest
}

// 查询管理员权限列表
func (c *AdminController) QueryAdminRoles(ctx *gin.Context) {
	app := request.New(ctx)

	var req QueryAdminRolesRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	data, count, err := service.AdminService.QueryAdminRoles(req.Page, req.Limit)
	if err != nil {
		app.Response(err)
		return
	}

	vos := []*AdminRoleVo{}
	for _, n := range data {
		vos = append(vos, NewAdminRoleVo(n))
	}

	app.Response(nil, NewListResult(count, vos))
}

type SaveAdminUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RoleId   uint   `json:"roleId"`   // 角色
	Name     string `json:"name"`     // 昵称
	Disabled bool   `json:"disabled"` // 是否禁用
	Password string `json:"password"`
}

// 保存管理员
func (c *AdminController) SaveAdminUser(ctx *gin.Context) {
	app := request.New(ctx)

	var req SaveAdminUserRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	var user *model.AdminUser

	if req.ID > 0 {
		var err error
		user, err = service.AdminService.GetAdminById(req.ID)
		if err != nil {
			app.Response(err)
			return
		}
		if user == nil {
			app.Response(ecode.ErrNotFind)
			return
		}
	} else {
		user = &model.AdminUser{}
		if req.Password == "" {
			app.Response(errors.New("请输入密码"))
			return
		}
	}

	user.Name = req.Name
	user.RoleId = req.RoleId
	user.Username = req.Username
	user.Disabled = req.Disabled
	if req.Password != "" {
		tmp := sqltypes.NewPassword(req.Password)
		user.Password = &tmp
	}

	err := service.AdminService.SaveAdminUser(user)
	app.Response(err)
}

type SaveAdminRoleRequest struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`        // 角色名称
	Permissions sqltypes.StringArray `json:"permissions"` // 权限列表
	Disabled    bool                 `json:"disabled"`    // 是否禁用
}

func (c *AdminController) SaveAdminRole(ctx *gin.Context) {
	app := request.New(ctx)

	var req SaveAdminRoleRequest
	if err := ctx.Bind(&req); err != nil {
		app.Response(err)
		return
	}

	var role *model.AdminRole
	if req.ID > 0 {
		var err error
		role, err = service.AdminService.GetAdminRoleById(req.ID)
		if err != nil {
			app.Response(err)
			return
		}
		if role == nil {
			app.Response(ecode.ErrNotFind)
			return
		}
	} else {
		role = &model.AdminRole{}
	}

	role.Name = req.Name
	role.Permissions = req.Permissions
	role.Disabled = req.Disabled

	err := service.AdminService.SaveAdminRole(role)
	app.Response(err)
}
