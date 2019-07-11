package iid

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteTime(t *testing.T) {
	var time uint32 = 1

	for time > 0 {
		buf := make([]byte, 8)
		writeTime(buf, time)

		i := binary.BigEndian.Uint64(buf)
		assert.Equal(t, uint64(time), i >> (32 - dateOffset))

		time = time << 1
	}

}
