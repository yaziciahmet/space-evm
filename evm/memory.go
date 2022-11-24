package space_evm

import (
	"github.com/holiman/uint256"
)

// EVM Memory is responsible from the volatile
// byte addressed memory of the current execution
type Memory []byte

func NewMemory() *Memory {
	return &Memory{}
}

// Store a byte in memory starting at offset
func (m *Memory) store1(offset uint64, val byte) {
	m.extend(offset + 1)
	(*m)[offset] = val
}

// Store 32 bytes in memory starting from offset, left-padded with zeroes
func (m *Memory) store32(offset uint64, val *uint256.Int) {
	m.extend(offset + 32)
	buff := val.Bytes32()
	copy((*m)[offset:offset+32], buff[:])
}

// Extend memory with zeroes to the given size's closest multiple of 32
func (m *Memory) extend(size uint64) {
	newSize := ceil32(size)
	// convert to signed to avoid subtraction overflow
	sizeDiff := int64(newSize) - int64(m.ByteLen())
	if sizeDiff > 0 {
		*m = append(*m, make([]byte, sizeDiff)...)
	}
}

func (m *Memory) WordLen() int {
	return m.ByteLen() / 32
}

func (m *Memory) ByteLen() int {
	return len(*m)
}
