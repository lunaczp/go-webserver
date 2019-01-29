package main

import (
	"app/controller/index"
	"app/controller/qr"
	"app/controller/user"
	"core/config"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//load config
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("LoadConfig:", err)
	}

	var port = config.Config().Global.Port
	var portStr = ":" + strconv.Itoa(port)

	//setup route
	setupRoute()

	//start server
	log.Println("server started.")
	err = http.ListenAndServe(portStr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func setupRoute() {
	http.Handle("/favicon.ico", http.HandlerFunc(index.Ico))
	http.Handle("/", http.HandlerFunc(index.Index))
	http.Handle("/qr", http.HandlerFunc(qr.QR))
	http.Handle("/user", http.HandlerFunc(user.User))
}
