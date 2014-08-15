package dmxconsoles

type dmxConsoles struct {
	name string
}

func Create() (dmxConsoles){
	var dmx dmxConsoles
	dmx.name = "Consoles"
	return dmx
}
