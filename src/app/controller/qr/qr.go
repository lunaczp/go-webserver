package qr

import (
	"core/tpl"
	"html/template"
	"log"
	"net/http"
)

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println("request /qr")

	var content = tpl.GetTplContent("qr")
	var qrT = template.Must(template.New("qr").Parse(content))
	qrT.Execute(w, req.FormValue("s"))
}
