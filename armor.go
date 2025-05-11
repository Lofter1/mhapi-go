package wildsapi

import (
	"encoding/json"
)

type Armor struct {
	ID          int              `json:"id"`          // The armor's ID
	Name        string           `json:"name"`        // The armor's name
	Description string           `json:"description"` // The armor's description
	Kind        ArmorKind        `json:"kind"`        // The slot the armor is worn in
	Rank        Rank             `json:"rank"`        // The armor's hunter rank group
	Rarity      int              `json:"rarity"`      // The armor's rarity value
	Defense     ArmorDefens      `json:"defense"`     // The armor's defense values at certain breakpoints
	Resistances ArmorResistances `json:"resistances"` // The armor's elemental resistances
	Slots       []DecorationSlot `json:"slots"`       // The decoration slots supported by the armor; the position in the array indicates which slot is being defined (i.e. the first element is the first slot), and the value is the maximum level of allowed decoration
	Skills      []SkillRank      `json:"skills"`      // An array of SkillRanks granted by the armor
	ArmorSet    ArmorSet         `json:"armorSet"`    // The set that the armor belongs to, if any
	Crafting    ArmorCrafting    `json:"crafting"`    // Crafting information for the armor
}

type ArmorCrafting struct {
	ID        int            `json:"id"`        // The crafting data ID
	ZennyCost int            `json:"zennyCost"` // The amount of zenny it costs to craft the armor
	Materials []CraftingCost `json:"materials"` // The amount of zenny it costs to craft the armor

}

type ArmorDefens struct {
	Base int `json:"base"` // The armor's base defense value
	Max  int `json:"max"`  // The armor's maximum, un-augmented defense value
}

type ArmorResistances struct {
	Fire    int `json:"fire"`    // Fire resistance
	Water   int `json:"water"`   // Water resistance
	Ice     int `json:"ice"`     // Ice resistance
	Thunder int `json:"thunder"` // Thunder resistance
	Dragon  int `json:"dragon"`  // Dragon resistance
}

type ArmorKind string

const (
	ArmorKindArms  ArmorKind = "arms"
	ArmorKindChest ArmorKind = "chest"
	ArmorKindHead  ArmorKind = "head"
	ArmorKindLegs  ArmorKind = "legs"
	ArmorKindWaist ArmorKind = "waist"
)

func (c *Client) FetchArmor(query QueryOptions) ([]Armor, error) {
	resp, err := c.fetch("armor", query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var armor []Armor
	err = json.NewDecoder(resp.Body).Decode(&armor)

	return armor, err
}

func (c *Client) FetchArmorById(id int) (*Armor, error) {
	resp, err := c.fetchById("armor", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var armor Armor
	err = json.NewDecoder(resp.Body).Decode(&armor)

	return &armor, err
}
