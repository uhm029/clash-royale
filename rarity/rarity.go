package rarity

import (
	"github.com/asukakenji/clash-royale/rarity/internal"

	"github.com/asukakenji/clash-royale/attribute"
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
	for i := range internal.Rarities {
		f(Rarity(i))
	}
}

func (r Rarity) Id() int {
	return internal.Rarities[r].Id
}

func (r Rarity) String() string {
	return internal.Rarities[r].Name
}

func (r Rarity) Name() string {
	return internal.Rarities[r].Name
}

func (r Rarity) CardsReq() []int {
	return internal.Rarities[r].CardsReq
}

func (r Rarity) GoldReq() []int {
	return internal.Rarities[r].GoldReq
}

func (r Rarity) ExpGain() []int {
	return internal.Rarities[r].ExpGain
}

func (r Rarity) HasAttribute(attr attribute.Attribute) bool {
	switch attr {
	case attribute.CARDS_REQ:
		return true
	case attribute.GOLD_REQ:
		return true
	case attribute.EXP_GAIN:
		return true
	default:
		return false
	}
}

func (r Rarity) Value(attr attribute.Attribute) []int {
	switch attr {
	case attribute.CARDS_REQ:
		return r.CardsReq()
	case attribute.GOLD_REQ:
		return r.GoldReq()
	case attribute.EXP_GAIN:
		return r.ExpGain()
	default:
		return nil
	}
}

func (r Rarity) MaxLevel() int {
	return len(r.CardsReq())
}