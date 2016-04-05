package lib

type Card map[Attribute]interface{}

func (c *Card) HasAttribute(attr Attribute) bool {
	if _, ok := (*c)[attr]; ok {
		return true
	}
	return (*c)[RARITY].(*Rarity).HasAttribute(attr)
}

func (c *Card) GetFixedAttributes() []*FixedAttribute {
	fas := []*FixedAttribute{}
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
	uas := []*UpgradableAttribute{}
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
	if cattr, ok := (*c)[attr]; ok {
		return cattr
	}
	return (*c)[RARITY].(*Rarity).GetValue(attr)
}

func (c *Card) GetFormattedValue(fattr *FixedAttribute) string {
	if value := c.GetValue(fattr); value != nil {
		return fattr.FormatValue(value)
	}
	return ""
}

func (c *Card) GetFormattedValues(uattr *UpgradableAttribute) []string {
	if values := c.GetValue(uattr); values != nil {
		return uattr.FormatValues(values)
	}
	return nil
}

func (c *Card) GetMaxLevel() int {
	return (*c)[RARITY].(*Rarity).GetMaxLevel()
}

var CARDS = [...]Card{
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

	// --- Rare Troops ---
	GIANT,
	MUSKETEER,
	MINI_PEKKA,
	VALKYRIE,
	HOG_RIDER,
	WIZARD,

	// --- Epic Troops ---
	WITCH,
	SKELETON_ARMY,
	BABY_DRAGON,
	PRINCE,
	GIANT_SKELETON,
	BALLOON,
	PEKKA,
	GOLEM,

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
