package lib

import (
	"sort"
)

// Rarity
type Rarity struct {
	id       int
	name     string // The name of the rarity
	cardsReq []int  // The number of cards needed for upgrading card at level "i"
	goldReq  []int  // The amount of gold needed for upgrading card at level "i"
	expGain  []int  // The amount of experience gained when upgrading card at level "i"
	goldCost []int  // The amount of gold needed to buy the "i + 1"-th card from the shop
}

func (r *Rarity) Id() int {
	return r.id
}

func (r *Rarity) String() string {
	return r.name
}

func (r *Rarity) Name() string {
	return r.name
}

func (r *Rarity) CardsReq() []int {
	return r.cardsReq
}

func (r *Rarity) GoldReq() []int {
	return r.goldReq
}

func (r *Rarity) ExpGain() []int {
	return r.expGain
}

func (r *Rarity) HasAttribute(attr Attribute) bool {
	switch attr {
	case CARDS_REQ:
		return true
	case GOLD_REQ:
		return true
	case EXP_GAIN:
		return true
	default:
		return false
	}
}

func (r *Rarity) Value(attr Attribute) interface{} {
	switch attr {
	case CARDS_REQ:
		return r.cardsReq
	case GOLD_REQ:
		return r.goldReq
	case EXP_GAIN:
		return r.expGain
	default:
		return nil
	}
}

func (r *Rarity) MaxLevel() int {
	return len(r.cardsReq)
}

// static
var (
	rarities = []*Rarity{}
)

// constructor
func newRarity(id int, name string, cardsReq, goldReq, expGain, goldCost []int) *Rarity {
	r := &Rarity{
		id,
		name,
		cardsReq,
		goldReq,
		expGain,
		goldCost,
	}
	rarities = append(rarities, r)
	return r
}

func ForEachRarity(f func(*Rarity)) {
	for _, r := range rarities {
		f(r)
	}
}

var (
	COMMON = newRarity(
		0,
		"Common",
		[]int{0, 2, 4, 10, 20, 50, 100, 200, 400, 1000, 2000, 4000},
		[]int{0, 5, 20, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		[]int{0, 4, 5, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		[]int{
			3, 4, 5, 6, 7, 8, 10, 12, 14, 17,
			20, 24, 29, 35, 42, 50, 60, 72, 86, 103,
			124, 149, 179, 210, 258, 310, 372, 446, 535, 642,
			770, 924, 1109, 1331, 1597, 1916, 2299, 2759, 3311, 3973,
			4768, 5722, 6866, 8239, 9887, 11864, 14237, 17084, 20501, 24601,
		},
	)
	RARE = newRarity(
		1,
		"Rare",
		[]int{0, 2, 4, 10, 20, 50, 100, 200, 400, 1000},
		[]int{0, 50, 150, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		[]int{0, 6, 10, 25, 50, 100, 200, 400, 800, 1600},
		[]int{
			40, 56, 78, 109, 153, 214, 300, 420, 588, 823,
			1152, 1613, 2258, 3161, 4425, 6195, 8673, 12142, 16999, 23799,
		},
	)
	EPIC = newRarity(
		2,
		"Epic",
		[]int{0, 2, 4, 10, 20, 50, 100, 200},
		[]int{0, 400, 1000, 2000, 4000, 8000, 20000, 50000},
		[]int{0, 25, 50, 100, 200, 400, 800, 1600},
		[]int{
			2000, 3600, 6480, 11664, 20995, 37791, 68024, 122443, 220397, 396715,
		},
	)
	LEGENDARY = newRarity(
		3,
		"Legendary",
		[]int{0, 2, 4, 10, 20, 50},
		[]int{0, 5000, 20000, 50000, 100000, 250000},
		[]int{0, 200, 400, 800, 1600, 3200},
		[]int{},
	)
)

// Initialization
type raritySlice []*Rarity

func (s raritySlice) Len() int {
	return len(s)
}

func (s raritySlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s raritySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(raritySlice(rarities))
}
