package handler

import (
	"awesomeProject/util"
	"fmt"
	"net/http"
)

func ErrorHandler(rw http.ResponseWriter, r *http.Request) {
	htmlContent, _ := util.LoadHtmlFromResourceFolder("error.html")
	fmt.Fprint(rw, htmlContent)
}
