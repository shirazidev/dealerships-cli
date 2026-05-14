package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	command := flag.String("command", "", "commands to run: list, get, create, edit, status")
	region := flag.String("region", "", "Region name e.g. tehran, isfahan, shiraz")
	flag.Parse()

	if *command == "" || *region == "" {
		fmt.Println("Usage: ./dealership-cli -command=<cmd> -region=<region>")
		fmt.Println("Commands: list, get, create, edit, status")
		os.Exit(1)
	}

	switch *command {
	case "list":
		listDealerships(*region)
	case "get":
		getDealership(*region)
	case "create":
		createDealership(*region)
	case "edit":
		editDealership(*region)
	case "status":
		statusDealerships(*region)
	default:
		fmt.Printf("Unknown command: %s\n", *command)
		os.Exit(1)
	}
}
