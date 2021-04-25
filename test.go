package main

//
//import (
//	"flag"
//	"golang.org/x/sys/windows"
//	"golang.org/x/sys/windows/svc/mgr"
//	"log"
//	"net/http"
//)
//
////var flagServiceInstall = flag.Bool("install", false, "Install service")
//
//func main() {
//
//}
//
//func httpService() {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		writer.Write([]byte("<h1>Hello world</h1>"))
//	})
//
//	_ = http.ListenAndServe("0.0.0.0:8080", nil)
//}
//
//func install() {
//	m, err := mgr.Connect()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	defer m.Disconnect()
//
//	s, err := m.CreateService("Service name", "program path",
//		mgr.Config{
//			ServiceType:      windows.SERVICE_WIN32_OWN_PROCESS,
//			StartType:        windows.SERVICE_AUTO_START,
//			ServiceStartName: "Service name",
//			DisplayName:      "Service alias name",
//			Description:      "Service Description",
//			DelayedAutoStart: false,
//		},
//		"-t -r ....",
//	)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	s.Close()
//}