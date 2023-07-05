package operation

import "fmt"

var Operation bool = true

func Set(operation bool) {
	Operation = operation
	if operation {
		fmt.Println("---------> [OP]: Operation resumed")
		return
	} else {
		fmt.Println("---------> [OP]: Operation paused")
	}
}

func Get() bool {
	return Operation
}
