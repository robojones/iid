package iid

import (
	"github.com/pkg/errors"
)

// Writes the provided t into the first 33 bit of b
func writeTime(b []byte, t uint32) {
	if len(b) < 5 {
		panic(errors.New("length of b is less than 5"))
	}

	b[0] = byte(t >> 25)
	b[1] = byte(t >> 17)
	b[2] = byte(t >> 9)
	b[3] = byte(t >> 1)
	// unset the first bit
	b[4] = b[4] << 1 >> 1
	// set the first bit to the value of the last bit of t
	b[4] = b[4] | byte(t << 7)
}
