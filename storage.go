package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func getFilePath(region string) string {
	return "data/" + region + ".csv"
}

func loadDealerships(region string) ([]Dealership, error) {
	path := getFilePath(region)

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Dealership{}, nil
		}
		return nil, err
	}
	defer file.Close()

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
func saveDealerships(region string, dealerships []Dealership) error {
	path := getFilePath(region)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

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
