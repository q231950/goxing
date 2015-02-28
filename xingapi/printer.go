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
	color.Printf("c", fmt.Sprintf("%s", command))
}

func  PrintMessageWithParam(message string, param string) {
	print(message)
	color.Print("m", fmt.Sprintf("%s%s", param, "\n"))
}

func PrintUser(user User) {
	color.Printf("", fmt.Sprintf("-----------------------------------\n%s <%s>:\n", user.DisplayName, user.Id))
	color.Printf("d", fmt.Sprintf("Email address:\t\t%s\nDate of birth:\t\t%s\n", user.ActiveEmail, user.Birthdate))
}

func PrintUserOneLine(user User) {
	color.Printf("m", fmt.Sprintf("%s\n", user.DisplayName))
}