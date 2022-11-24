package space_evm

import (
	"errors"
	"fmt"
	"testing"
)

var interpreterRunTests = []genericTest{
	{
		s: "trapesys example 1",
		in: interpreterRunTestIn{
			code:     hexToBytes("60016020526002606452600361ff0052600362ffffff526005601053"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("ab2744998886b708acadc0a32428d0aa1953e83924383d21c6de5dac852ccbcc"),
				ConsumedGas:  538445872,
				GasRefund:    MaxUint64 - 538445872,
			},
			stack: &Stack{},
		},
	},
	{
		s: "trapesys example 2",
		in: interpreterRunTestIn{
			code:     hexToBytes("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00016000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00026020527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff05604052"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("b9a07dba38aa24923a611fced9d2eede3bfbfa281e5e498d60f4bd99e5ce6a15"),
				ConsumedGas:  58,
				GasRefund:    MaxUint64 - 58,
			},
			stack: &Stack{},
		},
	},
	{
		s: "trapesys example 3",
		in: interpreterRunTestIn{
			code:     hexToBytes("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000a6000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000a6020527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0a604052"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("afe1e714d2cd3ed5b0fa0a04ee95cd564b955ab8661c5665588758b48b66e263"),
				ConsumedGas:  4875,
				GasRefund:    MaxUint64 - 4875,
			},
			stack: &Stack{},
		},
	},
	{
		s: "push1",
		in: interpreterRunTestIn{
			code:     hexToBytes("6001"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  3,
				GasRefund:    MaxUint64 - 3,
			},
			stack: &Stack{*u256(1)},
		},
	},
	{
		s: "push2",
		in: interpreterRunTestIn{
			code:     hexToBytes("610001"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  3,
				GasRefund:    MaxUint64 - 3,
			},
			stack: &Stack{*u256(1)},
		},
	},
	{
		s: "push3",
		in: interpreterRunTestIn{
			code:     hexToBytes("62000001"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  3,
				GasRefund:    MaxUint64 - 3,
			},
			stack: &Stack{*u256(1)},
		},
	},
	{
		s: "push32",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f0000000000000000000000000000000000000000000000000000000000000001"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  3,
				GasRefund:    MaxUint64 - 3,
			},
			stack: &Stack{*u256(1)},
		},
	},
	{
		s: "add",
		in: interpreterRunTestIn{
			code:     hexToBytes("600161000101"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  9,
				GasRefund:    MaxUint64 - 9,
			},
			stack: &Stack{*u256(2)},
		},
	},
	{
		s: "mul",
		in: interpreterRunTestIn{
			code:     hexToBytes("6100016200000102"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  11,
				GasRefund:    MaxUint64 - 11,
			},
			stack: &Stack{*u256(1)},
		},
	},
	{
		s: "sdiv",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f00000000000000000000000000000000000000000000000000000000000000017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff05"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  11,
				GasRefund:    MaxUint64 - 11,
			},
			stack: &Stack{*MaxUint256},
		},
	},
	{
		s: "exp",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f00000000000000000000000000000000000000000000000000000000000000047f00000000000000000000000000000000000000000000000000000000000000020a"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: EmptyMemHash,
				ConsumedGas:  66,
				GasRefund:    MaxUint64 - 66,
			},
			stack: &Stack{*u256(16)},
		},
	},
	{
		s: "mstore",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000152"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5"),
				ConsumedGas:  15,
				GasRefund:    MaxUint64 - 15,
			},
			stack: &Stack{},
		},
	},
	{
		s: "mstore8",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000153"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563"),
				ConsumedGas:  12,
				GasRefund:    MaxUint64 - 12,
			},
			stack: &Stack{},
		},
	},
	{
		s: "all opcodes",
		in: interpreterRunTestIn{
			code:     hexToBytes("60016100020162000002027f00000000000000000000000000000000000000000000000000000000000000030a6000526020600253"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("7e5130c77fb0eb5f9b14f9fc0a37569fb344b1ba6c214486998eb3444bcf4998"),
				ConsumedGas:  98,
				GasRefund:    MaxUint64 - 98,
			},
			stack: &Stack{},
		},
	},
	{
		s: "complex big code",
		in: interpreterRunTestIn{
			code:     hexToBytes("7f000000000000000000000000000000000000000000000000000000000000000f7f00000000000000000000000000000000000000000000000000000000000003e80a6001526200000f7f00000000000000000000000000000000000000000000000000000000000005dc01603f5261007b7f00000000000000000000000000000000000000000000000000000000002396a453"),
			gasLimit: MaxUint64,
		},
		exp: interpreterRunTestExp{
			runRes: &RunResult{
				HashedMemory: hexToBytes("7936548dcf1970696184a2a10f10cd6144bc36786b2492c64602171747439274"),
				ConsumedGas:  10594474,
				GasRefund:    MaxUint64 - 10594474,
			},
			stack: &Stack{},
		},
	},
	{
		s: "invalid opcode",
		in: interpreterRunTestIn{
			code:     hexToBytes("09"),
			gasLimit: MaxUint64,
		},
		exp:        ErrInvalidOpcode(0x09),
		shouldFail: true,
	},
	{
		s: "stack underflow",
		in: interpreterRunTestIn{
			code:     hexToBytes("600501"),
			gasLimit: MaxUint64,
		},
		exp:        ErrStackUnderflow,
		shouldFail: true,
	},
	{
		s: "gas uint overflow",
		in: interpreterRunTestIn{
			code:     hexToBytes("60017f0000000000000000000000000000000000000000000000000000001fffffffe052"),
			gasLimit: MaxUint64,
		},
		exp:        ErrGasUintOverflow,
		shouldFail: true,
	},
	{
		s: "out of gas",
		in: interpreterRunTestIn{
			code:     hexToBytes("6001610002"),
			gasLimit: 3,
		},
		exp:        ErrOutOfGas,
		shouldFail: true,
	},
	{
		s: "not enough bytes to read",
		in: interpreterRunTestIn{
			code:     hexToBytes("60016102"),
			gasLimit: MaxUint64,
		},
		exp:        ErrNotEnoughBytesToRead(0x61),
		shouldFail: true,
	},
}

func Test_Interpreter_Run(t *testing.T) {
	anyTestFailed := false
	for _, test := range interpreterRunTests {
		in := NewInterpreter(Moon)
		testIn := test.in.(interpreterRunTestIn)
		runRes := in.Run(testIn.code, testIn.gasLimit)
		if !test.shouldFail {
			test.act = interpreterRunTestExp{
				runRes: runRes,
				stack:  in.runState.Stack,
			}
		} else {
			// remove "evm error: " prefix
			test.act = errors.New(runRes.EvmError.Error()[11:])
		}
		msg, failed := test.Check()
		anyTestFailed = anyTestFailed || failed
		fmt.Print(msg)
	}
	if anyTestFailed {
		t.FailNow()
	}
}
