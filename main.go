package main

import (
	"os"

	"github.com/JosipVodafone/userbeat/cmd"

	_ "github.com/JosipVodafone/userbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
