package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Record struct {
	Name		string
	Surname		string
	Number		string
	LastAcess	string
}

var myData []Record


func readCSVFile(filepath string) ([][]string, error) {
	// Check if file exist or not?
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	// This defer ensure f.Close() will be executed right before the end of main function
	defer f.Close()

	// Load all data in csv file at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = ','
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAcess}
		err := csvwriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Println("csvData input output")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Read CSV data 
	for _, line := range lines {
		temp := Record{
			Name: line[0],
			Surname: line[1],
			Number: line[2],
			LastAcess: line[3],
		}
		myData = append(myData, temp)
	}
	err = saveCSVFile(output)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}