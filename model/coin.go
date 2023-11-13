package model

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name    string      // 名称
	Network CoinNetwork // 网络
	Logo    string      // 图标
}

type CoinNetwork string

const (
	CoinNetworkETH = CoinNetwork("eth")
	CoinNetworkBSC = CoinNetwork("bsc")
	CoinNetworkARB = CoinNetwork("arb")
	CoinNetworkTRX = CoinNetwork("trx")
)
