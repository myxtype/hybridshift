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
