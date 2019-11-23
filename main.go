package main

import (
	"fmt"
	"os"

	"github.com/qorbani/mysql-toolkit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
