package iid

import (
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const fixedTime uint32 = 1562658708

type mockReader struct {}
func (*mockReader) Read(b []byte) (n int, e error) {
	return len(b), nil
}

func TestMain(m *testing.M) {
	// Mock all inputs.
	RandReader = &mockReader{}
	Timestamp = func () uint32 {
		return fixedTime
	}

	os.Exit(m.Run())
}


func TestEncoding(t *testing.T) {
	// Verify that the encoding is sortable.
	for i := 1; i < len(encoder); i++ {
		assert.True(t, encoder[i - 1] < encoder[i])
	}
}

func TestNew(t *testing.T) {
	id := New()
	idTime := binary.BigEndian.Uint32(id)
	assert.True(t, idTime == (fixedTime >> 1))
}

func ExampleNew() {
	id := New()
	fmt.Println(id)

	// Output: Ad7YmV-----
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
