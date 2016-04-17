package card

import (
	"github.com/asukakenji/clash-royale/Type"
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attr"
	"github.com/asukakenji/clash-royale/rarity"

	"sort"
)

// Card
type Card struct {
	id       int
	fieldMap map[attr.Attribute]interface{}
}

func (c *Card) Name() string {
	return (c.fieldMap)[attr.Name].(string)
}

func (c *Card) Arena() arena.Arena {
	return (c.fieldMap)[attr.Arena].(arena.Arena)
}

func (c *Card) Rarity() rarity.Rarity {
	return (c.fieldMap)[attr.Rarity].(rarity.Rarity)
}

func (c *Card) Type() Type.Type {
	return (c.fieldMap)[attr.Type].(Type.Type)
}

func (c *Card) Description() string {
	return (c.fieldMap)[attr.Desc].(string)
}

func (c *Card) Elixir() int {
	return (c.fieldMap)[attr.Elixir].(int)
}

func (c *Card) MaxLevel() int {
	return c.Rarity().MaxLevel()
}

func (c *Card) HasAttribute(attr attr.Attribute) bool {
	if _, ok := c.fieldMap[attr]; ok {
		return true
	}
	return c.Rarity().HasAttribute(attr)
}

func (c *Card) Value(attr attr.Attribute) interface{} {
	if value, ok := c.fieldMap[attr]; ok {
		return value
	}
	return c.Rarity().Value(attr)
}

func (c *Card) FormattedValue(fattr attr.Fixed) string {
	if value := c.Value(fattr); value != nil {
		return fattr.FormatValue(value)
	}
	return ""
}

func (c *Card) FormattedValues(uattr attr.Upgradable) []string {
	max := c.MaxLevel()
	if values := c.Value(uattr); values != nil {
		return uattr.FormatValues(values)[0:max:max]
	}
	return nil
}

func (c *Card) ForEachFixedAttribute(f func(attr.Fixed)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attr.ForEachFixed(func(attr attr.Fixed) {
		if c.HasAttribute(attr) {
			f(attr)
		}
	})
}

func (c *Card) ForEachUpgradableAttribute(f func(attr.Upgradable)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attr.ForEachUpgradable(func(attr attr.Upgradable) {
		if c.HasAttribute(attr) {
			f(attr)
		}
	})
}

// static
var (
	cards = []*Card{}
)

// constructor
func newCard(id int, fieldMap map[attr.Attribute]interface{}) *Card {
	// "Compile" the "Generated"s to "Upgradable"s
	for k, v := range fieldMap {
		if attr, ok := k.(*attr.Generated); ok {
			// Generate values for the Generated
			fieldMap[attr.TargetAttribute()] = attr.GenerateValues(v)
		}
	}

	c := &Card{
		id,
		fieldMap,
	}
	cards = append(cards, c)
	return c
}

func ForEachCard(f func(*Card)) {
	for _, c := range cards {
		f(c)
	}
}

func ForEachCardOfArena(a arena.Arena, f func(*Card)) {
	for _, c := range cards {
		if c.Arena() == a {
			f(c)
		}
	}
}

func ForEachCardOfRarity(r rarity.Rarity, f func(*Card)) {
	for _, c := range cards {
		if c.Rarity() == r {
			f(c)
		}
	}
}

func ForEachCardOfType(t Type.Type, f func(*Card)) {
	for _, c := range cards {
		if c.Type() == t {
			f(c)
		}
	}
}

// Initialization
type cardSlice []*Card

func (s cardSlice) Len() int {
	return len(s)
}

func (s cardSlice) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s cardSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func init() {
	sort.Sort(cardSlice(cards))
}
