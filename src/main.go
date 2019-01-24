package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var addr = flag.String("addr", ":9090", "http service address")

const tplDiv = "src/view/"


func main() {
	flag.Parse()
	http.Handle("/favicon.ico", http.HandlerFunc(Ico))
	http.Handle("/", http.HandlerFunc(Index))
	http.Handle("/qr", http.HandlerFunc(QR))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Ico(w http.ResponseWriter, req *http.Request) {
	log.Println("request /favicon.ico")
	w.Write(nil);
}

func Index(w http.ResponseWriter, req *http.Request) {
	log.Println("request /")

	var content = getTplContent("index.html");
	var indexT = template.Must(template.New("index").Parse(content))
	indexT.Execute(w, nil)
}

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println("request /qr")

	var content = getTplContent("qr.html");
	var qrT = template.Must(template.New("qr").Parse(content))
	qrT.Execute(w, nil)
}

func getTplContent(filename string) string {
	filename, _ = filepath.Abs(tplDiv  + filename);
	log.Println(filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

