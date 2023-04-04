package block

//go:generate go run ./generator/blockentities/main.go
type Entity interface {
	ID() string
	IsValidBlock(block Block) bool
}

type (
	FurnaceEntity           struct{}
	ChestEntity             struct{}
	TrappedChestEntity      struct{}
	EnderChestEntity        struct{}
	JukeboxEntity           struct{}
	DispenserEntity         struct{}
	DropperEntity           struct{}
	SignEntity              struct{}
	HangingSignEntity       struct{}
	MobSpawnerEntity        struct{}
	PistonEntity            struct{}
	BrewingStandEntity      struct{}
	EnchantingTableEntity   struct{}
	EndPortalEntity         struct{}
	BeaconEntity            struct{}
	SkullEntity             struct{}
	DaylightDetectorEntity  struct{}
	HopperEntity            struct{}
	ComparatorEntity        struct{}
	BannerEntity            struct{}
	StructureBlockEntity    struct{}
	EndGatewayEntity        struct{}
	CommandBlockEntity      struct{}
	ShulkerBoxEntity        struct{}
	BedEntity               struct{}
	ConduitEntity           struct{}
	BarrelEntity            struct{}
	SmokerEntity            struct{}
	BlastFurnaceEntity      struct{}
	LecternEntity           struct{}
	BellEntity              struct{}
	JigsawEntity            struct{}
	CampfireEntity          struct{}
	BeehiveEntity           struct{}
	SculkSensorEntity       struct{}
	SculkCatalystEntity     struct{}
	SculkShriekerEntity     struct{}
	ChiseledBookshelfEntity struct{}
	SuspiciousSandEntity    struct{}
	DecoratedPotEntity      struct{}
)

type EntityType int32

var EntityTypes map[string]EntityType

func init() {
	EntityTypes = make(map[string]EntityType, len(EntityList))
	for i, v := range EntityList {
		EntityTypes[v.ID()] = EntityType(i)
	}
}
