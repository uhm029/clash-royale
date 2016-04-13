package lib

// Card
type Card struct {
	fixedAttributes      []*FixedAttribute
	upgradableAttributes []*UpgradableAttribute
	fieldMap             map[Attribute]interface{}
}

func newCard(fieldMap map[Attribute]interface{}) *Card {
	fixedAttributes := []*FixedAttribute{}
	upgradableAttributes := []*UpgradableAttribute{}

	// "Compile" the "GeneratedAttribute"s to "UpgradableAttribute"s
	for k, v := range fieldMap {
		if attr, ok := k.(*GeneratedAttribute); ok {
			// Generate values for the GeneratedAttribute
			fieldMap[attr.uattr] = attr.generateValues(v)
		}
	}

	// Note:
	// It is necessary to iterate ATTRIBUTES instead of fieldMap,
	// since the order of the keys in fieldMap is random.
	for _, k := range ATTRIBUTES {
		if _, ok := fieldMap[k]; ok {
			switch attr := k.(type) {
			case *FixedAttribute:
				fixedAttributes = append(fixedAttributes, attr)
			case *UpgradableAttribute:
				upgradableAttributes = append(upgradableAttributes, attr)
			}
		}
	}

	upgradableAttributes = append(upgradableAttributes, CARDS_REQ, GOLD_REQ, EXP_GAIN)

	return &Card{
		fixedAttributes,
		upgradableAttributes,
		fieldMap,
	}
}

func (c *Card) GetName() string {
	return (c.fieldMap)[NAME].(string)
}

func (c *Card) GetArena() *Arena {
	return (c.fieldMap)[ARENA].(*Arena)
}

func (c *Card) GetRarity() *Rarity {
	return (c.fieldMap)[RARITY].(*Rarity)
}

func (c *Card) GetType() *Type {
	return (c.fieldMap)[TYPE].(*Type)
}

func (c *Card) GetDescription() string {
	return (c.fieldMap)[DESC].(string)
}

func (c *Card) GetCost() int {
	return (c.fieldMap)[COST].(int)
}

func (c *Card) GetMaxLevel() int {
	return c.GetRarity().GetMaxLevel()
}

func (c *Card) HasAttribute(attr Attribute) bool {
	if _, ok := c.fieldMap[attr]; ok {
		return true
	}
	return c.GetRarity().HasAttribute(attr)
}

func (c *Card) GetFixedAttributes() []*FixedAttribute {
	return c.fixedAttributes
}

func (c *Card) GetUpgradableAttributes() []*UpgradableAttribute {
	return c.upgradableAttributes
}

func (c *Card) GetValue(attr Attribute) interface{} {
	if value, ok := c.fieldMap[attr]; ok {
		return value
	}
	return c.GetRarity().GetValue(attr)
}

func (c *Card) GetFormattedValue(fattr *FixedAttribute) string {
	if value := c.GetValue(fattr); value != nil {
		return fattr.formatValue(value)
	}
	return ""
}

func (c *Card) GetFormattedValues(uattr *UpgradableAttribute) []string {
	max := c.GetMaxLevel()
	if values := c.GetValue(uattr); values != nil {
		return uattr.formatValues(values)[0:max:max]
	}
	return nil
}

var CARDS = [...]*Card{
	// --- Common Troops ---
	KNIGHT,
	ARCHERS,
	BOMBER,
	GOBLINS,
	SPEAR_GOBLINS,
	SKELETONS,
	MINIONS,
	BARBARIANS,
	MINION_HORDE,
	ROYALE_GIANT,

	// --- Rare Troops ---
	GIANT,
	MUSKETEER,
	MINI_PEKKA,
	VALKYRIE,
	HOG_RIDER,
	WIZARD,
	THREE_MUSKETEERS,

	// --- Epic Troops ---
	WITCH,
	SKELETON_ARMY,
	BABY_DRAGON,
	PRINCE,
	GIANT_SKELETON,
	BALLOON,
	PEKKA,
	GOLEM,
	DARK_PRINCE,

	// --- Legendary Troops ---
	ICE_WIZARD,
	PRINCESS,

	// --- Common Buildings ---
	CANNON,
	TESLA,
	MORTAR,

	// --- Rare Buildings ---
	GOBLIN_HUT,
	BOMB_TOWER,
	TOMBSTONE,
	BARBARIAN_HUT,
	INFERNO_TOWER,
	ELIXIR_COLLECTOR,

	// --- Epic Buildings ---
	X_BOW,

	// --- Common Spells ---
	ARROWS,
	ZAP,

	// --- Rare Spells ---
	FIREBALL,
	ROCKET,

	// --- Epic Spells ---
	LIGHTNING,
	GOBLIN_BARREL,
	RAGE,
	FREEZE,
	MIRROR,
	POISON,
}
