package dmxscripts

type dmxScripts struct {
	name string
}

func Create() (dmxScripts){
	var rv dmxScripts
	rv.name = "Scripts"
	return rv
}
