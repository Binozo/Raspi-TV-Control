package main

import (
	"Raspi-TV-Control/pkg/handler/apihandler"
	"Raspi-TV-Control/pkg/system"
)

func main() {
	system.Init()
	apihandler.RegisterRoutes()
}
