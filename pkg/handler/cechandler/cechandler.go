package cechandler

import (
	"Raspi-TV-Control/pkg/constants/enums/powerstatus"
	"errors"
	"github.com/Binozo/cec.go"
	"time"
)

var (
	con           *cec.Connection
	address       int = 0
	tvIsConnected     = false
	maxRetries        = 10
)

func Init() error {
	connection, err := cec.Open("", "cec.go", false)
	if err != nil {
		return err
	}
	con = connection
	for _, device := range GetActiveDevices() {
		address = device.LogicalAddress // setting a default address
	}
	CheckIfTVIsConnectedAndSetup()
	return nil
}

// CheckIfTVIsConnectedAndSetup Checks if a TV is connected and sets the default address if so.
func CheckIfTVIsConnectedAndSetup() bool {
	for deviceType, device := range GetActiveDevices() {
		if deviceType == "TV" && device.LogicalAddress == 0 {
			address = device.LogicalAddress
			tvIsConnected = true
			return true
		}
	}
	return false
}

func GetLibInfo() string {
	return con.GetLibInfo()
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
