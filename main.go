package main

import "C"

import (
	"github.com/Ungigdu/contract_go_clib_python/rpc_caller"
)

//export Query
func Query(key string) *C.char {
	return C.CString(rpc_caller.QueryMap(key))
}

//export Update
func Update(privateKey, mapKey, value string) {
	rpc_caller.InsertMap(privateKey,mapKey,value)
}

//export Watch
func Watch()  {
	rpc_caller.WatchUpdate()
}

func main(){}