package wildsapi

import (
	"encoding/json"
)

type Decoration struct {
	ID          int            `json:"id"`          // the decoration's ID
	GameID      GameId         `json:"gameId"`      // The decoration's ID from the game files
	Name        string         `json:"name"`        // The decoration's name
	Description string         `json:"description"` // The decoration's description
	Slot        int            `json:"slot"`        // The minimum level of the slot the decoration can be placed in
	Rarity      int            `json:"rarity"`      // The decoration's rarity
	Kind        DecorationKind `json:"kind"`        // What equipment group the decoration is allowed to be used on
	Skills      []SkillRank    `json:"skills"`      // The skills granted by this decoration
}

func (c *Client) FetchDecorations(query QueryOptions) ([]Decoration, error) {
	resp, err := c.fetch("decorations", query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var decorations []Decoration
	err = json.NewDecoder(resp.Body).Decode(&decorations)

	return decorations, err
}

func (c *Client) FetchDecorationsById(id int) (*Decoration, error) {
	resp, err := c.fetchById("decorations", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var decoration Decoration
	err = json.NewDecoder(resp.Body).Decode(&decoration)

	return &decoration, err
}
