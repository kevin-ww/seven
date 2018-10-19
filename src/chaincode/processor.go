package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func test(stub shim.ChaincodeStubInterface, target interface{} , impl interface{}) peer.Response{

	funcName, _ := stub.GetFunctionAndParameters()
	//
	//if args == nil || len(args) == 0 {
	//	fmt.Printf("require at least one arg as chaincode function payload \n")
	//}
	//
	//fmt.Printf("invoking %v with args: %v \n", funcName, args[0])
	//
	////var token *Token
	//t, e := UnMarshal([]byte(args[0]), target)
	//
	////token := t.(Token)
	//
	//if e != nil {
	//	EncodeResponse(nil, e)
	//}
	//
	//
	////impl
	//
	//tokenLedger := &TokenImpl{
	//	ledgerDB: ledgerDB{
	//		admin:  "admin",
	//		bucket: "token-bucket",
	//		stub:   stub,
	//	},
	//}
	//
	//if funcName == `get` {
	//	return EncodeResponse(tokenLedger.get(&token))
	//} else if funcName == `create` {
	//	return EncodeResponse(tokenLedger.create(&token))
	//} else if funcName == `has` {
	//	return EncodeResponse(tokenLedger.has(&token))
	//} else if funcName == `update` {
	//	return EncodeResponse(tokenLedger.update(&token))
	//}

	return EncodeResponse(nil, errors.New(funcName+` not provided in `+CCName))
}


func core(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()

	if args == nil || len(args) == 0 {
		fmt.Printf("require at least one arg as chaincode function payload \n")
	}

	fmt.Printf("invoking %v with args: %v \n", funcName, args[0])

	//var token *Token
	t, e := UnMarshal([]byte(args[0]), &Token{})

	token := t.(Token)

	if e != nil {
		EncodeResponse(nil, e)
	}

	tokenLedger := &TokenImpl{
		ledgerDB: ledgerDB{
			admin:  "admin",
			bucket: "token-bucket",
			stub:   stub,
		},
	}

	if funcName == `get` {
		return EncodeResponse(tokenLedger.get(&token))
	} else if funcName == `create` {
		return EncodeResponse(tokenLedger.create(&token))
	} else if funcName == `has` {
		return EncodeResponse(tokenLedger.has(&token))
	} else if funcName == `update` {
		return EncodeResponse(tokenLedger.update(&token))
	}

	return EncodeResponse(nil, errors.New(funcName+` not provided in `+CCName))
}
