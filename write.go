package iid

import (
	"github.com/pkg/errors"
)

// dateOffset is the number of bits by which the date is offset to the right.
const dateOffset = 1

// writeTime writes the provided t into b with an offset of 1 to the right.
func writeTime(b []byte, t uint32) {
	if len(b) < 5 {
		panic(errors.New("length of b is less than 5"))
	}

	b[0] = byte(t >> (24 + dateOffset))
	b[1] = byte(t >> (16 + dateOffset))
	b[2] = byte(t >> (8 + dateOffset))
	b[3] = byte(t >> dateOffset)

	// set the msb of the fourth byte to the value of the lsb of t
	lsb := byte(t << (8 - dateOffset))
	b[4] = b[4] << dateOffset >> dateOffset
	b[4] = b[4] | lsb
}
