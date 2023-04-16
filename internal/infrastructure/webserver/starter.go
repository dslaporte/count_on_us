package webserver

type WebServerStarter struct {
	webServer WebServerInterface
}

func NewWebServerStarter(webServer WebServerInterface) *WebServerStarter {
	return &WebServerStarter{
		webServer: webServer,
	}
}

func (w *WebServerStarter) Start() {
	w.webServer.Start()
}
