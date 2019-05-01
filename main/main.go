package main

import (
	"footmap/app/controller"
	"footmap/app/utils"
	"net/http"
)


func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	http.HandleFunc("/user/",controller.Verify)
	http.HandleFunc("/record/",controller.Verify)

	err := server.ListenAndServe()
	utils.CheckError(err)
}