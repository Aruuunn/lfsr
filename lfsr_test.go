package lfsr

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	lfsr, err := NewLFSR("01101000010", 9)
	assert.Nil(t, err)

	assert.Equal(t, lfsr.Length(), 11)

	output := ""

	for i := 0; i < 10; i++ {
		bit := lfsr.Step()
		output += strconv.Itoa(bit)
	}

	assert.Equal(t, "1100100100", output)
}

func TestGenerate(t *testing.T) {
	lfsr, err := NewLFSR("01101000010", 9)
	assert.Nil(t, err)

	expected := []int{25, 4, 30, 27, 18, 26, 28, 24, 23, 29}

	for i := 0; i < 10; i++ {
		r := lfsr.Generate(5)
		assert.Equal(t, expected[i], r)
	}
}
