package util

import "fmt"

func PrintStringSlice(slice []string) {
	for _, element := range slice {
		fmt.Println(element)
	}
}

func PrintTotalAssets() {
	slice, _ := TotalHTMLPagesList()
	fmt.Println("<===>Total HTML pages count:", len(slice))
	PrintStringSlice(slice)
	slice, _ = TotalCSSFilesList()
	fmt.Println("<===>Total JS count:", len(slice))
	PrintStringSlice(slice)
	slice, _ = TotalJSScriptsList()
	fmt.Println("<===>Total CSS count:", len(slice))
	PrintStringSlice(slice)
}
