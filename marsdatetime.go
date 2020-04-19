package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/message"
)

func main() {
	if len(os.Args) < 2 {
		err, _ := json.Marshal(&ResultError{"Expected argument DateTime not provided. Use as format 2019-12-27T15:22:22Z"})
		fmt.Println(string(err))
		os.Exit(1)
	}

	datetime := os.Args[1]

	marsDate, err := EarthToMarsDate(datetime)
	if err != nil {
		err, _ := json.Marshal(err)
		fmt.Println(string(err))
		os.Exit(1)
	}

	result, _ := json.Marshal(marsDate)
	fmt.Println(string(result))
}

// MarsDate output response data
type MarsDate struct {
	MSD string `json:"msd"`
	MTC string `json:"mtc"`
}

// ResultError error messages
type ResultError struct {
	Message string `json:"error"`
}

// EarthToMarsDate algorithm based on nasa and mars-clock by James Tauber
// https://www.giss.nasa.gov/tools/mars24/help/algorithm.html
// http://jtauber.github.io/mars-clock/
func EarthToMarsDate(datetime string) (*MarsDate, *ResultError) {
	if !strings.Contains(datetime, "T") {
		return nil, &ResultError{"Invalid DateTime format provided. " + fmt.Sprintf("Use RFC3339 for DateTime format: %q", "2019-12-27T15:04:03")}
	}

	if !strings.Contains(datetime, "Z") {
		datetime = datetime + "Z"
	}
	t, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return nil, &ResultError{err.Error()}
	}

	// A-1. Get a starting Earth time
	millis := t.UTC().UnixNano() / 1000000
	// A-2. Convert millis to Julian Date
	jdUt := fmt.Sprintf("%.7f", 2440587.5+(float64(millis)/8.64e7))
	ut, _ := strconv.ParseFloat(jdUt, 64)
	jdTt := fmt.Sprintf("%.9f", ut+69.184/86400)
	tt, _ := strconv.ParseFloat(jdTt, 64)
	j2000 := tt - 2451545.0
	// C-2. Determine Coordinated Mars Time
	msd := ((j2000 - 4.5) / 1.027491252) + 44796.0 - 0.00096
	mtc := math.Mod(24*msd, 24)
	// format mtc hours to standard hours:minutes:seconds
	mtcH := mtc * 3600
	hh := math.Floor(mtcH / 3600)
	mtcHm := math.Mod(mtcH, 3600)
	mm := math.Floor(mtcHm / 60)
	ss := math.Round(math.Mod(mtcHm, 60))

	printer := message.NewPrinter(message.MatchLanguage("en"))
	marsDate := MarsDate{
		MSD: printer.Sprintf("%.5f", msd),
		MTC: fmt.Sprintf("%02d:%02d:%02d", int(hh), int(mm), int(ss)),
	}

	return &marsDate, nil
}
