package rpc_caller

import "C"
import (
	"context"
	"github.com/Ungigdu/contract_go_clib_python/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func GetOpts(blockNumber uint64) *bind.CallOpts {
	var opts = new(bind.CallOpts)
	if blockNumber == 0 {
		opts = nil
	}else{
		opts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
	return opts
}

func QueryMap(key string) string {
	result,err := Demo().Map(GetOpts(0),key)
	if err != nil {
		logger.Error("query error : ", err)
		return ""
	}else{
		return result
	}
}

func InsertMap(privateKey, mapKey, value string){
	pk := RecoverFromPrivateKeyString(privateKey)
	opts:=bind.NewKeyedTransactor(pk)
	tx,err := Demo().UpdateMap(opts,mapKey,value)
	if err != nil {
		logger.Error("pack transaction error : ", err)
	}
	receipt, err := bind.WaitMined(context.Background(), conn.Client, tx)
	if err != nil {
		logger.Error("mine update map error", err)
	}else{
		logger.Info("commit on block : ", receipt.BlockNumber.String())
	}
}

func WatchUpdate() {
	logs := make(chan *contract.SimpleDemoUpdate)
	sub,err := Demo().WatchUpdate(nil,logs)
	if err != nil {
		logger.Error("watch failed : ",err)
	}
	for {
		select {
			case e := <-sub.Err():
				logger.Error("subscript watch update runtime err",e)
				return
			case log := <-logs:
				logger.Info("detected update, key:", log.Key, "operator : ", log.Operator.String())
		}
	}
}