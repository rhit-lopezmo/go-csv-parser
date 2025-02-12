package csvParser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// map from a header to a slice of data
type CSVData map[string][]string

func (csvData CSVData) GetEntry(pos int) []string {
	entry := []string{}

	for _, val := range csvData {
		entry = append(entry, val[pos])
	}

	return entry
}

// takes in a CSV file path + outputs a struct of the headers and the data as strings
func CSVDataInit(filename string) (CSVData, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", filename)
		return nil, err
	}

	// create a reader
	reader := bufio.NewReader(file)

	// create a CSVData struct + slices to hold headers 
	csvData := make(CSVData)
	headers := []string{}
	
	// track if in headers line
	inHeaders := true

	for {
		line, err := reader.ReadString('\n')
		// handle errors
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Printf("Error reading from file: %s", err)
				return nil, err
			}
		}
		
		// trim new line char
		line = strings.TrimSuffix(line, "\n")

		// parse headers or parse data
		if inHeaders {
			headers = strings.Split(line, ",")
			inHeaders = false
		} else {
			currData := strings.Split(line, ",")
			for i, item := range currData {
				// get curr header and append the data to that header's slice (or creates a new one if nil)
				currHeader := headers[i]
				csvData[currHeader] = append(csvData[currHeader], item)
			}
		}
	}
		
	return csvData, nil
}

