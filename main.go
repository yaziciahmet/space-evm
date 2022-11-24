package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	space_evm "space/evm"
	"strconv"
)

func parseFlags(bytecode string, gas string) ([]byte, uint64, error) {
	if bytecode == "" {
		return nil, 0, errors.New("missing flag --bytecode")
	}

	gasLimit, err := strconv.Atoi(gas)
	if err != nil {
		return nil, 0, errors.New("gas should be a valid decimal")
	}

	code, err := hex.DecodeString(bytecode)
	if err != nil {
		return nil, 0, err
	}
	return code, uint64(gasLimit), nil
}

func main() {
	// bytecode required, gas optional
	var (
		bytecode string
		gas      string
	)
	flag.StringVar(&bytecode, "bytecode", "", "bytecode to be executed")
	flag.StringVar(&gas, "gas", "1000000000", "gas limit for the execution")
	flag.Parse()

	code, gasLimit, err := parseFlags(bytecode, gas)
	if err != nil {
		fmt.Println(err)
		return
	}

	evm := space_evm.NewEVM(space_evm.Moon)
	evm.RunCode(code, gasLimit)
}
