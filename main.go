package main

import (
	"fmt"
	_ "main/drivenadapters/repository"
	"main/driveradapters/cmd"
	"os"
)

func main() {
	if err := cmd.NewProcessor().Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
