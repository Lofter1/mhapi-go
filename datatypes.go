package wildsapi

type GameId int32
type Percent int

type Position struct {
	x float32
	y float32
	z float32
}

type Elderseal string

const (
	EldersealLow     Elderseal = "low"
	EldersealAvergae Elderseal = "average"
	EldersealHigh    Elderseal = "high"
)

type Element string

const (
	ElementDragon  Element = "dragon"
	ElementThunder Element = "thunder"
	ElementWater   Element = "water"
	ElementFire    Element = "fire"
	ElementIce     Element = "ice"
)

type Status string

const (
	StatusPoison    Status = "poison"
	StatusSleep     Status = "sleep"
	StatusParalysis Status = "paralysis"
	StatusStun      Status = "stun"
)

type Color string

const (
	ColorWhite       Color = "white"
	ColorGray        Color = "gray"
	ColorRose        Color = "rose"
	ColorPink        Color = "pink"
	ColorRed         Color = "red"
	ColorVermilion   Color = "vermilion"
	ColorOrange      Color = "orange"
	ColorBrown       Color = "brown"
	ColorIvory       Color = "ivory"
	ColorYellow      Color = "yellow"
	ColorLemon       Color = "lemon"
	ColorSageGreen   Color = "sage-green"
	ColorMossGreen   Color = "moss-green"
	ColorGreen       Color = "green"
	ColorEmerald     Color = "emerald"
	ColorSky         Color = "sky"
	ColorBlue        Color = "blue"
	ColorUltramarine Color = "ultramarine"
	ColorBluePurple  Color = "blue-purple"
	ColorPurple      Color = "purple"
	ColorDarkPurple  Color = "dark-purple"
)

type Rank string

const (
	RankLow    Rank = "low"
	RankHigh   Rank = "high"
	RankMaster Rank = "master"
)

type Effect string

const (
	EffectNoise   Effect = "noise"
	EffectFlash   Effect = "flash"
	EffectStun    Effect = "stun"
	EffectExhaust Effect = "exhaust"
)

type DamageKind string

const (
	DamageKindSevering   DamageKind = "severing"
	DamageKindBlunt      DamageKind = "blunt"
	DamageKindProjectile DamageKind = "projectile"
)

type DecorationKind string

const (
	DecorationKindWeapon DecorationKind = "weapon"
	DecorationKindArmor  DecorationKind = "armor"
)

// Decoration slots are represented by a positive integer indicating the max level of decoration that the slot will accept.
// Usually, these levels are provided as an array, with the number of slots on the item indicated by the number of elements in the array.
type DecorationSlot int

type AmmoKind string

const (
	AmmoKindNormal    AmmoKind = "normal"
	AmmoKindPierce    AmmoKind = "pierce"
	AmmoKindSpread    AmmoKind = "spread"
	AmmoKindSlicing   AmmoKind = "slicing"
	AmmoKindSticky    AmmoKind = "sticky"
	AmmoKindCluster   AmmoKind = "cluster"
	AmmoKindWyvern    AmmoKind = "wyvern"
	AmmoKindPoison    AmmoKind = "poison"
	AmmoKindParalysis AmmoKind = "paralysis"
	AmmoKindSleep     AmmoKind = "sleep"
	AmmoKindFlaming   AmmoKind = "flaming"
	AmmoKindWater     AmmoKind = "water"
	AmmoKindFreeze    AmmoKind = "freeze"
	AmmoKindThunder   AmmoKind = "thunder"
	AmmoKindDragon    AmmoKind = "dragon"
	AmmoKindRecover   AmmoKind = "recover"
	AmmoKindDemon     AmmoKind = "demon"
	AmmoKindArmor     AmmoKind = "armor"
	AmmoKindExhaust   AmmoKind = "exhaust"
	AmmoKindTranq     AmmoKind = "tranq"
)

type CraftingCost struct {
	Quantity int  `json:"quantity"` // The amount of the item required for the craft
	Item     Item `json:"item"`     // The item used as a material for the craft
}
