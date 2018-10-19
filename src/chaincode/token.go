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

type TokenImpl struct {
	ledgerDB
}

func NewImpl(l ledgerDB) *TokenImpl{
	return &TokenImpl{
		l,
	}
}

func (tI *TokenImpl) create(t *Token) (*Token, error) {
	return t, tI.ledgerDB.put(keyGen(t), t)
}

func (tI *TokenImpl) has(t *Token) (bool, error) {
	return tI.ledgerDB.exists(keyGen(t))
}

func (tI *TokenImpl) update(t *Token) (*Token, error) {
	return tI.create(t)
}

func (tI *TokenImpl) get(t *Token) (*Token, error) {
	r, e := tI.ledgerDB.get(keyGen(t), &Token{})
	if e != nil {
		return nil, e
	}
	return r.(*Token), nil
}
