package main

type Token struct {
	Payload
	Symbol      string
	Name        string
	TotalSupply uint64
	Decimals    uint8
}

func keyGen(t *Token) string {
	return t.Symbol
}

type TokenLedgerImpl struct {
	*ledgerStub
}

func (impl *TokenLedgerImpl) New(s *ledgerStub) interface{}{
	impl.ledgerStub=s
	return impl
}

//type nop int
//
//func (n *nop) New(s *ledgerStub) interface{} {
//	return &TokenLedgerImpl{
//		s,
//	}
//}

//type ledgerState interface {
//	New(s *ledgerStub) interface{}
//}

func (impl *TokenLedgerImpl) create(token *Token) (*Token, error) {
	return token, impl.ledgerStub.put(keyGen(token), token)
}

func (impl *TokenLedgerImpl) has(token *Token) (bool, error) {
	return impl.ledgerStub.exists(keyGen(token))
}

func (impl *TokenLedgerImpl) update(token *Token) (*Token, error) {
	return impl.create(token)
}

func (impl *TokenLedgerImpl) get(token *Token) (*Token, error) {
	res, e := impl.ledgerStub.get(keyGen(token), &Token{})
	if e != nil {
		return nil, e
	}
	return res.(*Token), nil
}
