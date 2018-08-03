package main

import (
	"fmt"
	"os"

	"github.com/JosipVodafone/userbeat/cmd"

	_ "github.com/JosipVodafone/userbeat/include"
)

func main() {
	fmt.Print("----------------------  Jove Starting User Beat Count beat MAIN-------------------------------")
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
