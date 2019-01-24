package main

import (
	"core/config"
	"core/tpl"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("LoadConfig:", err)
	}

	var port = config.Config().Global.Port
	var portStr = ":" + strconv.Itoa(port)

	http.Handle("/favicon.ico", http.HandlerFunc(Ico))
	http.Handle("/", http.HandlerFunc(Index))
	http.Handle("/qr", http.HandlerFunc(QR))

	log.Println("server started.")
	err = http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Ico(w http.ResponseWriter, req *http.Request) {
	log.Println("request /favicon.ico")
	w.Write(nil)
}

func Index(w http.ResponseWriter, req *http.Request) {
	log.Println("request /")

	var content = tpl.GetTplContent("index.html")
	var indexT = template.Must(template.New("index").Parse(content))
	indexT.Execute(w, nil)
}

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println("request /qr")

	var content = tpl.GetTplContent("qr.html")
	var qrT = template.Must(template.New("qr").Parse(content))
	qrT.Execute(w, nil)
}
