package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

// MinerTransaction inherits Transaction
type MinerTransaction struct {
	*Transaction
	Nonce uint32 // Random number to avoid hash collision
}

//// NewMinerTransaction creates an MinerTransaction
//func NewMinerTransaction(script []byte) *MinerTransaction {
//	mtx := &MinerTransaction{
//		Transaction:NewTransaction(),
//		Nonce:rand.Uint32(),
//	}
//	mtx.Type = Issue_Transaction
//	return mtx
//}

// HashString returns the transaction Id string
func (mtx *MinerTransaction) HashString() string {
	hash := crypto.Hash256(mtx.UnsignedRawTransaction())
	mtx.Hash, _ = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (mtx *MinerTransaction) UnsignedRawTransaction() []byte {
	buf := io.NewBufBinWriter()
	mtx.SerializeUnsigned(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (mtx *MinerTransaction) RawTransaction() []byte {
	buf := io.NewBufBinWriter()
	mtx.Serialize(buf.BinWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (mtx *MinerTransaction) RawTransactionString() string {
	return hex.EncodeToString(mtx.RawTransaction())
}

// FromHexString parses a hex string
func (mtx *MinerTransaction) FromHexString(rawTx string) (*MinerTransaction, error) {
	b, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, err
	}
	br := io.NewBinReaderFromBuf(b)
	mtx.Deserialize(br)
	if br.Err != nil {
		return nil, br.Err
	}
	return mtx, nil
}

// Deserialize implements Serializable interface.
func (mtx *MinerTransaction) Deserialize(br *io.BinReader) {
	mtx.DeserializeUnsigned(br)
	mtx.Transaction.DeserializeWitnesses(br)
}

func (mtx *MinerTransaction) DeserializeUnsigned(br *io.BinReader) {
	mtx.Transaction.DeserializeUnsigned1(br)
	mtx.DeserializeExclusiveData(br)
	mtx.Transaction.DeserializeUnsigned2(br)
}

func (mtx *MinerTransaction) DeserializeExclusiveData(br *io.BinReader) {
	br.ReadLE(&mtx.Nonce)
}

// Serialize implements Serializable interface.
func (mtx *MinerTransaction) Serialize(bw *io.BinWriter) {
	mtx.SerializeUnsigned(bw)
	mtx.SerializeWitnesses(bw)
}

func (mtx *MinerTransaction) SerializeUnsigned(bw *io.BinWriter)  {
	mtx.Transaction.SerializeUnsigned1(bw)
	mtx.SerializeExclusiveData(bw)
	mtx.SerializeUnsigned2(bw)
}

func (mtx *MinerTransaction) SerializeExclusiveData(bw *io.BinWriter)  {
	bw.WriteLE(mtx.Nonce)
}

