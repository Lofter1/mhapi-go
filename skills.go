package wildsapi

import (
	"encoding/json"
)

type Skill struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Ranks       []SkillRank `json:"ranks"`
	Kind        SkillKind   `json:"kind"`
	Icon        SkillIcon   `json:"icon"`
}

type SkillRank struct {
	Id          int    `json:"id"`          // The skill rank's ID
	Name        string `json:"name"`        // The rank's name; always null except for on set and group bonus skills
	Description string `json:"description"` // The rank's description
	Level       int    `json:"level"`       // The skill level of the rank

	Skill Skill `json:"skill"`
}

type SkillIcon struct {
	Id   GameId        `json:"id"`
	Kind SkillIconKind `json:"kind"`
}

type SkillKind string

const (
	SkillKindArmor  SkillKind = "armor"
	SkillKindWeapon SkillKind = "weapon"
	SkillKindSet    SkillKind = "set"
	SkillKindGroup  SkillKind = "group"
)

type SkillIconKind string

const (
	SkillIconKindAffinity   SkillIconKind = "affinity"
	SkillIconKindAttack     SkillIconKind = "attack"
	SkillIconKindDefense    SkillIconKind = "defense"
	SkillIconKindElement    SkillIconKind = "element"
	SkillIconKindGathering  SkillIconKind = "gathering"
	SkillIconKindGroup      SkillIconKind = "group"
	SkillIconKindHandicraft SkillIconKind = "handicraft"
	SkillIconKindHealth     SkillIconKind = "health"
	SkillIconKindItem       SkillIconKind = "item"
	SkillIconKindOffense    SkillIconKind = "offense"
	SkillIconKindRanged     SkillIconKind = "ranged"
	SkillIconKindSet        SkillIconKind = "set"
	SkillIconKindStamina    SkillIconKind = "stamina"
	SkillIconKindUtility    SkillIconKind = "utility"
)

func (c *Client) FetchSkills(query QueryOptions) ([]Skill, error) {
	resp, err := c.fetch("skills", query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var skills []Skill
	err = json.NewDecoder(resp.Body).Decode(&skills)

	return skills, err
}

func (c *Client) FetchSkillsById(id int) (*Skill, error) {
	resp, err := c.fetchById("skills", id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var skill Skill
	err = json.NewDecoder(resp.Body).Decode(&skill)

	return &skill, err
}
