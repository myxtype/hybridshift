package fixedfloat

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const baseURI = "https://fixedfloat.com"

type Client struct {
	apiKey    string
	apiSecret string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (c *Client) signature(data []byte) (string, error) {
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	_, err := mac.Write(data)
	if err != nil {
		return "", err
	}

	sig := hex.EncodeToString(mac.Sum(nil))
	return strings.ToLower(sig), nil
}

func (c *Client) request(p string, v interface{}) (*response, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseURI+p, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	sig, err := c.signature(data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-KEY", c.apiKey)
	req.Header.Add("X-API-SIGN", sig)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var r response
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
