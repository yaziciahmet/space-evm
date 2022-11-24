# <p align=center>SPACE EVM<p>
Hello there Earthling! It is 2080 here and I am here to tell my story from the future.

Mankind extinction has arrived and only few hundreds of people were able to survive and currently live in a spaceship. We developed farms and necessary storages to ensure the survival up here. But we lack one thing, which is a must for the proper living of the human civilization. You guessed it, it's blockchain. And me as a helpless human being, am the only one who can bring blockchain to our space civilization.

We are building the civilization from bottom to ground, therefore we must not make the mistakes our ancestors did. You guessed it again, it's the banks & government. This work of mine will ensure that in any field which involves exchange, there will be no pre-requisite to trust another party. The only truth will be the code itself!

I did not have big advancements yet, but I was able to do some doodling. I might not live long after I deliver you this, there are high-rank people who does not support what I am trying to achieve here. In such a case I must make sure this work of mine is inherited to another human being.

I will give you everyting you need to know to carry-on with this divine mission, <b>Space EVM</b>!

## Introduction
Space EVM is a 256-bit Virtual Machine capable of executing bytecode. Bytecode is a series of bytes that are interpreted and executed by the EVM Interpreter.

It currently contains 2 different types of memory, one is stack and other is memory. They are both limited to their execution, and do not persist after the execution is done. I am planning to bring another type of memory that persists in between executions, which is storage.

I did not seperate the project into multiple packages, because evm components do not mean anything outside of the EVM context, hence I put them all into single package.

In main file, you can find an example CLI application which uses Space EVM to execute bytecode.

## Opcodes
The only existing fork of this EVM is called the Moon Fork. It currently supports limited number of operations, but I believe this fork will be the basis for all the future forks.

Operations are represented with 1 byte. Following table shows the currently supported operations and relevant information about them.

Operation | Opcode | Read Value | Stack Input | Stack Output | Description
:---: | :---: | :---: | :---: | :---: | :---:
ADD | 01 | - | X \| Y | X + Y | addition
MUL | 02 | - | X \| Y | X * Y | multiplication
SDIV | 05 | - | X \| Y | X / Y | signed division
EXP | 0A | - | X \| Y | X ^ Y | exponentiation
MSTORE | 52 | - | X \| Y | - | store 32 bytes to memory
MSTORE8 | 53 | - | X \| Y | - | store 1 byte to memory
PUSH1 | 60 | 1 byte | - | value | push 1 byte value to stack
PUSH2 | 61 | 2 bytes | - | value | push 2 bytes value to stack
PUSH3 | 62 | 3 bytes | - | value | push 3 bytes value to stack
PUSH32 | 7F | 32 bytes | - | value | push 32 bytes value to stack

## Dependencies
- Install dependencies

  ```go get```

## Tests
- Run tests from root dir:

  ```go test ./evm -v```

## Main
main.go is an example CLI app for running bytecode using Space EVM.

- Run main:

  ```go run main.go --bytecode <bytecode> --gas <gas>```

**--bytecode (required):** bytecode to be executed, should contain only hex characters with no '0x' prefix.

**--gas (optional):** gas limit for the execution, it is a decimal, and it's default value is 1_000_000_000

## Examples
- ```go run main.go --bytecode 6001``` :

  ```
  --------------------------------------------------
  Memory Keccak256:     c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
  Total Gas Consumed:   3
  Gas Refund:           999999997
  --------------------------------------------------
  ```
- ```go run main.go --bytecode 600061000152``` :
  ```
  --------------------------------------------------
  Memory Keccak256:     ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5
  Total Gas Consumed:   15
  Gas Refund:           999999985
  --------------------------------------------------
  ```