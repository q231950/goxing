// printer.go 

package xingapi 

import (
	"github.com/str1ngs/ansi/color"
	"fmt"
	"net/http"
	)

type Printer struct {}

func PrintResponse(response *http.Response) {
	var colorCode string
	if (response.StatusCode == 200) {
		colorCode = "g"
	} else {
		colorCode = "r"
	}
	color.Printf(colorCode, fmt.Sprintf("%s\n", response.Status))
}

func PrintCommand(command string) {
	color.Printf("c", fmt.Sprintf("GET %s", command))
}