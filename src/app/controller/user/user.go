package user

import (
	"app/model/user"
	"core/tpl"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func User(w http.ResponseWriter, req *http.Request) {
	log.Println("request /user")

	var content = tpl.GetTplContent("user")
	var indexT = template.Must(template.New("user").Parse(content))

	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		id = 0
	}

	u, err := user.GetUserById(id)
	indexT.Execute(w, u)
}
