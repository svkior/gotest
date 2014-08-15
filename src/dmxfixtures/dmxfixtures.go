package dmxfixtures

type dmxFixtures struct {
	name string
}

func Create() (dmxFixtures){
	var dmx dmxFixtures
	dmx.name = "Fixtures"
	return dmx
}
