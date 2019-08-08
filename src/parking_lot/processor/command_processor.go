package processor

import (
	"fmt"
	"parking_lot/constants"
	"parking_lot/models"
	"parking_lot/park"
	"parking_lot/util"
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
	parkingLotObj := park.GetInstance()
	switch command {
	case constants.CommandCreate:
		capacity, err := util.StringToInt(params[0])
		if err != nil {
			fmt.Println("Error while parsing Capacity of parking lot", params[0], err)
			return
		}
		parkingLotObj.Create(capacity)
	case constants.CommandPark:
		regNo := params[0]
		licenceNo := params[1]
		vehicle := models.CreateVehicle(licenceNo, regNo)
		parkingLotObj.Park(vehicle)
	case constants.CommandLeave:
		slot, err := util.StringToInt(params[0])
		if err != nil {
			fmt.Println("Error while parsing slot for unparking", params[0], err )
			return
		}
		parkingLotObj.Leave(slot)
	case constants.CommandStatus:
		parkingLotObj.Status()
	case constants.CommandColorToReg:
		color := params[0]
		parkingLotObj.ColorToVehicle(color)
	case constants.CommandColorToSlot:
		color := params[0]
		parkingLotObj.ColorToSlot(color)
	case constants.CommandSlotToReg:
		slot, err := util.StringToInt(params[0])
		if err != nil {
			fmt.Println("Error while parsing slot for fetching vehicle registration number", params[0], err )
			return
		}
		parkingLotObj.GetVehicleOnSlot(slot)
	}
}
