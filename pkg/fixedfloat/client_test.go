package fixedfloat

import (
	"encoding/json"
	"testing"
)

func TestClient_GetAllCurrencies(t *testing.T) {
	cl := NewClient("MlwZ8yPHY2j34Vzuv9FiMWZdu0e57svhqxKDjCvY", "eyAcOG0KXiFuCyyKQ0W8EhSNYBTtGTILwSUTM2ra")

	data, err := cl.GetAllCurrencies()
	if err != nil {
		t.Fatal(err)
	}

	for _, n := range data {
		b, _ := json.Marshal(n)
		t.Log(string(b))
	}
}

func TestClient_GetPrices(t *testing.T) {
	cl := NewClient("MlwZ8yPHY2j34Vzuv9FiMWZdu0e57svhqxKDjCvY", "eyAcOG0KXiFuCyyKQ0W8EhSNYBTtGTILwSUTM2ra")

	data, err := cl.GetPrices(&GetPricesRequest{
		Type:      "float",
		FromCcy:   "USDTTRC",
		ToCcy:     "BTC",
		Direction: "to",
		Amount:    "10",
	})
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(data)
	t.Log(string(b))
}
