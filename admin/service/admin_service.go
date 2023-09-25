package service

import (
	"fmt"
	"frame/model"
	"frame/pkg/sql/sqltypes"
	"frame/pkg/utils"
	"frame/service"
	"frame/store/db"
	"github.com/spf13/cast"
	"mime/multipart"
	"time"
)

type adminService struct{}

var AdminService = new(adminService)

// 添加日志记录
func (s *adminService) AddLog(adminId uint, notes, ip string) error {
	log := &model.AdminLog{
		AdminId: adminId,
		Notes:   notes,
		Ip:      ip,
	}
	return db.Shared().AddAdminLog(log)
}

// 上传文件
func (s *adminService) Upload(h *multipart.FileHeader, admin *model.AdminUser) (string, error) {
	// 文件后缀
	fileType := utils.ResolveFileType(h.Filename)

	now := time.Now()

	// 文件路径
	fileObj := fmt.Sprintf("admin/%d/%d/%s.%s", admin.ID, now.Year(), cast.ToString(now.UnixNano()), fileType)

	return service.UploadService.UploadFromFileHeader(h, fileObj)
}

// 修改密码
func (s *adminService) UpdatePassword(admin *model.AdminUser, newPass string) error {
	pass := sqltypes.NewPassword(newPass)
	admin.Password = &pass

	return db.Shared().SaveAdminUser(admin)
}

// 检查此管理员账户是否有权限
func (s *adminService) CheckAdminRole(admin *model.AdminUser, permit string) (bool, error) {
	role, err := db.Shared().GetAdminRoleById(admin.RoleId)
	if err != nil {
		return false, err
	}

	if role == nil || role.Disabled {
		return false, nil
	}

	// 全部权限
	if utils.InStringSlice("*", role.Permissions) {
		return true, nil
	}

	return utils.InStringSlice(permit, role.Permissions), nil
}

func (s *adminService) GetAdminById(id uint) (*model.AdminUser, error) {
	return db.Shared().GetAdminUserById(id)
}

func (s *adminService) SaveAdminUser(admin *model.AdminUser) error {
	return db.Shared().SaveAdminUser(admin)
}

// 查询管理员列表
func (s *adminService) QueryAdminUsers(id, roleId int64, username string, page, limit int) ([]*model.AdminUser, int64, error) {
	return db.Shared().QueryAdminUsers(id, roleId, username, page, limit)
}

// 查询管理员角色
func (s *adminService) QueryAdminRoles(page, limit int) ([]*model.AdminRole, int64, error) {
	return db.Shared().QueryAdminRoles(page, limit)
}

func (s *adminService) GetAdminRoleById(id uint) (*model.AdminRole, error) {
	return db.Shared().GetAdminRoleById(id)
}

// 保存管理员权限
func (s *adminService) SaveAdminRole(v *model.AdminRole) error {
	return db.Shared().SaveAdminRole(v)
}

// 初始化管理员
func (s *adminService) InitAdmin() error {
	tx, err := db.Shared().BeginTx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	pwd := sqltypes.NewPassword("123456")

	role := &model.AdminRole{
		Name:        "超级管理员",
		Permissions: []string{"*"},
	}

	if err := tx.SaveAdminRole(role); err != nil {
		return err
	}

	if err := tx.SaveAdminUser(&model.AdminUser{
		Username: "admin",
		Password: &pwd,
		RoleId:   role.ID,
		Name:     "超级管理员",
	}); err != nil {
		return err
	}

	return tx.Commit()
}
