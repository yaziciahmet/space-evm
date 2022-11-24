package space_evm

import (
	"errors"
	"fmt"
	"testing"
)

var byteSliceToUint256Tests = []genericTest{
	{s: "convert 0", in: []byte{0}, exp: u256(0)},
	{s: "convert 1", in: []byte{1}, exp: u256(1)},
	{s: "convert 255", in: []byte{255}, exp: u256(255)},
	{s: "convert 256", in: []byte{1, 0}, exp: u256(256)},
	{
		s:   "convert 2^248",
		in:  u256Hex("0x100000000000000000000000000000000000000000000000000000000000000").Bytes(),
		exp: u256Hex("0x100000000000000000000000000000000000000000000000000000000000000"),
	},
	{
		s:   "convert max256",
		in:  u256Hex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff").Bytes(),
		exp: u256Hex("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
	},
	{
		s:          "uint256 overflow",
		in:         []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		exp:        errors.New("hex number > 256 bits"),
		shouldFail: true,
	},
}

func Test_Common_ByteSliceToUint256(t *testing.T) {
	anyTestFailed := false
	for _, test := range byteSliceToUint256Tests {
		if !test.shouldFail {
			test.act, _ = byteSliceToUint256(test.in.([]byte))
		} else {
			_, test.act = byteSliceToUint256(test.in.([]byte))
		}
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var ceil32Tests = []genericTest{
	{s: "ceil 0 to 0", in: uint64(0), exp: uint64(0)},
	{s: "ceil 1 to 32", in: uint64(1), exp: uint64(32)},
	{s: "ceil 32 to 32", in: uint64(31), exp: uint64(32)},
	{s: "ceil 150 to 160", in: uint64(150), exp: uint64(160)},
	{s: "max uint64 overflow", in: MaxUint64, exp: uint64(0)},
}

func Test_Common_Ceil32(t *testing.T) {
	anyTestFailed := false
	for _, test := range ceil32Tests {
		test.act = ceil32(test.in.(uint64))
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}
