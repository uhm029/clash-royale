package card

import (
	"github.com/asukakenji/clash-royale/arena"
	"github.com/asukakenji/clash-royale/attribute"
	"github.com/asukakenji/clash-royale/rarity"
	"github.com/asukakenji/clash-royale/types"

	"sort"
)

// Card
type Card struct {
	id       int
	fieldMap map[attribute.Attribute]interface{}
}

func (c *Card) Name() string {
	return (c.fieldMap)[attribute.Name].(string)
}

func (c *Card) Arena() arena.Arena {
	return (c.fieldMap)[attribute.Arena].(arena.Arena)
}

func (c *Card) Rarity() rarity.Rarity {
	return (c.fieldMap)[attribute.Rarity].(rarity.Rarity)
}

func (c *Card) Type() types.Type {
	return (c.fieldMap)[attribute.Type].(types.Type)
}

func (c *Card) Description() string {
	return (c.fieldMap)[attribute.Desc].(string)
}

func (c *Card) Elixir() int {
	return (c.fieldMap)[attribute.Elixir].(int)
}

func (c *Card) MaxLevel() int {
	return c.Rarity().MaxLevel()
}

func (c *Card) HasAttribute(attr attribute.Attribute) bool {
	if _, ok := c.fieldMap[attr]; ok {
		return true
	}
	return c.Rarity().HasAttribute(attr)
}

func (c *Card) Value(attr attribute.Attribute) interface{} {
	if value, ok := c.fieldMap[attr]; ok {
		return value
	}
	return c.Rarity().Value(attr)
}

func (c *Card) FormattedValue(fattr attribute.Fixed) string {
	if value := c.Value(fattr); value != nil {
		return fattr.FormatValue(value)
	}
	return ""
}

func (c *Card) FormattedValues(uattr attribute.Upgradable) []string {
	max := c.MaxLevel()
	if values := c.Value(uattr); values != nil {
		return uattr.FormatValues(values)[0:max:max]
	}
	return nil
}

func (c *Card) ForEachFixedAttribute(f func(attribute.Fixed)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attribute.ForEachFixed(func(attr attribute.Fixed) {
		if c.HasAttribute(attr) {
			f(attr)
		}
	})
}

func (c *Card) ForEachUpgradableAttribute(f func(attribute.Upgradable)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	attribute.ForEachUpgradable(func(attr attribute.Upgradable) {
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
func newCard(id int, fieldMap map[attribute.Attribute]interface{}) *Card {
	// "Compile" the "Generated"s to "Upgradable"s
	for k, v := range fieldMap {
		if attr, ok := k.(*attribute.Generated); ok {
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

func ForEachCardOfType(t types.Type, f func(*Card)) {
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
