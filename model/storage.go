package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func GetFilePath(region string) string {
	return "data/" + region + ".csv"
}

func LoadDealerships(region string) ([]Dealership, error) {
	path := GetFilePath(region)

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Dealership{}, nil
		}
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var dealerships []Dealership

	for _, row := range records {
		id, _ := strconv.Atoi(row[0])
		empCount, _ := strconv.Atoi(row[5])
		dealerships = append(dealerships, Dealership{
			ID:             id,
			Name:           row[1],
			Address:        row[2],
			Phone:          row[3],
			MembershipDate: row[4],
			EmployeeCount:  empCount,
		})
	}
	return dealerships, nil
}
func SaveDealerships(region string, dealerships []Dealership) error {
	path := GetFilePath(region)
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, d := range dealerships {
		row := []string{
			strconv.Itoa(d.ID),
			d.Name,
			d.Address,
			d.Phone,
			d.MembershipDate,
			strconv.Itoa(d.EmployeeCount),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
