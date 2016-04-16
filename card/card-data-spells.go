package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attribute"
	"github.com/asukakenji/clash-royale/lib"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/types"
)

var (
	// --- Common Spells ---
	ARROWS = newCard(2000, map[attribute.Attribute]interface{}{
		attribute.Name:     "Arrows",
		attribute.Arena:    arena.Arena0,
		attribute.Rarity:   rarity.Common,
		attribute.Type:     types.Spell,
		attribute.Desc:     `Arrows pepper a large area, damaging everyone hit. Reduced damage to Crown Towers.`,
		attribute.Elixir:   3,
		attribute.BaseADam: 115,
		attribute.CTDam:    []int{46, 51, 56, 61, 67, 74, 81, 89, 98, 107, 118, 130},
		attribute.Radius:   4,
	})
	ZAP = newCard(2050, map[attribute.Attribute]interface{}{
		attribute.Name:     "Zap",
		attribute.Arena:    arena.Arena5,
		attribute.Rarity:   rarity.Common,
		attribute.Type:     types.Spell,
		attribute.Desc:     `Zaps enemies, briefly stunning them and dealing damage inside a small radius. Reduced damage to Crown Towers.`,
		attribute.Elixir:   2,
		attribute.BaseADam: 80,
		attribute.CTDam:    []int{32, 36, 39, 43, 47, 52, 56, 62, 68, 75, 82, 90},
		attribute.DurF:     1,
		attribute.Radius:   2.5,
	})

	// --- Rare Spells ---
	FIREBALL = newCard(2100, map[attribute.Attribute]interface{}{
		attribute.Name:     "Fireball",
		attribute.Arena:    arena.Arena0,
		attribute.Rarity:   rarity.Rare,
		attribute.Type:     types.Spell,
		attribute.Desc:     `Annnnnd... Fireball. Incinerates a small area, dealing high damage. Reduced damage to Crown Towers.`,
		attribute.Elixir:   4,
		attribute.BaseADam: 325,
		attribute.CTDam:    []int{130, 143, 158, 173, 190, 208, 229, 251, 276, 303},
		attribute.Radius:   2.5,
	})
	ROCKET = newCard(2130, map[attribute.Attribute]interface{}{
		attribute.Name:     "Rocket",
		attribute.Arena:    arena.Arena3,
		attribute.Rarity:   rarity.Rare,
		attribute.Type:     types.Spell,
		attribute.Desc:     `Deals high damage to a small area. Looks really awesome doing it. Reduced damage to Crown Towers.`,
		attribute.Elixir:   6,
		attribute.BaseADam: 700,
		attribute.CTDam:    []int{280, 308, 339, 373, 409, 448, 493, 541, 594, 653},
		attribute.Radius:   2,
	})

	// --- Epic Spells ---
	LIGHTNING = newCard(2210, map[attribute.Attribute]interface{}{
		attribute.Name:    "Lightning",
		attribute.Arena:   arena.Arena1,
		attribute.Rarity:  rarity.Epic,
		attribute.Type:    types.Spell,
		attribute.Desc:    `Bolts of lightning hit up to three enemy troops or buildings with the most hitpoints in the target area. Reduced damage to Crown Towers.`,
		attribute.Elixir:  6,
		attribute.Count:   3,
		attribute.BaseDam: 650,
		attribute.CTDam:   []int{260, 286, 315, 346, 380, 416, 458, 502},
		attribute.DurF:    1.5,
		attribute.Radius:  3.5,
	})
	GOBLIN_BARREL = newCard(2211, map[attribute.Attribute]interface{}{
		attribute.Name:      "Goblin Barrel",
		attribute.Arena:     arena.Arena1,
		attribute.Rarity:    rarity.Epic,
		attribute.Type:      types.Spell,
		attribute.Desc:      `Spawns three Goblins anywhere on the Arena. It's going to be a thrilling ride, boys!`,
		attribute.Elixir:    4,
		attribute.ADam:      lib.V50[0:8:8],
		attribute.CTDam:     []int{20, 22, 24, 27, 30, 32, 36, 39},
		attribute.Radius:    1.5,
		attribute.BaseGobLV: 6,
		attribute.GobCount:  3,
	})
	RAGE = newCard(2230, map[attribute.Attribute]interface{}{
		attribute.Name:   "Rage",
		attribute.Arena:  arena.Arena3,
		attribute.Rarity: rarity.Epic,
		attribute.Type:   types.Spell,
		attribute.Desc:   `Increases troop movement and attack speed by 40%. Troop buildings and summoners deploy troop faster. Chaaaarge!`,
		attribute.Elixir: 3,
		attribute.DurU:   []float64{8, 8.5, 9, 9.5, 10, 10.5, 11, 11.5},
		attribute.Radius: 5,
	})
	FREEZE = newCard(2240, map[attribute.Attribute]interface{}{
		attribute.Name:   "Freeze",
		attribute.Arena:  arena.Arena4,
		attribute.Rarity: rarity.Epic,
		attribute.Type:   types.Spell,
		attribute.Desc:   `Freezes troops and buildings, making them unable to move or attack. Everybody chill.`,
		attribute.Elixir: 4,
		attribute.DurU:   []float64{5, 5.3, 5.6, 5.9, 6.2, 6.5, 6.8, 7.1},
		attribute.Radius: 3,
	})
	MIRROR = newCard(2250, map[attribute.Attribute]interface{}{
		attribute.Name:     "Mirror",
		attribute.Arena:    arena.Arena5,
		attribute.Rarity:   rarity.Epic,
		attribute.Type:     types.Spell,
		attribute.Desc:     `Mirrors your last card played for +1 Elixir`,
		attribute.Elixir:   lib.X,
		attribute.BaseMCLV: 5,
		attribute.BaseMRLV: 3,
		attribute.BaseMELV: 1,
		attribute.MLLV:     []int{1, 1, 1, 1, 2, 3, 4, 5},
	})
	POISON = newCard(2251, map[attribute.Attribute]interface{}{
		attribute.Name:   "Poison",
		attribute.Arena:  arena.Arena5,
		attribute.Rarity: rarity.Epic,
		attribute.Type:   types.Spell,
		attribute.Desc:   `Covers the target area in a sticky toxin, damaging and slowing down troops and buildings. Remember: solvent abuse can kill!`,
		attribute.Elixir: 4,
		attribute.DPS:    []int{42, 46, 50, 55, 62, 67, 73, 81},
		attribute.CTDPS:  []int{17, 19, 20, 22, 25, 27, 30, 33},
		attribute.DurF:   10,
		attribute.Radius: 3.5,
	})
)
