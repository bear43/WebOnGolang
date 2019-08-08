package main

import (
	"awesomeProject/handler"
	"awesomeProject/util"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting...")
	handler.DebugMode = true
	util.PrintTotalAssets()
	http.HandleFunc("/", handler.DefaultHandler)
	http.HandleFunc("/error", handler.ErrorHandler)
	http.HandleFunc("/register", handler.DefaultHandler)
	http.HandleFunc("/login", handler.DefaultHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./resource/assets/"))))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
