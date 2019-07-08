package iid

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsErrInvalid(t *testing.T) {
	assert.True(t, IsErrInvalid(ErrInvalid))
	assert.True(t, IsErrInvalid(errors.Wrap(ErrInvalid, "wow")))
	assert.False(t, IsErrInvalid(nil))
	testError := errors.New("test")
	assert.False(t, IsErrInvalid(testError))
}
