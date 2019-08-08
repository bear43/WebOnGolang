package handler

import (
	"awesomeProject/util"
	"fmt"
	"net/http"
)

var DebugMode bool

func DefaultHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}
	if DebugMode {
		fmt.Println("Request ", r.URL.Path)
	}
	htmlContent, err := util.LoadHtmlFromResourceFolder(r.URL.Path[1:] + ".html")
	if err != nil {
		http.Redirect(rw, r, "/error", 302)
	} else {
		fmt.Fprint(rw, htmlContent)
	}
}
