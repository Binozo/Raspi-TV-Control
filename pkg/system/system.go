package system

import (
	"Raspi-TV-Control/pkg/constants"
	"Raspi-TV-Control/pkg/handler/cechandler"
	"fmt"
	"runtime"
	"time"
)

var (
	uptime time.Time
)

func Init() {
	uptime = time.Now()
}

func GetUptimeDate() time.Time {
	return uptime
}

func GetFormattedUptimeDate() string {
	return uptime.Format("2006-01-02 15:04:05")
}

func GetUptime() time.Duration {
	return time.Since(uptime)
}

func GetInfo() map[string]string {
	note := "Target TV has been detected"
	if !cechandler.CheckIfTVIsConnectedAndSetup() {
		note = "Couldn't detect any TV"
	}
	libInfo := cechandler.GetLibInfo()
	return map[string]string{
		"OS":          runtime.GOOS,
		"Version":     constants.VERSION,
		"Uptime_Date": GetFormattedUptimeDate(),
		"Uptime":      fmt.Sprintf("%f", GetUptime().Seconds()),
		"Note":        note,
		"LibInfo":     libInfo,
	}
}
