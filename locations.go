package wildsapi

import (
	"encoding/json"
)

type Location struct {
	ID        int    `json:"id"`        // The location's ID
	GameID    GameId `json:"gameId"`    // The ID used by the game files to identify this location
	Name      string `json:"name"`      // The location's name
	ZoneCount int    `json:"zoneCount"` // The number of zones in the location
	Camps     []Camp `json:"camps"`     // An array of camps in the location

}

type Camp struct {
	ID       int      `json:"id"`       // The camp's ID
	Name     string   `json:"name"`     // The camp's name
	Zone     int      `json:"zone"`     // Which zone the camp is located in
	Floor    int      `json:"floor"`    // Which floor of the map the camp is located on
	Risk     Risk     `json:"risk"`     // How likely the camp is to be found by a monster
	Position Position `json:"position"` // The map coordinates that the camp is located at
}

type Risk string

const (
	RiskDangerous Risk = "dangerous"
	RiskInsecure  Risk = "insecure"
	RiskSafe      Risk = "safe"
)

func (c *Client) FetchLocations(query QueryOptions) ([]Location, error) {
	resp, err := c.fetch("locations", query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []Location
	err = json.NewDecoder(resp.Body).Decode(&locations)

	return locations, err
}

func (c *Client) FetchLocationsById(id int) (*Location, error) {
	resp, err := c.fetchById("locations", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location Location
	err = json.NewDecoder(resp.Body).Decode(&location)

	return &location, err
}
