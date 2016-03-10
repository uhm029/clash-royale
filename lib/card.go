package lib

type Card map[Attribute]interface{}

func (c *Card) GetValue(attr Attribute) interface{} {
	if cattr, ok := (*c)[attr]; ok {
		return cattr
	}
	return (*c)[RARITY].(*Rarity).GetValue(attr);
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
}
