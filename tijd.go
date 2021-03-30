package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func readConfig() map[string]interface{} {

	var configuration map[string]interface{}

	hdir, _ := os.UserHomeDir()
	data, err := ioutil.ReadFile(hdir + "/.timeytime.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		fmt.Println("Unmarshal: %v", err)
	}

	return configuration
}

func parseLocations(timeCurrent, timeInUTC time.Time, locations map[string]interface{}) bool {

	for loc, zone := range locations {
		location, err := time.LoadLocation(zone.(string))
		timeLoc := timeInUTC.In(location)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\t\t%s (%v)\n", loc, timeLoc.Format("15:04:05 -0700"), timeLoc.Location())
	}

	return true
}

func main() {

	// Open config file
	configuration := readConfig()

	// Fetch current time in utc
	timeCurrent := time.Now()
	timeInUTC := timeCurrent.In(time.UTC)

	// Print current time in local timezone
	fmt.Printf("Your time: \t%s (%s)\n", timeCurrent.Format("15:04:05 -0700"), timeCurrent.Location())
	fmt.Printf("UTC time: \t%s (%s)\n\n", timeInUTC.Format("15:04:05 -0700"), timeInUTC.Location())

	// Parse foreign locations
	locs := configuration["Locations"].(map[string]interface{})
	parseLocations(timeCurrent, timeInUTC, locs)

}
