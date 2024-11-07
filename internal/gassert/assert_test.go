package gassert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	assert.NotPanics(t, func() {
		MustBePositive(1)
	})

	assert.Panics(t, func() {
		MustBePositive(-1)
	})
}
