package park

import (
	"parking_lot/models"
)

func (p *ParkingLot) isParkingAvailable() bool {
	if p.availableSlot == 0 {
		return false
	}
	return true
}

func (p *ParkingLot) putColorToVehicleMap(slot int, vehicle *models.Vehicle) {
	color := vehicle.Color
	licenceNo := vehicle.LicenceNumber
	if vehicleMap, ok := p.colorRegNoMap[color]; !ok {
		p.colorRegNoMap[color] = map[int]string{slot: licenceNo}
	} else {
		vehicleMap[slot] = licenceNo
		p.colorRegNoMap[color] = vehicleMap
	}
}

func (p *ParkingLot) removeColorToVehicleMap(slot int) {
	if vehicle, _ := p.slotVehicleMap[slot]; vehicle != nil {
		color := vehicle.Color
		slotRegMap := p.colorRegNoMap[color]
		delete(slotRegMap, slot)
	}
}

func (p *ParkingLot) getVehicleOnSlot(slot int) *models.Vehicle {
	return p.slotVehicleMap[slot]
}
