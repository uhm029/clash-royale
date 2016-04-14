package lib

import (
	"sort"
)

// Card
type Card struct {
	id       int
	fieldMap map[Attribute]interface{}
}

func (c *Card) Name() string {
	return (c.fieldMap)[NAME].(string)
}

func (c *Card) Arena() *Arena {
	return (c.fieldMap)[ARENA].(*Arena)
}

func (c *Card) Rarity() *Rarity {
	return (c.fieldMap)[RARITY].(*Rarity)
}

func (c *Card) Type() *Type {
	return (c.fieldMap)[TYPE].(*Type)
}

func (c *Card) Description() string {
	return (c.fieldMap)[DESC].(string)
}

func (c *Card) Cost() int {
	return (c.fieldMap)[COST].(int)
}

func (c *Card) MaxLevel() int {
	return c.Rarity().MaxLevel()
}

func (c *Card) HasAttribute(attr Attribute) bool {
	if _, ok := c.fieldMap[attr]; ok {
		return true
	}
	return c.Rarity().HasAttribute(attr)
}

func (c *Card) Value(attr Attribute) interface{} {
	if value, ok := c.fieldMap[attr]; ok {
		return value
	}
	return c.Rarity().Value(attr)
}

func (c *Card) FormattedValue(fattr *FixedAttribute) string {
	if value := c.Value(fattr); value != nil {
		return fattr.FormatValue(value)
	}
	return ""
}

func (c *Card) FormattedValues(uattr *UpgradableAttribute) []string {
	max := c.MaxLevel()
	if values := c.Value(uattr); values != nil {
		return uattr.FormatValues(values)[0:max:max]
	}
	return nil
}

func (c *Card) ForEachFixedAttribute(f func(*FixedAttribute)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	ForEachFixedAttribute(func(attr *FixedAttribute) {
		if c.HasAttribute(attr) {
			f(attr)
		}
	})
}

func (c *Card) ForEachUpgradableAttribute(f func(*UpgradableAttribute)) {
	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	ForEachUpgradableAttribute(func(attr *UpgradableAttribute) {
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
func newCard(id int, fieldMap map[Attribute]interface{}) *Card {
	// "Compile" the "GeneratedAttribute"s to "UpgradableAttribute"s
	for k, v := range fieldMap {
		if attr, ok := k.(*GeneratedAttribute); ok {
			// Generate values for the GeneratedAttribute
			fieldMap[attr.uattr] = attr.GenerateValues(v)
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

func ForEachCardOfArena(a *Arena, f func(*Card)) {
	for _, c := range cards {
		if c.Arena() == a {
			f(c)
		}
	}
}

func ForEachCardOfRarity(r *Rarity, f func(*Card)) {
	for _, c := range cards {
		if c.Rarity() == r {
			f(c)
		}
	}
}

func ForEachCardOfType(t *Type, f func(*Card)) {
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
