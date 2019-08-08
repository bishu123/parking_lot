package models

type Vehicle struct {
	LicenceNumber string
	Color string
}

func (v *Vehicle)GetLicenceNumber() string{
	return v.LicenceNumber
}

func (v *Vehicle)GetColor() string {
	return v.Color
}

func CreateVehicle(LicenceNumber, Color string) *Vehicle{
	return &Vehicle{LicenceNumber:LicenceNumber,Color:Color}
}