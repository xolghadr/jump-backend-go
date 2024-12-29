package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Flight struct {
	AirplaneType string
	Origin       string
	Destination  string
	DestTime     time.Time
	OrgTime      time.Time
}

var airplanes map[string]int
var cities map[string]int
var flights []*Flight

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	loops, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	airplanes = make(map[string]int, loops)
	for i := 0; i < loops; i++ {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())
		velocity, _ := strconv.Atoi(inputs[1])
		airplanes[inputs[0]] = velocity
	}

	scanner.Scan()
	loops, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	cities = make(map[string]int, loops)
	for i := 0; i < loops; i++ {
		scanner.Scan()
		inputs := strings.Fields(scanner.Text())
		velocity, _ := strconv.Atoi(inputs[1])
		cities[inputs[0]] = velocity
	}

	scanner.Scan()
	loops, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	flights = make([]*Flight, 0, loops)
	for i := 0; i < loops; i++ {
		scanner.Scan()
		flight, _ := parseFlightData(scanner.Text())
		flights = append(flights, flight)
	}

	scanner.Scan()
	input := scanner.Text()

	if input == "admin" {
		processAdmin(flights)
	} else {
		processCity(input, flights)
	}
}

func parseFlightData(flightData string) (*Flight, error) {
	re := regexp.MustCompile(`([A-Za-z0-9\-\s]+)\s+([A-Za-z\-\s]+)\((\w{3},\s+\w{3}\s+\d{1,2},\s+\d{4}\s+\d{1,2}:\d{2}\s+[APM]+)\)\s+=>\s+([A-Za-z\-\s]+)` +
		`|([A-Za-z0-9\-\s]+)\s+([A-Za-z\-\s]+)\s+=>\s+([A-Za-z\-\s]+)\((\w{3},\s+\w{3}\s+\d{1,2},\s+\d{4}\s+\d{1,2}:\d{2}\s+[APM]+)\)`)

	matches := re.FindStringSubmatch(flightData)
	if len(matches) < 6 {
		return nil, fmt.Errorf("invalid flight data format")
	}
	var origin, destination, airplaneType, timeStrDest, timeStrOrg string

	if matches[2] == "Tehran" {
		airplaneType = matches[1]
		origin = matches[2]
		timeStrOrg = matches[3]
		destination = matches[4]
	} else if matches[7] == "Tehran" {
		airplaneType = matches[5]
		origin = matches[6]
		destination = matches[7]
		timeStrDest = matches[8]
	} else {
		return nil, fmt.Errorf("tehran is not involved in the flight")
	}
	var parsedOrgTime, parsedDestTime time.Time
	var err error
	if timeStrOrg != "" {
		parsedOrgTime, err = time.Parse("Mon, Jan 2, 2006 3:04 PM", timeStrOrg)
		parsedDestTime = calculateDestinationTime(parsedOrgTime, cities[destination], airplanes[airplaneType])
	} else if timeStrDest != "" {
		parsedDestTime, err = time.Parse("Mon, Jan 2, 2006 3:04 PM", timeStrDest)
		parsedOrgTime = calculateOriginTime(parsedDestTime, cities[origin], airplanes[airplaneType])
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %v", err)
	}

	return &Flight{
		AirplaneType: airplaneType,
		Origin:       origin,
		Destination:  destination,
		DestTime:     parsedDestTime,
		OrgTime:      parsedOrgTime,
	}, nil
}

func processAdmin(flights []*Flight) {
	var maxConcurrentFlights int
	for _, flight := range flights {
		var maxi int
		var ceil, floor time.Time
		if flight.Destination == "Tehran" {
			ceil = flight.DestTime.Add(1 * time.Second)
			floor = flight.DestTime.Add(-1 * time.Second)
		} else if flight.Origin == "Tehran" {
			ceil = flight.OrgTime.Add(1 * time.Second)
			floor = flight.OrgTime.Add(-1 * time.Second)
		}
		tempFlights := make([]*Flight, 0, len(flights))
		for i := 0; i < len(flights); i++ {
			if flights[i].Destination == "Tehran" {
				if flights[i].DestTime.Add(-10*time.Minute).Before(ceil) &&
					flights[i].DestTime.Add(5*time.Minute).After(floor) {
					maxi++
					tempFlights = append(tempFlights, flights[i])
				}
			} else if flights[i].Origin == "Tehran" {
				if flights[i].OrgTime.Add(-5*time.Minute).Before(ceil) &&
					flights[i].OrgTime.Add(5*time.Minute).After(floor) {
					maxi++
					tempFlights = append(tempFlights, flights[i])
				}
			}
		}
		maxConcurrentFlights = int(math.Max(float64(maxi), float64(maxConcurrentFlights)))
	}
	fmt.Print(maxConcurrentFlights)
}

func processCity(city string, flights []*Flight) {
	sort.Slice(flights, func(i, j int) bool {
		var aTime, bTime time.Time
		if flights[i].Destination == city {
			aTime = flights[i].DestTime
		} else {
			aTime = flights[i].OrgTime
		}
		if flights[j].Destination == city {
			bTime = flights[j].DestTime
		} else {
			bTime = flights[j].OrgTime
		}
		return aTime.Before(bTime)
	})

	timeFormat := "Mon, Jan 2, 2006 3:04 PM"
	for _, flight := range flights {
		if flight.Destination == city {
			text := fmt.Sprintf("%s %s %s %s(%s)", flight.AirplaneType, flight.Origin, "=>", city, flight.DestTime.Format(timeFormat))
			fmt.Println(text)
		} else if flight.Origin == city {
			text := fmt.Sprintf("%s %s(%s) %s %s", flight.AirplaneType, city, flight.OrgTime.Format(timeFormat), "=>", flight.Destination)
			fmt.Println(text)
		}
	}
}

func calculateOriginTime(tehranTime time.Time, distance int, speed int) time.Time {
	estimatedTimeSec := 3600.0 * float64(distance) / float64(speed)
	estimatedTimeDuration := time.Duration(estimatedTimeSec) * -1.0 * time.Second
	return tehranTime.Add(estimatedTimeDuration)
}
func calculateDestinationTime(tehranTime time.Time, distance int, speed int) time.Time {
	estimatedTimeSec := 3600.0 * float64(distance) / float64(speed)
	estimatedTimeDuration := time.Duration(estimatedTimeSec) * time.Second
	return tehranTime.Add(estimatedTimeDuration)
}
