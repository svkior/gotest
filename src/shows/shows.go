package shows

type shows struct {
	name string
}

func Create() (shows){
	var sh shows
	sh.name = "Shows"
	return sh
}
