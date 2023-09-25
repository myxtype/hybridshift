package rest

import (
	"frame/model"
	"github.com/shopspring/decimal"
	"time"
)

type ListVo struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

func NewListVo(list interface{}, total int64) *ListVo {
	return &ListVo{
		List:  list,
		Total: total,
	}
}

type UserVo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Address   string    `json:"address"`  // 钱包地址
	Nonce     uint64    `json:"nonce"`    // 登录的Nonce
	Disabled  bool      `json:"disabled"` // 禁用
}

func NewUserVo(v *model.User) *UserVo {
	return &UserVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Address:   v.Address,
		Nonce:     v.Nonce,
		Disabled:  v.Disabled,
	}
}

type WalletVo struct {
	Point decimal.Decimal `json:"point"` // 积分数量
	Token decimal.Decimal `json:"token"` // 代币数量
}

func NewWalletVo(v *model.Wallet) *WalletVo {
	vo := &WalletVo{
		Point: v.Point,
		Token: v.Token,
	}
	if v != nil {
		vo.Point = v.Point
		vo.Token = v.Token
	}
	return vo
}

type UserMe struct {
	*UserVo
	Account *WalletVo `json:"account"`
}

type UserOauthVo struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UserId        uint      `json:"userId"`
	OauthToken    string    `json:"oauthToken"`    // 授权Token
	OauthSecret   string    `json:"oauthSecret"`   // 授权Secret
	OauthId       string    `json:"oauthId"`       // 账号
	OauthName     string    `json:"oauthName"`     // 显示名称
	OauthUserName string    `json:"oauthUserName"` // 用户名
}

func NewUserOauthVo(v *model.UserOauth) *UserOauthVo {
	return &UserOauthVo{
		ID:            v.ID,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		UserId:        v.UserId,
		OauthToken:    v.OauthToken,
		OauthSecret:   v.OauthSecret,
		OauthId:       v.OauthId,
		OauthName:     v.OauthName,
		OauthUserName: v.OauthUserName,
	}
}

// Task 任务
type TaskVo struct {
	ID              uint             `json:"id"`
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
	UserId          uint             `json:"userId"`
	TotalToken      decimal.Decimal  `json:"totalToken"`      // 总币数
	RemainingToken  decimal.Decimal  `json:"remainingToken"`  // 剩余代币数
	TwitterId       string           `json:"twitterId"`       // 推特账户ID，指定推特账号
	TwitterUserName string           `json:"twitterUserName"` // 推特用户名
	TweetURL        string           `json:"tweetUrl"`        // 推特URL，指定推特
	Status          model.TaskStatus `json:"status"`          // 状态
}

func NewTaskVo(v *model.Task) *TaskVo {
	return &TaskVo{
		ID:              v.ID,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
		UserId:          v.UserId,
		TotalToken:      v.TotalToken,
		RemainingToken:  v.RemainingToken,
		TwitterId:       v.TwitterId,
		TwitterUserName: v.TwitterUserName,
		TweetURL:        v.TweetURL,
		Status:          v.Status,
	}
}
