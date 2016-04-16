package card

import (
	"github.com/asukakenji/clash-royale/Type"
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attribute"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/targets"
)

var (
	// --- Common Buildings ---
	CANNON = newCard(1030, map[attribute.Attribute]interface{}{
		attribute.Name:    "Cannon",
		attribute.Arena:   arena.Arena3,
		attribute.Rarity:  rarity.Common,
		attribute.Type:    Type.Building,
		attribute.Desc:    `Defensive building. Shoots cannonballs with deadly effect, but cannot target flying troops.`,
		attribute.Elixir:  3,
		attribute.BaseHP:  450,
		attribute.DPS:     []int{75, 82, 90, 98, 108, 120, 131, 143, 158, 173, 191, 210},
		attribute.BaseDam: 60,
		attribute.HSpeed:  0.8,
		attribute.Targets: targets.Ground,
		attribute.Range:   6,
		attribute.DTime:   1,
		attribute.LTime:   30,
	})
	TESLA = newCard(1040, map[attribute.Attribute]interface{}{
		attribute.Name:    "Tesla",
		attribute.Arena:   arena.Arena4,
		attribute.Rarity:  rarity.Common,
		attribute.Type:    Type.Building,
		attribute.Desc:    `Defensive building. Whenever it's not zapping the enemy, the power of Electrickery is best kept grounded.`,
		attribute.Elixir:  4,
		attribute.BaseHP:  400,
		attribute.DPS:     []int{80, 88, 96, 106, 116, 128, 140, 153, 168, 185, 203, 223},
		attribute.BaseDam: 64,
		attribute.HSpeed:  0.8,
		attribute.Targets: targets.AirAndGround,
		attribute.Range:   6,
		attribute.DTime:   1,
		attribute.LTime:   40,
	})
	MORTAR = newCard(1060, map[attribute.Attribute]interface{}{
		attribute.Name:     "Mortar",
		attribute.Arena:    arena.Arena6,
		attribute.Rarity:   rarity.Common,
		attribute.Type:     Type.Building,
		attribute.Desc:     `Defensive building with a long range. Shoots exploding shells that deal area damage. Cannot shoot at targets that get very close!`,
		attribute.Elixir:   4,
		attribute.BaseHP:   600,
		attribute.DPS:      []int{24, 26, 29, 31, 35, 38, 42, 46, 50, 55, 61, 67},
		attribute.BaseADam: 120,
		attribute.HSpeed:   5,
		attribute.Targets:  targets.Ground,
		attribute.Range:    12,
		attribute.DTime:    3,
		attribute.LTime:    30,
	})

	// --- Rare Buildings ---
	GOBLIN_HUT = newCard(1110, map[attribute.Attribute]interface{}{
		attribute.Name:      "Goblin Hut",
		attribute.Arena:     arena.Arena1,
		attribute.Rarity:    rarity.Rare,
		attribute.Type:      Type.Building,
		attribute.Desc:      `Building that spawns Spear Goblins. But don't look inside. You don't want to see how they are made.`,
		attribute.Elixir:    5,
		attribute.BaseHP:    700,
		attribute.BaseSgoLV: 3,
		attribute.SSpeed:    4.9,
		attribute.DTime:     1,
		attribute.LTime:     60,
	})
	BOMB_TOWER = newCard(1120, map[attribute.Attribute]interface{}{
		attribute.Name:     "Bomb Tower",
		attribute.Arena:    arena.Arena2,
		attribute.Rarity:   rarity.Rare,
		attribute.Type:     Type.Building,
		attribute.Desc:     `Defensive building that houses a Bomber. Deals area damage to anything dumb enough to stand near it.`,
		attribute.Elixir:   5,
		attribute.BaseHP:   900,
		attribute.DPS:      []int{62, 68, 75, 83, 91, 100, 110, 120, 132, 145},
		attribute.BaseADam: 100,
		attribute.HSpeed:   1.6,
		attribute.Targets:  targets.Ground,
		attribute.Range:    6.5,
		attribute.DTime:    1,
		attribute.LTime:    60,
	})
	TOMBSTONE = newCard(1121, map[attribute.Attribute]interface{}{
		attribute.Name:      "Tombstone",
		attribute.Arena:     arena.Arena2,
		attribute.Rarity:    rarity.Rare,
		attribute.Type:      Type.Building,
		attribute.Desc:      `Troop building that periodically deploys Skeletons to fight the enemy. When destroyed, spawns 4 Skeletons. Creepy!`,
		attribute.Elixir:    3,
		attribute.BaseHP:    200,
		attribute.BaseSkeLV: 3,
		attribute.SSpeed:    2.9,
		attribute.DTime:     1,
		attribute.LTime:     40,
	})
	BARBARIAN_HUT = newCard(1130, map[attribute.Attribute]interface{}{
		attribute.Name:      "Barbarian Hut",
		attribute.Arena:     arena.Arena3,
		attribute.Rarity:    rarity.Rare,
		attribute.Type:      Type.Building,
		attribute.Desc:      `Troop building that periodically deploys Barbarians to fight the enemy. Time to make the Barbarians.`,
		attribute.Elixir:    7,
		attribute.BaseHP:    1100,
		attribute.BaseBarLV: 3,
		attribute.SSpeed:    14,
		attribute.DTime:     1,
		attribute.LTime:     60,
	})
	INFERNO_TOWER = newCard(1160, map[attribute.Attribute]interface{}{
		attribute.Name:     "Inferno Tower",
		attribute.Arena:    arena.Arena6,
		attribute.Rarity:   rarity.Rare,
		attribute.Type:     Type.Building,
		attribute.Desc:     `Defensive building, roasts targets for damage that increases over time. Burns through even the biggest and toughest enemies!`,
		attribute.Elixir:   5,
		attribute.BaseHP:   800,
		attribute.DPSL:     []int{50, 55, 60, 65, 72, 80, 87, 95, 105, 115},
		attribute.DPSH:     []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330},
		attribute.BaseDamL: 20,
		attribute.BaseDamH: 400,
		attribute.HSpeed:   0.4,
		attribute.Targets:  targets.AirAndGround,
		attribute.Range:    6.5,
		attribute.DTime:    1,
		attribute.LTime:    40,
	})
	ELIXIR_COLLECTOR = newCard(1161, map[attribute.Attribute]interface{}{
		attribute.Name:   "Elixir Collector",
		attribute.Arena:  arena.Arena6,
		attribute.Rarity: rarity.Rare,
		attribute.Type:   Type.Building,
		attribute.Desc:   `You gotta spend Elixir to make Elixir.`,
		attribute.Elixir: 5,
		attribute.BaseHP: 800,
		attribute.PSpeed: 9.8,
		attribute.DTime:  1,
		attribute.LTime:  70,
	})

	// --- Epic Buildings ---
	X_BOW = newCard(1230, map[attribute.Attribute]interface{}{
		attribute.Name:    "X-Bow",
		attribute.Arena:   arena.Arena3,
		attribute.Rarity:  rarity.Epic,
		attribute.Type:    Type.Building,
		attribute.Desc:    `Nice tower you got there. Would be a shame if this X-Bow whittled it down from this side of the arena...`,
		attribute.Elixir:  6,
		attribute.BaseHP:  850,
		attribute.DPS:     []int{66, 73, 80, 86, 96, 106, 116, 126},
		attribute.BaseDam: 20,
		attribute.HSpeed:  0.3,
		attribute.Targets: targets.Ground,
		attribute.Range:   12,
		attribute.DTime:   5,
		attribute.LTime:   40,
	})
)
