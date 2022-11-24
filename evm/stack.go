package space_evm

import (
	"github.com/holiman/uint256"
)

// EVM Stack is responsible from the run stack
// which holds items with size of 32 bytes
type Stack []uint256.Int

const StackMaxHeight = 1024

func NewStack() *Stack {
	return &Stack{}
}

func (st *Stack) push(val *uint256.Int) error {
	if st.Size() >= StackMaxHeight {
		return ErrStackOverflow
	}
	*st = append(*st, *val)
	return nil
}

func (st *Stack) pop() (*uint256.Int, error) {
	if st.Size() == 0 {
		return nil, ErrStackUnderflow
	}
	index := st.Size() - 1
	val := (*st)[index]
	*st = (*st)[:index]
	return &val, nil
}

// Returns a reference to n'th item in the stack.
// Reference is returned to benefit in-place update.
func (st *Stack) peek(n int) (*uint256.Int, error) {
	if st.Size()-1 < n {
		return nil, ErrStackUnderflow
	}
	return &(*st)[st.Size()-n-1], nil
}

func (st *Stack) Size() int {
	return len(*st)
}
