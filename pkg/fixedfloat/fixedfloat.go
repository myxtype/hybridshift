package fixedfloat

import "errors"

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

type GetPricesRequest struct {
	Type      OrderType `json:"type"` // float/fixed
	FromCcy   string    `json:"fromCcy"`
	ToCcy     string    `json:"toCcy"`
	Direction Direction `json:"direction"`
	Amount    string    `json:"amount"`
	Ccies     *bool     `json:"ccies,omitempty"`
	Usd       *bool     `json:"usd,omitempty"`
	Refcode   *string   `json:"refcode,omitempty"`
	Afftax    *string   `json:"afftax,omitempty"`
}

func (c *Client) GetPrices(req *GetPricesRequest) (data *ExchangePrice, err error) {
	rsp, err := c.request("/api/v2/price", req)
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
