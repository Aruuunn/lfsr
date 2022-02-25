package lfsr

import (
	"fmt"
	"strconv"
	"strings"
)

type LFSR struct {
	state, tap, n int
}

func parseBinaryString(s string) (int, error) {
	parsed := 0

	for i := range s {
		if s[i] != byte('0') && s[i] != byte('1') {
			return 0, fmt.Errorf("parse: unexpected '%v'", string(s[i]))
		}

		parsed *= 2

		if s[i] == byte('1') {
			parsed += 1
		}
	}

	return parsed, nil
}

// NewLFSR creates an LFSR with the specified seed and tap
// seed must be a valid binary string made with only zeros and ones.
func NewLFSR(seed string, tap int) (*LFSR, error) {
	state, err := parseBinaryString(seed)

	if err != nil {
		return nil, err
	}

	lfsr := &LFSR{
		state: int(state),
		tap:   tap,
		n:     len(seed),
	}

	return lfsr, nil
}

// Length returns the number of bits in the LFSR.
func (lfsr *LFSR) Length() int {
	return lfsr.n
}

// BitAt returns bit i as 0 or 1.
func (lfsr *LFSR) BitAt(i int) int {
	if ((1 << (i - 1)) & lfsr.state) != 0 {
		return 1
	}

	return 0
}

// Step simulates one step of this LFSR and return the new bit as 0 or 1
func (lfsr *LFSR) Step() int {
	shiftBit := lfsr.BitAt(lfsr.n) ^ lfsr.BitAt(lfsr.tap)
	lfsr.state = (lfsr.state << 1)

	if lfsr.BitAt(lfsr.n+1) == 1 {
		lfsr.state ^= (1 << (lfsr.n))
	}

	lfsr.state |= shiftBit

	return shiftBit
}

// Generate simulates k steps of this LFSR and return the k bits as a k-bit integer
func (lfsr *LFSR) Generate(k int) int {
	result := 0

	for i := 0; i < k; i++ {
		result *= 2
		result += lfsr.Step()
	}

	return result
}

// String returns a string representation of this LFSR
func (lfsr *LFSR) String() string {
	binStr := strconv.FormatInt(int64(lfsr.state), 2)
	return fmt.Sprintf("LFSR(%s%v)", strings.Repeat("0", lfsr.n-len(binStr)), binStr)
}
