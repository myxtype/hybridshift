package ethsign

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func EcRecover(data hexutil.Bytes, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}

	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	hash := accounts.TextHash(data)
	rpk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}

func Sign(pri *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	hash := accounts.TextHash(data)
	sig, err := crypto.Sign(hash, pri)
	if err != nil {
		return nil, err
	}

	sig[crypto.RecoveryIDOffset] += 27
	return sig, nil
}

func VerifySign(sig []byte, account common.Address, data []byte) (bool, error) {
	rcAccount, err := EcRecover(data, sig)
	if err != nil {
		return false, err
	}
	return rcAccount.String() == account.String(), nil
}
