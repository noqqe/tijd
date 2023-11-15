package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Reset       = "\033[0m"
	Background  = "\033[38;5;59m"
	CurrentLine = "\033[38;5;60m"
	Foreground  = "\033[38;5;231m"
	Comment     = "\033[38;5;103m"
	Cyan        = "\033[38;5;159m"
	Green       = "\033[38;5;120m"
	Orange      = "\033[38;5;222m"
	Pink        = "\033[38;5;212m"
	Purple      = "\033[38;5;183m"
	Red         = "\033[38;5;210m"
	Yellow      = "\033[38;5;229m"
	Version     = "unknown"
)

func readConfig() map[string]interface{} {

	var configuration map[string]interface{}

	hdir, _ := os.UserHomeDir()
	data, err := os.ReadFile(hdir + "/.tijd.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return configuration
}

func parseLocations(timeInUTC time.Time, locations map[string]interface{}) bool {

	for loc, zone := range locations {
		location, err := time.LoadLocation(zone.(string))
		timeLoc := timeInUTC.In(location)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\t\t%s%s %s(%v)%s\n", loc, Green, timeLoc.Format("15:04:05 -0700"), Comment, timeLoc.Location(), Reset)
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
	fmt.Printf("%stijd %s%s\n\n", Purple, Version, Reset)
	fmt.Printf("Your time: \t%s%s %s(%s)%s\n", Pink, timeCurrent.Format("15:04:05 -0700"), Comment, timeCurrent.Location(), Reset)
	fmt.Printf("UTC time: \t%s%s %s(%s)%s\n\n", Yellow, timeInUTC.Format("15:04:05 -0700"), Comment, timeInUTC.Location(), Reset)

	// Parse foreign locations
	locs := configuration["Locations"].(map[string]interface{})
	parseLocations(timeInUTC, locs)

}
