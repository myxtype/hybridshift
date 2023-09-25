package service

import (
	"frame/conf"
	"frame/model"
	"frame/pkg/jwttool"
	"frame/store/db"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type authService struct{}

var AuthService = new(authService)

// 通过用户名密码登录
func (s *authService) Login(username, password string) (*model.AdminUser, string, error) {
	admin, err := db.Shared().GetAdminUserByUsername(username)
	if err != nil {
		return nil, "", err
	}

	if admin == nil {
		return nil, "", errors.New("用户不存在")
	}
	if !admin.Password.Check(password) {
		return nil, "", errors.New("密码错误")
	}

	token, err := s.LoginFromAdmin(admin)
	if err != nil {
		return nil, "", err
	}

	return admin, token, err
}

// 登录
func (s *authService) LoginFromAdmin(admin *model.AdminUser) (string, error) {
	if admin.Disabled {
		return "", errors.New("账号已被禁用")
	}

	// 更新登录版本
	admin.LoginVersion++
	if err := db.Shared().UpdateAdminUserLoginVersion(admin); err != nil {
		return "", err
	}

	token, err := s.GenerateToken(admin)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 生成Token
func (s *authService) GenerateToken(user *model.AdminUser) (string, error) {
	now := time.Now()
	claim := jwttool.AuthClaims{
		UserId:       user.ID,
		PasswordHash: user.Password.Hash,
		ExpiresAt:    now.AddDate(0, 1, 0).Unix(),
		IssuedAt:     now.Unix(),
		Issuer:       "admin",
		Version:      user.LoginVersion,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(conf.GetConfig().JwtSecret.Admin))
}

// 解析Token
func (s *authService) ParseToken(tokenStr string) (*jwttool.AuthClaims, error) {
	var claim jwttool.AuthClaims

	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GetConfig().JwtSecret.Admin), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("认证错误")
	}

	return &claim, nil
}

// 检查Token
func (s *authService) CheckToken(tokenStr string) (*model.AdminUser, error) {
	claims, err := s.ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	// 获取用户
	user, err := db.Shared().GetAdminUserById(claims.UserId)
	if err != nil {
		return nil, err
	}

	// 判断是否禁用此用户
	if user == nil || user.Disabled {
		return nil, errors.New("账号已被禁用")
	}
	// 多设备登录
	if user.LoginVersion != claims.Version {
		return nil, errors.New("账号已被其他设备登录")
	}
	// 登录密码
	if user.Password.Hash != claims.PasswordHash {
		return nil, errors.New("账号密码已被修改")
	}

	return user, nil
}
