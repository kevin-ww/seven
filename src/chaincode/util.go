package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//
//func EncodeResponse(data interface{}, err error) peer.Response {
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//	bytes, _ := json.Marshal(data)
//	return shim.Success(bytes)
//}

func UnMarshal(data []byte, target interface{}) (interface{}, error) {
	e := json.Unmarshal(data, target)
	if e != nil {
		return nil, e
	}
	return target, nil
}

func Marshal(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}


func EncodeResponse(data interface{}, err error) peer.Response {
	if err != nil {
		return shim.Error(err.Error())
	}

	bytes, _ := json.Marshal(data)

	return shim.Success(bytes)
}