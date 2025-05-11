package wildsapi

import (
	"encoding/json"
)

type Item struct {
	ID          int          `json:"id"`          // The item's ID
	GameID      GameId       `json:"gameId"`      // The ID used by the game files to identify the item
	Name        string       `json:"name"`        // The item's name
	Description string       `json:"description"` // The item's description
	Rarity      int          `json:"rarity"`      // The item's rarity
	CarryLimit  int          `json:"carryLimit"`  // The maximum number of the item that can be carried at once
	Value       int          `json:"value"`       // The value of the item when sold to a vendor
	Recipes     []ItemRecipe `json:"recipes"`     // An array of crafting recipes that produce the item
	Icon        ItemIcon     `json:"icon"`        // Icon information

}

type ItemRecipe struct {
	ID     int    `json:"id"`     // The ID of the recipe
	Amount int    `json:"amount"` // The number of items produced by the recipe
	Inputs []Item `json:"inputs"` // An array of items consumed by the recipe
}

type ItemIcon struct {
	ID      int          `json:"id"`      // The ID used by the game files to identify the icon
	Kind    ItemIconKind `json:"kind"`    // The string representation of the icon; will be used in the future to identify asset files
	ColorID GameId       `json:"colorId"` // The ID used by the game files to identify the icon's color
	Color   Color        `json:"color"`   // The string representation of the icon's color

}

type ItemIconKind string

const (
	ItemIconKindBone        ItemIconKind = "bone"
	ItemIconKindBug         ItemIconKind = "bug"
	ItemIconKindCertificate ItemIconKind = "certificate"
	ItemIconKindClaw        ItemIconKind = "claw"
	ItemIconKindCrystal     ItemIconKind = "crystal"
	ItemIconKindExtract     ItemIconKind = "extract"
	ItemIconKindGem         ItemIconKind = "gem"
	ItemIconKindHide        ItemIconKind = "hide"
	ItemIconKindHoney       ItemIconKind = "honey"
	ItemIconKindMedulla     ItemIconKind = "medulla"
	ItemIconKindOre         ItemIconKind = "ore"
	ItemIconKindPlate       ItemIconKind = "plate"
	ItemIconKindPowder      ItemIconKind = "powder"
	ItemIconKindQuestion    ItemIconKind = "question"
	ItemIconKindScale       ItemIconKind = "scale"
	ItemIconKindShell       ItemIconKind = "shell"
	ItemIconKindSkull       ItemIconKind = "skull"
	ItemIconKindTail        ItemIconKind = "tail"
	ItemIconKindWing        ItemIconKind = "wing"
)

func (c *Client) FetchItems(query QueryOptions) ([]Item, error) {
	resp, err := c.fetch("items", query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items []Item
	err = json.NewDecoder(resp.Body).Decode(&items)

	return items, err
}

func (c *Client) FetchItemById(id int) (*Item, error) {
	resp, err := c.fetchById("items", id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var item Item
	err = json.NewDecoder(resp.Body).Decode(&item)

	return &item, err
}
