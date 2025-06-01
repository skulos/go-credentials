package colours

import (
	"github.com/fatih/color"
)

var (
	KeyColor   = color.New(color.FgBlue).SprintFunc()
	ValueColor = color.New(color.FgGreen).SprintFunc()
	WarnColor  = color.New(color.FgYellow).SprintFunc()
	ErrorColor = color.New(color.FgRed).SprintFunc()
)
