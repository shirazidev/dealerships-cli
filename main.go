package main

import (
	"bufio"
	"dealerships-cli/handler"
	"flag"
	"fmt"
	"os"
)

func main() {
	command := flag.String("command", "", "commands to run: list, get, create, edit, status")
	region := flag.String("region", "", "Region name e.g. tehran, isfahan, shiraz")
	flag.Parse()

	if *command != "" && *region != "" {
		runCommand(*command, *region)
	} else {
		for {

			fmt.Println("Enter command (list, get, create, edit, status) or 'exit' to quit:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			*command = scanner.Text()
			if *command == "exit" || *command == "quit" {
				break
			}
			fmt.Println("Enter region (e.g., tehran, isfahan, shiraz):")
			scanner.Scan()
			*region = scanner.Text()

			runCommand(*command, *region)

		}
	}

}
func runCommand(command string, region string) {
	switch command {
	case "list":
		handler.ListDealerships(region)
	case "get":
		handler.GetDealership(region)
	case "create":
		handler.CreateDealership(region)
	case "edit":
		handler.EditDealership(region)
	case "status":
		handler.StatusDealerships(region)
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
