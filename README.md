# CSV to JSON Converter using Go

This Go program converts comma-separated values (csv) to JavaScript Object Notation (JSON)

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Code Explanation](#code-explanation)

## Introduction

This CSV to JSON converter is created for MSDS-431 assignment. It is made to specifically convert housing data from a CSV format into JSON Lines format, which is especially useful for large datasets, as each line is it's own JSON object.

The program works with a specific dataset structure. Example CSV file can be found [here](housesInput.csv)

### Description of the dataset

| Column Name  | Data Type | Description                                                 |
|--------------|-----------|-------------------------------------------------------------|
| Value        | Integer   | The median value of the houses in the area in USD           |
| Income       | Float     | The median income of the area in tens of thousands in USD   |
| Age          | Integer   | The median age of the houses in the area in years           |
| Rooms        | Integer   | The total number of rooms in the area                       |
| Bedrooms     | Integer   | The total number of bedrooms in the area                    |
| Population   | Integer   | The population of the area                                  |
| Households   | Integer   | The number of households in the area                        |

### Example CSV data and JSON Lines file

Here is an example data from the [housing dataset](housesInput.csv):

| Value  | Income | Age | Rooms | Bedrooms | Population | Households |
|--------|--------|-----|-------|----------|------------|------------|
| 452600 | 8.3252 | 41  | 880   | 129      | 322        | 126        |
| 358500 | 8.3014 | 21  | 7099  | 1106     | 2401       | 1138       |
| 352100 | 7.2574 | 52  | 1467  | 190      | 496        | 177        |

Running the program will return a file that looks like:

```json
{"value":452600,"income":8.3252,"age":41,"rooms":880,"bedrooms":129,"pop":322,"hh":126}
{"value":358500,"income":8.3014,"age":21,"rooms":7099,"bedrooms":1106,"pop":2401,"hh":1138}
{"value":352100,"income":7.2574,"age":52,"rooms":1467,"bedrooms":190,"pop":496,"hh":177}
...
```

## Features

- Converts CSV data to JSON Lines format.
- Validate CSV headers to make sure the input is correct.
- Deals with data types like floats and integers.
- Wraps the JSON in a top-level object to ensure valid JSON output.
- Generates an output file with structured JSON data.

## Installation

1. Make sure you have [Go installed](https://go.dev/doc/install).
2. Clone this repo to your local machine:
    ```bash
    git clone https://github.com/hamodikk/csvToJlConverter.git
    ```
3. Navigate to the project directory
    ```bash
    cd <project-directory>
    ```

## Usage

Use the following command in your terminal or Powershell to run the program:
```bash
.\csvtojl.exe <input.csv> <output.jl>
```

Example Usage:
```bash
.\csvtojl.exe .\housesInput.csv .\housesOutput.jl
```

### Code Explanation

1. Import dependencies. We need these libraries for the program to function correctly.
```go
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
```
2. Create a function that will handle errors.
```go
func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
```
3. Create the struct that will hold individual house data.
    Make sure the types are written correctly.
    ```go
    type House struct {
	Value      int     `json:"value"`
	Income     float64 `json:"income"`
	Age        int     `json:"age"`
	Rooms      int     `json:"rooms"`
	Bedrooms   int     `json:"bedrooms"`
	Population int     `json:"pop"`
	Households int     `json:"hh"`
    }
    ```
4. Create a function that checks the input file headers.
```go
func validateHeaders(headers []string) bool {
	expectedHeaders := []string{
		"value", "income", "age", "rooms", "bedrooms", "pop", "hh",
	}
	return strings.Join(headers, ",") == strings.Join(expectedHeaders, ",")
}
```
5. Create the main function
- Check for correct use of the command-line arguments.
```go
func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: csvtojl <input.csv> <output.jl>")
	}
```
- Assigning input and output arguments allow us to run the program from the command line with our input file.
```go
    inputFile := os.Args[1]
    outputFile := os.Args[2]
```
- Open the input file and create the output file.
```go
    csvFile, err := os.Open(inputFile)
    checkError(err)
    defer csvFile.Close()

    jsonFile, err := os.Create(outputFile)
    checkError(err)
    defer jsonFile.Close()
```
- Read the headers and check headers for correctness
```go
    // Init csv reader
    reader := csv.NewReader(csvFile)

    // Read the headers
    header, err := reader.Read()
    checkError(err)

    // Check that the headers are correct
    if !validateHeaders(header) {
        log.Fatal("Unexpected CSV headers.")
    }
```
- Read the data
```go
    records, err := reader.ReadAll()
    checkError(err)
```
- Loop through the read data and type it, followed by extracting the data in a way that fits the House struct. Lastly, append the struct to the houses slice.
```go
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
```
- Convert the response object to JSON and write it to the output file.
```go
		// Convert the House object to JSON
		jsonData, err := json.Marshal(house)
		checkError(err)

		// Write the JSON to the file, followed by a new line to create the jl format
		fmt.Fprintln(jsonFile, string(jsonData))
	}
	fmt.Printf("File converted successfully! JSON lines in %s\n", outputFile)
```