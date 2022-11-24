package space_evm

import (
	"fmt"
	"testing"

	"github.com/holiman/uint256"
)

var mathTestCases = [][]*uint256.Int{
	{u256(0), u256(0)},
	{u256(0), u256(1)},
	{u256(1), u256(1)},
	{u256(8), u256(0)},
	{u256(3), u256(5)},
	{u256(20), u256(10)},
	{u256(15), u256(1000)},
	{MaxUint256, u256(1)},
	{MaxUint256, MaxUint256},
}

var opAddTests = []genericTest{
	{s: "0 + 0 is 0", exp: u256(0)},
	{s: "1 + 0 is 1", exp: u256(1)},
	{s: "1 + 1 is 2", exp: u256(2)},
	{s: "0 + 8 is 8", exp: u256(8)},
	{s: "5 + 3 is 8", exp: u256(8)},
	{s: "10 + 20 is 30", exp: u256(30)},
	{s: "1000 + 15 is 1015", exp: u256(1015)},
	{s: "1 + max256 is 0", exp: u256(0)},
	{s: "max256 + max256 is max256-1", exp: u256(0).SubUint64(MaxUint256, 1)},
}

func Test_Op_Add(t *testing.T) {
	anyTestFailed := false
	for i, test := range opAddTests {
		runSt := &RunState{}
		runSt.Stack = NewStack()
		runSt.Stack.push(mathTestCases[i][0])
		runSt.Stack.push(mathTestCases[i][1])
		opAdd(runSt)
		test.act, _ = runSt.Stack.peek(0)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var opMulTests = []genericTest{
	{s: "0 x 0 is 0", exp: u256(0)},
	{s: "1 x 0 is 0", exp: u256(0)},
	{s: "1 x 1 is 1", exp: u256(1)},
	{s: "0 x 8 is 0", exp: u256(0)},
	{s: "5 x 3 is 15", exp: u256(15)},
	{s: "10 x 20 is 200", exp: u256(200)},
	{s: "1000 x 15 is 15000", exp: u256(15000)},
	{s: "1 x max256 is max256", exp: MaxUint256},
	{s: "max256 x max256 is 1", exp: u256(1)},
}

func Test_Op_Mul(t *testing.T) {
	anyTestFailed := false
	for i, test := range opMulTests {
		runSt := &RunState{}
		runSt.Stack = NewStack()
		runSt.Stack.push(mathTestCases[i][0])
		runSt.Stack.push(mathTestCases[i][1])
		opMul(runSt)
		test.act, _ = runSt.Stack.peek(0)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var opSDivTests = []genericTest{
	{s: "0 / 0 is 0", exp: u256(0)},
	{s: "1 / 0 is 0", exp: u256(0)},
	{s: "1 / 1 is 1", exp: u256(1)},
	{s: "0 / 8 is 0", exp: u256(0)},
	{s: "5 / 3 is 1", exp: u256(1)},
	{s: "10 / 20 is 0", exp: u256(0)},
	{s: "1000 / 15 is 66", exp: u256(66)},
	// for SDiv: max256 = -1
	{s: "1 / max256 is max256", exp: MaxUint256},
	{s: "max256 / max256 is 1", exp: u256(1)},
}

func Test_Op_SDiv(t *testing.T) {
	anyTestFailed := false
	for i, test := range opSDivTests {
		runSt := &RunState{}
		runSt.Stack = NewStack()
		runSt.Stack.push(mathTestCases[i][0])
		runSt.Stack.push(mathTestCases[i][1])
		opSDiv(runSt)
		test.act, _ = runSt.Stack.peek(0)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var opExpTests = []genericTest{
	{s: "0 ^ 0 is 1", exp: u256(1)},
	{s: "1 ^ 0 is 1", exp: u256(1)},
	{s: "1 ^ 1 is 1", exp: u256(1)},
	{s: "0 ^ 8 is 0", exp: u256(0)},
	{s: "5 ^ 3 is 125", exp: u256(125)},
	{s: "10 ^ 20 is big", exp: u256Hex("0x56bc75e2d63100000")},
	{s: "1000 ^ 15 is very big", exp: u256Hex("0x2cd76fe086b93ce2f768a00b22a00000000000")},
	{s: "1 ^ max256 is 1", exp: u256(1)},
	{s: "max256 ^ max256 is 1", exp: MaxUint256},
}

func Test_Op_Exp(t *testing.T) {
	anyTestFailed := false
	for i, test := range opExpTests {
		runSt := &RunState{}
		runSt.Stack = NewStack()
		runSt.Stack.push(mathTestCases[i][0])
		runSt.Stack.push(mathTestCases[i][1])
		opExp(runSt)
		test.act, _ = runSt.Stack.peek(0)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

var opPushTests = []genericTest{
	{
		s:   "push1",
		in:  genRunState("02", 0x60, []uint64{}, []byte{}),
		exp: u256(2),
	},
	{
		s:   "push2",
		in:  genRunState("05ff", 0x61, []uint64{}, []byte{}),
		exp: u256(1535),
	},
	{
		s:   "push3",
		in:  genRunState("010203", 0x62, []uint64{}, []byte{}),
		exp: u256(66051),
	},
	{
		s:   "push32",
		in:  genRunState("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0x7f, []uint64{}, []byte{}),
		exp: MaxUint256,
	},
	{
		s:          "not enough bytes to read",
		in:         genRunState("ff", 0x61, []uint64{}, []byte{}),
		exp:        ErrNotEnoughBytesToRead(0x61),
		shouldFail: true,
	},
}

func Test_Op_Push(t *testing.T) {
	anyTestFailed := false
	for _, test := range opPushTests {
		runSt := test.in.(*RunState)
		err := opPush(runSt)
		if !test.shouldFail {
			test.act, _ = runSt.Stack.peek(0)
		} else {
			test.act = err
		}
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

// opMStore does not have an extensive test 
// cases due to almost all of the main cases 
// being already tested in memory_test.go
var opMStoreTests = []genericTest{
	{
		s: "overwrite memory",
		in: genRunState("", 0x52, []uint64{MaxUint64, 16}, genZeroMem(2)),
		exp: hexToBytes("00000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff00000000000000000000000000000000"),
	},
}

func Test_Op_MStore(t *testing.T) {
	anyTestFailed := false
	for _, test := range opMStoreTests {
		runSt := test.in.(*RunState)
		opMStore(runSt)
		test.act = []byte(*runSt.Memory)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

// opMStore8 does not have an extensive test 
// cases due to almost all of the main cases 
// being already tested in memory_test.go
var opMStore8Tests = []genericTest{
	{
		s: "overwrite memory with overflowing uint8",
		in: genRunState("", 0x53, []uint64{MaxUint64, 16}, genZeroMem(2)),
		exp: hexToBytes("00000000000000000000000000000000ff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	},
}

func Test_Op_MStore8(t *testing.T) {
	anyTestFailed := false
	for _, test := range opMStore8Tests {
		runSt := test.in.(*RunState)
		opMStore8(runSt)
		test.act = []byte(*runSt.Memory)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}