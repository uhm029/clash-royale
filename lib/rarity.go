package lib

type RarityAttribute string

const (
	CARDS_REQ = RarityAttribute("Cards Required")
	GOLD_REQ  = RarityAttribute("Gold Required")
	EXP_GAIN  = RarityAttribute("Experience Gained")
)

var RARITY_ATTRIBUTES = [...]RarityAttribute{
	CARDS_REQ,
	GOLD_REQ,
	EXP_GAIN,
}

type Rarity struct {
	Name  string
	Cards []int
	Gold  []int
	Exp   []int
}

func (r *Rarity) String() string {
	return r.Name
}

var (
	COMMON = &Rarity{
		Name:  "Common",
		Cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X, X, X},
		Gold:  []int{0, 5, 20, 50, 150, 400, 1000, 2000, X, X, X, X},
		Exp:   []int{0, 4, 5, 6, 10, 25, 50, 100, X, X, X, X},
	}
	RARE = &Rarity{
		Name:  "Rare",
		Cards: []int{0, 2, 4, 10, 20, 50, 100, 200, X, X},
		Gold:  []int{0, 50, 150, 400, 1000, 2000, X, X, X, X},
		Exp:   []int{0, 6, 10, 25, 50, 100, X, X, X, X},
	}
	EPIC = &Rarity{
		Name:  "Epic",
		Cards: []int{0, 2, 4, 10, 20, 50, 100, 200},
		Gold:  []int{0, 400, 1000, 2000, X, X, X, X},
		Exp:   []int{0, 25, 50, 100, X, X, X, X},
	}
	LEGENDARY = &Rarity{
		Name:  "Legendary",
		Cards: []int{0, 2, 4, 10, 20, 50},
		Gold:  []int{0, 2000, X, X, X, X},
		Exp:   []int{0, 100, X, X, X, X},
	}
)

var RARITIES = [...]*Rarity{
	COMMON,
	RARE,
	EPIC,
	LEGENDARY,
}
