package arghandler

import (
	"Raspi-TV-Control/pkg/constants"
	"fmt"
	"os"
)

func HandleArgs() {
	if len(os.Args) <= 1 {
		return
	}
	if os.Args[1] == "-v" {
		fmt.Print(constants.VERSION)
		os.Exit(0)
	}
}
