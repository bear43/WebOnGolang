package handler

import (
	"awesomeProject/util"
	"fmt"
	"net/http"
)

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	htmlContent, err := util.LoadHtmlFromResourceFolder("index.html")
	if err != nil {
		http.Redirect(rw, r, "/error", 302)
	} else {
		fmt.Fprint(rw, htmlContent)
	}
}
