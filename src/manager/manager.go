package manager

type sysManager struct {
	name string
	sysLevel uint8
}

func Create() (sysManager){
	var sm sysManager
	sm.name = "System Manager"
	sm.sysLevel = 0
	return sm
}
