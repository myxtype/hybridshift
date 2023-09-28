package fixedfloat

import (
	"encoding/json"
	"github.com/shopspring/decimal"
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

	data, err := cl.GetPrices(GetPricesParams{
		Type:      "float",
		FromCcy:   "USDTTRC",
		ToCcy:     "BTC",
		Direction: "to",
		Amount:    decimal.NewFromInt(10),
		Ccies:     true,
	})
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(data)
	t.Log(string(b))
}

func TestClient_CreateOrder(t *testing.T) {
	cl := NewClient("MlwZ8yPHY2j34Vzuv9FiMWZdu0e57svhqxKDjCvY", "eyAcOG0KXiFuCyyKQ0W8EhSNYBTtGTILwSUTM2ra")

	data, err := cl.CreateOrder(CreateOrderParams{
		Type:      OrderTypeFloat,
		FromCcy:   "USDTBSC",
		ToCcy:     "USDTARBITRUM",
		Direction: DirectionTo,
		Amount:    decimal.NewFromInt(100),
		ToAddress: "0x8566e819D540F31F5009E9fD974B4B05c1e5d77A",
	})
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(data)
	t.Log(string(b))
}

func TestClient_GetOrder(t *testing.T) {
	cl := NewClient("MlwZ8yPHY2j34Vzuv9FiMWZdu0e57svhqxKDjCvY", "eyAcOG0KXiFuCyyKQ0W8EhSNYBTtGTILwSUTM2ra")

	data, err := cl.GetOrder(GetOrderParams{
		ID:    "F6HGHC",
		Token: "Y2xIOniviwF9M7fVKVYR3KZ5yzHsZo0ULyYrLLzW",
	})
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(data)
	t.Log(string(b))
}
