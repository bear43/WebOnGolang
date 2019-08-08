package main

import (
	"awesomeProject/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting...")
	handler.DebugMode = true
	http.HandleFunc("/", handler.DefaultHandler)
	http.HandleFunc("/error", handler.ErrorHandler)
	http.HandleFunc("/register", handler.DefaultHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./resource/assets/"))))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
