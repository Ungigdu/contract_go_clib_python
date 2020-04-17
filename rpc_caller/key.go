package rpc_caller

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverFromPrivateKeyByes(d []byte)  *ecdsa.PrivateKey{
	privateKey,err := crypto.ToECDSA(d)
	if err == nil {
		return privateKey
	}else{
		logger.Error("can't recover private key : ",err)
		return nil
	}
}

func RecoverFromPrivateKeyString(p string) *ecdsa.PrivateKey{
	hex,err := hexutil.Decode(p)
	if err == nil {
		return RecoverFromPrivateKeyByes(hex)
	} else {
		logger.Error("wrong format : " ,err)
		return nil
	}
}