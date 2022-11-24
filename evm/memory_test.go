package space_evm

import (
	"fmt"
	"testing"
)

var memoryExtendTests = []genericTest{
	{s: "extend 1 to 32", in: uint64(1), exp: 32},
	{s: "extend 31 to 32", in: uint64(31), exp: 32},
	{s: "extend 32 to 32", in: uint64(32), exp: 32},
	{s: "extend big number", in: uint64(3232323210), exp: 3232323232},
	// can not stably test the overflow case, because
	// the maximum size of an array is smh. in between
	// uint32.max and uint64.max and not a constant number
}

func Test_Memory_Extend(t *testing.T) {
	anyTestFailed := false
	for _, test := range memoryExtendTests {
		mem := NewMemory()
		mem.extend(test.in.(uint64))
		test.act = mem.ByteLen()
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var memoryStore1Tests = []genericTest{
	{
		s:   "store 1 at offset 0",
		in:  []memoryStoreTestIn{{0, uint8(1)}},
		exp: hexToBytes("0100000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 31",
		in:  []memoryStoreTestIn{{31, uint8(1)}},
		exp: hexToBytes("0000000000000000000000000000000000000000000000000000000000000001"),
	},
	{
		s:   "store 1 at offset 32",
		in:  []memoryStoreTestIn{{32, uint8(1)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "store 255 at offset 78",
		in:  []memoryStoreTestIn{{78, uint8(255)}},
		exp: hexToBytes("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ff0000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 0, and 2 at offset 31",
		in:  []memoryStoreTestIn{{0, uint8(1)}, {31, uint8(2)}},
		exp: hexToBytes("0100000000000000000000000000000000000000000000000000000000000002"),
	},
	{
		s:   "store 1 at offset 0, and 2 at offset 32",
		in:  []memoryStoreTestIn{{0, uint8(1)}, {32, uint8(2)}},
		exp: hexToBytes("01000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 31, and 255 at offset 31",
		in:  []memoryStoreTestIn{{31, uint8(1)}, {31, uint8(255)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000ff"),
	},
}

func Test_Memory_Store1(t *testing.T) {
	anyTestFailed := false
	for _, test := range memoryStore1Tests {
		mem := NewMemory()
		testIn := test.in.([]memoryStoreTestIn)
		memStore(mem, testIn, false)
		test.act = []byte(*mem)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var memoryStore32Tests = []genericTest{
	{
		s:   "store 1 at offset 0",
		in:  []memoryStoreTestIn{{0, u256(1)}},
		exp: hexToBytes("0000000000000000000000000000000000000000000000000000000000000001"),
	},
	{
		s:   "store 1 at offset 1",
		in:  []memoryStoreTestIn{{1, u256(1)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 32",
		in:  []memoryStoreTestIn{{32, u256(1)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"),
	},
	{
		s:   "store max256 at offset 78",
		in:  []memoryStoreTestIn{{78, MaxUint256}},
		exp: hexToBytes("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 0, and 2 at offset 1",
		in:  []memoryStoreTestIn{{0, u256(1)}, {1, u256(2)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "store 1 at offset 0, and 2 at offset 32",
		in:  []memoryStoreTestIn{{0, u256(1)}, {32, u256(2)}},
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"),
	},
	{
		s:   "store 1 at offset 32, and max256 at offset 32",
		in:  []memoryStoreTestIn{{32, u256(1)}, {32, MaxUint256}},
		exp: hexToBytes("0000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
	},
}

func Test_Memory_Store32(t *testing.T) {
	anyTestFailed := false
	for _, test := range memoryStore32Tests {
		mem := NewMemory()
		testIn := test.in.([]memoryStoreTestIn)
		memStore(mem, testIn, true)
		test.act = []byte(*mem)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}
