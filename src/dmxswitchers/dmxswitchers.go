package dmxswitchers

type dmxSwitchers struct {
	name string
}

func Create() (dmxSwitchers) {
	var dmx dmxSwitchers
	dmx.name = "Switchers"
	return dmx
}
