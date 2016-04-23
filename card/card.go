package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/typ"
)

// Card
type Card int8

const (
	// --- Common Troops ---
	Knight Card = iota
	Archers
	Bomber
	Goblins
	SpearGoblins
	Skeletons
	Minions
	Barbarians
	MinionHorde
	RoyaleGiant
	// --- Rare Troops ---
	Giant
	Musketeer
	MiniPekka
	Valkyrie
	HogRider
	Wizard
	ThreeMusketeers
	// --- Epic Troops ---
	Witch
	SkeletonArmy
	BabyDragon
	Prince
	GiantSkeleton
	Balloon
	Pekka
	Golem
	DarkPrince
	// --- Legendary Troops ---
	IceWizard
	Princess

	// --- Common Buildings ---
	Cannon
	Tesla
	Mortar
	// --- Rare Buildings ---
	GobllinHut
	BombTower
	Tombstone
	BarbarianHut
	InfernoTower
	ElixirCollector
	// --- Epic Buildings ---
	XBow

	// --- Common Spells ---
	Arrows
	Zap
	// --- Rare Spells ---
	Fireball
	Rocket
	// --- Epic Spells ---
	Lightning
	GoblinBarrel
	Rage
	Freeze
	Mirror
	Poison
)

func ForEach(f func(Card)) {
	for i := range cards {
		f(Card(i))
	}
}

func ForEachOfArena(a arena.Arena, f func(Card)) {
	ForEach(func(c Card) {
		if c.Arena() == a {
			f(c)
		}
	})
}

func ForEachOfRarity(r rarity.Rarity, f func(Card)) {
	ForEach(func(c Card) {
		if c.Rarity() == r {
			f(c)
		}
	})
}

func ForEachOfType(t typ.Type, f func(Card)) {
	ForEach(func(c Card) {
		if c.Type() == t {
			f(c)
		}
	})
}

func (c Card) Name() string {
	return cards[c][attr.Name].(string)
}

func (c Card) Arena() arena.Arena {
	return cards[c][attr.Arena].(arena.Arena)
}

func (c Card) Rarity() rarity.Rarity {
	return cards[c][attr.Rarity].(rarity.Rarity)
}

func (c Card) Type() typ.Type {
	return cards[c][attr.Type].(typ.Type)
}

func (c Card) Description() string {
	return cards[c][attr.Desc].(string)
}

func (c Card) Elixir() int {
	return cards[c][attr.Elixir].(int)
}

func (c Card) MaxLevel() int {
	return c.Rarity().MaxLevel()
}

func (c Card) HasAttribute(a attr.Attribute) bool {
	if _, ok := cards[c][a]; ok {
		return true
	}
	return c.Rarity().HasAttribute(a)
}

func (c Card) Value(a attr.Attribute) interface{} {
	if value, ok := cards[c][a]; ok {
		return value
	}
	return c.Rarity().Value(a)
}

func (c Card) Values(a attr.Upgradable) []interface{} {
	return c.Value(a).([]interface{})
}

func (c Card) ValueAtLevel(a attr.Attribute, level int) interface{} {
	switch a := a.(type) {
	case attr.Fixed:
		return c.Value(a)
	case attr.Upgradable:
		return c.Value(a).([]interface{})[level-1]
	case attr.Generated:
		return c.ValueAtLevel(a.TargetAttribute(), level)
	default:
		panic("Unknown attribute type")
	}
}

func (c Card) FormattedValue(a attr.Fixed) string {
	if value := c.Value(a); value != nil {
		return a.FormatValue(value)
	}
	return ""
}

func (c Card) FormattedValues(a attr.Upgradable) []string {
	if values := c.Values(a); values != nil {
		return a.FormatValues(values)
	}
	return nil
}

func (c Card) FormattedValueAtLevel(a attr.Attribute, level int) string {
	switch a := a.(type) {
	case attr.Fixed:
		return c.FormattedValue(a)
	case attr.Upgradable:
		return c.FormattedValues(a)[level-1]
	case attr.Generated:
		return c.FormattedValueAtLevel(a.TargetAttribute(), level)
	default:
		panic("Unknown attribute type")
	}
}

func (c Card) ForEachFixedAttribute(f func(attr.Fixed)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attr.ForEachFixed(func(a attr.Fixed) {
		if c.HasAttribute(a) {
			f(a)
		}
	})
}

func (c Card) ForEachUpgradableAttribute(f func(attr.Upgradable)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attr.ForEachUpgradable(func(a attr.Upgradable) {
		if c.HasAttribute(a) {
			f(a)
		}
	})
}

/////////////
// Private //
/////////////

type card map[attr.Attribute]interface{}

var cards = append(troops[:], append(buildings[:], spells[:]...)...)

// constructor
func (c card) init() card {
	max := c[attr.Rarity].(rarity.Rarity).MaxLevel()
	// "Compile" the "Generated"s to "Upgradable"s
	for k, v := range c {
		if attr, ok := k.(attr.Generated); ok {
			c[attr.TargetAttribute()] = attr.GenerateValues(v, max)
		}
	}

	return c
}
