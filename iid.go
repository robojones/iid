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
	iidLen    = 8
	strLen    = 11
	randIndex = 4
	encoder   = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	postfix   = "="
)

var enc = base64.NewEncoding(encoder)

var RandReader = rand.Reader

var Timestamp = func() uint32 {
	return uint32(time.Now().Unix())
}

// New generates a completely new Iid containing the current timestamp and four cryptographically secure random bytes.
func New() Iid {
	b := make([]byte, iidLen)

	// Set the last four bytes to random values.
	if _, err := RandReader.Read(b[randIndex:]); err != nil {
		panic(errors.Wrap(err, "generate new iid"))
	}

	// Write current UNIX time in ms to the first four bytes.
	writeTime(b, Timestamp())

	return Iid(b)
}

// FromString imports an existing Iid from its base64url encoded string representation.
func FromString(s string) (Iid, error) {
	b, err := enc.DecodeString(s + postfix)

	if err != nil || len(b) != iidLen {
		return nil, errors.Wrapf(ErrInvalid, "parse id %s", s)
	}

	return Iid(b), nil
}

// FromUint64 imports an existing Iid from its uint64 representation.
func FromUint64(i uint64) (Iid) {
	b := make([]byte, iidLen)

	binary.BigEndian.PutUint64(b, i)

	return Iid(b)
}

// FromInt imports an existing Iid from its int64 representation.
func FromInt(i int64) (Iid) {
	return FromUint64(uint64(i))
}

// Iid represents time sortable ID which can be exported as a base64url encoded string or uint64.
type Iid []byte

// String returns an 11 character/byte long string representing the Iid.
func (i Iid) String() string {
	return enc.EncodeToString(i)[:strLen]
}

// Uint64 returns a uint64 representing the Iid.
func (i Iid) Uint64() uint64 {
	return binary.BigEndian.Uint64(i)
}

// Int returns a int64 representing the Iid.
func (i Iid) Int() int64 {
	return int64(i.Uint64())
}
