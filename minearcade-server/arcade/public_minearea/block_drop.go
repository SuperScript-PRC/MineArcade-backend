package public_minearea

const (
	MINEAREA_DROPITEM_EMPTY = iota
	MINEAREA_DROPITEM_EXP_ORB
	MINEAREA_DROPITEM_COAL
	MINEAREA_DROPITEM_RAW_IRON
	MINEAREA_DROPITEM_RAW_GOLD
	MINEAREA_DROPITEM_DIAMOND
	MINEAREA_DROPITEM_EMERALD
	MINEAREA_DROPITEM_LAPIS_LAZULI
	MINEAREA_DROPITEM_REDSTONE_DUST
)

func GetDropFromBlock(blockID byte) byte {
	switch blockID {
	case Stone:
		return MINEAREA_DROPITEM_EMPTY
	case CoalOre:
		return MINEAREA_DROPITEM_COAL
	case IronOre:
		return MINEAREA_DROPITEM_RAW_IRON
	case GoldOre:
		return MINEAREA_DROPITEM_RAW_GOLD
	case DiamondOre:
		return MINEAREA_DROPITEM_DIAMOND
	case EmeraldOre:
		return MINEAREA_DROPITEM_EMERALD

	}
	return MINEAREA_DROPITEM_EMPTY
}
