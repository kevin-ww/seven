package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	testcc "github.com/s7techlab/cckit/testing"
	"testing"
)

func TestInit(t *testing.T) {
	tknCCStub := testcc.NewMockStub(`token`, NewCC())
	res := tknCCStub.Init(`kevin init the chaincode`)
	fmt.Printf("%v", res)
}

func TestGet(t *testing.T) {
	tknCCStub := testcc.NewMockStub(`token`, NewCC())
	token := &Token{
		Symbol: "kvntoken",
	}
	mockStrPayload := string(Marshal(*token))
	res := tknCCStub.Invoke(`get`, mockStrPayload)
	fmt.Printf("%v\n", res)

}

func TestCreate(t *testing.T) {
	tknCCStub := testcc.NewMockStub(`token`, NewCC())
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
	mockStrPayload := string(Marshal(*token))
	//res := tknCCStub.Invoke(`get`, mockStrPayload)
	res := tknCCStub.Invoke(`create`, mockStrPayload)
	fmt.Printf("%v %v %v\n", res.Status, res.Message, string(res.Payload))
}

func TestCreateThenGet(t *testing.T) {

	tknCCStub := testcc.NewMockStub(`token`, NewCC())
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
	mockStrPayload := string(Marshal(*token))
	//res := tknCCStub.Invoke(`get`, mockStrPayload)
	res := tknCCStub.Invoke(`create`, mockStrPayload)
	fmt.Printf("%v %v %v\n", res.Status, res.Message, string(res.Payload))

	q := &Token{
		Symbol: "kvntoken",
	}
	get := string(Marshal(*q))
	res = tknCCStub.Invoke(`get`, get)
	fmt.Printf("%v %v\n", res, string(res.Payload))

}

func TestEncodingResponse(t *testing.T) {

}

func enc(data interface{}, err error) peer.Response {
	if err != nil {
		return shim.Error(err.Error())
	}
	bytes, _ := json.Marshal(data)
	return shim.Success(bytes)
}
