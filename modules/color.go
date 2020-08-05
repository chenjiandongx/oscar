package modules

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintGreenWithBold(a string, others ...interface{}) {
	color.New(color.FgGreen).Add(color.Bold).PrintFunc()(a)
	fmt.Println(others...)
}

func PrintGreen(format string, a ...interface{}) {
	color.New(color.FgGreen).PrintfFunc()(fmt.Sprintf(format, a...))
}

func PrintYellow(format string, a ...interface{}) {
	color.New(color.FgYellow).PrintfFunc()(fmt.Sprintf(format, a...))
}

func GreenBoldString(format string, a ...interface{}) string {
	return color.New(color.FgGreen).Add(color.Bold).SprintfFunc()(format, a...)
}

func YellowBoldString(format string, a ...interface{}) string {
	return color.New(color.FgYellow).Add(color.Bold).SprintfFunc()(format, a...)
}

func RedString(format string, a ...interface{}) string {
	return color.New(color.FgRed).SprintfFunc()(format, a...)
}

func BoldString(format string, a ...interface{}) string {
	return color.New(color.Bold).SprintfFunc()(format, a...)
}
