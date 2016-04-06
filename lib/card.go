package lib

type Card struct {
	fieldMap map[Attribute]interface{}
}

func NewCard(fieldMap map[Attribute]interface{}) *Card {
	return &Card{fieldMap}
}

func (c *Card) GetName() string {
	return (c.fieldMap)[NAME].(string)
}

func (c *Card) GetRarity() *Rarity {
	return (c.fieldMap)[RARITY].(*Rarity)
}

func (c *Card) GetType() Type {
	return (c.fieldMap)[TYPE].(Type)
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
	fas := make([]*FixedAttribute, 0)
	for _, attr := range ATTRIBUTES {
		if c.HasAttribute(attr) {
			if fa, ok := attr.(*FixedAttribute); ok {
				fas = append(fas, fa)
			}
		}
	}
	return fas
}

func (c *Card) GetUpgradableAttributes() []*UpgradableAttribute {
	uas := make([]*UpgradableAttribute, 0)
	for _, attr := range ATTRIBUTES {
		if c.HasAttribute(attr) {
			if ua, ok := attr.(*UpgradableAttribute); ok {
				uas = append(uas, ua)
			}
		}
	}
	return uas
}

func (c *Card) GetValue(attr Attribute) interface{} {
	if value, ok := c.fieldMap[attr]; ok {
		return value
	}
	return c.GetRarity().GetValue(attr)
}

func (c *Card) GetFormattedValue(fattr *FixedAttribute) string {
	if value := c.GetValue(fattr); value != nil {
		return fattr.FormatValue(value)
	}
	return ""
}

func (c *Card) GetFormattedValues(uattr *UpgradableAttribute) []string {
	max := c.GetMaxLevel()
	if values := c.GetValue(uattr); values != nil {
		return uattr.FormatValues(values)[0:max:max]
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
