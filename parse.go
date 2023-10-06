package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseFileStl(path string) ([]Facet, error) {

	// read file
    file, err := os.Open(path)
    if err != nil {
		return nil, err
    }

	// close file when finish function
    defer file.Close()

	// scanner de file
	scanner := bufio.NewScanner(file)

	// variable for final result
	var arr []Facet

	// variable for manipule facet before save in arr
	var element Facet

	// variable for identify line of vertex
	index := 0

	// loop of lines of file
    for scanner.Scan() {

		// get line of text
        line := strings.TrimSpace(scanner.Text())

		// verify if line is begin of facet
		if strings.HasPrefix(line, "outer loop") {
			index = 0
		}

		// verify if line is one of three of vertex
        if strings.HasPrefix(line, "vertex") {
            parts := strings.Fields(line)
            x, _ := strconv.ParseFloat(parts[1], 64)
            y, _ := strconv.ParseFloat(parts[2], 64)
            z, _ := strconv.ParseFloat(parts[3], 64)

			// create vertex with three values
			vertex := Vertex{x, y, z}

			// verify which line of vertex is reading
			switch index {
				case 0: element.V1 = vertex 
				case 1: element.V2 = vertex 
				case 2: element.V3 = vertex 
				default: panic("More than 3 lines of vertex!")
			}
			index = index + 1
        }

		// verify if line is end of facet, then save in arr
		if strings.HasPrefix(line, "endloop") {
			arr = append(arr, element)
		}	
    }

    return arr, nil;
}