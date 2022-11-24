package space_evm

func opAdd(runState *RunState) error {
	x, err1 := runState.Stack.pop()
	y, err2 := runState.Stack.peek(0)
	if err1 != nil || err2 != nil {
		return err2
	}
	y.Add(x, y)
	return nil
}

func opMul(runState *RunState) error {
	x, err1 := runState.Stack.pop()
	y, err2 := runState.Stack.peek(0)
	if err1 != nil || err2 != nil {
		return err2
	}
	y.Mul(x, y)
	return nil
}

func opSDiv(runState *RunState) error {
	x, err1 := runState.Stack.pop()
	y, err2 := runState.Stack.peek(0)
	if err1 != nil || err2 != nil {
		return err2
	}
	y.SDiv(x, y)
	return nil
}

func opExp(runState *RunState) error {
	x, err1 := runState.Stack.pop()
	y, err2 := runState.Stack.peek(0)
	if err1 != nil || err2 != nil {
		return err2
	}
	y.Exp(x, y)
	return nil
}

func opPush(runState *RunState) error {
	// push opcodes are in range 0x60 to 0x7f, hence
	// (opcode - 0x5f) gives the byte len of the value to push
	n := int(runState.Opcode) - 0x5f
	nBytes := runState.ReadNextBytes(n)
	if nBytes == nil {
		return ErrNotEnoughBytesToRead(runState.Opcode)
	}

	val, err := byteSliceToUint256(nBytes)
	if err != nil {
		return err
	}

	err = runState.Stack.push(val)
	if err != nil {
		return err
	}
	return nil
}

func opMStore(runState *RunState) error {
	offset, err1 := runState.Stack.pop()
	val, err2 := runState.Stack.pop()
	if err1 != nil || err2 != nil {
		// pop only returns stack underflow error,
		// hence it is ok to return the last error
		return err2
	}
	runState.Memory.store32(offset.Uint64(), val)
	return nil
}

func opMStore8(runState *RunState) error {
	offset, err1 := runState.Stack.pop()
	val, err2 := runState.Stack.pop()
	if err1 != nil || err2 != nil {
		return err2
	}
	runState.Memory.store1(offset.Uint64(), byte(val.Uint64()))
	return nil
}
