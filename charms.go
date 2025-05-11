package wildsapi

import (
	"encoding/json"
)

type Charm struct {
	ID     int         `json:"id"`     // The charm's ID
	GameID GameId      `json:"gameId"` // The ID used in the game files to identify the charm
	Ranks  []CharmRank `json:"ranks"`  // An array levels the charm can be obtained at
}

type CharmRank struct {
	ID          int               `json:"id"`          // The ID of the charm rank
	Name        string            `json:"name"`        // The rank's name
	Description string            `json:"description"` // The rank's description
	Level       int               `json:"level"`       // The level of the rank
	Rarity      int               `json:"rarity"`      // The rank's rarity
	Skills      []SkillRank       `json:"skills"`      // An array of skills that are granted by this charm
	Crafting    CharmRankCrafting `json:"crafting"`    // Crafting info for this charm
}

type CharmRankCrafting struct {
	Craftable bool           `json:"craftable"` // Indicates if the rank can be crafted directly; false indicates that the rank must be upgraded from the previous
	ZennyCost int            `json:"zennyCost"` // The amount of zenny the crafting operation costs
	Materials []CraftingCost `json:"materials"` // An array of materials used to craft the charm
}

func (c *Client) FetchCharms(query QueryOptions) ([]Charm, error) {
	resp, err := c.fetch("charms", query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var charms []Charm
	err = json.NewDecoder(resp.Body).Decode(&charms)

	return charms, err
}

func (c *Client) FetchCharmsById(id int) (*Charm, error) {
	resp, err := c.fetchById("charms", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var charm Charm
	err = json.NewDecoder(resp.Body).Decode(&charm)

	return &charm, err
}
