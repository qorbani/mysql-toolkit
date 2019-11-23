package main

import (
	"fmt"
)

var (
	// Version will inject by Makefile
	Version = "0.0.0"
	// Build will inject by Makefile
	Build = "N/A"
	// BuildTime will inject by Makefile
	BuildTime = "Now"
)

func main() {
	fmt.Printf("MySQL Toolkit v%s (%s)\nBuild Time: %s\n", Version, Build, BuildTime)
}
