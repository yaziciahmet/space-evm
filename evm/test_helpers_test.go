// This file contains helper functions and data for testing
package space_evm

import (
	"encoding/hex"
	"fmt"
	"math"
	"reflect"

	"github.com/holiman/uint256"
)

// Constants
var MaxUint64 = uint64(math.MaxUint64)
var MaxUint256, _ = uint256.FromHex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
var EmptyMemHash = keccak256([]byte{})

// genericTest is the struct used to hold the test data
// in an organized format. It also helps generate messages
// based on the expected and actual result comparison.
type genericTest struct {
	s          string
	in         interface{}
	exp        interface{}
	act        interface{}
	shouldFail bool
}

func (t genericTest) Check() (string, bool) {
	if reflect.DeepEqual(t.exp, t.act) {
		return fmt.Sprintf("\t✔ %s\n", t.s), false
	} else {
		return fmt.Sprintf("\033[31m\t✖ %s\n\t\texp: %#v\n\t\tgot: %#v\n\033[39m", t.s, t.exp, t.act), true
	}
}

// stackTestExp is the struct used to hold
// expected results of stack operations
type stackTestExp struct {
	size int
	top  *uint256.Int
}

// memoryStoreTestInput is the struct used to hold
// inputs of the memory store operation
type memoryStoreTestIn struct {
	offset uint64
	val    interface{}
}

// interpreterRunTestIn is the struct used to hold
// inputs of the interpreter run operation
type interpreterRunTestIn struct {
	code     []byte
	gasLimit uint64
}

// interpreterRunTestExp is the struct used to hold
// expected results of the interpreter run operation
type interpreterRunTestExp struct {
	runRes *RunResult
	stack  *Stack
}

// Common helpers
func stackPush(s *Stack, el uint64) (*uint256.Int, error) {
	el256 := uint256.NewInt(el)
	err := s.push(el256)
	return el256, err
}

func populateStack(s *Stack, els ...uint64) error {
	for _, el := range els {
		_, err := stackPush(s, el)
		if err != nil {
			return err
		}
	}
	return nil
}

func stackPopN(s *Stack, n uint64) error {
	for ; n > 0; n-- {
		_, err := s.pop()
		if err != nil {
			return err
		}
	}
	return nil
}

func memStore(m *Memory, testIn []memoryStoreTestIn, is32 bool) {
	for _, in := range testIn {
		if is32 {
			m.store32(in.offset, in.val.(*uint256.Int))
		} else {
			m.store1(in.offset, in.val.(uint8))
		}
	}
}

func genZeroMem(wordSize int) []byte {
	var strMem string
	for ; wordSize > 0; wordSize-- {
		strMem += "0000000000000000000000000000000000000000000000000000000000000000"
	}
	return hexToBytes(strMem)
}

func genRunState(code string, opcode byte, stackValues []uint64, mem []byte) *RunState {
	runSt := NewRunState(hexToBytes(code), MaxUint64)
	runSt.Opcode = opcode
	populateStack(runSt.Stack, stackValues...)
	runSt.Memory = (*Memory)(&mem)
	wordLen := uint64(runSt.Memory.WordLen())
	runSt.HighestMemoryGasCost = 3*wordLen + wordLen*wordLen/512
	return runSt
}

func genU64Slice(size uint64) []uint64 {
	res := make([]uint64, size)
	for i := range res {
		res[i] = uint64(i) * 2
	}
	return res
}

func hexToBytes(str string) []byte {
	buff, _ := hex.DecodeString(str)
	return buff
}

func u256(val uint64) *uint256.Int {
	return uint256.NewInt(val)
}

func u256Hex(h string) *uint256.Int {
	v, _ := uint256.FromHex(h)
	return v
}
