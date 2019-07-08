package iid

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	before := time.Now().Unix()
	id := New()
	after := time.Now().Unix()

	idTime := int64(binary.BigEndian.Uint32(id))

	assert.True(t, idTime >= before)
	assert.True(t, idTime <= after)
}

func TestFromString(t *testing.T) {
	ex := New()
	s := enc.EncodeToString(ex)[:strLen]

	id, err := FromString(s)

	assert.Nil(t, err)
	assert.Equal(t, ex, id)
}

func TestFromUint64(t *testing.T) {
	ex := New()
	i := binary.BigEndian.Uint64(ex)

	id := FromUint64(i)

	assert.Equal(t, ex, id)
}

func TestIid_String(t *testing.T) {
	ex := New()

	str := ex.String()

	id, err := FromString(str)
	assert.Nil(t, err)
	assert.Equal(t, ex, id)
}

func TestIid_Uint64(t *testing.T) {
	ex := New()

	i := ex.Uint64()

	id := FromUint64(i)
	assert.Equal(t, ex, id)
}
