package cechandler

import (
	"Raspi-TV-Control/pkg/constants/enums/powerstatus"
	"errors"
	"github.com/chbmuc/cec"
	"time"
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

func VolumeUp() error {
	return con.VolumeUp()
}

func VolumeDown() error {
	return con.VolumeDown()
}

func Mute() error {
	return con.Mute()
}

func SendKey(keyName string) error {
	key := cec.GetKeyCodeByName(keyName)
	if key == -1 {
		return errors.New("invalid key name")
	}
	err := con.KeyPress(address, key)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 10)
	err = con.KeyRelease(address)
	return err
}
