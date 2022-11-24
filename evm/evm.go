package space_evm

// EVM is the Ethereum Virtual Machine
// which is capable of executing bytecode
type EVM struct {
	Fork        EVMFork
	interpreter *Interpreter
}

// Create an EVM instance
func NewEVM(fork EVMFork) *EVM {
	return &EVM{
		Fork:        fork,
		interpreter: NewInterpreter(fork),
	}
}

// Run the code with the given gasLimit, and
// display the results to command line.
// Instruction set to interpret the code
// is determined by the given EVM fork.
func (evm *EVM) RunCode(code []byte, gasLimit uint64) {
	result := evm.interpreter.Run(code, gasLimit)
	result.Display()
}
