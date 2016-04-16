package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attribute"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/targets"
	"github.com/asukakenji/clash-royale/types"
)

var (
	// --- Common Buildings ---
	CANNON = newCard(1030, map[attribute.Attribute]interface{}{
		attribute.NAME:     "Cannon",
		attribute.ARENA:    arena.ARENA_3,
		attribute.RARITY:   rarity.COMMON,
		attribute.TYPE:     types.BUILDING,
		attribute.DESC:     `Defensive building. Shoots cannonballs with deadly effect, but cannot target flying troops.`,
		attribute.COST:     3,
		attribute.BASE_HP:  450,
		attribute.DPS:      []int{75, 82, 90, 98, 108, 120, 131, 143, 158, 173, 191, 210},
		attribute.BASE_DAM: 60,
		attribute.HSPD:     0.8,
		attribute.TGTS:     targets.GROUND,
		attribute.RNG:      6,
		attribute.DTIME:    1,
		attribute.LTIME:    30,
	})
	TESLA = newCard(1040, map[attribute.Attribute]interface{}{
		attribute.NAME:     "Tesla",
		attribute.ARENA:    arena.ARENA_4,
		attribute.RARITY:   rarity.COMMON,
		attribute.TYPE:     types.BUILDING,
		attribute.DESC:     `Defensive building. Whenever it's not zapping the enemy, the power of Electrickery is best kept grounded.`,
		attribute.COST:     4,
		attribute.BASE_HP:  400,
		attribute.DPS:      []int{80, 88, 96, 106, 116, 128, 140, 153, 168, 185, 203, 223},
		attribute.BASE_DAM: 64,
		attribute.HSPD:     0.8,
		attribute.TGTS:     targets.AIR_AND_GROUND,
		attribute.RNG:      6,
		attribute.DTIME:    1,
		attribute.LTIME:    40,
	})
	MORTAR = newCard(1060, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Mortar",
		attribute.ARENA:     arena.ARENA_6,
		attribute.RARITY:    rarity.COMMON,
		attribute.TYPE:      types.BUILDING,
		attribute.DESC:      `Defensive building with a long range. Shoots exploding shells that deal area damage. Cannot shoot at targets that get very close!`,
		attribute.COST:      4,
		attribute.BASE_HP:   600,
		attribute.DPS:       []int{24, 26, 29, 31, 35, 38, 42, 46, 50, 55, 61, 67},
		attribute.BASE_ADAM: 120,
		attribute.HSPD:      5,
		attribute.TGTS:      targets.GROUND,
		attribute.RNG:       12,
		attribute.DTIME:     3,
		attribute.LTIME:     30,
	})

	// --- Rare Buildings ---
	GOBLIN_HUT = newCard(1110, map[attribute.Attribute]interface{}{
		attribute.NAME:        "Goblin Hut",
		attribute.ARENA:       arena.ARENA_1,
		attribute.RARITY:      rarity.RARE,
		attribute.TYPE:        types.BUILDING,
		attribute.DESC:        `Building that spawns Spear Goblins. But don't look inside. You don't want to see how they are made.`,
		attribute.COST:        5,
		attribute.BASE_HP:     700,
		attribute.BASE_SGO_LV: 3,
		attribute.SSPD:        4.9,
		attribute.DTIME:       1,
		attribute.LTIME:       60,
	})
	BOMB_TOWER = newCard(1120, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Bomb Tower",
		attribute.ARENA:     arena.ARENA_2,
		attribute.RARITY:    rarity.RARE,
		attribute.TYPE:      types.BUILDING,
		attribute.DESC:      `Defensive building that houses a Bomber. Deals area damage to anything dumb enough to stand near it.`,
		attribute.COST:      5,
		attribute.BASE_HP:   900,
		attribute.DPS:       []int{62, 68, 75, 83, 91, 100, 110, 120, 132, 145},
		attribute.BASE_ADAM: 100,
		attribute.HSPD:      1.6,
		attribute.TGTS:      targets.GROUND,
		attribute.RNG:       6.5,
		attribute.DTIME:     1,
		attribute.LTIME:     60,
	})
	TOMBSTONE = newCard(1121, map[attribute.Attribute]interface{}{
		attribute.NAME:        "Tombstone",
		attribute.ARENA:       arena.ARENA_2,
		attribute.RARITY:      rarity.RARE,
		attribute.TYPE:        types.BUILDING,
		attribute.DESC:        `Troop building that periodically deploys Skeletons to fight the enemy. When destroyed, spawns 4 Skeletons. Creepy!`,
		attribute.COST:        3,
		attribute.BASE_HP:     200,
		attribute.BASE_SKE_LV: 3,
		attribute.SSPD:        2.9,
		attribute.DTIME:       1,
		attribute.LTIME:       40,
	})
	BARBARIAN_HUT = newCard(1130, map[attribute.Attribute]interface{}{
		attribute.NAME:        "Barbarian Hut",
		attribute.ARENA:       arena.ARENA_3,
		attribute.RARITY:      rarity.RARE,
		attribute.TYPE:        types.BUILDING,
		attribute.DESC:        `Troop building that periodically deploys Barbarians to fight the enemy. Time to make the Barbarians.`,
		attribute.COST:        7,
		attribute.BASE_HP:     1100,
		attribute.BASE_BAR_LV: 3,
		attribute.SSPD:        14,
		attribute.DTIME:       1,
		attribute.LTIME:       60,
	})
	INFERNO_TOWER = newCard(1160, map[attribute.Attribute]interface{}{
		attribute.NAME:      "Inferno Tower",
		attribute.ARENA:     arena.ARENA_6,
		attribute.RARITY:    rarity.RARE,
		attribute.TYPE:      types.BUILDING,
		attribute.DESC:      `Defensive building, roasts targets for damage that increases over time. Burns through even the biggest and toughest enemies!`,
		attribute.COST:      5,
		attribute.BASE_HP:   800,
		attribute.DPSL:      []int{50, 55, 60, 65, 72, 80, 87, 95, 105, 115},
		attribute.DPSH:      []int{1000, 1100, 1210, 1330, 1460, 1600, 1760, 1930, 2120, 2330},
		attribute.BASE_DAML: 20,
		attribute.BASE_DAMH: 400,
		attribute.HSPD:      0.4,
		attribute.TGTS:      targets.AIR_AND_GROUND,
		attribute.RNG:       6.5,
		attribute.DTIME:     1,
		attribute.LTIME:     40,
	})
	ELIXIR_COLLECTOR = newCard(1161, map[attribute.Attribute]interface{}{
		attribute.NAME:    "Elixir Collector",
		attribute.ARENA:   arena.ARENA_6,
		attribute.RARITY:  rarity.RARE,
		attribute.TYPE:    types.BUILDING,
		attribute.DESC:    `You gotta spend Elixir to make Elixir.`,
		attribute.COST:    5,
		attribute.BASE_HP: 800,
		attribute.PSPD:    9.8,
		attribute.DTIME:   1,
		attribute.LTIME:   70,
	})

	// --- Epic Buildings ---
	X_BOW = newCard(1230, map[attribute.Attribute]interface{}{
		attribute.NAME:     "X-Bow",
		attribute.ARENA:    arena.ARENA_3,
		attribute.RARITY:   rarity.EPIC,
		attribute.TYPE:     types.BUILDING,
		attribute.DESC:     `Nice tower you got there. Would be a shame if this X-Bow whittled it down from this side of the arena...`,
		attribute.COST:     6,
		attribute.BASE_HP:  850,
		attribute.DPS:      []int{66, 73, 80, 86, 96, 106, 116, 126},
		attribute.BASE_DAM: 20,
		attribute.HSPD:     0.3,
		attribute.TGTS:     targets.GROUND,
		attribute.RNG:      12,
		attribute.DTIME:    5,
		attribute.LTIME:    40,
	})
)
