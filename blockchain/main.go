package blockchain

import (
	"alertbot/rpc"
	"fmt"
	"strconv"
)

var BlockHeight string = ""
var CurrentBlock int = 0
var CurrentSignature string = ""

func InitBlock() {
	Height, err := rpc.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}

	BlockHeight = Height
	CurrentBlock, err = strconv.Atoi(Height)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetCurrentBlockHeight() int {
	return CurrentBlock
}

func IncreaseCurrentBlock() {
	CurrentBlock++
}

func GetBlockHeight() string {
	return BlockHeight
}

func SetCurrentSignature(signature string) {
	CurrentSignature = signature
}

func GetCurrentSignature() string {
	return CurrentSignature
}
