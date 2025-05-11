package wildsapi

import (
	"encoding/json"
	"fmt"
)

type Weapon struct {
	ID           int              `json:"id"`                   // The weapon's ID
	GameID       GameId           `json:"gameId"`               // The ID used by the game files to identify the weapon; unique only for weapons with the same kind
	Kind         WeaponKind       `json:"kind"`                 // The weapon type
	Name         string           `json:"name"`                 // The weapon's name
	Rarity       int              `json:"rarity"`               // The weapon's rarity
	Description  string           `json:"description"`          // The weapon description
	Damage       WeaponDamage     `json:"damage"`               // An object describing the damage the weapon deals
	Specials     []WeaponSpecial  `json:"specials"`             // An array of objects describing element or status damage dealt by the weapon
	Sharpness    *Sharpness       `json:"sharpness,omitempty"`  // The base sharpness of the weapon; not present on bows or bowguns
	Handicraft   []int            `json:"handicraft,omitempty"` // An array of breakpoints for the Handicraft skill; see Handicraft for more information (not present on bows or bowguns)
	Skills       []SkillRank      `json:"skills"`               // An array of skills granted by the weapon
	DefenseBonus int              `json:"defenseBonus"`         // Additional defense granted by the weapon
	Elderseal    Elderseal        `json:"elderseal"`            // The elderseal strength of the weapon; null indicates the weapon does not apply elderseal
	Affinity     Percent          `json:"affinity"`             // The base affinity of the weapon; can be negative
	Slots        []DecorationSlot `json:"slots"`                // An array of decoration slots on the weapon
	Crafting     WeaponCrafting   `json:"crafting"`             // Crafting information for the weapon
	Series       *WeaponSeries    `json:"series"`               // Which crafting series the weapon belongs to; null if it does not belong to a series (e.g. Artian weapons)

	Coatings     []BowCoating       `json:"coatings,omitempty"`     // An array of coatings supported by the weapon
	Phial        *Phial             `json:"phial"`                  // The type of phial used by the weapon.
	Shell        *GunlanceShell     `json:"shell,omitempty"`        //The type of shell used by the weapon
	ShellLevel   *int               `json:"shellLevel,omitempty"`   //The level of the weapon's shell
	Ammo         []BowgunAmmo       `json:"ammo,omitempty"`         // An array of ammo and capacities for the weapon; ammo the weapon cannot use will not be included in the array
	SpecialAmmo  *BowgunSpecialAmmo `json:"specialAmmo,omitempty"`  // Indicates the type of special ammo available to the weapon
	Melody       *HuntingHornMelody `json:"melody,omitempty"`       // The weapon's note and melody information
	EchoBubble   *HuntingHornBubble `json:"echoBubble,omitempty"`   // The echo bubble used by the hunting horn
	EchoWave     *HuntingHornWave   `json:"echoWave"`               // The echo wave used by the hunting horn; some horns do not have an echo wave
	KinsectLevel *int               `json:"kinsectLevel,omitempty"` // The kinsect level modifier for the weapon

}

type BowgunAmmo struct {
	Kind     AmmoKind `json:"kind"`            // The ammo type
	Level    int      `json:"level"`           // The ammo level
	Capacity int      `json:"capacity"`        // The number of shots before reloading
	Rapid    *bool    `json:"rapid,omitempty"` // Indicates if the ammo has the "Rapid Fire" modifier. Only relelvant on Light Bowgun
}

type WeaponCrafting struct {
	ID                int            `json:"id"`                // The ID
	Craftable         bool           `json:"craftable"`         // Indicates if the weapon can be directly crafted; false indicates that the weapon must be upgraded from the previous weapon
	Previous          *Weapon        `json:"previous"`          // The previous weapon in the crafting tree, or null if there isn't one
	Branches          []Weapon       `json:"branches"`          // An array of weapons that the weapon can be upgraded into
	CraftingZennyCost int            `json:"craftingZennyCost"` // The amount of zenny required to craft the weapon
	CraftingMaterials []CraftingCost `json:"craftingMaterials"` // An array of materials required to craft the weapon
	UpgradeZennyCost  int            `json:"upgradeZennyCost"`  // The amount of zenny required to upgrade the previous weapon into this one
	UpgradeMaterials  []CraftingCost `json:"upgradeMaterials"`  // An array of materials required to upgrade the previous weapon into this one
	Row               int            `json:"row"`               // Which row of the crafting tree the weapon can be found in
	Column            int            `json:"column"`            // Which column of the crafting tree the weapon can be found in
}

type WeaponDamage struct {
	Raw     int `json:"raw"`     // The raw (true) damage value
	Display int `json:"display"` // The damage value displayed in-game

}

type WeaponSeries struct {
	ID     int    `json:"id"`     // The ID of the series
	GameID GameId `json:"gameId"` // The ID used by the game files to identify the series
	Name   string `json:"name"`   // The Series name
}

type HuntingHornBubble struct {
	ID     int                    `json:"id"`             // The ID
	GameID GameId                 `json:"gameId"`         // The ID used by the game files to identify the echo bubble
	Kind   *HuntingHornBubbleKind `json:"kind,omitempty"` // The type of effect granted by the echo bubble
	Name   string                 `json:"name"`           // The name of the echo bubble effect
}

type HuntingHornBubbleKind string

const (
	HuntingHornBubbleKindEvasion  HuntingHornBubbleKind = "evasion"
	HuntingHornBubbleKindRegen    HuntingHornBubbleKind = "regen"
	HuntingHornBubbleKindStamina  HuntingHornBubbleKind = "stamina"
	HuntingHornBubbleKindDamage   HuntingHornBubbleKind = "damage"
	HuntingHornBubbleKindDefense  HuntingHornBubbleKind = "defense"
	HuntingHornBubbleKindImmunity HuntingHornBubbleKind = "immunity"
)

type HuntingHornWave struct {
	ID     int                  `json:"id"`             // The ID
	GameID GameId               `json:"gameId"`         // The ID used by the game files to identify the echo wave
	Kind   *HuntingHornWaveKind `json:"kind,omitempty"` // The type of effect granted by the echo wave
	Name   string               `json:"name"`           // The name of the echo wave effect
}

type HuntingHornWaveKind string

const (
	HuntingHornWaveKindBlunt    HuntingHornWaveKind = "blunt"
	HuntingHornWaveKindSlash    HuntingHornWaveKind = "slash"
	HuntingHornWaveKindFire     HuntingHornWaveKind = "fire"
	HuntingHornWaveKindWater    HuntingHornWaveKind = "water"
	HuntingHornWaveKindThunder  HuntingHornWaveKind = "thunder"
	HuntingHornWaveKindIce      HuntingHornWaveKind = "ice"
	HuntingHornWaveKindDragon   HuntingHornWaveKind = "dragon"
	HuntingHornWaveKindPoison   HuntingHornWaveKind = "poison"
	HuntingHornWaveKindParalyze HuntingHornWaveKind = "paralyze"
	HuntingHornWaveKindSleep    HuntingHornWaveKind = "sleep"
	HuntingHornWaveKindBlast    HuntingHornWaveKind = "blast"
)

type HuntingHornMelody struct {
	ID     int               `json:"id"`     // The ID
	GameID GameId            `json:"gameId"` // The ID used by the game files to identify the melody (note set)
	Notes  []HuntingHornNote `json:"notes"`  // An array of notes used in the melody
	Songs  []HuntingHornSong `json:"songs"`  // An array of songs that can be played by the weapon
}

type HuntingHornSong struct {
	ID       int               `json:"id"`       // The ID
	EffectID int               `json:"effectId"` // An identifier for the effect granted by the song; all songs that grant the same effect will share the same effectId
	Sequence []HuntingHornNote `json:"sequence"` // An array of notes that make up the song, in the order they must be played
	Name     string            `json:"name"`     // The name of the song, e.g. "Attack Up (S)"
}

type Sharpness struct {
	Red    int `json:"red"`    // The durability of the red (first) bar segment
	Orange int `json:"orange"` // The durability of the orange (second) bar segment
	Yellow int `json:"yellow"` // The durability of the yellow (third) bar segment
	Green  int `json:"green"`  // The durability of the gree (fourth) bar segment
	Blue   int `json:"blue"`   // The durability of the blue (fifth) bar segment
	White  int `json:"white"`  // The durability of the white (sixth) bar segment
	Purple int `json:"purple"` // The durability of the purple (seventh) bar segment
}

type WeaponSpecial struct {
	ID     int          `json:"id"`     // The ID
	Damage WeaponDamage `json:"damage"` // An object describing the damage dealt by the special attribute
	Hidden bool         `json:"hidden"` // Indicates that this special must be activated by the Free Element/Ammo Up skill
	Kind   SpecialKind  `json:"kind"`   // The discriminant

	Element *Element `json:"element,omitempty"` // The element dealt by the special attribute
	Status  *Status  `json:"status,omitempty"`  // The status effect dealt by the special attribte
}

type BowCoating string

const (
	BowCoatingBlast      BowCoating = "blast"
	BowCoatingCloseRange BowCoating = "close-range"
	BowCoatingExhaust    BowCoating = "exhaust"
	BowCoatingParalysis  BowCoating = "paralysis"
	BowCoatingPierce     BowCoating = "pierce"
	BowCoatingPoison     BowCoating = "poison"
	BowCoatingPower      BowCoating = "power"
	BowCoatingSleep      BowCoating = "sleep"
)

type HuntingHornNote string

const (
	HuntingHornNotePurple HuntingHornNote = "purple"
	HuntingHornNoteRed    HuntingHornNote = "red"
	HuntingHornNoteOrange HuntingHornNote = "orange"
	HuntingHornNoteYellow HuntingHornNote = "yellow"
	HuntingHornNoteGreen  HuntingHornNote = "green"
	HuntingHornNoteBlue   HuntingHornNote = "blue"
	HuntingHornNoteAqua   HuntingHornNote = "aqua"
	HuntingHornNoteWhite  HuntingHornNote = "white"
)

type WeaponKind string

const (
	WeaponKindBow          WeaponKind = "bow"
	WeaponKindChargeBlade  WeaponKind = "charge-blade"
	WeaponKindDualBlades   WeaponKind = "dual-blades"
	WeaponKindGreatSword   WeaponKind = "great-sword"
	WeaponKindGunlance     WeaponKind = "gunlance"
	WeaponKindHammer       WeaponKind = "hammer"
	WeaponKindHeavyBowgun  WeaponKind = "heavy-bowgun"
	WeaponKindHuntingHorn  WeaponKind = "hunting-horn"
	WeaponKindInsectGlaive WeaponKind = "insect-glaive"
	WeaponKindLance        WeaponKind = "lance"
	WeaponKindLightBowgun  WeaponKind = "light-bowgun"
	WeaponKindLongSword    WeaponKind = "long-sword"
	WeaponKindSwitchAxe    WeaponKind = "switch-axe"
	WeaponKindSwordShield  WeaponKind = "sword-shield"
)

type GunlanceShell string

const (
	GunlanceShellLong   GunlanceShell = "long"
	GunlanceShellNormal GunlanceShell = "normal"
	GunlanceShellWide   GunlanceShell = "wide"
)

type BowgunSpecialAmmo string

const (
	BowgunSpecialAmmoWyvernblast BowgunSpecialAmmo = "wyvernblast"
	BowgunSpecialAmmoAdhesive    BowgunSpecialAmmo = "adhesive"
)

type SpecialKind string

const (
	SpecialKindElement SpecialKind = "element"
	SpecialKindStatus  SpecialKind = "status"
)

type Phial struct {
	PhialType PhialKind     `json:"kind"`             // The phial's type
	Damage    *WeaponDamage `json:"damage,omitempty"` // The damage dealt by the phial; not present for power and element types or if weapon type is Charge Blade
}

type PhialKind string

const (
	SwitchAxePhialDragon    PhialKind = "dragon"
	SwitchAxePhialExhaust   PhialKind = "exhaust"
	SwitchAxePhialPoison    PhialKind = "poison"
	SwitchAxePhialPower     PhialKind = "power"
	SwitchAxePhialParalyze  PhialKind = "paralyze"
	SwitchAxePhialElement   PhialKind = "element"
	ChargeBladePhialImpact  PhialKind = "impact"
	ChargeBladePhialElement PhialKind = "element"
)

// Custom Phial unmarshal method
func (p *Phial) UnmarshalJSON(data []byte) error {
	// Try as a string (short form)
	var kindOnly string
	if err := json.Unmarshal(data, &kindOnly); err == nil {
		p.PhialType = PhialKind(kindOnly)
		p.Damage = nil
		return nil
	}

	// Try as full object
	type Alias Phial
	var obj Alias
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("Phial: unrecognized format: %s", string(data))
	}

	*p = Phial(obj)
	return nil
}

func (c *Client) FetchWeapons(query QueryOptions) ([]Weapon, error) {
	resp, err := c.fetch("weapons", query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weapons []Weapon
	err = json.NewDecoder(resp.Body).Decode(&weapons)

	return weapons, err
}

func (c *Client) FetchWeaponsById(id int) (*Weapon, error) {
	resp, err := c.fetchById("weapons", id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weapon Weapon
	err = json.NewDecoder(resp.Body).Decode(&weapon)

	return &weapon, err
}
