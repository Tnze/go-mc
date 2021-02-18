// Package entity stores information about entities in Minecraft.
package entity

// ID describes the numeric ID of an entity.
type ID uint32

// Category groups like entities.
type Category uint8

// Valid entity categories.
const (
	Unknown Category = iota
	Blocks
	Immobile
	Vehicles
	Drops
	Projectiles
	PassiveMob
	HostileMob
)

// Entity describes information about a type of entity.
type Entity struct {
	ID          ID
	InternalID  uint32
	DisplayName string
	Name        string

	Width  float64
	Height float64

	Type     string
	Category Category
}

//goland:noinspection ALL
var (
	AreaEffectCloud      = Entity{ID: 0, InternalID: 0, DisplayName: "Area Effect Cloud", Name: "area_effect_cloud", Width: 6, Height: 0.5, Type: "mob", Category: Immobile}
	ArmorStand           = Entity{ID: 1, InternalID: 1, DisplayName: "Armor Stand", Name: "armor_stand", Width: 0.5, Height: 1.975, Type: "mob", Category: Immobile}
	Arrow                = Entity{ID: 2, InternalID: 2, DisplayName: "Arrow", Name: "arrow", Width: 0.5, Height: 0.5, Type: "mob", Category: Projectiles}
	Bat                  = Entity{ID: 3, InternalID: 3, DisplayName: "Bat", Name: "bat", Width: 0.5, Height: 0.9, Type: "mob", Category: PassiveMob}
	Bee                  = Entity{ID: 4, InternalID: 4, DisplayName: "Bee", Name: "bee", Width: 0.7, Height: 0.6, Type: "UNKNOWN", Category: Unknown}
	Blaze                = Entity{ID: 5, InternalID: 5, DisplayName: "Blaze", Name: "blaze", Width: 0.6, Height: 1.8, Type: "mob", Category: HostileMob}
	Boat                 = Entity{ID: 6, InternalID: 6, DisplayName: "Boat", Name: "boat", Width: 1.375, Height: 0.5625, Type: "mob", Category: Vehicles}
	Cat                  = Entity{ID: 7, InternalID: 7, DisplayName: "Cat", Name: "cat", Width: 0.6, Height: 0.7, Type: "mob", Category: PassiveMob}
	CaveSpider           = Entity{ID: 8, InternalID: 8, DisplayName: "Cave Spider", Name: "cave_spider", Width: 0.7, Height: 0.5, Type: "mob", Category: HostileMob}
	Chicken              = Entity{ID: 9, InternalID: 9, DisplayName: "Chicken", Name: "chicken", Width: 0.4, Height: 0.7, Type: "mob", Category: PassiveMob}
	Cod                  = Entity{ID: 10, InternalID: 10, DisplayName: "Cod", Name: "cod", Width: 0.5, Height: 0.3, Type: "mob", Category: PassiveMob}
	Cow                  = Entity{ID: 11, InternalID: 11, DisplayName: "Cow", Name: "cow", Width: 0.9, Height: 1.4, Type: "mob", Category: PassiveMob}
	Creeper              = Entity{ID: 12, InternalID: 12, DisplayName: "Creeper", Name: "creeper", Width: 0.6, Height: 1.7, Type: "mob", Category: HostileMob}
	Dolphin              = Entity{ID: 13, InternalID: 13, DisplayName: "Dolphin", Name: "dolphin", Width: 0.9, Height: 0.6, Type: "mob", Category: PassiveMob}
	Donkey               = Entity{ID: 14, InternalID: 14, DisplayName: "Donkey", Name: "donkey", Width: 1.39648, Height: 1.5, Type: "mob", Category: PassiveMob}
	DragonFireball       = Entity{ID: 15, InternalID: 15, DisplayName: "Dragon Fireball", Name: "dragon_fireball", Width: 1, Height: 1, Type: "mob", Category: Projectiles}
	Drowned              = Entity{ID: 16, InternalID: 16, DisplayName: "Drowned", Name: "drowned", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	ElderGuardian        = Entity{ID: 17, InternalID: 17, DisplayName: "Elder Guardian", Name: "elder_guardian", Width: 1.9975, Height: 1.9975, Type: "mob", Category: HostileMob}
	EndCrystal           = Entity{ID: 18, InternalID: 18, DisplayName: "End Crystal", Name: "end_crystal", Width: 2, Height: 2, Type: "mob", Category: Unknown}
	EnderDragon          = Entity{ID: 19, InternalID: 19, DisplayName: "Ender Dragon", Name: "ender_dragon", Width: 16, Height: 8, Type: "mob", Category: HostileMob}
	Enderman             = Entity{ID: 20, InternalID: 20, DisplayName: "Enderman", Name: "enderman", Width: 0.6, Height: 2.9, Type: "mob", Category: HostileMob}
	Endermite            = Entity{ID: 21, InternalID: 21, DisplayName: "Endermite", Name: "endermite", Width: 0.4, Height: 0.3, Type: "mob", Category: HostileMob}
	Evoker               = Entity{ID: 22, InternalID: 22, DisplayName: "Evoker", Name: "evoker", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	EvokerFangs          = Entity{ID: 23, InternalID: 23, DisplayName: "Evoker Fangs", Name: "evoker_fangs", Width: 0.5, Height: 0.8, Type: "mob", Category: HostileMob}
	ExperienceOrb        = Entity{ID: 24, InternalID: 24, DisplayName: "Experience Orb", Name: "experience_orb", Width: 0.5, Height: 0.5, Type: "mob", Category: Unknown}
	EyeOfEnder           = Entity{ID: 25, InternalID: 25, DisplayName: "Eye of Ender", Name: "eye_of_ender", Width: 0.25, Height: 0.25, Type: "mob", Category: Unknown}
	FallingBlock         = Entity{ID: 26, InternalID: 26, DisplayName: "Falling Block", Name: "falling_block", Width: 0.98, Height: 0.98, Type: "mob", Category: Blocks}
	FireworkRocket       = Entity{ID: 27, InternalID: 27, DisplayName: "Firework Rocket", Name: "firework_rocket", Width: 0.25, Height: 0.25, Type: "mob", Category: Unknown}
	Fox                  = Entity{ID: 28, InternalID: 28, DisplayName: "Fox", Name: "fox", Width: 0.6, Height: 0.7, Type: "mob", Category: Unknown}
	Ghast                = Entity{ID: 29, InternalID: 29, DisplayName: "Ghast", Name: "ghast", Width: 4, Height: 4, Type: "mob", Category: HostileMob}
	Giant                = Entity{ID: 30, InternalID: 30, DisplayName: "Giant", Name: "giant", Width: 3.6, Height: 12, Type: "mob", Category: HostileMob}
	Guardian             = Entity{ID: 31, InternalID: 31, DisplayName: "Guardian", Name: "guardian", Width: 0.85, Height: 0.85, Type: "mob", Category: HostileMob}
	Hoglin               = Entity{ID: 32, InternalID: 32, DisplayName: "Hoglin", Name: "hoglin", Width: 1.39648, Height: 1.4, Type: "UNKNOWN", Category: Unknown}
	Horse                = Entity{ID: 33, InternalID: 33, DisplayName: "Horse", Name: "horse", Width: 1.39648, Height: 1.6, Type: "mob", Category: PassiveMob}
	Husk                 = Entity{ID: 34, InternalID: 34, DisplayName: "Husk", Name: "husk", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	Illusioner           = Entity{ID: 35, InternalID: 35, DisplayName: "Illusioner", Name: "illusioner", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	IronGolem            = Entity{ID: 36, InternalID: 36, DisplayName: "Iron Golem", Name: "iron_golem", Width: 1.4, Height: 2.7, Type: "mob", Category: PassiveMob}
	Item                 = Entity{ID: 37, InternalID: 37, DisplayName: "Item", Name: "item", Width: 0.25, Height: 0.25, Type: "mob", Category: Drops}
	ItemFrame            = Entity{ID: 38, InternalID: 38, DisplayName: "Item Frame", Name: "item_frame", Width: 0.5, Height: 0.5, Type: "mob", Category: Immobile}
	Fireball             = Entity{ID: 39, InternalID: 39, DisplayName: "Fireball", Name: "fireball", Width: 1, Height: 1, Type: "mob", Category: Projectiles}
	LeashKnot            = Entity{ID: 40, InternalID: 40, DisplayName: "Leash Knot", Name: "leash_knot", Width: 0.5, Height: 0.5, Type: "mob", Category: Immobile}
	LightningBolt        = Entity{ID: 41, InternalID: 41, DisplayName: "Lightning Bolt", Name: "lightning_bolt", Width: 0, Height: 0, Type: "mob", Category: Unknown}
	Llama                = Entity{ID: 42, InternalID: 42, DisplayName: "Llama", Name: "llama", Width: 0.9, Height: 1.87, Type: "mob", Category: PassiveMob}
	LlamaSpit            = Entity{ID: 43, InternalID: 43, DisplayName: "Llama Spit", Name: "llama_spit", Width: 0.25, Height: 0.25, Type: "mob", Category: Projectiles}
	MagmaCube            = Entity{ID: 44, InternalID: 44, DisplayName: "Magma Cube", Name: "magma_cube", Width: 2.04, Height: 2.04, Type: "mob", Category: HostileMob}
	Minecart             = Entity{ID: 45, InternalID: 45, DisplayName: "Minecart", Name: "minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	ChestMinecart        = Entity{ID: 46, InternalID: 46, DisplayName: "Minecart with Chest", Name: "chest_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	CommandBlockMinecart = Entity{ID: 47, InternalID: 47, DisplayName: "Minecart with Command Block", Name: "command_block_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	FurnaceMinecart      = Entity{ID: 48, InternalID: 48, DisplayName: "Minecart with Furnace", Name: "furnace_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	HopperMinecart       = Entity{ID: 49, InternalID: 49, DisplayName: "Minecart with Hopper", Name: "hopper_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	SpawnerMinecart      = Entity{ID: 50, InternalID: 50, DisplayName: "Minecart with Spawner", Name: "spawner_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	TntMinecart          = Entity{ID: 51, InternalID: 51, DisplayName: "Minecart with TNT", Name: "tnt_minecart", Width: 0.98, Height: 0.7, Type: "mob", Category: Vehicles}
	Mule                 = Entity{ID: 52, InternalID: 52, DisplayName: "Mule", Name: "mule", Width: 1.39648, Height: 1.6, Type: "mob", Category: PassiveMob}
	Mooshroom            = Entity{ID: 53, InternalID: 53, DisplayName: "Mooshroom", Name: "mooshroom", Width: 0.9, Height: 1.4, Type: "mob", Category: PassiveMob}
	Ocelot               = Entity{ID: 54, InternalID: 54, DisplayName: "Ocelot", Name: "ocelot", Width: 0.6, Height: 0.7, Type: "mob", Category: PassiveMob}
	Painting             = Entity{ID: 55, InternalID: 55, DisplayName: "Painting", Name: "painting", Width: 0.5, Height: 0.5, Type: "mob", Category: Immobile}
	Panda                = Entity{ID: 56, InternalID: 56, DisplayName: "Panda", Name: "panda", Width: 1.3, Height: 1.25, Type: "mob", Category: PassiveMob}
	Parrot               = Entity{ID: 57, InternalID: 57, DisplayName: "Parrot", Name: "parrot", Width: 0.5, Height: 0.9, Type: "mob", Category: PassiveMob}
	Phantom              = Entity{ID: 58, InternalID: 58, DisplayName: "Phantom", Name: "phantom", Width: 0.9, Height: 0.5, Type: "mob", Category: HostileMob}
	Pig                  = Entity{ID: 59, InternalID: 59, DisplayName: "Pig", Name: "pig", Width: 0.9, Height: 0.9, Type: "mob", Category: PassiveMob}
	Piglin               = Entity{ID: 60, InternalID: 60, DisplayName: "Piglin", Name: "piglin", Width: 0.6, Height: 1.95, Type: "UNKNOWN", Category: Unknown}
	PiglinBrute          = Entity{ID: 61, InternalID: 61, DisplayName: "Piglin Brute", Name: "piglin_brute", Width: 0.6, Height: 1.95, Type: "UNKNOWN", Category: Unknown}
	Pillager             = Entity{ID: 62, InternalID: 62, DisplayName: "Pillager", Name: "pillager", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	PolarBear            = Entity{ID: 63, InternalID: 63, DisplayName: "Polar Bear", Name: "polar_bear", Width: 1.4, Height: 1.4, Type: "mob", Category: PassiveMob}
	Tnt                  = Entity{ID: 64, InternalID: 64, DisplayName: "Primed TNT", Name: "tnt", Width: 0.98, Height: 0.98, Type: "mob", Category: Blocks}
	Pufferfish           = Entity{ID: 65, InternalID: 65, DisplayName: "Pufferfish", Name: "pufferfish", Width: 0.7, Height: 0.7, Type: "mob", Category: PassiveMob}
	Rabbit               = Entity{ID: 66, InternalID: 66, DisplayName: "Rabbit", Name: "rabbit", Width: 0.4, Height: 0.5, Type: "mob", Category: PassiveMob}
	Ravager              = Entity{ID: 67, InternalID: 67, DisplayName: "Ravager", Name: "ravager", Width: 1.95, Height: 2.2, Type: "mob", Category: HostileMob}
	Salmon               = Entity{ID: 68, InternalID: 68, DisplayName: "Salmon", Name: "salmon", Width: 0.7, Height: 0.4, Type: "mob", Category: PassiveMob}
	Sheep                = Entity{ID: 69, InternalID: 69, DisplayName: "Sheep", Name: "sheep", Width: 0.9, Height: 1.3, Type: "mob", Category: PassiveMob}
	Shulker              = Entity{ID: 70, InternalID: 70, DisplayName: "Shulker", Name: "shulker", Width: 1, Height: 1, Type: "mob", Category: HostileMob}
	ShulkerBullet        = Entity{ID: 71, InternalID: 71, DisplayName: "Shulker Bullet", Name: "shulker_bullet", Width: 0.3125, Height: 0.3125, Type: "mob", Category: Projectiles}
	Silverfish           = Entity{ID: 72, InternalID: 72, DisplayName: "Silverfish", Name: "silverfish", Width: 0.4, Height: 0.3, Type: "mob", Category: HostileMob}
	Skeleton             = Entity{ID: 73, InternalID: 73, DisplayName: "Skeleton", Name: "skeleton", Width: 0.6, Height: 1.99, Type: "mob", Category: HostileMob}
	SkeletonHorse        = Entity{ID: 74, InternalID: 74, DisplayName: "Skeleton Horse", Name: "skeleton_horse", Width: 1.39648, Height: 1.6, Type: "mob", Category: PassiveMob}
	Slime                = Entity{ID: 75, InternalID: 75, DisplayName: "Slime", Name: "slime", Width: 2.04, Height: 2.04, Type: "mob", Category: HostileMob}
	SmallFireball        = Entity{ID: 76, InternalID: 76, DisplayName: "Small Fireball", Name: "small_fireball", Width: 0.3125, Height: 0.3125, Type: "mob", Category: Projectiles}
	SnowGolem            = Entity{ID: 77, InternalID: 77, DisplayName: "Snow Golem", Name: "snow_golem", Width: 0.7, Height: 1.9, Type: "mob", Category: PassiveMob}
	Snowball             = Entity{ID: 78, InternalID: 78, DisplayName: "Snowball", Name: "snowball", Width: 0.25, Height: 0.25, Type: "mob", Category: Projectiles}
	SpectralArrow        = Entity{ID: 79, InternalID: 79, DisplayName: "Spectral Arrow", Name: "spectral_arrow", Width: 0.5, Height: 0.5, Type: "mob", Category: Projectiles}
	Spider               = Entity{ID: 80, InternalID: 80, DisplayName: "Spider", Name: "spider", Width: 1.4, Height: 0.9, Type: "mob", Category: HostileMob}
	Squid                = Entity{ID: 81, InternalID: 81, DisplayName: "Squid", Name: "squid", Width: 0.8, Height: 0.8, Type: "mob", Category: PassiveMob}
	Stray                = Entity{ID: 82, InternalID: 82, DisplayName: "Stray", Name: "stray", Width: 0.6, Height: 1.99, Type: "mob", Category: HostileMob}
	Strider              = Entity{ID: 83, InternalID: 83, DisplayName: "Strider", Name: "strider", Width: 0.9, Height: 1.7, Type: "UNKNOWN", Category: Unknown}
	Egg                  = Entity{ID: 84, InternalID: 84, DisplayName: "Thrown Egg", Name: "egg", Width: 0.25, Height: 0.25, Type: "mob", Category: Projectiles}
	EnderPearl           = Entity{ID: 85, InternalID: 85, DisplayName: "Thrown Ender Pearl", Name: "ender_pearl", Width: 0.25, Height: 0.25, Type: "mob", Category: Projectiles}
	ExperienceBottle     = Entity{ID: 86, InternalID: 86, DisplayName: "Thrown Bottle o' Enchanting", Name: "experience_bottle", Width: 0.25, Height: 0.25, Type: "mob", Category: Unknown}
	Potion               = Entity{ID: 87, InternalID: 87, DisplayName: "Potion", Name: "potion", Width: 0.25, Height: 0.25, Type: "mob", Category: Projectiles}
	Trident              = Entity{ID: 88, InternalID: 88, DisplayName: "Trident", Name: "trident", Width: 0.5, Height: 0.5, Type: "mob", Category: Unknown}
	TraderLlama          = Entity{ID: 89, InternalID: 89, DisplayName: "Trader Llama", Name: "trader_llama", Width: 0.9, Height: 1.87, Type: "mob", Category: PassiveMob}
	TropicalFish         = Entity{ID: 90, InternalID: 90, DisplayName: "Tropical Fish", Name: "tropical_fish", Width: 0.5, Height: 0.4, Type: "mob", Category: PassiveMob}
	Turtle               = Entity{ID: 91, InternalID: 91, DisplayName: "Turtle", Name: "turtle", Width: 1.2, Height: 0.4, Type: "mob", Category: PassiveMob}
	Vex                  = Entity{ID: 92, InternalID: 92, DisplayName: "Vex", Name: "vex", Width: 0.4, Height: 0.8, Type: "mob", Category: HostileMob}
	Villager             = Entity{ID: 93, InternalID: 93, DisplayName: "Villager", Name: "villager", Width: 0.6, Height: 1.95, Type: "mob", Category: PassiveMob}
	Vindicator           = Entity{ID: 94, InternalID: 94, DisplayName: "Vindicator", Name: "vindicator", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	WanderingTrader      = Entity{ID: 95, InternalID: 95, DisplayName: "Wandering Trader", Name: "wandering_trader", Width: 0.6, Height: 1.95, Type: "mob", Category: PassiveMob}
	Witch                = Entity{ID: 96, InternalID: 96, DisplayName: "Witch", Name: "witch", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	Wither               = Entity{ID: 97, InternalID: 97, DisplayName: "Wither", Name: "wither", Width: 0.9, Height: 3.5, Type: "mob", Category: HostileMob}
	WitherSkeleton       = Entity{ID: 98, InternalID: 98, DisplayName: "Wither Skeleton", Name: "wither_skeleton", Width: 0.7, Height: 2.4, Type: "mob", Category: HostileMob}
	WitherSkull          = Entity{ID: 99, InternalID: 99, DisplayName: "Wither Skull", Name: "wither_skull", Width: 0.3125, Height: 0.3125, Type: "mob", Category: Projectiles}
	Wolf                 = Entity{ID: 100, InternalID: 100, DisplayName: "Wolf", Name: "wolf", Width: 0.6, Height: 0.85, Type: "mob", Category: PassiveMob}
	Zoglin               = Entity{ID: 101, InternalID: 101, DisplayName: "Zoglin", Name: "zoglin", Width: 1.39648, Height: 1.4, Type: "UNKNOWN", Category: Unknown}
	Zombie               = Entity{ID: 102, InternalID: 102, DisplayName: "Zombie", Name: "zombie", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	ZombieHorse          = Entity{ID: 103, InternalID: 103, DisplayName: "Zombie Horse", Name: "zombie_horse", Width: 1.39648, Height: 1.6, Type: "mob", Category: PassiveMob}
	ZombieVillager       = Entity{ID: 104, InternalID: 104, DisplayName: "Zombie Villager", Name: "zombie_villager", Width: 0.6, Height: 1.95, Type: "mob", Category: HostileMob}
	ZombifiedPiglin      = Entity{ID: 105, InternalID: 105, DisplayName: "Zombified Piglin", Name: "zombified_piglin", Width: 0.6, Height: 1.95, Type: "UNKNOWN", Category: Unknown}
	Player               = Entity{ID: 106, InternalID: 106, DisplayName: "Player", Name: "player", Width: 0.6, Height: 1.8, Type: "mob", Category: Unknown}
	FishingBobber        = Entity{ID: 107, InternalID: 107, DisplayName: "Fishing Bobber", Name: "fishing_bobber", Width: 0.25, Height: 0.25, Type: "mob", Category: Unknown}
)

// ByID is an index of minecraft entities by their ID.
var ByID = map[ID]*Entity{
	0:   &AreaEffectCloud,
	1:   &ArmorStand,
	2:   &Arrow,
	3:   &Bat,
	4:   &Bee,
	5:   &Blaze,
	6:   &Boat,
	7:   &Cat,
	8:   &CaveSpider,
	9:   &Chicken,
	10:  &Cod,
	11:  &Cow,
	12:  &Creeper,
	13:  &Dolphin,
	14:  &Donkey,
	15:  &DragonFireball,
	16:  &Drowned,
	17:  &ElderGuardian,
	18:  &EndCrystal,
	19:  &EnderDragon,
	20:  &Enderman,
	21:  &Endermite,
	22:  &Evoker,
	23:  &EvokerFangs,
	24:  &ExperienceOrb,
	25:  &EyeOfEnder,
	26:  &FallingBlock,
	27:  &FireworkRocket,
	28:  &Fox,
	29:  &Ghast,
	30:  &Giant,
	31:  &Guardian,
	32:  &Hoglin,
	33:  &Horse,
	34:  &Husk,
	35:  &Illusioner,
	36:  &IronGolem,
	37:  &Item,
	38:  &ItemFrame,
	39:  &Fireball,
	40:  &LeashKnot,
	41:  &LightningBolt,
	42:  &Llama,
	43:  &LlamaSpit,
	44:  &MagmaCube,
	45:  &Minecart,
	46:  &ChestMinecart,
	47:  &CommandBlockMinecart,
	48:  &FurnaceMinecart,
	49:  &HopperMinecart,
	50:  &SpawnerMinecart,
	51:  &TntMinecart,
	52:  &Mule,
	53:  &Mooshroom,
	54:  &Ocelot,
	55:  &Painting,
	56:  &Panda,
	57:  &Parrot,
	58:  &Phantom,
	59:  &Pig,
	60:  &Piglin,
	61:  &PiglinBrute,
	62:  &Pillager,
	63:  &PolarBear,
	64:  &Tnt,
	65:  &Pufferfish,
	66:  &Rabbit,
	67:  &Ravager,
	68:  &Salmon,
	69:  &Sheep,
	70:  &Shulker,
	71:  &ShulkerBullet,
	72:  &Silverfish,
	73:  &Skeleton,
	74:  &SkeletonHorse,
	75:  &Slime,
	76:  &SmallFireball,
	77:  &SnowGolem,
	78:  &Snowball,
	79:  &SpectralArrow,
	80:  &Spider,
	81:  &Squid,
	82:  &Stray,
	83:  &Strider,
	84:  &Egg,
	85:  &EnderPearl,
	86:  &ExperienceBottle,
	87:  &Potion,
	88:  &Trident,
	89:  &TraderLlama,
	90:  &TropicalFish,
	91:  &Turtle,
	92:  &Vex,
	93:  &Villager,
	94:  &Vindicator,
	95:  &WanderingTrader,
	96:  &Witch,
	97:  &Wither,
	98:  &WitherSkeleton,
	99:  &WitherSkull,
	100: &Wolf,
	101: &Zoglin,
	102: &Zombie,
	103: &ZombieHorse,
	104: &ZombieVillager,
	105: &ZombifiedPiglin,
	106: &Player,
	107: &FishingBobber,
}
