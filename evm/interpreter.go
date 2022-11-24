package space_evm

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// RunState handles the state of the interpreter run
type RunState struct {
	Code                 []byte
	Stack                *Stack
	Memory               *Memory
	HighestMemoryGasCost uint64
	RemainingGas         uint64
	ConsumedGas          uint64
	ProgramCounter       int
	Opcode               byte
}

func NewRunState(code []byte, gasLimit uint64) *RunState {
	return &RunState{
		Code:         code,
		Stack:        NewStack(),
		Memory:       NewMemory(),
		RemainingGas: gasLimit,
	}
}

// Returns next n bytes starting from the current program counter
func (runSt *RunState) ReadNextBytes(n int) []byte {
	pc := runSt.ProgramCounter
	// check if there is enough bytes to read
	if pc+n > len(runSt.Code) {
		return nil
	}
	nBytes := runSt.Code[pc : pc+n]
	runSt.ProgramCounter += n
	return nBytes
}

func (runSt *RunState) useGas(gas uint64) bool {
	if gas > runSt.RemainingGas {
		return false
	}
	runSt.RemainingGas -= gas
	runSt.ConsumedGas += gas
	return true
}

// RunResult is used to track and display the result of the interpreter run
type RunResult struct {
	HashedMemory []byte
	ConsumedGas  uint64
	GasRefund    uint64
	EvmError     error
}

func NewRunResult() *RunResult {
	return &RunResult{}
}

func (res *RunResult) setResult(runState *RunState) {
	res.ConsumedGas = runState.ConsumedGas
	res.GasRefund = runState.RemainingGas
	res.HashedMemory = keccak256([]byte(*runState.Memory))
}

func (res *RunResult) setError(err error) {
	res.EvmError = errors.New("evm error: " + err.Error())
}

func (res *RunResult) Display() {
	fmt.Println("--------------------------------------------------")
	if res.EvmError == nil {
		fmt.Printf("%-22s%v\n", "Memory Keccak256:", hex.EncodeToString(res.HashedMemory))
		fmt.Printf("%-22s%v\n", "Total Gas Consumed:", res.ConsumedGas)
		fmt.Printf("%-22s%v\n", "Gas Refund:", res.GasRefund)
	} else {
		fmt.Println(res.EvmError)
	}
	fmt.Println("--------------------------------------------------")
}

// EVM Interpreter is responsible for executing the given bytecode with the
// given gasLimit with respect to instruction set it holds. Instruction set
// is automatically determined by the selected fork of the EVM.
type Interpreter struct {
	runState  *RunState
	runResult *RunResult
	jumpTable *JumpTable
}

func NewInterpreter(fork EVMFork) *Interpreter {
	var jumpTable *JumpTable

	switch fork {
	case Moon:
		jumpTable = newMoonInstructionSet()
	default:
		jumpTable = newMoonInstructionSet()
	}

	return &Interpreter{
		jumpTable: jumpTable,
	}
}

func (in *Interpreter) Run(code []byte, gasLimit uint64) *RunResult {
	in.runState = NewRunState(code, gasLimit)
	in.runResult = NewRunResult()
	// Main execution loop of interpreter. Continues
	// until encountering end of the code or error.
	for pc := 0; pc < len(code); {
		opcode := code[pc]
		opInfo := in.jumpTable.getOpInfo(opcode)
		// return error if opcode invalid
		if opInfo.handler == nil {
			in.runResult.setError(ErrInvalidOpcode(opcode))
			break
		}
		in.runState.Opcode = opcode
		in.runState.ProgramCounter += 1

		gas := opInfo.constGas
		if opInfo.dynGasHandler != nil {
			// dynamic gas handlers can return some
			// of the errors prior to execution
			dynGas, err := opInfo.dynGasHandler(in.runState)
			if err != nil {
				in.runResult.setError(err)
				break
			}
			gas += dynGas
		}
		// use gas, return error if not enough gas
		if !in.runState.useGas(gas) {
			in.runResult.setError(ErrOutOfGas)
			break
		}

		// execute operation
		err := opInfo.handler(in.runState)
		if err != nil {
			in.runResult.setError(err)
			break
		}
		pc = in.runState.ProgramCounter
	}
	in.runResult.setResult(in.runState)
	return in.runResult
}
