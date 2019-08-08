package util

import (
	"io/ioutil"
	"os"
)

const defaultResourceFolder string = "./resource/"
const defaultHtmlFolder string = "html/"
const defaultAssetsFolder string = "assets/"
const defaultJSFolder string = "js/"
const defaultCSSFolder string = "css/"

var ResourceFolder string = defaultResourceFolder
var HtmlFolder string = defaultHtmlFolder
var AssetsFolder string = defaultAssetsFolder
var JSFolder string = defaultJSFolder
var CSSFolder string = defaultCSSFolder

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

func TotalFilesInDirectory(directory string) ([]string, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func TotalHTMLPagesList() ([]string, error) {
	htmlFiles, err := TotalFilesInDirectory(ResourceFolder + HtmlFolder)
	return htmlFiles, err
}

func TotalJSScriptsList() ([]string, error) {
	htmlFiles, err := TotalFilesInDirectory(ResourceFolder + AssetsFolder + JSFolder)
	return htmlFiles, err
}

func TotalCSSFilesList() ([]string, error) {
	htmlFiles, err := TotalFilesInDirectory(ResourceFolder + AssetsFolder + CSSFolder)
	return htmlFiles, err
}
