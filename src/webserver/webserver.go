package webserver

type webServer struct {
	name string
	port uint
}

func CreateWebServer() (webServer){
	var web webServer
	web.name = "Web Server"
	web.port = 3000
	return web
}
