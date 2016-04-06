package lib

var (
	// Common
	ARROWS = NewCard(map[Attribute]interface{}{
		NAME:   "Arrows",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   SPELL,
		DESC:   `Arrows pepper a large area, damaging everyone hit. Reduced damage to Crown Towers.`,
		COST:   3,
		ADAM:   V115[0:12:12],
		CTDAM:  []int{46, 51, 56, 61, 67, 74, 81, 89, 98, 107, 118, 130},
		RAD:    4,
	})
	ZAP = NewCard(map[Attribute]interface{}{
		NAME:   "Zap",
		ARENA:  ARENA_5,
		RARITY: COMMON,
		TYPE:   SPELL,
		DESC:   `Zaps enemies, briefly stunning them and dealing damage inside a small radius. Reduced damage to Crown Towers.`,
		COST:   2,
		ADAM:   V80[0:12:12],
		CTDAM:  []int{32, 36, 39, 43, 47, 52, 56, 62, 68, 75, 82, 90},
		DUR_F:  1,
		RAD:    2.5,
	})

	// Rare
	FIREBALL = NewCard(map[Attribute]interface{}{
		NAME:   "Fireball",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   SPELL,
		DESC:   `Annnnnd... Fireball. Incinerates a small area, dealing high damage. Reduced damage to Crown Towers.`,
		COST:   4,
		ADAM:   V325[0:10:10],
		CTDAM:  []int{130, 143, 158, 173, 190, 208, 229, 251, 276, 303},
		RAD:    2.5,
	})
	ROCKET = NewCard(map[Attribute]interface{}{
		NAME:   "Rocket",
		ARENA:  ARENA_3,
		RARITY: RARE,
		TYPE:   SPELL,
		DESC:   `Deals high damage to a small area. Looks really awesome doing it. Reduced damage to Crown Towers.`,
		COST:   6,
		ADAM:   V700[0:10:10],
		CTDAM:  []int{280, 308, 339, 373, 409, 448, 493, 541, 594, 653},
		RAD:    2,
	})

	// Epic
	LIGHTNING = NewCard(map[Attribute]interface{}{
		NAME:   "Lightning",
		ARENA:  ARENA_1,
		RARITY: EPIC,
		TYPE:   SPELL,
		DESC:   `Bolts of lightning hit up to three enemy troops or buildings with the most hitpoints in the target area. Reduced damage to Crown Towers.`,
		COST:   6,
		COUNT:  3,
		DAM:    V650[0:8:8],
		CTDAM:  []int{260, 286, 315, 346, 380, 416, 458, 502},
		DUR_F:  1.5,
		RAD:    3.5,
	})
	GOBLIN_BARREL = NewCard(map[Attribute]interface{}{
		NAME:      "Goblin Barrel",
		ARENA:     ARENA_1,
		RARITY:    EPIC,
		TYPE:      SPELL,
		DESC:      `Spawns three Goblins anywhere on the Arena. It's going to be a thrilling ride, boys!`,
		COST:      4,
		ADAM:      V50[0:8:8],
		CTDAM:     []int{20, 22, 24, 27, 30, 32, 36, 39},
		RAD:       1.5,
		GOB_LV:    []int{6, 7, 8, 9, 10, 11, 12, 13},
		GOB_COUNT: 3,
	})
	RAGE = NewCard(map[Attribute]interface{}{
		NAME:   "Rage",
		ARENA:  ARENA_3,
		RARITY: EPIC,
		TYPE:   SPELL,
		DESC:   `Increases troop movement and attack speed by 40%. Troop buildings and summoners deploy troop faster. Chaaaarge!`,
		COST:   3,
		DUR_U:  []float64{8, 8.5, 9, 9.5, 10, 10.5, 11, 11.5},
		RAD:    5,
	})
	FREEZE = NewCard(map[Attribute]interface{}{
		NAME:   "Freeze",
		ARENA:  ARENA_4,
		RARITY: EPIC,
		TYPE:   SPELL,
		DESC:   `Freezes troops and buildings, making them unable to move or attack. Everybody chill.`,
		COST:   4,
		DUR_U:  []float64{5, 5.3, 5.6, 5.9, 6.2, 6.5, 6.8, 7.1},
		RAD:    3,
	})
	MIRROR = NewCard(map[Attribute]interface{}{
		NAME:   "Mirror",
		ARENA:  ARENA_5,
		RARITY: EPIC,
		TYPE:   SPELL,
		DESC:   `Mirrors your last card played for +1 Elixir`,
		COST:   X,
		MC_LV:  []int{5, 6, 7, 8, 9, 10, 11, 12},
		MR_LV:  []int{3, 4, 5, 6, 7, 8, 9, 10},
		ME_LV:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		ML_LV:  []int{1, 1, 1, 1, 2, 3, 4, 5},
	})
	POISON = NewCard(map[Attribute]interface{}{
		NAME:   "Poison",
		ARENA:  ARENA_5,
		RARITY: EPIC,
		TYPE:   SPELL,
		DESC:   `Covers the target area in a sticky toxin, damaging and slowing down troops and buildings. Remember: solvent abuse can kill!`,
		COST:   4,
		DPS:    []int{42, 46, 50, 55, 62, 67, 73, 81},
		CTDPS:  []int{17, 19, 20, 22, 25, 27, 30, 33},
		DUR_F:  10,
		RAD:    3.5,
	})
)
