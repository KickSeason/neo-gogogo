package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBytes(t *testing.T) {
	var b = make([]byte, 0)
	r := ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1}
	r = ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1, 2}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{2, 1}, r)

	b = []byte{1, 2, 3}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{1, 2, 3}, b)
	assert.Equal(t, []byte{3, 2, 1}, r)
}

//func Test(t *testing.T)  {
//	//var v = int((0x30 - 27) & ^byte(4)) // 0001_0101 & ^ 0000_0100 = 0001_0001 = 17
//	//	//assert.Equal(t, 0xfb, v)
//
//	p := 34 // 0010_0010
//	q := 20 // 0001_0100
//	//
//	assert.Equal(t, 34, p& ^q)
//}

//func TestHashToInt(t *testing.T) {
//	s := "Hello World"
//	encoded := []byte(s);
//	keccak := sha3.NewLegacyKeccak256()
//	keccak.Write(encoded)
//	hash := keccak.Sum(nil)
//
//	bi := HashToInt(hash)
//
//	assert.Equal(t, 0, bi)
//}