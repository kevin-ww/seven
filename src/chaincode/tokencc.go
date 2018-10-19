package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

type ChainCode struct {
	name    string
	handler func(fn string, arg string) peer.Response
}

const CCName = `AcctChainCode`

func (c *ChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Initial ..." + CCName))
}

func (c *ChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return process(stub)
}

func process(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, args := stub.GetFunctionAndParameters()

	if args == nil || len(args) == 0 {
		fmt.Printf("require at least one arg as chaincode function payload \n")
	}

	fmt.Printf("invoking %v with args: %v \n", funcName, args[0])

	//var token *Token
	t, e := UnMarshal([]byte(args[0]), &Token{})

	token := t.(*Token)

	if e != nil {
		EncodeResponse(nil, e)
	}

	tokenState := &TokenImpl{
		ledgerDB: ledgerDB{
			admin:  "",
			bucket: "",
			stub:   stub,
		},
	}

	if funcName == `get` {
		return EncodeResponse(tokenState.get(token))
	} else if funcName == `create` {
		return EncodeResponse(tokenState.create(token))
	} else if funcName == `has` {
		return EncodeResponse(tokenState.has(token))
	} else if funcName == `update` {
		return EncodeResponse(tokenState.update(token))
	}

	return EncodeResponse(nil, errors.New(funcName+` not provided in `+CCName))
}

func NewCC() *ChainCode {
	return &ChainCode{
		name:    CCName,
		handler: nil,
	}
}

func main() {

	if err := shim.Start(NewCC()); err != nil {
		fmt.Printf("Error starting %s: %s", CCName, err)
	}
}


