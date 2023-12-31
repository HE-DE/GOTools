package utils

import (
	"github.com/schollz/progressbar/v3"
)

func CreateBar(max int) *progressbar.ProgressBar {
	bar := progressbar.NewOptions(max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(25),
		progressbar.OptionSetDescription("[cyan][reset]正在检查..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	return bar
}
