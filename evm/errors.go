package space_evm

import (
	"errors"
	"fmt"
)

var (
	ErrStackOverflow   = errors.New("stack overflow")
	ErrStackUnderflow  = errors.New("stack underflow")
	ErrGasUintOverflow = errors.New("gas uint64 overflow")
	ErrOutOfGas        = errors.New("out of gas")
)

func ErrInvalidOpcode(opcode byte) error {
	return fmt.Errorf("invalid opcode %02d", opcode)
}

func ErrNotEnoughBytesToRead(opcode byte) error {
	return fmt.Errorf("not enough bytes to read for opcode %02d", opcode)
}
