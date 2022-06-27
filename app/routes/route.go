package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Positions []Position `json:"position"`
}

type Position struct {
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID string `json:"routeId"`
	ClientID string `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool `json:"finished"`
}

func NewRoute() *Route {
 return &Route{}
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route id not informed")
	}

	file, error := os.Open("coordinates/" + route.ID + ".txt")

	if error != nil {
		return error
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{
			Lat: lat,
			Long: long,
		})
	}

	return nil
}

func (route *Route) ExportJSONPositions() ([]string, error) {
	var partialRoute PartialRoutePosition
	var result []string

	total := len(route.Positions)
	for k, v := range route.Positions {
		partialRoute.ID = route.ID
		partialRoute.ClientID = route.ClientID
		partialRoute.Position = []float64{v.Lat, v.Long}
		partialRoute.Finished = false

		if total - 1 == k {
			partialRoute.Finished = true
		}

		jsonRoute, err := json.Marshal(partialRoute)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
