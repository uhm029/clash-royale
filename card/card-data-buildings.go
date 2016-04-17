package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/targets"
	"github.com/asukakenji/clash-royale/typ"
)

var (
	// --- Common Buildings ---
	CANNON = newCard(1030, map[attr.Attribute]interface{}{
		attr.Name:    "Cannon",
		attr.Arena:   arena.Arena3,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Building,
		attr.Desc:    `Defensive building. Shoots cannonballs with deadly effect, but cannot target flying troops.`,
		attr.Elixir:  3,
		attr.BaseHP:  450,
		attr.DPS:     []int{75, 82, 90, 98, 108, 120, 131, 143, 158, 173, 191, 210},
		attr.BaseDam: 60,
		attr.HSpeed:  0.8,
		attr.Targets: targets.Ground,
		attr.Range:   6,
		attr.DTime:   1,
		attr.LTime:   30,
	})
	TESLA = newCard(1040, map[attr.Attribute]interface{}{
		attr.Name:    "Tesla",
		attr.Arena:   arena.Arena4,
		attr.Rarity:  rarity.Common,
		attr.Type:    typ.Building,
		attr.Desc:    `Defensive building. Whenever it's not zapping the enemy, the power of Electrickery is best kept grounded.`,
		attr.Elixir:  4,
		attr.BaseHP:  400,
		attr.DPS:     []int{80, 88, 96, 106, 116, 128, 140, 153, 168, 185, 203, 223},
		attr.BaseDam: 64,
		attr.HSpeed:  0.8,
		attr.Targets: targets.AirAndGround,
		attr.Range:   6,
		attr.DTime:   1,
		attr.LTime:   40,
	})
	MORTAR = newCard(1060, map[attr.Attribute]interface{}{
		attr.Name:     "Mortar",
		attr.Arena:    arena.Arena6,
		attr.Rarity:   rarity.Common,
		attr.Type:     typ.Building,
		attr.Desc:     `Defensive building with a long range. Shoots exploding shells that deal area damage. Cannot shoot at targets that get very close!`,
		attr.Elixir:   4,
		attr.BaseHP:   600,
		attr.DPS:      []int{24, 26, 29, 31, 35, 38, 42, 46, 50, 55, 61, 67},
		attr.BaseADam: 120,
		attr.HSpeed:   5,
		attr.Targets:  targets.Ground,
		attr.Range:    12,
		attr.DTime:    3,
		attr.LTime:    30,
	})

	// --- Rare Buildings ---
	GOBLIN_HUT = newCard(1110, map[attr.Attribute]interface{}{
		attr.Name:      "Goblin Hut",
		attr.Arena:     arena.Arena1,
		attr.Rarity:    rarity.Rare,
		attr.Type:      typ.Building,
		attr.Desc:      `Building that spawns Spear Goblins. But don't look inside. You don't want to see how they are made.`,
		attr.Elixir:    5,
		attr.BaseHP:    700,
		attr.BaseSgoLV: 3,
		attr.SSpeed:    4.9,
		attr.DTime:     1,
		attr.LTime:     60,
	})
	BOMB_TOWER = newCard(1120, map[attr.Attribute]interface{}{
		attr.Name:     "Bomb Tower",
		attr.Arena:    arena.Arena2,
		attr.Rarity:   rarity.Rare,
		attr.Type:     typ.Building,
		attr.Desc:     `Defensive building that houses a Bomber. Deals area damage to anything dumb enough to stand near it.`,
		attr.Elixir:   5,
		attr.BaseHP:   900,
		attr.DPS:      []int{62, 68, 75, 83, 91, 100, 110, 120, 132, 145},
		attr.BaseADam: 100,
		attr.HSpeed:   1.6,
		attr.Targets:  targets.Ground,
		attr.Range:    6.5,
		attr.DTime:    1,
		attr.LTime:    60,
	})
	TOMBSTONE = newCard(1121, map[attr.Attribute]interface{}{
		attr.Name:      "Tombstone",
		attr.Arena:     arena.Arena2,
		attr.Rarity:    rarity.Rare,
		attr.Type:      typ.Building,
		attr.Desc:      `Troop building that periodically deploys Skeletons to fight the enemy. When destroyed, spawns 4 Skeletons. Creepy!`,
		attr.Elixir:    3,
		attr.BaseHP:    200,
		attr.BaseSkeLV: 3,
		attr.SSpeed:    2.9,
		attr.DTime:     1,
		attr.LTime:     40,
	})
	BARBARIAN_HUT = newCard(1130, map[attr.Attribute]interface{}{
		attr.Name:      "Barbarian Hut",
		attr.Arena:     arena.Arena3,
		attr.Rarity:    rarity.Rare,
		attr.Type:      typ.Building,
		attr.Desc:      `Troop building that periodically deploys Barbarians to fight the enemy. Time to make the Barbarians.`,
		attr.Elixir:    7,
		attr.BaseHP:    1100,
		attr.BaseBarLV: 3,
		attr.SSpeed:    14,
		attr.DTime:     1,
		attr.LTime:     60,
	})
	INFERNO_TOWER = newCard(1160, map[attr.Attribute]interface{}{
		attr.Name:     "Inferno Tower",
		attr.Arena:    arena.Arena6,
		attr.Rarity:   rarity.Rare,
		attr.Type:     typ.Building,
		attr.Desc:     `Defensive building, roasts targets for damage that increases over time. Burns through even the biggest and toughest enemies!`,
		attr.Elixir:   5,
		attr.BaseHP:   800,
		attr.DPSL:     []int{50, 55, 60, 65, 72, 80, 87, 95, 105, 115},
		attr.DPSH:     []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330},
		attr.BaseDamL: 20,
		attr.BaseDamH: 400,
		attr.HSpeed:   0.4,
		attr.Targets:  targets.AirAndGround,
		attr.Range:    6.5,
		attr.DTime:    1,
		attr.LTime:    40,
	})
	ELIXIR_COLLECTOR = newCard(1161, map[attr.Attribute]interface{}{
		attr.Name:   "Elixir Collector",
		attr.Arena:  arena.Arena6,
		attr.Rarity: rarity.Rare,
		attr.Type:   typ.Building,
		attr.Desc:   `You gotta spend Elixir to make Elixir.`,
		attr.Elixir: 5,
		attr.BaseHP: 800,
		attr.PSpeed: 9.8,
		attr.DTime:  1,
		attr.LTime:  70,
	})

	// --- Epic Buildings ---
	X_BOW = newCard(1230, map[attr.Attribute]interface{}{
		attr.Name:    "X-Bow",
		attr.Arena:   arena.Arena3,
		attr.Rarity:  rarity.Epic,
		attr.Type:    typ.Building,
		attr.Desc:    `Nice tower you got there. Would be a shame if this X-Bow whittled it down from this side of the arena...`,
		attr.Elixir:  6,
		attr.BaseHP:  850,
		attr.DPS:     []int{66, 73, 80, 86, 96, 106, 116, 126},
		attr.BaseDam: 20,
		attr.HSpeed:  0.3,
		attr.Targets: targets.Ground,
		attr.Range:   12,
		attr.DTime:   5,
		attr.LTime:   40,
	})
)
