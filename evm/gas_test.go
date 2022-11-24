package space_evm

import (
	"fmt"
	"testing"
)

// stackValue in genRunState is the offset to read
var memoryExpansionGasCostTests = []genericTest{
	{
		s:   "gas for 0 to 1 word",
		in:  genRunState("", 0x52, []uint64{0}, genZeroMem(0)),
		exp: uint64(3),
	},
	{
		s:   "gas for 1 to 2 word",
		in:  genRunState("", 0x52, []uint64{32}, genZeroMem(1)),
		exp: uint64(3),
	},
	{
		s:   "gas for 0 to 5 word",
		in:  genRunState("", 0x52, []uint64{128}, genZeroMem(0)),
		exp: uint64(15),
	},
	{
		s:   "gas for 10 to 25 word",
		in:  genRunState("", 0x52, []uint64{24 * 32}, genZeroMem(10)),
		exp: uint64(46),
	},
	{
		s:   "gas for 1 to 512 word",
		in:  genRunState("", 0x52, []uint64{511 * 32}, genZeroMem(1)),
		exp: uint64(2045),
	},
	{
		s:          "gas uint64 overflow",
		in:         genRunState("", 0x52, []uint64{0x1FFFFFFFE0}, genZeroMem(0)),
		exp:        ErrGasUintOverflow,
		shouldFail: true,
	},
}

func Test_Gas_MemoryExpansionGasCost(t *testing.T) {
	anyTestFailed := false
	for _, test := range memoryExpansionGasCostTests {
		if !test.shouldFail {
			test.act, _ = memoryExpansionGasCost(test.in.(*RunState))
		} else {
			_, test.act = memoryExpansionGasCost(test.in.(*RunState))
		}
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}

// stackValues first element is the exp
var expGasCostTests = []genericTest{
	{
		s: "gas for 1 byte",
		in: genRunState("", 0x0a, []uint64{1, 0}, genZeroMem(0)),
		exp: uint64(50),
	},
	{
		s: "gas for 2 byte",
		in: genRunState("", 0x0a, []uint64{256, 0}, genZeroMem(0)),
		exp: uint64(100),
	},
	{
		s: "gas for 8 byte",
		in: genRunState("", 0x0a, []uint64{MaxUint64, 0}, genZeroMem(0)),
		exp: uint64(400),
	},
	{
		s: "gas for 32 byte",
		// exp value will be pushed into stack later
		in: genRunState("", 0x0a, []uint64{}, genZeroMem(0)),
		exp: uint64(1600),
	},
}

func Test_Gas_ExpGasCost(t *testing.T) {
	anyTestFailed := false
	for _, test := range expGasCostTests {
		runSt := test.in.(*RunState)
		if runSt.Stack.Size() == 0 {
			runSt.Stack.push(MaxUint256)
			runSt.Stack.push(u256(0))
		}
		test.act, _ = expGasCost(runSt)
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}
