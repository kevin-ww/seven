package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"strings"
)

//type ledger interface {
//	put(k string, v interface{}) error
//	get(k string, target interface{}) (interface{}, error)
//	exists(k string) (bool, error)
//	ledgerKey(k string) string
//}

type ledgerStub struct {
	admin  string
	bucket string
	stub   shim.ChaincodeStubInterface
}

type ledgerState interface {
	New(s *ledgerStub) interface{}
}

func (l *ledgerStub) ledgerKey(k string) string {
	return strings.Join([]string{k, l.bucket, l.admin}, "|")
}

func (l *ledgerStub) put(k string, v interface{}) error {
	bytes, _ := json.Marshal(v)
	return l.stub.PutState(l.ledgerKey(k), bytes)
}

func (l *ledgerStub) get(k string, target interface{}) (interface{}, error) {
	bytes, e := l.stub.GetState(l.ledgerKey(k))
	if e != nil {
		return nil, e
	}
	if bytes == nil {
		return nil, errors.New(`no such record in ledger`)
	}
	e = json.Unmarshal(bytes, target)
	return target, e
}

func (l *ledgerStub) exists(k string) (bool, error) {
	bytes, e := l.stub.GetState(l.ledgerKey(k))
	if e != nil || bytes == nil {
		return false, e
	}

	return true, nil
}
