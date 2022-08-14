package cechandler

import (
	"Raspi-TV-Control/pkg/constants/enums/powerstatus"
	"github.com/chbmuc/cec"
)

var (
	con     *cec.Connection
	address int = 0
)

func Init() error {
	connection, err := cec.Open("", "cec.go")
	if err != nil {
		return err
	}
	con = connection
	for _, device := range GetActiveDevices() {
		address = device.LogicalAddress // setting a default address
	}
	return nil
}

func GetActiveDevices() map[string]cec.Device {
	return con.List()
}

func GetPowerStatus() string {
	return con.GetDevicePowerStatus(address)
}

func SetPowerStatus(status string) error {
	if status == powerstatus.POWERSTATUS_ON {
		return con.PowerOn(address)
	} else if status == powerstatus.POWERSTATUS_STANDBY {
		return con.Standby(address)
	}
	return nil
}
