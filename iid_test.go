package iid

import (
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
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

func ExampleNew() {
	id := New()
	fmt.Println(id)
}

func TestFromString(t *testing.T) {
	ex := New()
	s := enc.EncodeToString(ex)[:strLen]

	id, err := FromString(s)

	assert.Nil(t, err)
	assert.Equal(t, ex, id)
}

func ExampleFromString() {
	str := "MHDSedbNhXB"
	s, err := FromString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.String())
	// Output: MHDSedbNhXB
}

func TestFromUint64(t *testing.T) {
	ex := New()
	i := binary.BigEndian.Uint64(ex)

	id := FromUint64(i)

	assert.Equal(t, ex, id)
}

func ExampleFromUint64() {
	var i uint64 = 6711382541547442289
	id := FromUint64(i)

	fmt.Println(id.Uint64())
	// Output: 6711382541547442289
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
