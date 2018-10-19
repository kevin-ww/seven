package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"testing"
)

func TestCast(t *testing.T) {

	token := &Token{
		Payload: Payload{
			TxId: "txid001",
			TxTs: 0,
			Memo: "a test payload",
		},
		Symbol:      "kvntoken",
		Name:        "token of kevin",
		TotalSupply: 110,
		Decimals:    10,
	}

	data := Marshal(token)

	//
	impl := &TokenLedgerImpl{}
	target := &Token{}

	core(nil, impl, target, data, `create`)

}

func core(stub shim.ChaincodeStubInterface, impl interface{}, target interface{}, data []byte, invoke string) {

	//


	ledgerImpl := impl.(ledgerState).New(&ledgerStub{
		admin:  "",
		bucket: "",
		stub:   stub,
	})

	//ref
	val, e := UnMarshal(data, target)
	if e != nil{
		fmt.Printf("%v\n",e)
	}

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(val)
	reflect.ValueOf(ledgerImpl).MethodByName(invoke).Call(inputs)



	fmt.Printf("%v \n",ledgerImpl)

	fmt.Printf("%v  \n", impl)
	fmt.Printf("%v  \n", target)

	fmt.Printf("%v \n", string(data))

	fmt.Printf("%v  \n", invoke)

}
