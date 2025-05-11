package wildsapi

import (
	"encoding/json"
)

type ArmorSet struct {
	ID         int            `json:"id"`         // The armor set's ID
	GameID     GameId         `json:"gameId"`     // The armor set's ID from the game files
	Name       string         `json:"name"`       // The armor set's name
	Pieces     []Armor        `json:"pieces"`     // An array of armor pieces belonging to the armor set
	Bonus      *ArmorSetBonus `json:"bonus"`      // The skill granted by the set when a certain number of pieces from the same set are worn
	GroupBonus *ArmorSetBonus `json:"groupBonus"` // The skill granted by the set when a certain number of pieces from the same group are worn
}

type ArmorSetBonus struct {
	ID    int                `json:"id"`    // The bonus's ID
	Skill Skill              `json:"skill"` // The skill granted by the bonus
	Ranks []AmorSetBonusRank `json:"ranks"` // The different ranks of the bonus
}

type AmorSetBonusRank struct {
	ID     int       `json:"id"`     // The rank's ID
	Pieces int       `json:"pieces"` // The number of armor pieces from the set that must be worn to activate the bonus
	Skill  SkillRank `json:"skill"`  // The skill granted by this rank of the bonus
}

func (c *Client) FetchArmorSets(query QueryOptions) ([]ArmorSet, error) {
	resp, err := c.fetch("armor/sets", query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var armorSets []ArmorSet
	err = json.NewDecoder(resp.Body).Decode(&armorSets)

	return armorSets, err
}

func (c *Client) FetchArmorSetsById(id int) (*ArmorSet, error) {
	resp, err := c.fetchById("armor/sets", id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var armorset ArmorSet
	err = json.NewDecoder(resp.Body).Decode(&armorset)

	return &armorset, err
}
