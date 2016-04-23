package rarity

import (
	"github.com/asukakenji/clash-royale/attr"
)

// Rarity
type Rarity int8

const (
	Common Rarity = iota
	Rare
	Epic
	Legendary
)

func ForEach(f func(Rarity)) {
	for i := range rarities {
		f(Rarity(i))
	}
}

func (r Rarity) Id() int {
	return int(r)
}

func (r Rarity) String() string {
	return rarities[r].name
}

func (r Rarity) Name() string {
	return rarities[r].name
}

func (r Rarity) CardsReq() []interface{} {
	return rarities[r].cardsReq
}

func (r Rarity) GoldReq() []interface{} {
	return rarities[r].goldReq
}

func (r Rarity) ExpGain() []interface{} {
	return rarities[r].expGain
}

func (r Rarity) MaxLevel() int {
	return len(r.CardsReq())
}

func (r Rarity) HasAttribute(a attr.Attribute) bool {
	switch a {
	case attr.CardsReq:
		return true
	case attr.GoldReq:
		return true
	case attr.ExpGain:
		return true
	default:
		return false
	}
}

func (r Rarity) Value(a attr.Attribute) []interface{} {
	switch a {
	case attr.CardsReq:
		return r.CardsReq()
	case attr.GoldReq:
		return r.GoldReq()
	case attr.ExpGain:
		return r.ExpGain()
	default:
		return nil
	}
}

/////////////
// Private //
/////////////

type rarity struct {
	name     string        // The name of the rarity
	cardsReq []interface{} // The number of cards needed for upgrading card at level "i"
	goldReq  []interface{} // The amount of gold needed for upgrading card at level "i"
	expGain  []interface{} // The amount of experience gained when upgrading card at level "i"
	goldCost []interface{} // The amount of gold needed to buy the "i + 1"-th card from the shop
}

var rarities = []*rarity{
	&rarity{
		name:     "Common",
		cardsReq: []interface{}{0, 2, 4, 10, 20, 50, 100, 200, 400, 1000, 2000, 4000},
		goldReq:  []interface{}{0, 5, 20, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []interface{}{0, 4, 5, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []interface{}{
			3, 4, 5, 6, 7, 8, 10, 12, 14, 17,
			20, 24, 29, 35, 42, 50, 60, 72, 86, 103,
			124, 149, 179, 210, 258, 310, 372, 446, 535, 642,
			770, 924, 1109, 1331, 1597, 1916, 2299, 2759, 3311, 3973,
			4768, 5722, 6866, 8239, 9887, 11864, 14237, 17084, 20501, 24601,
		},
	},
	&rarity{
		name:     "Rare",
		cardsReq: []interface{}{0, 2, 4, 10, 20, 50, 100, 200, 400, 1000},
		goldReq:  []interface{}{0, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []interface{}{0, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []interface{}{
			40, 56, 78, 109, 153, 214, 300, 420, 588, 823,
			1152, 1613, 2258, 3161, 4425, 6195, 8673, 12142, 16999, 23799,
		},
	},
	&rarity{
		name:     "Epic",
		cardsReq: []interface{}{0, 2, 4, 10, 20, 50, 100, 200},
		goldReq:  []interface{}{0, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		expGain:  []interface{}{0, 25, 50, 100, 200, 400, 800, 1600},
		goldCost: []interface{}{
			2000, 3600, 6480, 11664, 20995, 37791, 68024, 122443, 220397, 396715,
		},
	},
	&rarity{
		name:     "Legendary",
		cardsReq: []interface{}{0, 2, 4, 10, 20, 50},
		goldReq:  []interface{}{0, 5000, 20000, 50000, 100000, 250000},
		expGain:  []interface{}{0, 200, 400, 800, 1600, 3200},
		goldCost: []interface{}{},
	},
}
