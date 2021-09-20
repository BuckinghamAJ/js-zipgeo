package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//	Look for the file zip-coordinates.csv and open it
	file, err := os.Open("zip-coordinates.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//	Read in the entire csv file
	log.Println("Reading in csv file ... ")
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//	Create a buffered writer
	log.Println("Creating ziptogeo.js ... ")
	f, err := os.Create("ziptogeo.js")
	if err != nil {
		log.Fatalf("There was a problem creating the output file: %v", err)
	}
	w := bufio.NewWriter(f)

	//	Add a line to declare the map
	_, err = w.WriteString("\n\nlet zipmap = new Map();\n\n")
	if err != nil {
		log.Fatalf("There was a problem writing to the output file: %v", err)
	}

	//	For each line in the CSV file
	for _, record := range records {
		//	Parse
		zip := record[0]
		lat := strings.TrimSpace(record[1])
		long := strings.TrimSpace(record[2])

		_, err = w.WriteString(fmt.Sprintf("zipmap.set('%s', {lat: %s, long: %s});\n", zip, lat, long))
		if err != nil {
			log.Fatalf("problem saving the zip/geo data: %v", err)
		}

		//	Print out the value
		fmt.Printf("Saving data for %s\n", zip)
	}

	//	Write our utility function and export lines:
	_, err = w.WriteString("\n\nfunction ZipToGeo(zipcode) {\n\treturn zipmap.get(zipcode);\n}\n\nexport default ZipToGeo;\n")
	if err != nil {
		log.Fatalf("There was a problem writing to the output file: %v", err)
	}

	//	Flush writes:
	log.Println("Finishing ... ")
	w.Flush()
	log.Println("Done.")
}
