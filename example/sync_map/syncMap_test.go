package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestT(t *testing.T) {
	for i := 0; i < 100; i++ {
		c := T()
		assert.Equal(t, int32(100), c, "do not equal")
	}
}
