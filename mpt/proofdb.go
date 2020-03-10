package mpt

import (
	"errors"

	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
)

//ProofDb a db to use for verify
type ProofDb struct {
	nodes map[string]([]byte)
}

//NewProofDb new instance of ProofDb from a string list
func NewProofDb(proof []string) *ProofDb {
	p := &ProofDb{}
	p.nodes = make(map[string]([]byte), len(proof))
	for _, v := range proof {
		data := helper.HexTobytes(v)
		hashstr := helper.BytesToHex(crypto.Hash256(data))
		p.nodes[hashstr] = data
	}
	return p
}

//Get for TrieDb
func (pd *ProofDb) Get(key []byte) (value []byte, err error) {
	keystr := helper.BytesToHex(key)
	if v, ok := pd.nodes[keystr]; ok {
		return v, nil
	}
	return nil, errors.New("cant find the value in ProofDb")
}
