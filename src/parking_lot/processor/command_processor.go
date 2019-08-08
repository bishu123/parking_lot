package processor

import (
	"fmt"
	"parking_lot/constants"
	"strings"
)

func processInput(inputStr string){
	splitParameter := strings.Split(inputStr, " ")
	if len(splitParameter) < 1 {
		fmt.Println(constants.ErrorInvalidCommand)
		return
	}
	command := constants.CommandType(splitParameter[0])
	params := splitParameter[1:]
	commandParams, ok := constants.CommandToArgMap[command]
	if !ok {
		fmt.Println(constants.ErrorInvalidCommand)
	}
	if commandParams != len(params){
		fmt.Println(constants.ErrorArgumentMisMatch)
	}
	processCommand(command, params)
}

func processCommand(command constants.CommandType, params []string) {
	switch command {
	case constants.CommandCreate:
	case constants.CommandPark:
	case constants.CommandLeave:
	case constants.CommandStatus:
	case constants.CommandColorToReg:
	case constants.CommandColorToSlot:
	case constants.CommandSlotToReg:
	}
}
