package lib

import (
	"fmt"
)

type CardAttribute interface {
	CardAttribute() // Tag
	String() string
}

type FixedCardAttribute struct {
	name        string
	FormatValue func(value interface{}) string
}

func (attr *FixedCardAttribute) CardAttribute() {
}

func (attr *FixedCardAttribute) String() string {
	return attr.name
}

type UpgradableCardAttribute struct {
	name         string
	FormatValues func(values interface{}) []string
}

func (attr *UpgradableCardAttribute) CardAttribute() {
}

func (attr *UpgradableCardAttribute) String() string {
	return attr.name
}

// --- Format Functions ---

func FormatString(value interface{}) string {
	return value.(string)
}

func FormatInt(value interface{}) string {
	if X == value.(int) {
		return ""
	}
	return fmt.Sprintf("%d", value.(int))
}

func FormatTime(value interface{}) string {
	switch value.(type) {
	case int:
		return fmt.Sprintf("%dsec", value)
	case float64:
		return fmt.Sprintf("%.1fsec", value)
	default:
		panic("Unknown value type")
	}
}

func FormatInts(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = FormatInt(v)
	}
	return strings
}

func FormatTimes(values interface{}) []string {
	ints := values.([]int)
	strings := make([]string, len(ints))
	for i, v := range ints {
		strings[i] = FormatTime(v)
	}
	return strings
}

var (
	NAME = &FixedCardAttribute{
		"Name",
		FormatString,
	}
	ARENA = &FixedCardAttribute{
		"Arena",
		func(value interface{}) string {
			return value.(*Arena).String()
		},
	}
	RARITY = &FixedCardAttribute{
		"Rarity",
		func(value interface{}) string {
			return value.(*Rarity).String()
		},
	}
	TYPE = &FixedCardAttribute{
		"Type",
		func(value interface{}) string {
			return string(value.(Type))
		},
	}
	DESC = &FixedCardAttribute{
		"Description",
		FormatString,
	}
	COST = &FixedCardAttribute{
		"Elixir Cost",
		FormatInt,
	}
	HP = &UpgradableCardAttribute{
		"Hitpoints",
		FormatInts,
	}
	DPS = &UpgradableCardAttribute{
		"Damage per Second",
		FormatInts,
	}
	DAM = &UpgradableCardAttribute{
		"Damage",
		FormatInts,
	}
	ADAM = &UpgradableCardAttribute{
		"Area Damage",
		FormatInts,
	}
	DDAM = &UpgradableCardAttribute{
		"Death Damage",
		FormatInts,
	}
	SKE_LV = &UpgradableCardAttribute{
		"Skeleton Level",
		FormatInts,
	}
	SGO_LV = &UpgradableCardAttribute{
		"Spear Goblin Level",
		FormatInts,
	}
	SSPD = &FixedCardAttribute{
		"Spawn Speed",
		FormatTime,
	}
	HSPD = &FixedCardAttribute{
		"Hit Speed",
		FormatTime,
	}
	TGTS = &FixedCardAttribute{
		"Targets",
		func(value interface{}) string {
			return string(value.(Targets))
		},
	}
	SPD = &FixedCardAttribute{
		"Speed",
		func(value interface{}) string {
			return string(value.(Speed))
		},
	}
	RNG = &FixedCardAttribute{
		"Range",
		func(value interface{}) string {
			switch value.(type) {
			case int:
				if value.(int) == MELEE {
					return "Melee"
				}
				return fmt.Sprintf("%d", value)
			case float64:
				return fmt.Sprintf("%.1f", value)
			default:
				panic("Unknown value type")
			}
		},
	}
	DTIME = &FixedCardAttribute{
		"Deploy Time",
		FormatTime,
	}
	COUNT = &FixedCardAttribute{
		"Count",
		func(value interface{}) string {
			return fmt.Sprintf("x %d", value.(int))
		},
	}
)

var CARD_ATTRIBUTES = [...]CardAttribute{
	NAME,
	ARENA,
	RARITY,
	TYPE,
	DESC,
	COST,
	HP,
	DPS,
	DAM,
	ADAM,
	DDAM,
	SKE_LV,
	SGO_LV,
	SSPD,
	HSPD,
	TGTS,
	SPD,
	RNG,
	DTIME,
	COUNT,
}

type Card map[CardAttribute]interface{}

var (
	KNIGHT = Card{
		NAME:   "Knight",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A tough melee fighter. The Barbarian's handsome, cultured cousin. Rumor has it that he was knighted based on the sheer awesomeness of his mustache alone.`,
		COST:   3,
		HP:     []int{600, 660, 726, 798, 876, 960, 1056, 1158, 1272, 1398, 1536, 1686},
		DPS:    []int{68, 74, 81, 90, 99, 109, 120, 130, 144, 158, 174, 190},
		DAM:    []int{75, 82, 90, 99, 109, 120, 132, 144, 159, 174, 192, 210},
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	ARCHERS = Card{
		NAME:   "Archers",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A pair of unarmored ranged attackers. They'll help you with ground and air unit attacks, but you're on your own with coloring your hair.`,
		COST:   3,
		HP:     []int{125, 137, 151, 166, 182, 200, 220, 241, 265, 291, 320, 351},
		DPS:    []int{33, 36, 40, 44, 48, 53, 58, 64, 70, 77, 85, 93},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		HSPD:   1.2,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  2,
	}
	BOMBER = Card{
		NAME:   "Bomber",
		ARENA:  ARENA_0,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Small, lightly protected skeleton that throws bombs. Deals area damage that can wipe out a swarm of enemies.`,
		COST:   3,
		HP:     []int{150, 165, 181, 199, 219, 240, 264, 289, 318, 349, 384, 421},
		DPS:    []int{52, 57, 63, 70, 76, 84, 92, 101, 111, 122, 134, 147},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, 193, 212, 233, 256, 281},
		HSPD:   1.9,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    5,
		DTIME:  1,
	}
	GOBLINS = Card{
		NAME:   "Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three fast, unarmored melee attackers. Small, fast, green and mean!`,
		COST:   2,
		HP:     []int{80, 88, 96, 106, 116, 128, 140, X, 169, 186, 204, 224},
		DPS:    []int{45, 50, 54, 60, 66, 72, 80, X, 96, 105, 116, 127},
		DAM:    []int{50, 55, 60, 66, 73, 80, 88, X, 106, 116, 128, 140},
		HSPD:   1.1,
		TGTS:   GROUND,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  3,
	}
	SPEAR_GOBLINS = Card{
		NAME:   "Spear Goblins",
		ARENA:  ARENA_1,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three unarmored ranged attackers. Who the heck taught these guys to throw spears!?! Who thought that was a good idea?!`,
		COST:   2,
		HP:     []int{52, 57, 62, 69, 75, 83, 91, 100, 110, 121, 133, 146},
		DPS:    []int{18, 20, 22, 23, 26, 29, 32, 35, 38, 42, 46, 51},
		DAM:    []int{24, 26, 29, 31, 35, 38, 42, 46, 50, 55, 61, 67},
		HSPD:   1.3,
		TGTS:   AIR_AND_GROUND,
		SPD:    VERY_FAST,
		RNG:    5.5,
		DTIME:  1,
		COUNT:  3,
	}
	SKELETONS = Card{
		NAME:   "Skeletons",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Four fast, very weak melee fighters. Swarm your enemies with this pile of bones!`,
		COST:   1,
		HP:     []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		DPS:    []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		DAM:    []int{30, 33, 36, 39, 43, 48, 52, 57, 63, 69, 76, 84},
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINIONS = Card{
		NAME:   "Minions",
		ARENA:  ARENA_2,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Three fast, unarmored flying attackers. Roses are red, minions are blue, they can fly, and will crush you!`,
		COST:   3,
		HP:     []int{90, 99, 108, 119, 131, 144, 158, 173, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  3,
	}
	BARBARIANS = Card{
		NAME:   "Barbarians",
		ARENA:  ARENA_3,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `A horde of melee attackers with mean mustaches and even meaner tempers.`,
		COST:   5,
		HP:     []int{300, X, X, 399, 438, 480, 528, 579, X, 699, 768, 843},
		DPS:    []int{50, X, X, 66, 72, 80, 88, 96, X, 116, 128, 140},
		DAM:    []int{75, X, X, 99, 109, 120, 132, 144, X, 174, 192, 210},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  4,
	}
	MINION_HORDE = Card{
		NAME:   "Minion Horde",
		ARENA:  ARENA_4,
		RARITY: COMMON,
		TYPE:   TROOP,
		DESC:   `Six fast, unarmored flying attackers. Three's a crowd, six is a horde!`,
		COST:   5,
		HP:     []int{90, 99, 108, 119, 131, 144, 158, 173, 190, 209, 230, 252},
		DPS:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		DAM:    []int{40, 44, 48, 53, 58, 64, 70, 77, 84, 93, 102, 112},
		HSPD:   1,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    2.5,
		DTIME:  1,
		COUNT:  6,
	}
	GIANT = Card{
		NAME:   "Giant",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Slow but durable, only attacks buildings. A real one-man wrecking crew!`,
		COST:   5,
		HP:     []int{2000, 2200, 2420, 2660, 2920, X, 3520, 3860, X, 4660},
		DPS:    []int{80, 88, 96, 106, 116, X, 140, 154, X, 186},
		DAM:    []int{120, 132, 145, 159, 175, X, 221, 231, X, 279},
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    SLOW,
		RNG:    MELEE,
		DTIME:  1,
	}
	MUSKETEER = Card{
		NAME:   "Musketeer",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Don't be fooled by her delicately coiffed hair, the Musketeer is a mean shot with her trusty boomstick.`,
		COST:   4,
		HP:     []int{340, 374, 411, 452, 496, X, X, 656, 720, 792},
		DPS:    []int{90, X, 110, 120, 132, X, X, 175, 192, 211},
		DAM:    []int{100, 110, 121, 133, 146, 160, 176, 193, 212, 233},
		HSPD:   1.1,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    6.5,
		DTIME:  1,
	}
	MINI_PEKKA = Card{
		NAME:   "Mini P.E.K.K.A",
		ARENA:  ARENA_0,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `The arena is a certified butterfly-free zone. No distractions for P.E.K.K.A, only destruction.`,
		COST:   4,
		HP:     []int{600, 660, 726, 798, 876, X, X, X, X, X},
		DPS:    []int{180, 198, 218, 240, 263, X, 317, 348, X, X},
		DAM:    []int{325, 357, 393, 432, 474, X, 572, 627, X, X},
		HSPD:   1.8,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	VALKYRIE = Card{
		NAME:   "Valkyrie",
		ARENA:  ARENA_1,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Tough melee fighter, deals area damage around her. Swarm or horde, no problem! She can take them all out with a few spins.`,
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544, 1696, 1864},
		DPS:    []int{73, X, X, X, X, X, X, 141, X, X},
		DAM:    []int{110, 121, 133, 146, 160, 176, 193, 212, 233, 256},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    MELEE,
		DTIME:  1,
	}
	HOG_RIDER = Card{
		NAME:   "Hog Rider",
		ARENA:  ARENA_4,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `Fast melee troop that targets buildings and can jump over the river. He followed the echoing call of "Hog Riderrrrr" all the way through the arena doors.`,
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544, 1696, 1864},
		DPS:    []int{106, 116, 128, 141, 155, X, X, 205, 226, 248},
		DAM:    []int{160, 176, 193, 212, 233, X, X, 308, 339, 372},
		HSPD:   1.5,
		TGTS:   BUILDINGS,
		SPD:    VERY_FAST,
		RNG:    MELEE,
		DTIME:  1,
	}
	WIZARD = Card{
		NAME:   "Wizard",
		ARENA:  ARENA_5,
		RARITY: RARE,
		TYPE:   TROOP,
		DESC:   `The most awesome man to ever set foot in the arena, the Wizard will blow you away with his handsomeness... and/or fireballs.`,
		COST:   5,
		HP:     []int{340, X, 411, 452, 496, X, X, 656, 720, 792},
		DPS:    []int{76, X, 92, 101, 111, X, X, 147, 161, 177},
		ADAM:   []int{130, X, 157, 172, 189, X, X, 250, 275, 302},
		HSPD:   1.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}
	WITCH = Card{
		NAME:   "Witch",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Summons skeletons, shoots destructo beams, has glowing pink eyes that unfortunately don't shoot lasers.`,
		COST:   5,
		HP:     []int{500, 550, 605, X, 730, 800, 880, 968},
		DPS:    []int{51, 55, 61, X, 74, 81, 90, X},
		ADAM:   []int{36, 39, 43, 48, 52, 57, 63, 69},
		SKE_LV: []int{6, 7, 8, 9, 10, 11, 12, X},
		SSPD:   7.5,
		HSPD:   0.7,
		TGTS:   AIR_AND_GROUND,
		SPD:    MEDIUM,
		RNG:    5.5,
		DTIME:  1,
	}
	SKELETON_ARMY = Card{
		NAME:   "Skeleton Army",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Summons an army of skeletons. Meet Larry and his friends Harry, Gerry, Terry, Mary, etc.`,
		COST:   4,
		HP:     []int{30, 33, 36, 39, 43, 48, 52, 57},
		DPS:    []int{30, 33, 36, 39, 43, 48, 52, 57},
		DAM:    []int{30, 33, 36, 39, 43, 48, 52, 57},
		HSPD:   1,
		TGTS:   GROUND,
		SPD:    FAST,
		RNG:    MELEE,
		DTIME:  1,
		COUNT:  20,
	}
	BABY_DRAGON = Card{
		NAME:   "Baby Dragon",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Flying troop that deals area damage. Baby dragons hatch cute, hungry, and ready for a barbeque.`,
		COST:   4,
		HP:     []int{800, 880, 968, 1064, 1168, 1280, 1408, 1544},
		DPS:    []int{55, 61, 67, 73, 81, 88, 97, 107},
		ADAM:   []int{100, 110, 121, 133, 146, 160, 176, 193},
		HSPD:   1.8,
		TGTS:   AIR_AND_GROUND,
		SPD:    FAST,
		RNG:    3.5,
		DTIME:  1,
	}
	PRINCE = Card{
		NAME:   "Prince",
		ARENA:  ARENA_0,
		RARITY: EPIC,
		TYPE:   TROOP,
		DESC:   `Don't let the little pony fool you. Once the Prince gets a running start, you WILL be trampled. Does 2x damage once he gets charging.`,
		COST:   5,
		HP:     []int{1100, 1210, 1331, 1463, 1606, 1760, 1936, 2123},
		DPS:    []int{146, 161, 177, 194, 214, 234, 258, 282},
		DAM:    []int{220, 242, 266, 292, 321, 352, 387, 424},
		HSPD:   1.5,
		TGTS:   GROUND,
		SPD:    MEDIUM,
		RNG:    2.5,
		DTIME:  1,
	}
	GIANT_SKELETON = Card{
		HP:   []int{2000, 2200, 2420, X, X, 3200, 3520, X},
		DPS:  []int{66, X, 80, X, X, 106, 117, X},
		DAM:  []int{100, 110, 121, X, 160, 176, X},
		DDAM: []int{720, 792, 871, X, 1152, 1267, X},
	}
	BALLOON = Card{
		HP:   []int{1000, 1100, X, X, 1460, X, 1760, X},
		DPS:  []int{200, 220, X, X, 292, X, 352, X},
		DAM:  []int{600, 660, X, X, 876, X, 1056, X},
		DDAM: []int{100, 110, X, X, 146, X, 176, X},
	}
	PEKKA = Card{
		HP:  []int{2600, 2860, X, X, X, X, 4576, X},
		DPS: []int{250, 275, X, X, X, X, 440, X},
		DAM: []int{450, 495, X, X, X, X, 792, X},
	}
	GOLEM = Card{
		HP:   []int{X, X, X, X, 4380, 4800, 5280, 5790},
		DPS:  []int{X, X, X, X, 108, 118, 130, 143},
		DAM:  []int{X, X, X, X, 271, 297, 327, 358},
		DDAM: []int{X, X, X, X, 271, 297, 327, 358},
	}
)

var CARDS = [...]Card{
	KNIGHT,
	ARCHERS,
	BOMBER,
	GOBLINS,
	SPEAR_GOBLINS,
	SKELETONS,
	MINIONS,
	BARBARIANS,
	MINION_HORDE,
	GIANT,
	MUSKETEER,
	MINI_PEKKA,
	VALKYRIE,
	HOG_RIDER,
	WIZARD,
	WITCH,
	SKELETON_ARMY,
	BABY_DRAGON,
	PRINCE,
}
