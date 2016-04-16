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
		attribute.NAME:      "Arrows",
		attribute.ARENA:     arena.Arena0,
		attribute.RARITY:    rarity.Common,
		attribute.TYPE:      types.Spell,
		attribute.DESC:      `Arrows pepper a large area, damaging everyone hit. Reduced damage to Crown Towers.`,
		attribute.COST:      3,
		attribute.BASE_ADAM: 115,
		attribute.CTDAM:     []int{46, 51, 56, 61, 67, 74, 81, 89, 98, 107, 118, 130},
		attribute.RAD:       4,
	})
	ZAP = newCard(2050, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Zap",
		attribute.ARENA:     arena.Arena5,
		attribute.RARITY:    rarity.Common,
		attribute.TYPE:      types.Spell,
		attribute.DESC:      `Zaps enemies, briefly stunning them and dealing damage inside a small radius. Reduced damage to Crown Towers.`,
		attribute.COST:      2,
		attribute.BASE_ADAM: 80,
		attribute.CTDAM:     []int{32, 36, 39, 43, 47, 52, 56, 62, 68, 75, 82, 90},
		attribute.DUR_F:     1,
		attribute.RAD:       2.5,
	})

	// --- Rare Spells ---
	FIREBALL = newCard(2100, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Fireball",
		attribute.ARENA:     arena.Arena0,
		attribute.RARITY:    rarity.Rare,
		attribute.TYPE:      types.Spell,
		attribute.DESC:      `Annnnnd... Fireball. Incinerates a small area, dealing high damage. Reduced damage to Crown Towers.`,
		attribute.COST:      4,
		attribute.BASE_ADAM: 325,
		attribute.CTDAM:     []int{130, 143, 158, 173, 190, 208, 229, 251, 276, 303},
		attribute.RAD:       2.5,
	})
	ROCKET = newCard(2130, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Rocket",
		attribute.ARENA:     arena.Arena3,
		attribute.RARITY:    rarity.Rare,
		attribute.TYPE:      types.Spell,
		attribute.DESC:      `Deals high damage to a small area. Looks really awesome doing it. Reduced damage to Crown Towers.`,
		attribute.COST:      6,
		attribute.BASE_ADAM: 700,
		attribute.CTDAM:     []int{280, 308, 339, 373, 409, 448, 493, 541, 594, 653},
		attribute.RAD:       2,
	})

	// --- Epic Spells ---
	LIGHTNING = newCard(2210, map[attribute.Attribute]interface{}{
		attribute.NAME:     "Lightning",
		attribute.ARENA:    arena.Arena1,
		attribute.RARITY:   rarity.Epic,
		attribute.TYPE:     types.Spell,
		attribute.DESC:     `Bolts of lightning hit up to three enemy troops or buildings with the most hitpoints in the target area. Reduced damage to Crown Towers.`,
		attribute.COST:     6,
		attribute.COUNT:    3,
		attribute.BASE_DAM: 650,
		attribute.CTDAM:    []int{260, 286, 315, 346, 380, 416, 458, 502},
		attribute.DUR_F:    1.5,
		attribute.RAD:      3.5,
	})
	GOBLIN_BARREL = newCard(2211, map[attribute.Attribute]interface{}{
		attribute.NAME:        "Goblin Barrel",
		attribute.ARENA:       arena.Arena1,
		attribute.RARITY:      rarity.Epic,
		attribute.TYPE:        types.Spell,
		attribute.DESC:        `Spawns three Goblins anywhere on the Arena. It's going to be a thrilling ride, boys!`,
		attribute.COST:        4,
		attribute.ADAM:        lib.V50[0:8:8],
		attribute.CTDAM:       []int{20, 22, 24, 27, 30, 32, 36, 39},
		attribute.RAD:         1.5,
		attribute.BASE_GOB_LV: 6,
		attribute.GOB_COUNT:   3,
	})
	RAGE = newCard(2230, map[attribute.Attribute]interface{}{
		attribute.NAME:   "Rage",
		attribute.ARENA:  arena.Arena3,
		attribute.RARITY: rarity.Epic,
		attribute.TYPE:   types.Spell,
		attribute.DESC:   `Increases troop movement and attack speed by 40%. Troop buildings and summoners deploy troop faster. Chaaaarge!`,
		attribute.COST:   3,
		attribute.DUR_U:  []float64{8, 8.5, 9, 9.5, 10, 10.5, 11, 11.5},
		attribute.RAD:    5,
	})
	FREEZE = newCard(2240, map[attribute.Attribute]interface{}{
		attribute.NAME:   "Freeze",
		attribute.ARENA:  arena.Arena4,
		attribute.RARITY: rarity.Epic,
		attribute.TYPE:   types.Spell,
		attribute.DESC:   `Freezes troops and buildings, making them unable to move or attack. Everybody chill.`,
		attribute.COST:   4,
		attribute.DUR_U:  []float64{5, 5.3, 5.6, 5.9, 6.2, 6.5, 6.8, 7.1},
		attribute.RAD:    3,
	})
	MIRROR = newCard(2250, map[attribute.Attribute]interface{}{
		attribute.NAME:       "Mirror",
		attribute.ARENA:      arena.Arena5,
		attribute.RARITY:     rarity.Epic,
		attribute.TYPE:       types.Spell,
		attribute.DESC:       `Mirrors your last card played for +1 Elixir`,
		attribute.COST:       lib.X,
		attribute.BASE_MC_LV: 5,
		attribute.BASE_MR_LV: 3,
		attribute.BASE_ME_LV: 1,
		attribute.ML_LV:      []int{1, 1, 1, 1, 2, 3, 4, 5},
	})
	POISON = newCard(2251, map[attribute.Attribute]interface{}{
		attribute.NAME:   "Poison",
		attribute.ARENA:  arena.Arena5,
		attribute.RARITY: rarity.Epic,
		attribute.TYPE:   types.Spell,
		attribute.DESC:   `Covers the target area in a sticky toxin, damaging and slowing down troops and buildings. Remember: solvent abuse can kill!`,
		attribute.COST:   4,
		attribute.DPS:    []int{42, 46, 50, 55, 62, 67, 73, 81},
		attribute.CTDPS:  []int{17, 19, 20, 22, 25, 27, 30, 33},
		attribute.DUR_F:  10,
		attribute.RAD:    3.5,
	})
)
