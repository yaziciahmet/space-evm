package space_evm

type handlerFunc func(*RunState) error
type dynGasHandlerFunc func(*RunState) (uint64, error)

type opInfo struct {
	name          string
	handler       handlerFunc
	constGas      uint64
	dynGasHandler dynGasHandlerFunc
}

// JumpTable contains info about given fork's valid opcodes
type JumpTable [256]opInfo

func newMoonInstructionSet() *JumpTable {
	return &JumpTable{
		0x01: {
			name:          "ADD",
			handler:       opAdd,
			constGas:      3,
			dynGasHandler: nil,
		},
		0x02: {
			name:          "MUL",
			handler:       opMul,
			constGas:      5,
			dynGasHandler: nil,
		},
		0x05: {
			name:          "SDIV",
			handler:       opSDiv,
			constGas:      5,
			dynGasHandler: nil,
		},
		0x0a: {
			name:          "EXP",
			handler:       opExp,
			constGas:      10,
			dynGasHandler: expGasCost,
		},
		0x52: {
			name:          "MSTORE",
			handler:       opMStore,
			constGas:      3,
			dynGasHandler: memoryExpansionGasCost,
		},
		0x53: {
			name:          "MSTORE8",
			handler:       opMStore8,
			constGas:      3,
			dynGasHandler: memoryExpansionGasCost,
		},
		0x60: {
			name:          "PUSH1",
			handler:       opPush,
			constGas:      3,
			dynGasHandler: nil,
		},
		0x61: {
			name:          "PUSH2",
			handler:       opPush,
			constGas:      3,
			dynGasHandler: nil,
		},
		0x62: {
			name:          "PUSH3",
			handler:       opPush,
			constGas:      3,
			dynGasHandler: nil,
		},
		0x7f: {
			name:          "PUSH32",
			handler:       opPush,
			constGas:      3,
			dynGasHandler: nil,
		},
	}
}

func (jt *JumpTable) getOpInfo(opcode byte) *opInfo {
	return &(*jt)[opcode]
}

// Enums for evm forks
type EVMFork int

const (
	Moon EVMFork = iota
)

// Return stringified enum
func (fork EVMFork) String() string {
	return []string{"Moon"}[fork]
}
