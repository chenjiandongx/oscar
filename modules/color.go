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

func CyanBoldString(format string, a ...interface{}) string {
	return color.New(color.FgCyan).Add(color.Bold).SprintfFunc()(format, a...)
}

func GreenBoldString(format string, a ...interface{}) string {
	return color.New(color.FgGreen).Add(color.Bold).SprintfFunc()(format, a...)
}

func YellowBoldString(format string, a ...interface{}) string {
	return color.New(color.FgYellow).Add(color.Bold).SprintfFunc()(format, a...)
}

func MagentaString(format string, a ...interface{}) string {
	return color.New(color.FgMagenta).SprintfFunc()(format, a...)
}

func GreenString(format string, a ...interface{}) string {
	return color.New(color.FgGreen).SprintfFunc()(format, a...)
}

func CyanString(format string, a ...interface{}) string {
	return color.New(color.FgCyan).SprintfFunc()(format, a...)
}

func RedString(format string, a ...interface{}) string {
	return color.New(color.FgRed).SprintfFunc()(format, a...)
}

func BlackString(format string, a ...interface{}) string {
	return color.New(color.FgBlack).SprintfFunc()(format, a...)
}

func BlueString(format string, a ...interface{}) string {
	return color.New(color.FgBlue).SprintfFunc()(format, a...)
}

func BoldString(format string, a ...interface{}) string {
	return color.New(color.Bold).SprintfFunc()(format, a...)
}
