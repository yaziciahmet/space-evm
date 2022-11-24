package space_evm

func memoryExpansionGasCost(runState *RunState) (uint64, error) {
	var size uint64
	if runState.Opcode == 0x52 {
		size = 32 // MSTORE
	} else {
		size = 1 // MSTORE8
	}

	offset, err := runState.Stack.peek(0)
	if err != nil {
		return 0, err
	}
	newMemByteLen := ceil32(offset.Uint64() + size)
	// any newMemByteLen above the constant number
	// 0x1FFFFFFFE0 causes square operation to overflow
	if newMemByteLen > 0x1FFFFFFFE0 {
		return 0, ErrGasUintOverflow
	}
	if newMemByteLen <= uint64(runState.Memory.ByteLen()) {
		return 0, nil
	}

	newMemWordLen := newMemByteLen / 32
	newMemCost := 3*newMemWordLen + newMemWordLen*newMemWordLen/512
	oldMemCost := runState.HighestMemoryGasCost
	runState.HighestMemoryGasCost = newMemCost

	return newMemCost - oldMemCost, nil
}

func expGasCost(runState *RunState) (uint64, error) {
	exp, err := runState.Stack.peek(1)
	if err != nil {
		return 0, err
	}
	return uint64(exp.ByteLen()) * 50, nil
}
