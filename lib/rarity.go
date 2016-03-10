package lib

type Rarity struct {
	name  string
	cards []int
	gold  []int
	exp   []int
}

func (r *Rarity) String() string {
	return r.name
}

func (r *Rarity) GetValue(attr Attribute) interface{} {
	switch attr {
	case CARDS_REQ:
		return r.cards
	case GOLD_REQ:
		return r.gold
	case EXP_GAIN:
		return r.exp
	default:
		return nil
	}
}

var (
	COMMON = &Rarity{
		name:  "Common",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X, X, X},
		gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 4, 5, 6, 10, 25, 50, 100, X, X, X, X},
	}
	RARE = &Rarity{
		name:  "Rare",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X},
		gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 6, 10, 25, 50, 100, X, X, X, X},
	}
	EPIC = &Rarity{
		name:  "Epic",
		cards: []int{0, 2, 4, 10, 20, 50, 100, 200},
		gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		exp:   []int{0, 25, 50, 100, X, X, X, X},
	}
	LEGENDARY = &Rarity{
		name:  "Legendary",
		cards: []int{0, 2, 4, 10, 20, 50},
		gold:  []int{0, 2000, X, X, X, X},
		exp:   []int{0, 100, X, X, X, X},
	}
)

var RARITIES = [...]*Rarity{
	COMMON,
	RARE,
	EPIC,
	LEGENDARY,
}
