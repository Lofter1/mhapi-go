package wildsapi

import (
	"encoding/json"
)

type Monster struct {
	ID             int                 `json:"id"`             // The monster's ID
	GameID         GameId              `json:"gameId"`         // The ID used by the game files to identify the monster
	Kind           MonsterKind         `json:"kind"`           // The monster's category
	Species        Species             `json:"species"`        // The monster's species
	Name           string              `json:"name"`           // The monster's name
	Size           MonsterSize         `json:"size"`           // The monster's base size and crown size breakpoints
	Description    string              `json:"description"`    // The monster's description
	Features       string              `json:"features"`       // The "features" section from the monster guide
	Tips           string              `json:"tips"`           // The "tips" section from the monster guide
	BaseHealth     int                 `json:"baseHealth"`     // The monster's base health
	Locations      []Location          `json:"locations"`      // The locations in which the monster can be found
	Resistances    []MonsterResistance `json:"resistances"`    // The monster's elemental and status resistances
	Weaknesses     []MonsterWeakness   `json:"weaknesses"`     // The monster's elemental and status weaknesses
	Rewards        []MonsterReward     `json:"rewards"`        // Items can be obtained from fighting the monster
	BreakableParts []MonsterPart       `json:"breakableParts"` // Monster parts that can be broken; corresponds to broken-part entries in rewards

}

type MonsterPart struct {
	ID   int     `json:"id"`
	Part *string `json:"part,omitempty"`
	Name string  `json:"name"`
}

type MonsterResistance struct {
	ID        int            `json:"id"`                  // The resistance's ID
	Condition string         `json:"condition,omitempty"` // The condition or status of the monster during which this resistance is active; a null value indicates the resistance is always active
	Kind      ResistanceKind `json:"kind"`                // The discriminant
	Effect    *Effect        `json:"effect,omitempty"`    // The effect the monster is resistant to
	Element   *Element       `json:"element,omitempty"`   // The element the monster is resistant to
	Status    *Status        `json:"status,omitempty"`    // The status the monster is resistant to
}

type MonsterReward struct {
	Item       Item              `json:"item"`       // The item given by the reward
	Conditions []RewardCondition `json:"conditions"` // An array of conditions required for the reward to be granted
}

type RewardCondition struct {
	Kind     RewardConditionKind `json:"kind"`     // The type of reward this is (such as a carve, wound, etc.)
	Rank     Rank                `json:"rank"`     // The hunter rank group for the reward
	Quantity int                 `json:"quantity"` // The base amount of the item that can be given by the reward
	Chance   Percent             `json:"chance"`   // How likely it is that the reward will be given
	Part     *string             `json:"part"`     // Only populated for broken-part rewards; if set, corresponds to an entry in Monster.BreakableParts indicating which part needs to be broken
}

type MonsterSize struct {
	Base   float64 `json:"base"`   // The monster's base size
	Mini   float64 `json:"mini"`   // The size the monster must be below to qualify as a mini crown
	Silver float64 `json:"silver"` // The size the monster must be above to qualify as a silver crown
	Gold   float64 `json:"gold"`   // The size the monster must be above to qualify as a gold crown
}

type MonsterWeakness struct {
	ID        int          `json:"id"`                // The weakness's ID
	Level     int          `json:"level"`             // How powerful the weakness is; higher values indicate a more severe weakness
	Condition *string      `json:"condition"`         // The condition or status of the monster during which this weakness applies; a null value indicates that the weakness is always active
	Kind      WeaknessKind `json:"kind"`              // The discriminant for the tagged union
	Element   *Element     `json:"element,omitempty"` // The element the monster is weak to
	Status    *Status      `json:"status,omitempty"`  // The status the monster is weak to
	Effect    *Effect      `json:"effect,omitempty"`  // The effect the monster is weak to
}

type MonsterKind string

const (
	MonsterKindLarge MonsterKind = "large"
	MonsterKindSmall MonsterKind = "small"
)

type Species string

const (
	SpeciesFlyingWyvern  Species = "flying-wyvern"
	SpeciesFish          Species = "fish"
	SpeciesHerbivore     Species = "herbivore"
	SpeciesLynian        Species = "lynian"
	SpeciesNeopteron     Species = "neopteron"
	SpeciesCarapaceon    Species = "carapaceon"
	SpeciesFangedBeast   Species = "fanged-beast"
	SpeciesBirdWyvern    Species = "bird-wyvern"
	SpeciesPiscineWyvern Species = "piscine-wyvern"
	SpeciesLeviathan     Species = "leviathan"
	SpeciesBruteWyvern   Species = "brute-wyvern"
	SpeciesFangedWyvern  Species = "fanged-wyvern"
	SpeciesAmphibian     Species = "amphibian"
	SpeciesTemnoceran    Species = "temnoceran"
	SpeciesSnakeWyvern   Species = "snake-wyvern"
	SpeciesElderDragon   Species = "elder-dragon"
	SpeciesCephalopod    Species = "cephalopod"
	SpeciesConstruct     Species = "construct"
	SpeciesWingdrake     Species = "wingdrake"
	SpeciesDemiElder     Species = "demi-elder"
)

type ResistanceKind string

const (
	ResistanceKindEffect  ResistanceKind = "effect"
	ResistanceKindElement ResistanceKind = "element"
	ResistanceKindStatus  ResistanceKind = "status"
)

type WeaknessKind string

const (
	WeaknessKindEffect  WeaknessKind = "effect"
	WeaknessKindElement WeaknessKind = "element"
	WeaknessKindStatus  WeaknessKind = "status"
)

type RewardConditionKind string

const (
	RewardConditionKindCarve                  RewardConditionKind = "carve"
	RewardConditionKindCarveSevered           RewardConditionKind = "carve-severed"
	RewardConditionKindEndemicCapture         RewardConditionKind = "endemic-capture"
	RewardConditionKindTargetReward           RewardConditionKind = "target-reward"
	RewardConditionKindBrokenPart             RewardConditionKind = "broken-part"
	RewardConditionKindWoundDestroyed         RewardConditionKind = "wound-destroyed"
	RewardConditionKindCarveRotten            RewardConditionKind = "carve-rotten"
	RewardConditionKindSlingerGather          RewardConditionKind = "slinger-gather"
	RewardConditionKindCarveRottenSevered     RewardConditionKind = "carve-rotten-severed"
	RewardConditionKindTemperedWoundDestroyed RewardConditionKind = "tempered-wound-destroyed"
	RewardConditionKindCarveCrystallized      RewardConditionKind = "carve-crystallized"
)

func (c *Client) FetchMonsters(query QueryOptions) ([]Monster, error) {
	resp, err := c.fetch("monsters", query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var monsters []Monster
	err = json.NewDecoder(resp.Body).Decode(&monsters)

	return monsters, err
}

func (c *Client) FetchMonstersById(id int) (*Monster, error) {
	resp, err := c.fetchById("monsters", id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var monster Monster
	err = json.NewDecoder(resp.Body).Decode(&monster)

	return &monster, err
}
