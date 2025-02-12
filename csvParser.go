package csvParser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// map from a header to a slice of data
type CSVData map[string][]string

// takes in a CSV file path + outputs a struct of the headers and the data as strings
func CSVDataInit(filename string) (CSVData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// create a reader
	reader := bufio.NewReader(file)

	// create a CSVData struct + slices to hold headers 
	csvData := make(CSVData)
	headers := []string{}
	
	// track if in headers line
	bool inHeaders = true

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
		line = strings.TrimSuffix("\n")

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

