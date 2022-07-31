package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

// Struct that holds power and timestamp data read from GPX file.
type PowerAndTimestamp struct {
	power     int
	timestamp time.Time
}

// An array holding PowerAndTimestamp values.
var powerArray []PowerAndTimestamp

// Integer that holds the maxAvgPower while going over the sliding window.
var maxAvgPower int

// Calculate the average power of all objects in the powerArray.
func getAvgPower(powerArray []PowerAndTimestamp) int {
	sum := 0
	for _, item := range powerArray {
		sum = sum + item.power
	}
	avg := sum / len(powerArray)

	return avg
}

// Calculates the highest average power given the samplesize in seconds. Using the Sliding Window algorithm.
func slidingwindow(samplesize int, Tracks []gpx.GPXTrack) {
	for _, track := range Tracks {
		for _, segment := range track.Segments {
			// Loop over all the data points in the GPX file.
			for _, point := range segment.Points {
				for _, node := range point.Extensions.Nodes {
					// Filters out the "power" field from the Extensions node in the GPX XML.
					if node.XMLName.Local == "power" {
						power, _ := strconv.Atoi(node.Data)
						powerArray = append(powerArray, PowerAndTimestamp{power, point.Timestamp})
						sliceavgPower := getAvgPower(powerArray)
						if sliceavgPower > maxAvgPower {
							maxAvgPower = sliceavgPower
						}
						// If the sample size is reached (buffer filled) remove the first item in the list and continue looping.
						if len(powerArray) == samplesize {
							powerArray = powerArray[1:]
						}
					}
				}
			}
		}
	}
	// Print out the results.
	fmt.Printf("Best %vs power: %vw\n", samplesize, maxAvgPower)
}

func main() {
	// Parse flags.
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Please provide a GPX file path!")
		return
	}

	// Loading GPX file from arg0 and sampleSizeSec from arg1
	gpxFileArg := args[0]

	sampleSizeSec, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Cannot convert arg to integer: ", err)
		return
	}

	gpxFile, err := gpx.ParseFile(gpxFileArg)

	if err != nil {
		fmt.Println("Error opening gpx file: ", err)
		return
	}

	// Run the main func
	slidingwindow(sampleSizeSec, gpxFile.Tracks)
}
