package utils

import (
	"net/url"

	"github.com/common-nighthawk/go-figure"
)

var Version string = "1.0.0"

func PrintBanner() {
	ascii := figure.NewFigure("go-diydns", "", true)
	ascii.Print()

	println("\t\t\t\t\t\t\t" + Version + "\n")
}

func GetHostname(fullDomain string) string {
	u, err := url.Parse(fullDomain)
	if err != nil {
		panic(err)
	}

	return u.Hostname()
}
