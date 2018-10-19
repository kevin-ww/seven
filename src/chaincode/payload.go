package main

type Payload struct {
	//key  func(i interface{}) string
	TxId string
	TxTs int64
	Memo string
}

//type payload interface {
//	keyGen(p *Payload) string
//}
