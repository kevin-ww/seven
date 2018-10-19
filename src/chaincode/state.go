package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
	"strings"
)

type ledger interface {
	put(k string, v interface{}) error
	get(k string, target interface{}) (interface{}, error)
	has(k string) (bool, error)
}

type ledgerDB struct {
	admin  string
	bucket string
	stub   shim.ChaincodeStubInterface
	//stub *shim.MockStub
}

func (l *ledgerDB) ledgerKey(k string) string {
	return strings.Join([]string{k, l.bucket, l.admin}, "|")
}

func (l *ledgerDB) put(k string, v interface{}) error {
	bytes, _ := json.Marshal(v)
	return l.stub.PutState(l.ledgerKey(k), bytes)
}

func (l *ledgerDB) get(k string, target interface{}) (interface{}, error) {
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

func (l *ledgerDB) exists(k string) (bool, error) {
	bytes, e := l.stub.GetState(l.ledgerKey(k))
	if e != nil || bytes == nil {
		return false, e
	}

	return true, nil
}
