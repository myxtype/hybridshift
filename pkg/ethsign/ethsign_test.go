package ethsign

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 0xE28fe4eb29400346f37c6B60a9b33372C3147d7C
func TestSign(t *testing.T) {
	a := assert.New(t)

	privateKey, err := crypto.HexToECDSA("e32a9758b6322dee519103b677bc9665a2d5533dee74af584ec6c58e91d057c1")
	a.Nil(err)

	sig, err := Sign(privateKey, []byte(`User withdrawal
User: 0x9c2e491034EC23D1a0ac724B78e0635A5Ef3E10f
Amount: 1
Timestamp: 1690722664`))
	a.Nil(err)

	t.Log(hexutil.Encode(sig))

	account := crypto.PubkeyToAddress(privateKey.PublicKey)
	//account2 := common.HexToAddress("0x11f4d0A3c12e86B4b5F39B213F7E19D048276DAe")

	ok, err := VerifySign(sig, account, []byte(`User withdrawal
User: 0x9c2e491034EC23D1a0ac724B78e0635A5Ef3E10f
Amount: 1
Timestamp: 1690722664`))
	a.Nil(err)

	t.Log(ok)
}
