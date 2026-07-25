package handler

import (
	"bufio"
	"dealerships-cli/model"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)

}

func ListDealerships(region string) {
	dealerships, err := model.LoadDealerships(region)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(dealerships) == 0 {
		fmt.Printf("No dealer ships found in region %s.\n", region)
		return
	}
	fmt.Printf("=== Dealerships in %s ===\n", region)
	for _, d := range dealerships {
		fmt.Printf("[%d] %s — %s — %s\n", d.ID, d.Name, d.Address, d.Phone)
	}
}
func GetDealership(region string) {
	dealerships, err := model.LoadDealerships(region)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print("Please enter your dealership ID: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	for _, d := range dealerships {
		if d.ID == id {
			fmt.Println("==================")
			fmt.Printf("ID:              %d\n", d.ID)
			fmt.Printf("Name:            %s\n", d.Name)
			fmt.Printf("Address:         %s\n", d.Address)
			fmt.Printf("Phone:           %s\n", d.Phone)
			fmt.Printf("Membership Date: %s\n", d.MembershipDate)
			fmt.Printf("Employees:       %d\n", d.EmployeeCount)
			fmt.Println("==================")
			return
		}
	}
	fmt.Printf("No dealership found with ID %d in %s\n", id, region)

}
func StatusDealerships(region string) {
	dealerships, err := model.LoadDealerships(region)
	if err != nil {
		log.Fatal(err)
		return
	}
	totalEmployees := 0
	for _, d := range dealerships {
		totalEmployees += d.EmployeeCount
	}
	fmt.Printf("=== Status for %s ===\n", region)
	fmt.Printf("Total dealerships: %d\n", len(dealerships))
	fmt.Printf("Total employees:   %d\n", totalEmployees)
}
func CreateDealership(region string) {
	dealerships, err := model.LoadDealerships(region)
	if err != nil {
		log.Fatal(err)
		return
	}
	nextID := 1
	for _, d := range dealerships {
		if d.ID >= nextID {
			nextID = d.ID + 1
		}
	}
	name := ReadInput("Enter your name:")
	phone := ReadInput("Enter your phone:")
	address := ReadInput("Enter your address:")
	date := ReadInput("Enter your dealership date:")
	empCount := ReadInput("Enter your dealership employee count:")

	empCountInt, err := strconv.Atoi(empCount)
	if err != nil {
		log.Fatal(err)
		return
	}
	newDealership := model.Dealership{
		ID:             nextID,
		Name:           name,
		Address:        address,
		Phone:          phone,
		MembershipDate: date,
		EmployeeCount:  empCountInt,
	}
	dealerships = append(dealerships, newDealership)
	if err := model.SaveDealerships(region, dealerships); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("✅ Dealership created with ID %d in %s\n", nextID, region)
}
func EditDealership(region string) {
	dealerships, err := model.LoadDealerships(region)
	if err != nil {
		log.Fatal(err)
		return
	}
	idStr := ReadInput("Enter your ID:")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	index := -1
	for i, d := range dealerships {
		if d.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("No dealership found with ID " + idStr)
		return
	}
	d := dealerships[index]
	fmt.Printf("Editing dealership [%d] %s\n", d.ID, d.Name)
	fmt.Println("Press Enter to keep current value")
	if input := ReadInput(fmt.Sprintf("Name (%s):", d.Name)); input != "" {
		d.Name = input
	}
	if input := ReadInput(fmt.Sprintf("Address (%s):", d.Address)); input != "" {
		d.Address = input
	}
	if input := ReadInput(fmt.Sprintf("Phone (%s):", d.Phone)); input != "" {
		d.Phone = input
	}
	if input := ReadInput(fmt.Sprintf("Name (%s):", d.MembershipDate)); input != "" {
		d.MembershipDate = input
	}
	if input := ReadInput(fmt.Sprintf("Employees count (%d):", d.EmployeeCount)); input != "" {
		empCount, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
			return
		}
		d.EmployeeCount = empCount
	}
	dealerships[index] = d
	if err := model.SaveDealerships(region, dealerships); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("✅ Dealership %d updated successfully\n", id)
}
