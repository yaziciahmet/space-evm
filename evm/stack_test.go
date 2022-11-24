package space_evm

import (
	"fmt"
	"testing"
)

var stackPushTests = []genericTest{
	{s: "push 1 item", in: []uint64{4}, exp: stackTestExp{1, u256(4)}},
	{s: "push 2 items", in: []uint64{4, 6}, exp: stackTestExp{2, u256(6)}},
	{s: "push 5 items", in: []uint64{4, 6, 8, 10, 15}, exp: stackTestExp{5, u256(15)}},
	{s: "push 1024 items", in: genU64Slice(1024), exp: stackTestExp{1024, u256(2046)}},
	{s: "overflow 1025 items", in: genU64Slice(1025), exp: ErrStackOverflow, shouldFail: true},
}

func Test_Stack_Push(t *testing.T) {
	anyTestFailed := false
	for _, test := range stackPushTests {
		stack := NewStack()
		err := populateStack(stack, test.in.([]uint64)...)
		if !test.shouldFail {
			top, _ := stack.peek(0)
			test.act = stackTestExp{size: stack.Size(), top: top}
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

// input first item size to push, second item size to pop
var stackPopTests = []genericTest{
	{s: "pop 1 item from 1", in: []uint64{1, 1}, exp: stackTestExp{0, nil}},
	{s: "pop 2 item from 2", in: []uint64{2, 2}, exp: stackTestExp{0, nil}},
	{s: "pop 1 item from 2", in: []uint64{2, 1}, exp: stackTestExp{1, u256(0)}},
	{s: "pop 5 item from 10", in: []uint64{10, 5}, exp: stackTestExp{5, u256(8)}},
	{s: "pop 1024 item from 1024", in: []uint64{1024, 1024}, exp: stackTestExp{0, nil}},
	{s: "underflow pop 1 item from 0", in: []uint64{0, 1}, exp: ErrStackUnderflow, shouldFail: true},
	{s: "underflow pop 1025 item from 1024", in: []uint64{1024, 1025}, exp: ErrStackUnderflow, shouldFail: true},
}

func Test_Stack_Pop(t *testing.T) {
	anyTestFailed := false
	for _, test := range stackPopTests {
		stack := NewStack()
		testIn := test.in.([]uint64)
		populateStack(stack, genU64Slice(testIn[0])...)
		err := stackPopN(stack, testIn[1])
		if !test.shouldFail {
			top, _ := stack.peek(0)
			test.act = stackTestExp{size: stack.Size(), top: top}
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
