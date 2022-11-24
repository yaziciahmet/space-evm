package space_evm

import (
	"encoding/hex"
	"strings"

	"github.com/holiman/uint256"
	"golang.org/x/crypto/sha3"
)

// Convert byte slice to *uint256.Int, left-padded with zeroes
func byteSliceToUint256(buff []byte) (*uint256.Int, error) {
	hexStr := hex.EncodeToString(buff)
	// uint256.FromHex requires no left-padded zeroes
	hexStr = strings.TrimLeft(hexStr, "0")
	if hexStr == "" {
		hexStr = "0"
	}
	return uint256.FromHex("0x" + hexStr)
}

func keccak256(buff []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(buff)
	return hash.Sum(nil)
}

// Ceil value to closest multiple of 32
func ceil32(val uint64) uint64 {
	r := val % 32
	if r == 0 {
		return val
	}
	return val - r + 32
}
