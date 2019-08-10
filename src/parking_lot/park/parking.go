package park

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"parking_lot/constants"
	"parking_lot/models"
	"parking_lot/util"
)

var parkingLot *ParkingLot
var once sync.Once

func GetInstance() *ParkingLot {
	once.Do(func() {
		parkingLot = &ParkingLot{}
	})
	return parkingLot
}

type BaseParkingLoter interface {
	Create(capacity int) error
	Park(vehicle *models.Vehicle) error
	Leave(slot int) error
	Status()
}

type ParkingLot struct {
	maximumSlot    int
	availableSlot  int
	regNoSlotMap   map[string]int
	slotVehicleMap map[int]*models.Vehicle
	colorRegNoMap  map[string]map[int]string
	isCreated      bool
}

func (p *ParkingLot) Create(capacity int) error {
	if p.isCreated {
		parkingCreationErr := fmt.Errorf(string(constants.ErrorParkingAlreadyCreated))
		return errors.New(parkingCreationErr.Error())
	}
	p.isCreated = true
	p.maximumSlot = capacity
	p.availableSlot = capacity
	p.regNoSlotMap = make(map[string]int)
	p.slotVehicleMap = make(map[int]*models.Vehicle)
	p.colorRegNoMap = make(map[string]map[int]string)
	fmt.Printf("Created a parking lot with %v slots", capacity)
	return nil
}

func (p *ParkingLot) Park(vehicle *models.Vehicle) error {
	for slot := 0; slot < p.maximumSlot; slot++ {
		if value, ok := p.slotVehicleMap[slot]; value != nil && !ok {
			p.slotVehicleMap[slot] = vehicle
			p.putColorToVehicleMap(slot, vehicle)
			p.regNoSlotMap[vehicle.LicenceNumber] = slot
			p.availableSlot -= 1
			fmt.Println("Allocated slot number: ", slot)
			return nil
		}
	}
	return errors.New(string(constants.ErrorParkingFull))
}

func (p *ParkingLot) Leave(slot int) error {
	vehicle := p.getVehicleOnSlot(slot)
	if vehicle == nil {
		noVehicleError := fmt.Errorf("no Vehicle on Slot: %v", slot)
		return errors.New(noVehicleError.Error())
	} else {
		p.removeColorToVehicleMap(slot)
		delete(p.regNoSlotMap, vehicle.LicenceNumber)
		delete(p.slotVehicleMap, slot)
		p.availableSlot += 1
		fmt.Printf("Slot number %v is free", slot)
	}
	return nil
}

func (p *ParkingLot) Status() {
	// returning the result in the sorted occupied slot.
	fmt.Println("Slot No.		Registration No		Colour")
	for slot := 0; slot < p.maximumSlot; slot++ {
		if vehicle, ok := p.slotVehicleMap[slot]; ok {
			fmt.Println(slot, "\t\t", vehicle.LicenceNumber, "\t\t", vehicle.Color)
		}
	}
}

func (p *ParkingLot) ColorToVehicle(color string) error {
	var vehicleList []string
	if slotRegMap, ok := p.colorRegNoMap[color]; !ok {
		slotRegErr := fmt.Errorf("no vehicle with this color")
		return errors.New(slotRegErr.Error())
	} else {
		for _, regNo := range slotRegMap {
			vehicleList = append(vehicleList, regNo)
		}
	}
	if len(vehicleList) > 0 {
		fmt.Println(strings.Join(vehicleList, ","))
	}
	return nil
}

func (p *ParkingLot) ColorToSlot(color string) error {
	var slotList []string
	if slotRegMap, ok := p.colorRegNoMap[color]; !ok {
		slotRegErr := fmt.Errorf("no slot assigned to the vehicle")
		return errors.New(slotRegErr.Error())
	} else {
		for slot, _ := range slotRegMap {
			slotList = append(slotList, util.IntToString(slot))
		}
	}
	if len(slotList) > 0 {
		fmt.Println(strings.Join(slotList, ","))
	}
	return nil
}

func (p *ParkingLot) RegNoToSlot(regNo string) error {
	if slot, ok := p.regNoSlotMap[regNo]; !ok {
		fmt.Println(slot)
	} else {
		regNoErr := fmt.Errorf("Not Found")
		return errors.New(regNoErr.Error())
	}
}
