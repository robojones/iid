package iid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	var b uint8 = 0xFF

	buf := make([]byte, 5)
	writeTime(buf, 0xFFFFFFFF)

	assert.Equal(t, b >> 1, buf[0])
	assert.Equal(t, b, buf[1])
	assert.Equal(t, b, buf[2])
	assert.Equal(t, b, buf[3])
	assert.Equal(t, b << 7, buf[4])
}
