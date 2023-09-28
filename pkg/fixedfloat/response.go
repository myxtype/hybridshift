package fixedfloat

import "encoding/json"

type response struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func (r *response) Unmarshal(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}

// Currency
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

// ExchangePrice
type ExchangePrice struct {
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

// Order
type Order struct {
	ID     string      `json:"id"`
	Type   OrderType   `json:"type"`
	Email  string      `json:"email"`
	Status OrderStatus `json:"status"`
	Time   struct {
		Reg        int64 `json:"reg"`
		Start      int64 `json:"start"`
		Finish     int64 `json:"finish"`
		Update     int64 `json:"update"`
		Expiration int64 `json:"expiration"`
		Left       int64 `json:"left"`
	} `json:"time"`
	From struct {
		Code             string `json:"code"`
		Coin             string `json:"coin"`
		Network          string `json:"network"`
		Name             string `json:"name"`
		Alias            string `json:"alias"`
		Amount           string `json:"amount"`
		Address          string `json:"address"`
		AddressAlt       string `json:"addressAlt"`
		Tag              string `json:"tag"`
		TagName          string `json:"tagName"`
		ReqConfirmations int64  `json:"reqConfirmations"`
		MaxConfirmations int64  `json:"maxConfirmations"`
		Tx               struct {
			ID            string `json:"id"`
			Amount        string `json:"amount"`
			Fee           string `json:"fee"`
			CcyFee        string `json:"ccyfee"`
			TimeReg       int64  `json:"timeReg"`
			TimeBlock     int64  `json:"timeBlock"`
			Confirmations int64  `json:"confirmations"`
		} `json:"tx"`
	} `json:"from"`
	To struct {
		Code       string `json:"code"`
		Coin       string `json:"coin"`
		Network    string `json:"network"`
		Name       string `json:"name"`
		Alias      string `json:"alias"`
		Amount     string `json:"amount"`
		Address    string `json:"address"`
		AddressAlt string `json:"addressAlt"`
		Tag        string `json:"tag"`
		TagName    string `json:"tagName"`
		Tx         struct {
			ID            string `json:"id"`
			Amount        string `json:"amount"`
			Fee           string `json:"fee"`
			CcyFee        string `json:"ccyfee"`
			TimeReg       int64  `json:"timeReg"`
			TimeBlock     int64  `json:"timeBlock"`
			Confirmations int64  `json:"confirmations"`
		} `json:"tx"`
	} `json:"to"`
	Back struct {
		Code       string `json:"code"`
		Coin       string `json:"coin"`
		Network    string `json:"network"`
		Name       string `json:"name"`
		Alias      string `json:"alias"`
		Amount     string `json:"amount"`
		Address    string `json:"address"`
		AddressAlt string `json:"addressAlt"`
		Tag        string `json:"tag"`
		TagName    string `json:"tagName"`
		Tx         struct {
			ID            string `json:"id"`
			Amount        string `json:"amount"`
			Fee           string `json:"fee"`
			CcyFee        string `json:"ccyfee"`
			TimeReg       int64  `json:"timeReg"`
			TimeBlock     int64  `json:"timeBlock"`
			Confirmations int64  `json:"confirmations"`
		} `json:"tx"`
	} `json:"back"`
	Emergency struct {
		Status []EmergencyStatus `json:"status"`
		Choice EmergencyChoice   `json:"choice"`
		Repeat int64             `json:"repeat"`
	} `json:"emergency"`
	Token string `json:"token"`
}
