package main

import (
//	"log"

	"dmxouts"
//	"dmxins"
//	"dmxscripts"
//	"dmxfixtures"
//	"dmxconsoles"
//	"dmxswitchers"
//	"webserver"
	"cli"
//	"setups"
//	"shows"
//	"udpresolver"
//	"manager"
)


func main(){
	// DMX Outputs
	dmxo := dmxouts.Create()
	// DMX Inputs
//	dmxi := dmxins.Create()
	// DMX Scripts
//	dmxs := dmxscripts.Create()
	// DMX Fixtures
//	dmxf := dmxfixtures.Create()
	// DMX Controllers
//	dmxc := dmxconsoles.Create()
	// DMX Switchers
//	dmxsw := dmxswitchers.Create()

	// Setups
//	setUps := setups.Create()
	//Shows
//	sh := shows.Create()

	// WebServer
//	webs := webserver.CreateWebServer()
	// UDP Resolver
//	ur := udpresolver.Create()
	// Console Client
	cl := cli.Create()

	// System Manager
//	sm := manager.Create()


	dmxo.Run()		// Run DMX Outputs

	cl.RunningLoop()
}
