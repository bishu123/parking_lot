package park

import (
	"parking_lot/models"
	"sync"
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
	Create(capacity int)
	Park(vehicle *models.Vehicle)
	Leave(slot int)
	Status()
	GetVehicleOnSlot(slot int)
}

type ParkingLot struct {
	MaximumSlot   int
	AvailableSlot int
	regNoSlotMap  map[string]int
	slotCarMap    map[int]models.Vehicle
	colorRegNoMap map[string][]string
}

func (p *ParkingLot) Create(capacity int) {

}

func (p *ParkingLot) Park(v *models.Vehicle) {

}

func (p *ParkingLot) Leave(slot int) {

}

func (p *ParkingLot) Status() {

}

func (p *ParkingLot) GetVehicleOnSlot(slot int) {

}

func (p *ParkingLot) ColorToVehicle(color string) []*models.Vehicle {
	return nil
}

func(p *ParkingLot)ColorToSlot(color string)[]int{
	return nil
}
