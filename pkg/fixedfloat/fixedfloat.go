package fixedfloat

import (
	"errors"
	"github.com/shopspring/decimal"
)

func (c *Client) GetAllCurrencies() (data []*Currency, err error) {
	rsp, err := c.request("/api/v2/ccies", nil)
	if err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}
	if err := rsp.Unmarshal(&data); err != nil {
		return nil, err
	}
	return
}

type GetPricesParams struct {
	Type      OrderType       `json:"type"`
	FromCcy   string          `json:"fromCcy"`
	ToCcy     string          `json:"toCcy"`
	Direction Direction       `json:"direction"`
	Amount    decimal.Decimal `json:"amount"`
	Ccies     bool            `json:"ccies,omitempty"`
	Usd       bool            `json:"usd,omitempty"`
	Refcode   string          `json:"refcode,omitempty"`
	Afftax    string          `json:"afftax,omitempty"`
}

func (c *Client) GetPrices(params GetPricesParams) (data *ExchangePrice, err error) {
	rsp, err := c.request("/api/v2/price", params)
	if err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}
	if err := rsp.Unmarshal(&data); err != nil {
		return nil, err
	}
	return
}

type CreateOrderParams struct {
	Type      OrderType       `json:"type"`
	FromCcy   string          `json:"fromCcy"`
	ToCcy     string          `json:"toCcy"`
	Direction Direction       `json:"direction"`
	Amount    decimal.Decimal `json:"amount"`
	ToAddress string          `json:"toAddress"`
	Tag       string          `json:"tag,omitempty"`
	Refcode   string          `json:"refcode,omitempty"`
	Afftax    string          `json:"afftax,omitempty"`
}

func (c *Client) CreateOrder(params CreateOrderParams) (data *Order, err error) {
	rsp, err := c.request("/api/v2/create", params)
	if err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}
	if err := rsp.Unmarshal(&data); err != nil {
		return nil, err
	}
	return
}

type GetOrderParams struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (c *Client) GetOrder(params GetOrderParams) (data *Order, err error) {
	rsp, err := c.request("/api/v2/order", params)
	if err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}
	if err := rsp.Unmarshal(&data); err != nil {
		return nil, err
	}
	return
}

type EmergencyParams struct {
	ID      string          `json:"id"`
	Token   string          `json:"token"`
	Choice  EmergencyChoice `json:"choice"`
	Address string          `json:"address,omitempty"`
	Tag     string          `json:"tag,omitempty"`
}

func (c *Client) Emergency(params EmergencyParams) (data *bool, err error) {
	rsp, err := c.request("/api/v2/emergency", params)
	if err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}
	if err := rsp.Unmarshal(&data); err != nil {
		return nil, err
	}
	return
}
