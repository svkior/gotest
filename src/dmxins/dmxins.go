package dmxins

/*
 * Структура будет использоваться для работы со входами DMX
 */
type dmxIns struct {
	name string
}

func Create() (dmxIns){
	var rv dmxIns
	rv.name = "dmxIns"
	return rv
}
