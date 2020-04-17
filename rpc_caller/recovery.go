package rpc_caller

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/Ungigdu/contract_go_clib_python/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"sync"
)


var logger, _ = logging.GetLogger("rpc_caller")

var AccessPoints = []string{
	"wss://ropsten.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
	"wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
	"ws://75.135.96.248:3334",
}
var conn = Bas_Ethereum.NewConn()

var DemoAddress = "0x2A2A73aDBc72Cc4203568b3601F70dcf7bCA1DbF"

var demo *contract.SimpleDemo

type Conn struct {
	Client *ethclient.Client
	lock *sync.Mutex
}



func NewConn() *Conn {
	return &Conn{
		Client: nil,
		lock:   &sync.Mutex{},
	}
}

func (conn *Conn) GetClient() *ethclient.Client {
	conn.lock.Lock()
	defer conn.lock.Unlock()
	if conn.Client!=nil{
		return conn.Client
	}else{
		for _,s := range AccessPoints{
			c, err := ethclient.Dial(s)
			if err!=nil{
				logger.Error("can't get access through : " ,s)
				continue
			}else{
				conn.Client = c
				logger.Info("conn is ready")
				return conn.Client
			}
		}
	}
	logger.Fatal("can't get access to ethereum through any points")
	return nil
}

func CloseConn(){
	conn.Client.Close()
	conn.Client = nil
}

func (conn *Conn) Reset(){
	CloseConn()
	conn.GetClient()
}

func Demo() *contract.SimpleDemo {
	if demo == nil {
		if d,err := contract.NewSimpleDemo(common.HexToAddress(DemoAddress),conn.GetClient());err == nil {
			demo = d
		}else{
			logger.Fatal("can't recover demo", err)
		}
	}
	return demo
}
