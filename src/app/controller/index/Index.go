package index

import (
	"core/tpl"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	log.Println("request /")

	var content = tpl.GetTplContent("index")
	var indexT = template.Must(template.New("index").Parse(content))
	indexT.Execute(w, nil)
}

func Ico(w http.ResponseWriter, req *http.Request) {
	log.Println("request /favicon.ico")
	w.Write(nil)
}
