package main

import (
	"fmt"
	_ "main/drivenadapters/repository"
	_ "main/driveradapters/async"
	"main/driveradapters/cmd"
	"os"
)

func main() {
	if err := cmd.NewProcessor().Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
