package util

import (
	"os"
)

const defaultResourceFolder string = "./resource/"
const defaultHtmlFolder string = "html/"

var ResourceFolder string = defaultResourceFolder
var HtmlFolder string = defaultHtmlFolder

func LoadTxtFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	filesize := info.Size()
	byteArray := make([]byte, filesize)
	_, err = file.Read(byteArray)
	return string(byteArray), err
}

func LoadHtmlFromResourceFolder(htmlname string) (string, error) {
	htmlContent, err := LoadTxtFile(ResourceFolder + HtmlFolder + htmlname)
	return htmlContent, err
}
