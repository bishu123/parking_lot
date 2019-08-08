package constants

type CommandType string

const (
	CommandCreate      CommandType = "create_parking_lot"
	CommandPark        CommandType = "park"
	CommandLeave       CommandType = "leave"
	CommandStatus      CommandType = "status"
	CommandColorToReg  CommandType = "registration_numbers_for_cars_with_colour"
	CommandColorToSlot CommandType = "slot_numbers_for_cars_with_colour"
	CommandSlotToReg   CommandType = "slot_number_for_registration_number"
)

var CommandToArgMap = map[CommandType]int{
	CommandCreate:      1,
	CommandPark:        2,
	CommandLeave:       1,
	CommandStatus:      0,
	CommandColorToReg:  1,
	CommandColorToSlot: 1,
	CommandSlotToReg:   1,
}

type ErrorType string

const (
	ErrorInvalidCommand            ErrorType = "Invalid Command"
	ErrorParkingFull               ErrorType = "Parking Full"
	ErrorParkingNotInitialized     ErrorType = "Parking Not Initialized"
	ErrorParkingAlreadyInitialized ErrorType = "Paring Already Initialized"
	ErrorArgumentMisMatch          ErrorType = "Argument MisMatch"
)
