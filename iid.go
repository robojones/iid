// Package iid generates small, globally unique IDs.
package iid

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"github.com/pkg/errors"
	"time"
)

const (
	// IidLen is the length of the byte slice used for the Iid type.
	IidLen = 8
	// StrLen is the length of iid string representations in bytes.
	StrLen    = 11
	// randIndex is the index of the first byte which is initialized with random values during the creation of iids.
	// All following bytes are also initialized with random values.
	randIndex = 4
	// encoder is the base64 URL encoding used for the string serialization of iids.
	encoder   = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	// postfix is the character which needs to be appended to iid string
	// representations so they can be parsed by the encoding in Enc.
	postfix   = "="
)

// The encoding used for string serialization and deserialization of iids.
var Enc = base64.NewEncoding(encoder)

// RandReader is used to fill the 31 least significant bits of new iids with random values.
var RandReader = rand.Reader

// Timestamp is used to get the current UNIX time in seconds when generating new iids.
var Timestamp = func() uint32 {
	return uint32(time.Now().Unix())
}

// New generates a completely new Iid containing:
// 1 bit empty,
// 32 bit current timestamp,
// 31 bit cryptographically secure random bits.
func New() Iid {
	b := make([]byte, IidLen)

	// Set the last four bytes to random values.
	if _, err := RandReader.Read(b[randIndex:]); err != nil {
		panic(errors.Wrap(err, "generate new iid"))
	}

	writeTime(b, Timestamp())

	return Iid(b)
}

// FromString imports an existing Iid from its base64url encoded string representation.
func FromString(s string) (Iid, error) {
	b, err := Enc.DecodeString(s + postfix)

	if err != nil || len(b) != IidLen {
		return nil, errors.Wrapf(ErrInvalid, "parse id %s", s)
	}

	return Iid(b), nil
}

// FromUint64 imports an existing Iid from its uint64 representation.
func FromUint64(i uint64) (Iid) {
	b := make([]byte, IidLen)

	binary.BigEndian.PutUint64(b, i)

	return Iid(b)
}

// FromInt imports an existing Iid from its int64 representation.
func FromInt(i int64) (Iid) {
	return FromUint64(uint64(i))
}

// Iid represents time sortable ID which can be exported as a base64url encoded string, int64 or uint64.
type Iid []byte

// String returns an 11 byte long, base64url encoded string representing the Iid.
func (i Iid) String() string {
	return Enc.EncodeToString(i)[:StrLen]
}

// Uint64 returns a uint64 representing the Iid.
func (i Iid) Uint64() uint64 {
	return binary.BigEndian.Uint64(i)
}

// Int returns a int64 representing the Iid.
func (i Iid) Int() int64 {
	return int64(i.Uint64())
}
