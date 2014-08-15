package udpresolver

type udpResolver struct {
	name string
	port uint
}

func Create() (udpResolver) {
	var udr udpResolver
	udr.name = "UDP Resolver"
	udr.port = 4321
	return udr
}
