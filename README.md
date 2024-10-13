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
{
 "houses": [
  {
   "value": 452600,
   "income": 8.3252,
   "age": 41,
   "rooms": 880,
   "bedrooms": 129,
   "pop": 322,
   "hh": 126
  },
  {
   "value": 358500,
   "income": 8.3014,
   "age": 21,
   "rooms": 7099,
   "bedrooms": 1106,
   "pop": 2401,
   "hh": 1138
  },
  {
   "value": 352100,
   "income": 7.2574,
   "age": 52,
   "rooms": 1467,
   "bedrooms": 190,
   "pop": 496,
   "hh": 177
  }
 ]
}
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