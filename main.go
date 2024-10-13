package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Create the top-level structure for JSON output
type Response struct {
	Houses []House `json:"houses"`
}

// Helper function that checks for errors, terminate the program if an error occurs
func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// Struct representing a single house record
type House struct {
	Value      int     `json:"value"`
	Income     float64 `json:"income"`
	Age        int     `json:"age"`
	Rooms      int     `json:"rooms"`
	Bedrooms   int     `json:"bedrooms"`
	Population int     `json:"pop"`
	Households int     `json:"hh"`
}

// Helper function that makes sure the headers are in the csv file and correct
func validateHeaders(headers []string) bool {
	expectedHeaders := []string{
		"value", "income", "age", "rooms", "bedrooms", "pop", "hh",
	}
	return strings.Join(headers, ",") == strings.Join(expectedHeaders, ",")
}

func main() {
	// Check for correct use of the command-line arguments
	if len(os.Args) < 3 {
		log.Fatal("Usage: csvtojl <input.csv> <output.jl>")
	}

	// Assigns input and output arguments
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open the input file
	csvFile, err := os.Open(inputFile)
	checkError(err)
	defer csvFile.Close()

	// Create the output file
	jsonFile, err := os.Create(outputFile)
	checkError(err)
	defer jsonFile.Close()

	// Init csv reader
	reader := csv.NewReader(csvFile)

	// Read the headers
	header, err := reader.Read()
	checkError(err)

	// Check that the headers are correct
	if !validateHeaders(header) {
		log.Fatal("Unexpected CSV headers.")
	}

	// Read the rest of the rows
	records, err := reader.ReadAll()
	checkError(err)

	var houses []House

	// Extract the data line by line and fit it on to the struct
	for _, line := range records {
		// Convert data to their corresponding types
		value, err := strconv.ParseFloat(line[0], 64)
		checkError(err)

		income, err := strconv.ParseFloat(line[1], 64)
		checkError(err)

		age, err := strconv.Atoi(line[2])
		checkError(err)

		rooms, err := strconv.Atoi(line[3])
		checkError(err)

		bedrooms, err := strconv.Atoi(line[4])
		checkError(err)

		population, err := strconv.Atoi(line[5])
		checkError(err)

		households, err := strconv.Atoi(line[6])
		checkError(err)

		// Extract the data in the form of the struct
		house := House{
			Value:      int(value),
			Income:     income,
			Age:        int(age),
			Rooms:      int(rooms),
			Bedrooms:   int(bedrooms),
			Population: int(population),
			Households: int(households),
		}

		// Append the struct to the slice
		houses = append(houses, house)
	}

	// Create the top-level response object
	response := Response{Houses: houses}

	// Convert the response to JSON
	jsonData, err := json.MarshalIndent(response, "", " ")
	checkError(err)

	// Write the JSON to the output file
	_, err = jsonFile.WriteString(string(jsonData))
	checkError(err)

	fmt.Printf("File converted successfully! JSON lines in %s\n", outputFile)
}
