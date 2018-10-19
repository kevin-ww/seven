package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"reflect"
)

func test(stub shim.ChaincodeStubInterface, ccName string, target interface{}, impl interface{}) peer.Response {

	ledgerImpl := impl.(ledgerState).New(&ledgerStub{
		admin:  "admin",
		bucket: ccName,
		stub:   stub,
	})

	funcName, args := stub.GetFunctionAndParameters()

	//ref
	val, e := UnMarshal([]byte(args[0]), target)

	if e != nil{
		fmt.Printf("%v\n",e)
	}

	m := reflect.ValueOf(ledgerImpl).MethodByName(funcName)

	if m.Type() == nil{
		log.Fatal(`no such method :`+funcName)
	}


	//

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(val)

	response := m.Call(inputs)

	return shim.Error(``)
}
