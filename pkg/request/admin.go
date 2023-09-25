package request

import "frame/model"

// 获取管理员用户
func (r *AppRequest) GetAdminUser() *model.AdminUser {
	v, found := r.c.Get("__admin_user")
	if found {
		return v.(*model.AdminUser)
	}
	return nil
}

// 设置管理员用户
func (r *AppRequest) SetAdminUser(user *model.AdminUser) {
	r.c.Set("__admin_user", user)
}
