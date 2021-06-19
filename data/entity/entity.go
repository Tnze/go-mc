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


var (
	AreaEffectCloud      = Entity{ID: 0, InternalID: 0, DisplayName: "Area Effect Cloud", Name: "area_effect_cloud", Width: 6, Height: 0.5, Type: "other", Category: }
	ArmorStand           = Entity{ID: 1, InternalID: 1, DisplayName: "Armor Stand", Name: "armor_stand", Width: 0.5, Height: 1.975, Type: "living", Category: }
	Arrow                = Entity{ID: 2, InternalID: 2, DisplayName: "Arrow", Name: "arrow", Width: 0.5, Height: 0.5, Type: "projectile", Category: }
	Axolotl              = Entity{ID: 3, InternalID: 3, DisplayName: "Axolotl", Name: "axolotl", Width: 0.75, Height: 0.42, Type: "animal", Category: }
	Bat                  = Entity{ID: 4, InternalID: 4, DisplayName: "Bat", Name: "bat", Width: 0.5, Height: 0.9, Type: "ambient", Category: }
	Bee                  = Entity{ID: 5, InternalID: 5, DisplayName: "Bee", Name: "bee", Width: 0.7, Height: 0.6, Type: "animal", Category: }
	Blaze                = Entity{ID: 6, InternalID: 6, DisplayName: "Blaze", Name: "blaze", Width: 0.6, Height: 1.8, Type: "hostile", Category: }
	Boat                 = Entity{ID: 7, InternalID: 7, DisplayName: "Boat", Name: "boat", Width: 1.375, Height: 0.5625, Type: "other", Category: }
	Cat                  = Entity{ID: 8, InternalID: 8, DisplayName: "Cat", Name: "cat", Width: 0.6, Height: 0.7, Type: "animal", Category: }
	CaveSpider           = Entity{ID: 9, InternalID: 9, DisplayName: "Cave Spider", Name: "cave_spider", Width: 0.7, Height: 0.5, Type: "hostile", Category: }
	Chicken              = Entity{ID: 10, InternalID: 10, DisplayName: "Chicken", Name: "chicken", Width: 0.4, Height: 0.7, Type: "animal", Category: }
	Cod                  = Entity{ID: 11, InternalID: 11, DisplayName: "Cod", Name: "cod", Width: 0.5, Height: 0.3, Type: "water_creature", Category: }
	Cow                  = Entity{ID: 12, InternalID: 12, DisplayName: "Cow", Name: "cow", Width: 0.9, Height: 1.4, Type: "animal", Category: }
	Creeper              = Entity{ID: 13, InternalID: 13, DisplayName: "Creeper", Name: "creeper", Width: 0.6, Height: 1.7, Type: "hostile", Category: }
	Dolphin              = Entity{ID: 14, InternalID: 14, DisplayName: "Dolphin", Name: "dolphin", Width: 0.9, Height: 0.6, Type: "water_creature", Category: }
	Donkey               = Entity{ID: 15, InternalID: 15, DisplayName: "Donkey", Name: "donkey", Width: 1.3964844, Height: 1.5, Type: "animal", Category: }
	DragonFireball       = Entity{ID: 16, InternalID: 16, DisplayName: "Dragon Fireball", Name: "dragon_fireball", Width: 1, Height: 1, Type: "projectile", Category: }
	Drowned              = Entity{ID: 17, InternalID: 17, DisplayName: "Drowned", Name: "drowned", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	ElderGuardian        = Entity{ID: 18, InternalID: 18, DisplayName: "Elder Guardian", Name: "elder_guardian", Width: 1.9975, Height: 1.9975, Type: "hostile", Category: }
	EndCrystal           = Entity{ID: 19, InternalID: 19, DisplayName: "End Crystal", Name: "end_crystal", Width: 2, Height: 2, Type: "other", Category: }
	EnderDragon          = Entity{ID: 20, InternalID: 20, DisplayName: "Ender Dragon", Name: "ender_dragon", Width: 16, Height: 8, Type: "mob", Category: }
	Enderman             = Entity{ID: 21, InternalID: 21, DisplayName: "Enderman", Name: "enderman", Width: 0.6, Height: 2.9, Type: "hostile", Category: }
	Endermite            = Entity{ID: 22, InternalID: 22, DisplayName: "Endermite", Name: "endermite", Width: 0.4, Height: 0.3, Type: "hostile", Category: }
	Evoker               = Entity{ID: 23, InternalID: 23, DisplayName: "Evoker", Name: "evoker", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	EvokerFangs          = Entity{ID: 24, InternalID: 24, DisplayName: "Evoker Fangs", Name: "evoker_fangs", Width: 0.5, Height: 0.8, Type: "other", Category: }
	ExperienceOrb        = Entity{ID: 25, InternalID: 25, DisplayName: "Experience Orb", Name: "experience_orb", Width: 0.5, Height: 0.5, Type: "other", Category: }
	EyeOfEnder           = Entity{ID: 26, InternalID: 26, DisplayName: "Eye of Ender", Name: "eye_of_ender", Width: 0.25, Height: 0.25, Type: "other", Category: }
	FallingBlock         = Entity{ID: 27, InternalID: 27, DisplayName: "Falling Block", Name: "falling_block", Width: 0.98, Height: 0.98, Type: "other", Category: }
	FireworkRocket       = Entity{ID: 28, InternalID: 28, DisplayName: "Firework Rocket", Name: "firework_rocket", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	Fox                  = Entity{ID: 29, InternalID: 29, DisplayName: "Fox", Name: "fox", Width: 0.6, Height: 0.7, Type: "animal", Category: }
	Ghast                = Entity{ID: 30, InternalID: 30, DisplayName: "Ghast", Name: "ghast", Width: 4, Height: 4, Type: "mob", Category: }
	Giant                = Entity{ID: 31, InternalID: 31, DisplayName: "Giant", Name: "giant", Width: 3.6, Height: 12, Type: "hostile", Category: }
	GlowItemFrame        = Entity{ID: 32, InternalID: 32, DisplayName: "Glow Item Frame", Name: "glow_item_frame", Width: 0.5, Height: 0.5, Type: "other", Category: }
	GlowSquid            = Entity{ID: 33, InternalID: 33, DisplayName: "Glow Squid", Name: "glow_squid", Width: 0.8, Height: 0.8, Type: "water_creature", Category: }
	Goat                 = Entity{ID: 34, InternalID: 34, DisplayName: "Goat", Name: "goat", Width: 0.9, Height: 1.3, Type: "animal", Category: }
	Guardian             = Entity{ID: 35, InternalID: 35, DisplayName: "Guardian", Name: "guardian", Width: 0.85, Height: 0.85, Type: "hostile", Category: }
	Hoglin               = Entity{ID: 36, InternalID: 36, DisplayName: "Hoglin", Name: "hoglin", Width: 1.3964844, Height: 1.4, Type: "animal", Category: }
	Horse                = Entity{ID: 37, InternalID: 37, DisplayName: "Horse", Name: "horse", Width: 1.3964844, Height: 1.6, Type: "animal", Category: }
	Husk                 = Entity{ID: 38, InternalID: 38, DisplayName: "Husk", Name: "husk", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	Illusioner           = Entity{ID: 39, InternalID: 39, DisplayName: "Illusioner", Name: "illusioner", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	IronGolem            = Entity{ID: 40, InternalID: 40, DisplayName: "Iron Golem", Name: "iron_golem", Width: 1.4, Height: 2.7, Type: "mob", Category: }
	Item                 = Entity{ID: 41, InternalID: 41, DisplayName: "Item", Name: "item", Width: 0.25, Height: 0.25, Type: "other", Category: }
	ItemFrame            = Entity{ID: 42, InternalID: 42, DisplayName: "Item Frame", Name: "item_frame", Width: 0.5, Height: 0.5, Type: "other", Category: }
	Fireball             = Entity{ID: 43, InternalID: 43, DisplayName: "Fireball", Name: "fireball", Width: 1, Height: 1, Type: "projectile", Category: }
	LeashKnot            = Entity{ID: 44, InternalID: 44, DisplayName: "Leash Knot", Name: "leash_knot", Width: 0.375, Height: 0.5, Type: "other", Category: }
	LightningBolt        = Entity{ID: 45, InternalID: 45, DisplayName: "Lightning Bolt", Name: "lightning_bolt", Width: 0, Height: 0, Type: "other", Category: }
	Llama                = Entity{ID: 46, InternalID: 46, DisplayName: "Llama", Name: "llama", Width: 0.9, Height: 1.87, Type: "animal", Category: }
	LlamaSpit            = Entity{ID: 47, InternalID: 47, DisplayName: "Llama Spit", Name: "llama_spit", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	MagmaCube            = Entity{ID: 48, InternalID: 48, DisplayName: "Magma Cube", Name: "magma_cube", Width: 2.04, Height: 2.04, Type: "mob", Category: }
	Marker               = Entity{ID: 49, InternalID: 49, DisplayName: "Marker", Name: "marker", Width: 0, Height: 0, Type: "other", Category: }
	Minecart             = Entity{ID: 50, InternalID: 50, DisplayName: "Minecart", Name: "minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	ChestMinecart        = Entity{ID: 51, InternalID: 51, DisplayName: "Minecart with Chest", Name: "chest_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	CommandBlockMinecart = Entity{ID: 52, InternalID: 52, DisplayName: "Minecart with Command Block", Name: "command_block_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	FurnaceMinecart      = Entity{ID: 53, InternalID: 53, DisplayName: "Minecart with Furnace", Name: "furnace_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	HopperMinecart       = Entity{ID: 54, InternalID: 54, DisplayName: "Minecart with Hopper", Name: "hopper_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	SpawnerMinecart      = Entity{ID: 55, InternalID: 55, DisplayName: "Minecart with Spawner", Name: "spawner_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	TntMinecart          = Entity{ID: 56, InternalID: 56, DisplayName: "Minecart with TNT", Name: "tnt_minecart", Width: 0.98, Height: 0.7, Type: "other", Category: }
	Mule                 = Entity{ID: 57, InternalID: 57, DisplayName: "Mule", Name: "mule", Width: 1.3964844, Height: 1.6, Type: "animal", Category: }
	Mooshroom            = Entity{ID: 58, InternalID: 58, DisplayName: "Mooshroom", Name: "mooshroom", Width: 0.9, Height: 1.4, Type: "animal", Category: }
	Ocelot               = Entity{ID: 59, InternalID: 59, DisplayName: "Ocelot", Name: "ocelot", Width: 0.6, Height: 0.7, Type: "animal", Category: }
	Painting             = Entity{ID: 60, InternalID: 60, DisplayName: "Painting", Name: "painting", Width: 0.5, Height: 0.5, Type: "other", Category: }
	Panda                = Entity{ID: 61, InternalID: 61, DisplayName: "Panda", Name: "panda", Width: 1.3, Height: 1.25, Type: "animal", Category: }
	Parrot               = Entity{ID: 62, InternalID: 62, DisplayName: "Parrot", Name: "parrot", Width: 0.5, Height: 0.9, Type: "animal", Category: }
	Phantom              = Entity{ID: 63, InternalID: 63, DisplayName: "Phantom", Name: "phantom", Width: 0.9, Height: 0.5, Type: "mob", Category: }
	Pig                  = Entity{ID: 64, InternalID: 64, DisplayName: "Pig", Name: "pig", Width: 0.9, Height: 0.9, Type: "animal", Category: }
	Piglin               = Entity{ID: 65, InternalID: 65, DisplayName: "Piglin", Name: "piglin", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	PiglinBrute          = Entity{ID: 66, InternalID: 66, DisplayName: "Piglin Brute", Name: "piglin_brute", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	Pillager             = Entity{ID: 67, InternalID: 67, DisplayName: "Pillager", Name: "pillager", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	PolarBear            = Entity{ID: 68, InternalID: 68, DisplayName: "Polar Bear", Name: "polar_bear", Width: 1.4, Height: 1.4, Type: "animal", Category: }
	Tnt                  = Entity{ID: 69, InternalID: 69, DisplayName: "Primed TNT", Name: "tnt", Width: 0.98, Height: 0.98, Type: "other", Category: }
	Pufferfish           = Entity{ID: 70, InternalID: 70, DisplayName: "Pufferfish", Name: "pufferfish", Width: 0.7, Height: 0.7, Type: "water_creature", Category: }
	Rabbit               = Entity{ID: 71, InternalID: 71, DisplayName: "Rabbit", Name: "rabbit", Width: 0.4, Height: 0.5, Type: "animal", Category: }
	Ravager              = Entity{ID: 72, InternalID: 72, DisplayName: "Ravager", Name: "ravager", Width: 1.95, Height: 2.2, Type: "hostile", Category: }
	Salmon               = Entity{ID: 73, InternalID: 73, DisplayName: "Salmon", Name: "salmon", Width: 0.7, Height: 0.4, Type: "water_creature", Category: }
	Sheep                = Entity{ID: 74, InternalID: 74, DisplayName: "Sheep", Name: "sheep", Width: 0.9, Height: 1.3, Type: "animal", Category: }
	Shulker              = Entity{ID: 75, InternalID: 75, DisplayName: "Shulker", Name: "shulker", Width: 1, Height: 1, Type: "mob", Category: }
	ShulkerBullet        = Entity{ID: 76, InternalID: 76, DisplayName: "Shulker Bullet", Name: "shulker_bullet", Width: 0.3125, Height: 0.3125, Type: "projectile", Category: }
	Silverfish           = Entity{ID: 77, InternalID: 77, DisplayName: "Silverfish", Name: "silverfish", Width: 0.4, Height: 0.3, Type: "hostile", Category: }
	Skeleton             = Entity{ID: 78, InternalID: 78, DisplayName: "Skeleton", Name: "skeleton", Width: 0.6, Height: 1.99, Type: "hostile", Category: }
	SkeletonHorse        = Entity{ID: 79, InternalID: 79, DisplayName: "Skeleton Horse", Name: "skeleton_horse", Width: 1.3964844, Height: 1.6, Type: "animal", Category: }
	Slime                = Entity{ID: 80, InternalID: 80, DisplayName: "Slime", Name: "slime", Width: 2.04, Height: 2.04, Type: "mob", Category: }
	SmallFireball        = Entity{ID: 81, InternalID: 81, DisplayName: "Small Fireball", Name: "small_fireball", Width: 0.3125, Height: 0.3125, Type: "projectile", Category: }
	SnowGolem            = Entity{ID: 82, InternalID: 82, DisplayName: "Snow Golem", Name: "snow_golem", Width: 0.7, Height: 1.9, Type: "mob", Category: }
	Snowball             = Entity{ID: 83, InternalID: 83, DisplayName: "Snowball", Name: "snowball", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	SpectralArrow        = Entity{ID: 84, InternalID: 84, DisplayName: "Spectral Arrow", Name: "spectral_arrow", Width: 0.5, Height: 0.5, Type: "projectile", Category: }
	Spider               = Entity{ID: 85, InternalID: 85, DisplayName: "Spider", Name: "spider", Width: 1.4, Height: 0.9, Type: "hostile", Category: }
	Squid                = Entity{ID: 86, InternalID: 86, DisplayName: "Squid", Name: "squid", Width: 0.8, Height: 0.8, Type: "water_creature", Category: }
	Stray                = Entity{ID: 87, InternalID: 87, DisplayName: "Stray", Name: "stray", Width: 0.6, Height: 1.99, Type: "hostile", Category: }
	Strider              = Entity{ID: 88, InternalID: 88, DisplayName: "Strider", Name: "strider", Width: 0.9, Height: 1.7, Type: "animal", Category: }
	Egg                  = Entity{ID: 89, InternalID: 89, DisplayName: "Thrown Egg", Name: "egg", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	EnderPearl           = Entity{ID: 90, InternalID: 90, DisplayName: "Thrown Ender Pearl", Name: "ender_pearl", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	ExperienceBottle     = Entity{ID: 91, InternalID: 91, DisplayName: "Thrown Bottle o' Enchanting", Name: "experience_bottle", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	Potion               = Entity{ID: 92, InternalID: 92, DisplayName: "Potion", Name: "potion", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
	Trident              = Entity{ID: 93, InternalID: 93, DisplayName: "Trident", Name: "trident", Width: 0.5, Height: 0.5, Type: "projectile", Category: }
	TraderLlama          = Entity{ID: 94, InternalID: 94, DisplayName: "Trader Llama", Name: "trader_llama", Width: 0.9, Height: 1.87, Type: "animal", Category: }
	TropicalFish         = Entity{ID: 95, InternalID: 95, DisplayName: "Tropical Fish", Name: "tropical_fish", Width: 0.5, Height: 0.4, Type: "water_creature", Category: }
	Turtle               = Entity{ID: 96, InternalID: 96, DisplayName: "Turtle", Name: "turtle", Width: 1.2, Height: 0.4, Type: "animal", Category: }
	Vex                  = Entity{ID: 97, InternalID: 97, DisplayName: "Vex", Name: "vex", Width: 0.4, Height: 0.8, Type: "hostile", Category: }
	Villager             = Entity{ID: 98, InternalID: 98, DisplayName: "Villager", Name: "villager", Width: 0.6, Height: 1.95, Type: "passive", Category: }
	Vindicator           = Entity{ID: 99, InternalID: 99, DisplayName: "Vindicator", Name: "vindicator", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	WanderingTrader      = Entity{ID: 100, InternalID: 100, DisplayName: "Wandering Trader", Name: "wandering_trader", Width: 0.6, Height: 1.95, Type: "passive", Category: }
	Witch                = Entity{ID: 101, InternalID: 101, DisplayName: "Witch", Name: "witch", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	Wither               = Entity{ID: 102, InternalID: 102, DisplayName: "Wither", Name: "wither", Width: 0.9, Height: 3.5, Type: "hostile", Category: }
	WitherSkeleton       = Entity{ID: 103, InternalID: 103, DisplayName: "Wither Skeleton", Name: "wither_skeleton", Width: 0.7, Height: 2.4, Type: "hostile", Category: }
	WitherSkull          = Entity{ID: 104, InternalID: 104, DisplayName: "Wither Skull", Name: "wither_skull", Width: 0.3125, Height: 0.3125, Type: "projectile", Category: }
	Wolf                 = Entity{ID: 105, InternalID: 105, DisplayName: "Wolf", Name: "wolf", Width: 0.6, Height: 0.85, Type: "animal", Category: }
	Zoglin               = Entity{ID: 106, InternalID: 106, DisplayName: "Zoglin", Name: "zoglin", Width: 1.3964844, Height: 1.4, Type: "hostile", Category: }
	Zombie               = Entity{ID: 107, InternalID: 107, DisplayName: "Zombie", Name: "zombie", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	ZombieHorse          = Entity{ID: 108, InternalID: 108, DisplayName: "Zombie Horse", Name: "zombie_horse", Width: 1.3964844, Height: 1.6, Type: "animal", Category: }
	ZombieVillager       = Entity{ID: 109, InternalID: 109, DisplayName: "Zombie Villager", Name: "zombie_villager", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	ZombifiedPiglin      = Entity{ID: 110, InternalID: 110, DisplayName: "Zombified Piglin", Name: "zombified_piglin", Width: 0.6, Height: 1.95, Type: "hostile", Category: }
	Player               = Entity{ID: 111, InternalID: 111, DisplayName: "Player", Name: "player", Width: 0.6, Height: 1.8, Type: "player", Category: }
	FishingBobber        = Entity{ID: 112, InternalID: 112, DisplayName: "Fishing Bobber", Name: "fishing_bobber", Width: 0.25, Height: 0.25, Type: "projectile", Category: }
)

// ByID is an index of minecraft entities by their ID.
var ByID = map[ID]*Entity{
  0: &AreaEffectCloud,
  1: &ArmorStand,
  2: &Arrow,
  3: &Axolotl,
  4: &Bat,
  5: &Bee,
  6: &Blaze,
  7: &Boat,
  8: &Cat,
  9: &CaveSpider,
  10: &Chicken,
  11: &Cod,
  12: &Cow,
  13: &Creeper,
  14: &Dolphin,
  15: &Donkey,
  16: &DragonFireball,
  17: &Drowned,
  18: &ElderGuardian,
  19: &EndCrystal,
  20: &EnderDragon,
  21: &Enderman,
  22: &Endermite,
  23: &Evoker,
  24: &EvokerFangs,
  25: &ExperienceOrb,
  26: &EyeOfEnder,
  27: &FallingBlock,
  28: &FireworkRocket,
  29: &Fox,
  30: &Ghast,
  31: &Giant,
  32: &GlowItemFrame,
  33: &GlowSquid,
  34: &Goat,
  35: &Guardian,
  36: &Hoglin,
  37: &Horse,
  38: &Husk,
  39: &Illusioner,
  40: &IronGolem,
  41: &Item,
  42: &ItemFrame,
  43: &Fireball,
  44: &LeashKnot,
  45: &LightningBolt,
  46: &Llama,
  47: &LlamaSpit,
  48: &MagmaCube,
  49: &Marker,
  50: &Minecart,
  51: &ChestMinecart,
  52: &CommandBlockMinecart,
  53: &FurnaceMinecart,
  54: &HopperMinecart,
  55: &SpawnerMinecart,
  56: &TntMinecart,
  57: &Mule,
  58: &Mooshroom,
  59: &Ocelot,
  60: &Painting,
  61: &Panda,
  62: &Parrot,
  63: &Phantom,
  64: &Pig,
  65: &Piglin,
  66: &PiglinBrute,
  67: &Pillager,
  68: &PolarBear,
  69: &Tnt,
  70: &Pufferfish,
  71: &Rabbit,
  72: &Ravager,
  73: &Salmon,
  74: &Sheep,
  75: &Shulker,
  76: &ShulkerBullet,
  77: &Silverfish,
  78: &Skeleton,
  79: &SkeletonHorse,
  80: &Slime,
  81: &SmallFireball,
  82: &SnowGolem,
  83: &Snowball,
  84: &SpectralArrow,
  85: &Spider,
  86: &Squid,
  87: &Stray,
  88: &Strider,
  89: &Egg,
  90: &EnderPearl,
  91: &ExperienceBottle,
  92: &Potion,
  93: &Trident,
  94: &TraderLlama,
  95: &TropicalFish,
  96: &Turtle,
  97: &Vex,
  98: &Villager,
  99: &Vindicator,
  100: &WanderingTrader,
  101: &Witch,
  102: &Wither,
  103: &WitherSkeleton,
  104: &WitherSkull,
  105: &Wolf,
  106: &Zoglin,
  107: &Zombie,
  108: &ZombieHorse,
  109: &ZombieVillager,
  110: &ZombifiedPiglin,
  111: &Player,
  112: &FishingBobber,
}

