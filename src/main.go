package main

import (
	"core/config"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
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

	var content = getTplContent("index.html")
	var indexT = template.Must(template.New("index").Parse(content))
	indexT.Execute(w, nil)
}

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println("request /qr")

	var content = getTplContent("qr.html")
	var qrT = template.Must(template.New("qr").Parse(content))
	qrT.Execute(w, nil)
}

func getTplContent(filename string) string {
	var tplDir = config.Config().Global.TplDir
	filename, _ = filepath.Abs(tplDir + filename)
	log.Println(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
