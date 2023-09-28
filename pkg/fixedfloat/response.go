package fixedfloat

import "encoding/json"

type response struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func (r *response) DataUnmarshal(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}

type Currency struct {
	Code     string `json:"code"`
	Coin     string `json:"coin"`
	Network  string `json:"network"`
	Name     string `json:"name"`
	Recv     uint8  `json:"recv"`
	Send     uint8  `json:"send"`
	Tag      string `json:"tag"`
	Logo     string `json:"logo"`
	Color    string `json:"color"`
	Priority int    `json:"priority"`
}

type ExchangeRate struct {
	From struct {
		Code      string `json:"code"`
		Network   string `json:"network"`
		Coin      string `json:"coin"`
		Amount    string `json:"amount"`
		Rate      string `json:"rate"`
		Precision int    `json:"precision"`
		Min       string `json:"min"`
		Max       string `json:"max"`
		USD       string `json:"usd"`
		BTC       string `json:"btc"`
	} `json:"from"`
	To struct {
		Code      string `json:"code"`
		Network   string `json:"network"`
		Coin      string `json:"coin"`
		Amount    string `json:"amount"`
		Rate      string `json:"rate"`
		Precision int    `json:"precision"`
		Min       string `json:"min"`
		Max       string `json:"max"`
		USD       string `json:"usd"`
	} `json:"to"`
	Errors []string `json:"errors"`
	Ccies  []struct {
		Code string `json:"code"`
		Recv bool   `json:"recv"`
		Send bool   `json:"send"`
	} `json:"ccies"`
}
