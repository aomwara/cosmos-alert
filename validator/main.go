package validator

import (
	"alertbot/notification"
	"fmt"
)

var Validator string = "0x"
var StartWatchBlock int = 0
var MissBlock int = 0
var MissStatus bool = false
var AlertStatus bool = false

var Uptime = [100]bool{}

func SetValidator(validator string) {
	Validator = validator
}

func GetValidator() string {
	return Validator
}

func SetStartWatchBlock(blockHeight int) {
	StartWatchBlock = blockHeight
}

func GetStartWatchBlock() int {
	return StartWatchBlock
}

func SetUptime(status bool) {
	//set uptime status to first index and shift the rest
	for i := len(Uptime) - 1; i > 0; i-- {
		Uptime[i] = Uptime[i-1]
	}
	Uptime[0] = status
}

func GetUptime() [100]bool {
	return Uptime
}

func IncreaseMissBlock() {
	MissBlock++
	fmt.Println("---------------> [IN]: Miss Block (", MissBlock, ")")
	if MissBlock >= 5 {
		fmt.Println("Miss Block >= 5")

		if MissBlock >= 5 && AlertStatus == false {
			notification.Alert("[MISS]: Miss Block more than 5")
			AlertStatus = true
		}

		MissStatus = true
	}
}

func DecreaseMissBlock() {
	if MissBlock > 0 {
		MissBlock = 0
		fmt.Println("---------------> [DE]: Miss Block (", MissBlock, ")")
		notification.Alert("[RE]: Recovery Miss Block")
		AlertStatus = false
		fmt.Println("---------------> [RE]: Recovery Miss Block")
	}

	if MissStatus && MissBlock == 0 {
		MissStatus = false
		fmt.Println("---------------> [RE]: Recovery Miss Block")
	}

}

func GetMissBlock() int {
	return MissBlock
}

func ResetMissBlock() {
	MissBlock = 0
	fmt.Println("Recovery Miss Block")
}
