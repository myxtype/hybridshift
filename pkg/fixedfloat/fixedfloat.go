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

	if err := rsp.DataUnmarshal(&data); err != nil {
		return nil, err
	}

	return
}
