// Package item stores information about items in Minecraft.
package item

// ID describes the numeric ID of an item.
type ID uint32

// Item describes information about a type of item.
type Item struct {
	ID          ID
	DisplayName string
	Name        string
	StackSize   uint
}

var (
	Stone                           = Item{ID: 1, DisplayName: "Stone", Name: "stone", StackSize: 64}
	Granite                         = Item{ID: 2, DisplayName: "Granite", Name: "granite", StackSize: 64}
	PolishedGranite                 = Item{ID: 3, DisplayName: "Polished Granite", Name: "polished_granite", StackSize: 64}
	Diorite                         = Item{ID: 4, DisplayName: "Diorite", Name: "diorite", StackSize: 64}
	PolishedDiorite                 = Item{ID: 5, DisplayName: "Polished Diorite", Name: "polished_diorite", StackSize: 64}
	Andesite                        = Item{ID: 6, DisplayName: "Andesite", Name: "andesite", StackSize: 64}
	PolishedAndesite                = Item{ID: 7, DisplayName: "Polished Andesite", Name: "polished_andesite", StackSize: 64}
	Deepslate                       = Item{ID: 8, DisplayName: "Deepslate", Name: "deepslate", StackSize: 64}
	CobbledDeepslate                = Item{ID: 9, DisplayName: "Cobbled Deepslate", Name: "cobbled_deepslate", StackSize: 64}
	PolishedDeepslate               = Item{ID: 10, DisplayName: "Polished Deepslate", Name: "polished_deepslate", StackSize: 64}
	Calcite                         = Item{ID: 11, DisplayName: "Calcite", Name: "calcite", StackSize: 64}
	Tuff                            = Item{ID: 12, DisplayName: "Tuff", Name: "tuff", StackSize: 64}
	DripstoneBlock                  = Item{ID: 13, DisplayName: "Dripstone Block", Name: "dripstone_block", StackSize: 64}
	GrassBlock                      = Item{ID: 14, DisplayName: "Grass Block", Name: "grass_block", StackSize: 64}
	Dirt                            = Item{ID: 15, DisplayName: "Dirt", Name: "dirt", StackSize: 64}
	CoarseDirt                      = Item{ID: 16, DisplayName: "Coarse Dirt", Name: "coarse_dirt", StackSize: 64}
	Podzol                          = Item{ID: 17, DisplayName: "Podzol", Name: "podzol", StackSize: 64}
	RootedDirt                      = Item{ID: 18, DisplayName: "Rooted Dirt", Name: "rooted_dirt", StackSize: 64}
	CrimsonNylium                   = Item{ID: 19, DisplayName: "Crimson Nylium", Name: "crimson_nylium", StackSize: 64}
	WarpedNylium                    = Item{ID: 20, DisplayName: "Warped Nylium", Name: "warped_nylium", StackSize: 64}
	Cobblestone                     = Item{ID: 21, DisplayName: "Cobblestone", Name: "cobblestone", StackSize: 64}
	OakPlanks                       = Item{ID: 22, DisplayName: "Oak Planks", Name: "oak_planks", StackSize: 64}
	SprucePlanks                    = Item{ID: 23, DisplayName: "Spruce Planks", Name: "spruce_planks", StackSize: 64}
	BirchPlanks                     = Item{ID: 24, DisplayName: "Birch Planks", Name: "birch_planks", StackSize: 64}
	JunglePlanks                    = Item{ID: 25, DisplayName: "Jungle Planks", Name: "jungle_planks", StackSize: 64}
	AcaciaPlanks                    = Item{ID: 26, DisplayName: "Acacia Planks", Name: "acacia_planks", StackSize: 64}
	DarkOakPlanks                   = Item{ID: 27, DisplayName: "Dark Oak Planks", Name: "dark_oak_planks", StackSize: 64}
	CrimsonPlanks                   = Item{ID: 28, DisplayName: "Crimson Planks", Name: "crimson_planks", StackSize: 64}
	WarpedPlanks                    = Item{ID: 29, DisplayName: "Warped Planks", Name: "warped_planks", StackSize: 64}
	OakSapling                      = Item{ID: 30, DisplayName: "Oak Sapling", Name: "oak_sapling", StackSize: 64}
	SpruceSapling                   = Item{ID: 31, DisplayName: "Spruce Sapling", Name: "spruce_sapling", StackSize: 64}
	BirchSapling                    = Item{ID: 32, DisplayName: "Birch Sapling", Name: "birch_sapling", StackSize: 64}
	JungleSapling                   = Item{ID: 33, DisplayName: "Jungle Sapling", Name: "jungle_sapling", StackSize: 64}
	AcaciaSapling                   = Item{ID: 34, DisplayName: "Acacia Sapling", Name: "acacia_sapling", StackSize: 64}
	DarkOakSapling                  = Item{ID: 35, DisplayName: "Dark Oak Sapling", Name: "dark_oak_sapling", StackSize: 64}
	Bedrock                         = Item{ID: 36, DisplayName: "Bedrock", Name: "bedrock", StackSize: 64}
	Sand                            = Item{ID: 37, DisplayName: "Sand", Name: "sand", StackSize: 64}
	RedSand                         = Item{ID: 38, DisplayName: "Red Sand", Name: "red_sand", StackSize: 64}
	Gravel                          = Item{ID: 39, DisplayName: "Gravel", Name: "gravel", StackSize: 64}
	CoalOre                         = Item{ID: 40, DisplayName: "Coal Ore", Name: "coal_ore", StackSize: 64}
	DeepslateCoalOre                = Item{ID: 41, DisplayName: "Deepslate Coal Ore", Name: "deepslate_coal_ore", StackSize: 64}
	IronOre                         = Item{ID: 42, DisplayName: "Iron Ore", Name: "iron_ore", StackSize: 64}
	DeepslateIronOre                = Item{ID: 43, DisplayName: "Deepslate Iron Ore", Name: "deepslate_iron_ore", StackSize: 64}
	CopperOre                       = Item{ID: 44, DisplayName: "Copper Ore", Name: "copper_ore", StackSize: 64}
	DeepslateCopperOre              = Item{ID: 45, DisplayName: "Deepslate Copper Ore", Name: "deepslate_copper_ore", StackSize: 64}
	GoldOre                         = Item{ID: 46, DisplayName: "Gold Ore", Name: "gold_ore", StackSize: 64}
	DeepslateGoldOre                = Item{ID: 47, DisplayName: "Deepslate Gold Ore", Name: "deepslate_gold_ore", StackSize: 64}
	RedstoneOre                     = Item{ID: 48, DisplayName: "Redstone Ore", Name: "redstone_ore", StackSize: 64}
	DeepslateRedstoneOre            = Item{ID: 49, DisplayName: "Deepslate Redstone Ore", Name: "deepslate_redstone_ore", StackSize: 64}
	EmeraldOre                      = Item{ID: 50, DisplayName: "Emerald Ore", Name: "emerald_ore", StackSize: 64}
	DeepslateEmeraldOre             = Item{ID: 51, DisplayName: "Deepslate Emerald Ore", Name: "deepslate_emerald_ore", StackSize: 64}
	LapisOre                        = Item{ID: 52, DisplayName: "Lapis Lazuli Ore", Name: "lapis_ore", StackSize: 64}
	DeepslateLapisOre               = Item{ID: 53, DisplayName: "Deepslate Lapis Lazuli Ore", Name: "deepslate_lapis_ore", StackSize: 64}
	DiamondOre                      = Item{ID: 54, DisplayName: "Diamond Ore", Name: "diamond_ore", StackSize: 64}
	DeepslateDiamondOre             = Item{ID: 55, DisplayName: "Deepslate Diamond Ore", Name: "deepslate_diamond_ore", StackSize: 64}
	NetherGoldOre                   = Item{ID: 56, DisplayName: "Nether Gold Ore", Name: "nether_gold_ore", StackSize: 64}
	NetherQuartzOre                 = Item{ID: 57, DisplayName: "Nether Quartz Ore", Name: "nether_quartz_ore", StackSize: 64}
	AncientDebris                   = Item{ID: 58, DisplayName: "Ancient Debris", Name: "ancient_debris", StackSize: 64}
	CoalBlock                       = Item{ID: 59, DisplayName: "Block of Coal", Name: "coal_block", StackSize: 64}
	RawIronBlock                    = Item{ID: 60, DisplayName: "Block of Raw Iron", Name: "raw_iron_block", StackSize: 64}
	RawCopperBlock                  = Item{ID: 61, DisplayName: "Block of Raw Copper", Name: "raw_copper_block", StackSize: 64}
	RawGoldBlock                    = Item{ID: 62, DisplayName: "Block of Raw Gold", Name: "raw_gold_block", StackSize: 64}
	AmethystBlock                   = Item{ID: 63, DisplayName: "Block of Amethyst", Name: "amethyst_block", StackSize: 64}
	BuddingAmethyst                 = Item{ID: 64, DisplayName: "Budding Amethyst", Name: "budding_amethyst", StackSize: 64}
	IronBlock                       = Item{ID: 65, DisplayName: "Block of Iron", Name: "iron_block", StackSize: 64}
	CopperBlock                     = Item{ID: 66, DisplayName: "Block of Copper", Name: "copper_block", StackSize: 64}
	GoldBlock                       = Item{ID: 67, DisplayName: "Block of Gold", Name: "gold_block", StackSize: 64}
	DiamondBlock                    = Item{ID: 68, DisplayName: "Block of Diamond", Name: "diamond_block", StackSize: 64}
	NetheriteBlock                  = Item{ID: 69, DisplayName: "Block of Netherite", Name: "netherite_block", StackSize: 64}
	ExposedCopper                   = Item{ID: 70, DisplayName: "Exposed Copper", Name: "exposed_copper", StackSize: 64}
	WeatheredCopper                 = Item{ID: 71, DisplayName: "Weathered Copper", Name: "weathered_copper", StackSize: 64}
	OxidizedCopper                  = Item{ID: 72, DisplayName: "Oxidized Copper", Name: "oxidized_copper", StackSize: 64}
	CutCopper                       = Item{ID: 73, DisplayName: "Cut Copper", Name: "cut_copper", StackSize: 64}
	ExposedCutCopper                = Item{ID: 74, DisplayName: "Exposed Cut Copper", Name: "exposed_cut_copper", StackSize: 64}
	WeatheredCutCopper              = Item{ID: 75, DisplayName: "Weathered Cut Copper", Name: "weathered_cut_copper", StackSize: 64}
	OxidizedCutCopper               = Item{ID: 76, DisplayName: "Oxidized Cut Copper", Name: "oxidized_cut_copper", StackSize: 64}
	CutCopperStairs                 = Item{ID: 77, DisplayName: "Cut Copper Stairs", Name: "cut_copper_stairs", StackSize: 64}
	ExposedCutCopperStairs          = Item{ID: 78, DisplayName: "Exposed Cut Copper Stairs", Name: "exposed_cut_copper_stairs", StackSize: 64}
	WeatheredCutCopperStairs        = Item{ID: 79, DisplayName: "Weathered Cut Copper Stairs", Name: "weathered_cut_copper_stairs", StackSize: 64}
	OxidizedCutCopperStairs         = Item{ID: 80, DisplayName: "Oxidized Cut Copper Stairs", Name: "oxidized_cut_copper_stairs", StackSize: 64}
	CutCopperSlab                   = Item{ID: 81, DisplayName: "Cut Copper Slab", Name: "cut_copper_slab", StackSize: 64}
	ExposedCutCopperSlab            = Item{ID: 82, DisplayName: "Exposed Cut Copper Slab", Name: "exposed_cut_copper_slab", StackSize: 64}
	WeatheredCutCopperSlab          = Item{ID: 83, DisplayName: "Weathered Cut Copper Slab", Name: "weathered_cut_copper_slab", StackSize: 64}
	OxidizedCutCopperSlab           = Item{ID: 84, DisplayName: "Oxidized Cut Copper Slab", Name: "oxidized_cut_copper_slab", StackSize: 64}
	WaxedCopperBlock                = Item{ID: 85, DisplayName: "Waxed Block of Copper", Name: "waxed_copper_block", StackSize: 64}
	WaxedExposedCopper              = Item{ID: 86, DisplayName: "Waxed Exposed Copper", Name: "waxed_exposed_copper", StackSize: 64}
	WaxedWeatheredCopper            = Item{ID: 87, DisplayName: "Waxed Weathered Copper", Name: "waxed_weathered_copper", StackSize: 64}
	WaxedOxidizedCopper             = Item{ID: 88, DisplayName: "Waxed Oxidized Copper", Name: "waxed_oxidized_copper", StackSize: 64}
	WaxedCutCopper                  = Item{ID: 89, DisplayName: "Waxed Cut Copper", Name: "waxed_cut_copper", StackSize: 64}
	WaxedExposedCutCopper           = Item{ID: 90, DisplayName: "Waxed Exposed Cut Copper", Name: "waxed_exposed_cut_copper", StackSize: 64}
	WaxedWeatheredCutCopper         = Item{ID: 91, DisplayName: "Waxed Weathered Cut Copper", Name: "waxed_weathered_cut_copper", StackSize: 64}
	WaxedOxidizedCutCopper          = Item{ID: 92, DisplayName: "Waxed Oxidized Cut Copper", Name: "waxed_oxidized_cut_copper", StackSize: 64}
	WaxedCutCopperStairs            = Item{ID: 93, DisplayName: "Waxed Cut Copper Stairs", Name: "waxed_cut_copper_stairs", StackSize: 64}
	WaxedExposedCutCopperStairs     = Item{ID: 94, DisplayName: "Waxed Exposed Cut Copper Stairs", Name: "waxed_exposed_cut_copper_stairs", StackSize: 64}
	WaxedWeatheredCutCopperStairs   = Item{ID: 95, DisplayName: "Waxed Weathered Cut Copper Stairs", Name: "waxed_weathered_cut_copper_stairs", StackSize: 64}
	WaxedOxidizedCutCopperStairs    = Item{ID: 96, DisplayName: "Waxed Oxidized Cut Copper Stairs", Name: "waxed_oxidized_cut_copper_stairs", StackSize: 64}
	WaxedCutCopperSlab              = Item{ID: 97, DisplayName: "Waxed Cut Copper Slab", Name: "waxed_cut_copper_slab", StackSize: 64}
	WaxedExposedCutCopperSlab       = Item{ID: 98, DisplayName: "Waxed Exposed Cut Copper Slab", Name: "waxed_exposed_cut_copper_slab", StackSize: 64}
	WaxedWeatheredCutCopperSlab     = Item{ID: 99, DisplayName: "Waxed Weathered Cut Copper Slab", Name: "waxed_weathered_cut_copper_slab", StackSize: 64}
	WaxedOxidizedCutCopperSlab      = Item{ID: 100, DisplayName: "Waxed Oxidized Cut Copper Slab", Name: "waxed_oxidized_cut_copper_slab", StackSize: 64}
	OakLog                          = Item{ID: 101, DisplayName: "Oak Log", Name: "oak_log", StackSize: 64}
	SpruceLog                       = Item{ID: 102, DisplayName: "Spruce Log", Name: "spruce_log", StackSize: 64}
	BirchLog                        = Item{ID: 103, DisplayName: "Birch Log", Name: "birch_log", StackSize: 64}
	JungleLog                       = Item{ID: 104, DisplayName: "Jungle Log", Name: "jungle_log", StackSize: 64}
	AcaciaLog                       = Item{ID: 105, DisplayName: "Acacia Log", Name: "acacia_log", StackSize: 64}
	DarkOakLog                      = Item{ID: 106, DisplayName: "Dark Oak Log", Name: "dark_oak_log", StackSize: 64}
	CrimsonStem                     = Item{ID: 107, DisplayName: "Crimson Stem", Name: "crimson_stem", StackSize: 64}
	WarpedStem                      = Item{ID: 108, DisplayName: "Warped Stem", Name: "warped_stem", StackSize: 64}
	StrippedOakLog                  = Item{ID: 109, DisplayName: "Stripped Oak Log", Name: "stripped_oak_log", StackSize: 64}
	StrippedSpruceLog               = Item{ID: 110, DisplayName: "Stripped Spruce Log", Name: "stripped_spruce_log", StackSize: 64}
	StrippedBirchLog                = Item{ID: 111, DisplayName: "Stripped Birch Log", Name: "stripped_birch_log", StackSize: 64}
	StrippedJungleLog               = Item{ID: 112, DisplayName: "Stripped Jungle Log", Name: "stripped_jungle_log", StackSize: 64}
	StrippedAcaciaLog               = Item{ID: 113, DisplayName: "Stripped Acacia Log", Name: "stripped_acacia_log", StackSize: 64}
	StrippedDarkOakLog              = Item{ID: 114, DisplayName: "Stripped Dark Oak Log", Name: "stripped_dark_oak_log", StackSize: 64}
	StrippedCrimsonStem             = Item{ID: 115, DisplayName: "Stripped Crimson Stem", Name: "stripped_crimson_stem", StackSize: 64}
	StrippedWarpedStem              = Item{ID: 116, DisplayName: "Stripped Warped Stem", Name: "stripped_warped_stem", StackSize: 64}
	StrippedOakWood                 = Item{ID: 117, DisplayName: "Stripped Oak Wood", Name: "stripped_oak_wood", StackSize: 64}
	StrippedSpruceWood              = Item{ID: 118, DisplayName: "Stripped Spruce Wood", Name: "stripped_spruce_wood", StackSize: 64}
	StrippedBirchWood               = Item{ID: 119, DisplayName: "Stripped Birch Wood", Name: "stripped_birch_wood", StackSize: 64}
	StrippedJungleWood              = Item{ID: 120, DisplayName: "Stripped Jungle Wood", Name: "stripped_jungle_wood", StackSize: 64}
	StrippedAcaciaWood              = Item{ID: 121, DisplayName: "Stripped Acacia Wood", Name: "stripped_acacia_wood", StackSize: 64}
	StrippedDarkOakWood             = Item{ID: 122, DisplayName: "Stripped Dark Oak Wood", Name: "stripped_dark_oak_wood", StackSize: 64}
	StrippedCrimsonHyphae           = Item{ID: 123, DisplayName: "Stripped Crimson Hyphae", Name: "stripped_crimson_hyphae", StackSize: 64}
	StrippedWarpedHyphae            = Item{ID: 124, DisplayName: "Stripped Warped Hyphae", Name: "stripped_warped_hyphae", StackSize: 64}
	OakWood                         = Item{ID: 125, DisplayName: "Oak Wood", Name: "oak_wood", StackSize: 64}
	SpruceWood                      = Item{ID: 126, DisplayName: "Spruce Wood", Name: "spruce_wood", StackSize: 64}
	BirchWood                       = Item{ID: 127, DisplayName: "Birch Wood", Name: "birch_wood", StackSize: 64}
	JungleWood                      = Item{ID: 128, DisplayName: "Jungle Wood", Name: "jungle_wood", StackSize: 64}
	AcaciaWood                      = Item{ID: 129, DisplayName: "Acacia Wood", Name: "acacia_wood", StackSize: 64}
	DarkOakWood                     = Item{ID: 130, DisplayName: "Dark Oak Wood", Name: "dark_oak_wood", StackSize: 64}
	CrimsonHyphae                   = Item{ID: 131, DisplayName: "Crimson Hyphae", Name: "crimson_hyphae", StackSize: 64}
	WarpedHyphae                    = Item{ID: 132, DisplayName: "Warped Hyphae", Name: "warped_hyphae", StackSize: 64}
	OakLeaves                       = Item{ID: 133, DisplayName: "Oak Leaves", Name: "oak_leaves", StackSize: 64}
	SpruceLeaves                    = Item{ID: 134, DisplayName: "Spruce Leaves", Name: "spruce_leaves", StackSize: 64}
	BirchLeaves                     = Item{ID: 135, DisplayName: "Birch Leaves", Name: "birch_leaves", StackSize: 64}
	JungleLeaves                    = Item{ID: 136, DisplayName: "Jungle Leaves", Name: "jungle_leaves", StackSize: 64}
	AcaciaLeaves                    = Item{ID: 137, DisplayName: "Acacia Leaves", Name: "acacia_leaves", StackSize: 64}
	DarkOakLeaves                   = Item{ID: 138, DisplayName: "Dark Oak Leaves", Name: "dark_oak_leaves", StackSize: 64}
	AzaleaLeaves                    = Item{ID: 139, DisplayName: "Azalea Leaves", Name: "azalea_leaves", StackSize: 64}
	FloweringAzaleaLeaves           = Item{ID: 140, DisplayName: "Flowering Azalea Leaves", Name: "flowering_azalea_leaves", StackSize: 64}
	Sponge                          = Item{ID: 141, DisplayName: "Sponge", Name: "sponge", StackSize: 64}
	WetSponge                       = Item{ID: 142, DisplayName: "Wet Sponge", Name: "wet_sponge", StackSize: 64}
	Glass                           = Item{ID: 143, DisplayName: "Glass", Name: "glass", StackSize: 64}
	TintedGlass                     = Item{ID: 144, DisplayName: "Tinted Glass", Name: "tinted_glass", StackSize: 64}
	LapisBlock                      = Item{ID: 145, DisplayName: "Block of Lapis Lazuli", Name: "lapis_block", StackSize: 64}
	Sandstone                       = Item{ID: 146, DisplayName: "Sandstone", Name: "sandstone", StackSize: 64}
	ChiseledSandstone               = Item{ID: 147, DisplayName: "Chiseled Sandstone", Name: "chiseled_sandstone", StackSize: 64}
	CutSandstone                    = Item{ID: 148, DisplayName: "Cut Sandstone", Name: "cut_sandstone", StackSize: 64}
	Cobweb                          = Item{ID: 149, DisplayName: "Cobweb", Name: "cobweb", StackSize: 64}
	Grass                           = Item{ID: 150, DisplayName: "Grass", Name: "grass", StackSize: 64}
	Fern                            = Item{ID: 151, DisplayName: "Fern", Name: "fern", StackSize: 64}
	Azalea                          = Item{ID: 152, DisplayName: "Azalea", Name: "azalea", StackSize: 64}
	FloweringAzalea                 = Item{ID: 153, DisplayName: "Flowering Azalea", Name: "flowering_azalea", StackSize: 64}
	DeadBush                        = Item{ID: 154, DisplayName: "Dead Bush", Name: "dead_bush", StackSize: 64}
	Seagrass                        = Item{ID: 155, DisplayName: "Seagrass", Name: "seagrass", StackSize: 64}
	SeaPickle                       = Item{ID: 156, DisplayName: "Sea Pickle", Name: "sea_pickle", StackSize: 64}
	WhiteWool                       = Item{ID: 157, DisplayName: "White Wool", Name: "white_wool", StackSize: 64}
	OrangeWool                      = Item{ID: 158, DisplayName: "Orange Wool", Name: "orange_wool", StackSize: 64}
	MagentaWool                     = Item{ID: 159, DisplayName: "Magenta Wool", Name: "magenta_wool", StackSize: 64}
	LightBlueWool                   = Item{ID: 160, DisplayName: "Light Blue Wool", Name: "light_blue_wool", StackSize: 64}
	YellowWool                      = Item{ID: 161, DisplayName: "Yellow Wool", Name: "yellow_wool", StackSize: 64}
	LimeWool                        = Item{ID: 162, DisplayName: "Lime Wool", Name: "lime_wool", StackSize: 64}
	PinkWool                        = Item{ID: 163, DisplayName: "Pink Wool", Name: "pink_wool", StackSize: 64}
	GrayWool                        = Item{ID: 164, DisplayName: "Gray Wool", Name: "gray_wool", StackSize: 64}
	LightGrayWool                   = Item{ID: 165, DisplayName: "Light Gray Wool", Name: "light_gray_wool", StackSize: 64}
	CyanWool                        = Item{ID: 166, DisplayName: "Cyan Wool", Name: "cyan_wool", StackSize: 64}
	PurpleWool                      = Item{ID: 167, DisplayName: "Purple Wool", Name: "purple_wool", StackSize: 64}
	BlueWool                        = Item{ID: 168, DisplayName: "Blue Wool", Name: "blue_wool", StackSize: 64}
	BrownWool                       = Item{ID: 169, DisplayName: "Brown Wool", Name: "brown_wool", StackSize: 64}
	GreenWool                       = Item{ID: 170, DisplayName: "Green Wool", Name: "green_wool", StackSize: 64}
	RedWool                         = Item{ID: 171, DisplayName: "Red Wool", Name: "red_wool", StackSize: 64}
	BlackWool                       = Item{ID: 172, DisplayName: "Black Wool", Name: "black_wool", StackSize: 64}
	Dandelion                       = Item{ID: 173, DisplayName: "Dandelion", Name: "dandelion", StackSize: 64}
	Poppy                           = Item{ID: 174, DisplayName: "Poppy", Name: "poppy", StackSize: 64}
	BlueOrchid                      = Item{ID: 175, DisplayName: "Blue Orchid", Name: "blue_orchid", StackSize: 64}
	Allium                          = Item{ID: 176, DisplayName: "Allium", Name: "allium", StackSize: 64}
	AzureBluet                      = Item{ID: 177, DisplayName: "Azure Bluet", Name: "azure_bluet", StackSize: 64}
	RedTulip                        = Item{ID: 178, DisplayName: "Red Tulip", Name: "red_tulip", StackSize: 64}
	OrangeTulip                     = Item{ID: 179, DisplayName: "Orange Tulip", Name: "orange_tulip", StackSize: 64}
	WhiteTulip                      = Item{ID: 180, DisplayName: "White Tulip", Name: "white_tulip", StackSize: 64}
	PinkTulip                       = Item{ID: 181, DisplayName: "Pink Tulip", Name: "pink_tulip", StackSize: 64}
	OxeyeDaisy                      = Item{ID: 182, DisplayName: "Oxeye Daisy", Name: "oxeye_daisy", StackSize: 64}
	Cornflower                      = Item{ID: 183, DisplayName: "Cornflower", Name: "cornflower", StackSize: 64}
	LilyOfTheValley                 = Item{ID: 184, DisplayName: "Lily of the Valley", Name: "lily_of_the_valley", StackSize: 64}
	WitherRose                      = Item{ID: 185, DisplayName: "Wither Rose", Name: "wither_rose", StackSize: 64}
	SporeBlossom                    = Item{ID: 186, DisplayName: "Spore Blossom", Name: "spore_blossom", StackSize: 64}
	BrownMushroom                   = Item{ID: 187, DisplayName: "Brown Mushroom", Name: "brown_mushroom", StackSize: 64}
	RedMushroom                     = Item{ID: 188, DisplayName: "Red Mushroom", Name: "red_mushroom", StackSize: 64}
	CrimsonFungus                   = Item{ID: 189, DisplayName: "Crimson Fungus", Name: "crimson_fungus", StackSize: 64}
	WarpedFungus                    = Item{ID: 190, DisplayName: "Warped Fungus", Name: "warped_fungus", StackSize: 64}
	CrimsonRoots                    = Item{ID: 191, DisplayName: "Crimson Roots", Name: "crimson_roots", StackSize: 64}
	WarpedRoots                     = Item{ID: 192, DisplayName: "Warped Roots", Name: "warped_roots", StackSize: 64}
	NetherSprouts                   = Item{ID: 193, DisplayName: "Nether Sprouts", Name: "nether_sprouts", StackSize: 64}
	WeepingVines                    = Item{ID: 194, DisplayName: "Weeping Vines", Name: "weeping_vines", StackSize: 64}
	TwistingVines                   = Item{ID: 195, DisplayName: "Twisting Vines", Name: "twisting_vines", StackSize: 64}
	SugarCane                       = Item{ID: 196, DisplayName: "Sugar Cane", Name: "sugar_cane", StackSize: 64}
	Kelp                            = Item{ID: 197, DisplayName: "Kelp", Name: "kelp", StackSize: 64}
	MossCarpet                      = Item{ID: 198, DisplayName: "Moss Carpet", Name: "moss_carpet", StackSize: 64}
	MossBlock                       = Item{ID: 199, DisplayName: "Moss Block", Name: "moss_block", StackSize: 64}
	HangingRoots                    = Item{ID: 200, DisplayName: "Hanging Roots", Name: "hanging_roots", StackSize: 64}
	BigDripleaf                     = Item{ID: 201, DisplayName: "Big Dripleaf", Name: "big_dripleaf", StackSize: 64}
	SmallDripleaf                   = Item{ID: 202, DisplayName: "Small Dripleaf", Name: "small_dripleaf", StackSize: 64}
	Bamboo                          = Item{ID: 203, DisplayName: "Bamboo", Name: "bamboo", StackSize: 64}
	OakSlab                         = Item{ID: 204, DisplayName: "Oak Slab", Name: "oak_slab", StackSize: 64}
	SpruceSlab                      = Item{ID: 205, DisplayName: "Spruce Slab", Name: "spruce_slab", StackSize: 64}
	BirchSlab                       = Item{ID: 206, DisplayName: "Birch Slab", Name: "birch_slab", StackSize: 64}
	JungleSlab                      = Item{ID: 207, DisplayName: "Jungle Slab", Name: "jungle_slab", StackSize: 64}
	AcaciaSlab                      = Item{ID: 208, DisplayName: "Acacia Slab", Name: "acacia_slab", StackSize: 64}
	DarkOakSlab                     = Item{ID: 209, DisplayName: "Dark Oak Slab", Name: "dark_oak_slab", StackSize: 64}
	CrimsonSlab                     = Item{ID: 210, DisplayName: "Crimson Slab", Name: "crimson_slab", StackSize: 64}
	WarpedSlab                      = Item{ID: 211, DisplayName: "Warped Slab", Name: "warped_slab", StackSize: 64}
	StoneSlab                       = Item{ID: 212, DisplayName: "Stone Slab", Name: "stone_slab", StackSize: 64}
	SmoothStoneSlab                 = Item{ID: 213, DisplayName: "Smooth Stone Slab", Name: "smooth_stone_slab", StackSize: 64}
	SandstoneSlab                   = Item{ID: 214, DisplayName: "Sandstone Slab", Name: "sandstone_slab", StackSize: 64}
	CutSandstoneSlab                = Item{ID: 215, DisplayName: "Cut Sandstone Slab", Name: "cut_sandstone_slab", StackSize: 64}
	PetrifiedOakSlab                = Item{ID: 216, DisplayName: "Petrified Oak Slab", Name: "petrified_oak_slab", StackSize: 64}
	CobblestoneSlab                 = Item{ID: 217, DisplayName: "Cobblestone Slab", Name: "cobblestone_slab", StackSize: 64}
	BrickSlab                       = Item{ID: 218, DisplayName: "Brick Slab", Name: "brick_slab", StackSize: 64}
	StoneBrickSlab                  = Item{ID: 219, DisplayName: "Stone Brick Slab", Name: "stone_brick_slab", StackSize: 64}
	NetherBrickSlab                 = Item{ID: 220, DisplayName: "Nether Brick Slab", Name: "nether_brick_slab", StackSize: 64}
	QuartzSlab                      = Item{ID: 221, DisplayName: "Quartz Slab", Name: "quartz_slab", StackSize: 64}
	RedSandstoneSlab                = Item{ID: 222, DisplayName: "Red Sandstone Slab", Name: "red_sandstone_slab", StackSize: 64}
	CutRedSandstoneSlab             = Item{ID: 223, DisplayName: "Cut Red Sandstone Slab", Name: "cut_red_sandstone_slab", StackSize: 64}
	PurpurSlab                      = Item{ID: 224, DisplayName: "Purpur Slab", Name: "purpur_slab", StackSize: 64}
	PrismarineSlab                  = Item{ID: 225, DisplayName: "Prismarine Slab", Name: "prismarine_slab", StackSize: 64}
	PrismarineBrickSlab             = Item{ID: 226, DisplayName: "Prismarine Brick Slab", Name: "prismarine_brick_slab", StackSize: 64}
	DarkPrismarineSlab              = Item{ID: 227, DisplayName: "Dark Prismarine Slab", Name: "dark_prismarine_slab", StackSize: 64}
	SmoothQuartz                    = Item{ID: 228, DisplayName: "Smooth Quartz Block", Name: "smooth_quartz", StackSize: 64}
	SmoothRedSandstone              = Item{ID: 229, DisplayName: "Smooth Red Sandstone", Name: "smooth_red_sandstone", StackSize: 64}
	SmoothSandstone                 = Item{ID: 230, DisplayName: "Smooth Sandstone", Name: "smooth_sandstone", StackSize: 64}
	SmoothStone                     = Item{ID: 231, DisplayName: "Smooth Stone", Name: "smooth_stone", StackSize: 64}
	Bricks                          = Item{ID: 232, DisplayName: "Bricks", Name: "bricks", StackSize: 64}
	Bookshelf                       = Item{ID: 233, DisplayName: "Bookshelf", Name: "bookshelf", StackSize: 64}
	MossyCobblestone                = Item{ID: 234, DisplayName: "Mossy Cobblestone", Name: "mossy_cobblestone", StackSize: 64}
	Obsidian                        = Item{ID: 235, DisplayName: "Obsidian", Name: "obsidian", StackSize: 64}
	Torch                           = Item{ID: 236, DisplayName: "Torch", Name: "torch", StackSize: 64}
	EndRod                          = Item{ID: 237, DisplayName: "End Rod", Name: "end_rod", StackSize: 64}
	ChorusPlant                     = Item{ID: 238, DisplayName: "Chorus Plant", Name: "chorus_plant", StackSize: 64}
	ChorusFlower                    = Item{ID: 239, DisplayName: "Chorus Flower", Name: "chorus_flower", StackSize: 64}
	PurpurBlock                     = Item{ID: 240, DisplayName: "Purpur Block", Name: "purpur_block", StackSize: 64}
	PurpurPillar                    = Item{ID: 241, DisplayName: "Purpur Pillar", Name: "purpur_pillar", StackSize: 64}
	PurpurStairs                    = Item{ID: 242, DisplayName: "Purpur Stairs", Name: "purpur_stairs", StackSize: 64}
	Spawner                         = Item{ID: 243, DisplayName: "Spawner", Name: "spawner", StackSize: 64}
	OakStairs                       = Item{ID: 244, DisplayName: "Oak Stairs", Name: "oak_stairs", StackSize: 64}
	Chest                           = Item{ID: 245, DisplayName: "Chest", Name: "chest", StackSize: 64}
	CraftingTable                   = Item{ID: 246, DisplayName: "Crafting Table", Name: "crafting_table", StackSize: 64}
	Farmland                        = Item{ID: 247, DisplayName: "Farmland", Name: "farmland", StackSize: 64}
	Furnace                         = Item{ID: 248, DisplayName: "Furnace", Name: "furnace", StackSize: 64}
	Ladder                          = Item{ID: 249, DisplayName: "Ladder", Name: "ladder", StackSize: 64}
	CobblestoneStairs               = Item{ID: 250, DisplayName: "Cobblestone Stairs", Name: "cobblestone_stairs", StackSize: 64}
	Snow                            = Item{ID: 251, DisplayName: "Snow", Name: "snow", StackSize: 64}
	Ice                             = Item{ID: 252, DisplayName: "Ice", Name: "ice", StackSize: 64}
	SnowBlock                       = Item{ID: 253, DisplayName: "Snow Block", Name: "snow_block", StackSize: 64}
	Cactus                          = Item{ID: 254, DisplayName: "Cactus", Name: "cactus", StackSize: 64}
	Clay                            = Item{ID: 255, DisplayName: "Clay", Name: "clay", StackSize: 64}
	Jukebox                         = Item{ID: 256, DisplayName: "Jukebox", Name: "jukebox", StackSize: 64}
	OakFence                        = Item{ID: 257, DisplayName: "Oak Fence", Name: "oak_fence", StackSize: 64}
	SpruceFence                     = Item{ID: 258, DisplayName: "Spruce Fence", Name: "spruce_fence", StackSize: 64}
	BirchFence                      = Item{ID: 259, DisplayName: "Birch Fence", Name: "birch_fence", StackSize: 64}
	JungleFence                     = Item{ID: 260, DisplayName: "Jungle Fence", Name: "jungle_fence", StackSize: 64}
	AcaciaFence                     = Item{ID: 261, DisplayName: "Acacia Fence", Name: "acacia_fence", StackSize: 64}
	DarkOakFence                    = Item{ID: 262, DisplayName: "Dark Oak Fence", Name: "dark_oak_fence", StackSize: 64}
	CrimsonFence                    = Item{ID: 263, DisplayName: "Crimson Fence", Name: "crimson_fence", StackSize: 64}
	WarpedFence                     = Item{ID: 264, DisplayName: "Warped Fence", Name: "warped_fence", StackSize: 64}
	Pumpkin                         = Item{ID: 265, DisplayName: "Pumpkin", Name: "pumpkin", StackSize: 64}
	CarvedPumpkin                   = Item{ID: 266, DisplayName: "Carved Pumpkin", Name: "carved_pumpkin", StackSize: 64}
	JackOLantern                    = Item{ID: 267, DisplayName: "Jack o'Lantern", Name: "jack_o_lantern", StackSize: 64}
	Netherrack                      = Item{ID: 268, DisplayName: "Netherrack", Name: "netherrack", StackSize: 64}
	SoulSand                        = Item{ID: 269, DisplayName: "Soul Sand", Name: "soul_sand", StackSize: 64}
	SoulSoil                        = Item{ID: 270, DisplayName: "Soul Soil", Name: "soul_soil", StackSize: 64}
	Basalt                          = Item{ID: 271, DisplayName: "Basalt", Name: "basalt", StackSize: 64}
	PolishedBasalt                  = Item{ID: 272, DisplayName: "Polished Basalt", Name: "polished_basalt", StackSize: 64}
	SmoothBasalt                    = Item{ID: 273, DisplayName: "Smooth Basalt", Name: "smooth_basalt", StackSize: 64}
	SoulTorch                       = Item{ID: 274, DisplayName: "Soul Torch", Name: "soul_torch", StackSize: 64}
	Glowstone                       = Item{ID: 275, DisplayName: "Glowstone", Name: "glowstone", StackSize: 64}
	InfestedStone                   = Item{ID: 276, DisplayName: "Infested Stone", Name: "infested_stone", StackSize: 64}
	InfestedCobblestone             = Item{ID: 277, DisplayName: "Infested Cobblestone", Name: "infested_cobblestone", StackSize: 64}
	InfestedStoneBricks             = Item{ID: 278, DisplayName: "Infested Stone Bricks", Name: "infested_stone_bricks", StackSize: 64}
	InfestedMossyStoneBricks        = Item{ID: 279, DisplayName: "Infested Mossy Stone Bricks", Name: "infested_mossy_stone_bricks", StackSize: 64}
	InfestedCrackedStoneBricks      = Item{ID: 280, DisplayName: "Infested Cracked Stone Bricks", Name: "infested_cracked_stone_bricks", StackSize: 64}
	InfestedChiseledStoneBricks     = Item{ID: 281, DisplayName: "Infested Chiseled Stone Bricks", Name: "infested_chiseled_stone_bricks", StackSize: 64}
	InfestedDeepslate               = Item{ID: 282, DisplayName: "Infested Deepslate", Name: "infested_deepslate", StackSize: 64}
	StoneBricks                     = Item{ID: 283, DisplayName: "Stone Bricks", Name: "stone_bricks", StackSize: 64}
	MossyStoneBricks                = Item{ID: 284, DisplayName: "Mossy Stone Bricks", Name: "mossy_stone_bricks", StackSize: 64}
	CrackedStoneBricks              = Item{ID: 285, DisplayName: "Cracked Stone Bricks", Name: "cracked_stone_bricks", StackSize: 64}
	ChiseledStoneBricks             = Item{ID: 286, DisplayName: "Chiseled Stone Bricks", Name: "chiseled_stone_bricks", StackSize: 64}
	DeepslateBricks                 = Item{ID: 287, DisplayName: "Deepslate Bricks", Name: "deepslate_bricks", StackSize: 64}
	CrackedDeepslateBricks          = Item{ID: 288, DisplayName: "Cracked Deepslate Bricks", Name: "cracked_deepslate_bricks", StackSize: 64}
	DeepslateTiles                  = Item{ID: 289, DisplayName: "Deepslate Tiles", Name: "deepslate_tiles", StackSize: 64}
	CrackedDeepslateTiles           = Item{ID: 290, DisplayName: "Cracked Deepslate Tiles", Name: "cracked_deepslate_tiles", StackSize: 64}
	ChiseledDeepslate               = Item{ID: 291, DisplayName: "Chiseled Deepslate", Name: "chiseled_deepslate", StackSize: 64}
	BrownMushroomBlock              = Item{ID: 292, DisplayName: "Brown Mushroom Block", Name: "brown_mushroom_block", StackSize: 64}
	RedMushroomBlock                = Item{ID: 293, DisplayName: "Red Mushroom Block", Name: "red_mushroom_block", StackSize: 64}
	MushroomStem                    = Item{ID: 294, DisplayName: "Mushroom Stem", Name: "mushroom_stem", StackSize: 64}
	IronBars                        = Item{ID: 295, DisplayName: "Iron Bars", Name: "iron_bars", StackSize: 64}
	Chain                           = Item{ID: 296, DisplayName: "Chain", Name: "chain", StackSize: 64}
	GlassPane                       = Item{ID: 297, DisplayName: "Glass Pane", Name: "glass_pane", StackSize: 64}
	Melon                           = Item{ID: 298, DisplayName: "Melon", Name: "melon", StackSize: 64}
	Vine                            = Item{ID: 299, DisplayName: "Vines", Name: "vine", StackSize: 64}
	GlowLichen                      = Item{ID: 300, DisplayName: "Glow Lichen", Name: "glow_lichen", StackSize: 64}
	BrickStairs                     = Item{ID: 301, DisplayName: "Brick Stairs", Name: "brick_stairs", StackSize: 64}
	StoneBrickStairs                = Item{ID: 302, DisplayName: "Stone Brick Stairs", Name: "stone_brick_stairs", StackSize: 64}
	Mycelium                        = Item{ID: 303, DisplayName: "Mycelium", Name: "mycelium", StackSize: 64}
	LilyPad                         = Item{ID: 304, DisplayName: "Lily Pad", Name: "lily_pad", StackSize: 64}
	NetherBricks                    = Item{ID: 305, DisplayName: "Nether Bricks", Name: "nether_bricks", StackSize: 64}
	CrackedNetherBricks             = Item{ID: 306, DisplayName: "Cracked Nether Bricks", Name: "cracked_nether_bricks", StackSize: 64}
	ChiseledNetherBricks            = Item{ID: 307, DisplayName: "Chiseled Nether Bricks", Name: "chiseled_nether_bricks", StackSize: 64}
	NetherBrickFence                = Item{ID: 308, DisplayName: "Nether Brick Fence", Name: "nether_brick_fence", StackSize: 64}
	NetherBrickStairs               = Item{ID: 309, DisplayName: "Nether Brick Stairs", Name: "nether_brick_stairs", StackSize: 64}
	EnchantingTable                 = Item{ID: 310, DisplayName: "Enchanting Table", Name: "enchanting_table", StackSize: 64}
	EndPortalFrame                  = Item{ID: 311, DisplayName: "End Portal Frame", Name: "end_portal_frame", StackSize: 64}
	EndStone                        = Item{ID: 312, DisplayName: "End Stone", Name: "end_stone", StackSize: 64}
	EndStoneBricks                  = Item{ID: 313, DisplayName: "End Stone Bricks", Name: "end_stone_bricks", StackSize: 64}
	DragonEgg                       = Item{ID: 314, DisplayName: "Dragon Egg", Name: "dragon_egg", StackSize: 64}
	SandstoneStairs                 = Item{ID: 315, DisplayName: "Sandstone Stairs", Name: "sandstone_stairs", StackSize: 64}
	EnderChest                      = Item{ID: 316, DisplayName: "Ender Chest", Name: "ender_chest", StackSize: 64}
	EmeraldBlock                    = Item{ID: 317, DisplayName: "Block of Emerald", Name: "emerald_block", StackSize: 64}
	SpruceStairs                    = Item{ID: 318, DisplayName: "Spruce Stairs", Name: "spruce_stairs", StackSize: 64}
	BirchStairs                     = Item{ID: 319, DisplayName: "Birch Stairs", Name: "birch_stairs", StackSize: 64}
	JungleStairs                    = Item{ID: 320, DisplayName: "Jungle Stairs", Name: "jungle_stairs", StackSize: 64}
	CrimsonStairs                   = Item{ID: 321, DisplayName: "Crimson Stairs", Name: "crimson_stairs", StackSize: 64}
	WarpedStairs                    = Item{ID: 322, DisplayName: "Warped Stairs", Name: "warped_stairs", StackSize: 64}
	CommandBlock                    = Item{ID: 323, DisplayName: "Command Block", Name: "command_block", StackSize: 64}
	Beacon                          = Item{ID: 324, DisplayName: "Beacon", Name: "beacon", StackSize: 64}
	CobblestoneWall                 = Item{ID: 325, DisplayName: "Cobblestone Wall", Name: "cobblestone_wall", StackSize: 64}
	MossyCobblestoneWall            = Item{ID: 326, DisplayName: "Mossy Cobblestone Wall", Name: "mossy_cobblestone_wall", StackSize: 64}
	BrickWall                       = Item{ID: 327, DisplayName: "Brick Wall", Name: "brick_wall", StackSize: 64}
	PrismarineWall                  = Item{ID: 328, DisplayName: "Prismarine Wall", Name: "prismarine_wall", StackSize: 64}
	RedSandstoneWall                = Item{ID: 329, DisplayName: "Red Sandstone Wall", Name: "red_sandstone_wall", StackSize: 64}
	MossyStoneBrickWall             = Item{ID: 330, DisplayName: "Mossy Stone Brick Wall", Name: "mossy_stone_brick_wall", StackSize: 64}
	GraniteWall                     = Item{ID: 331, DisplayName: "Granite Wall", Name: "granite_wall", StackSize: 64}
	StoneBrickWall                  = Item{ID: 332, DisplayName: "Stone Brick Wall", Name: "stone_brick_wall", StackSize: 64}
	NetherBrickWall                 = Item{ID: 333, DisplayName: "Nether Brick Wall", Name: "nether_brick_wall", StackSize: 64}
	AndesiteWall                    = Item{ID: 334, DisplayName: "Andesite Wall", Name: "andesite_wall", StackSize: 64}
	RedNetherBrickWall              = Item{ID: 335, DisplayName: "Red Nether Brick Wall", Name: "red_nether_brick_wall", StackSize: 64}
	SandstoneWall                   = Item{ID: 336, DisplayName: "Sandstone Wall", Name: "sandstone_wall", StackSize: 64}
	EndStoneBrickWall               = Item{ID: 337, DisplayName: "End Stone Brick Wall", Name: "end_stone_brick_wall", StackSize: 64}
	DioriteWall                     = Item{ID: 338, DisplayName: "Diorite Wall", Name: "diorite_wall", StackSize: 64}
	BlackstoneWall                  = Item{ID: 339, DisplayName: "Blackstone Wall", Name: "blackstone_wall", StackSize: 64}
	PolishedBlackstoneWall          = Item{ID: 340, DisplayName: "Polished Blackstone Wall", Name: "polished_blackstone_wall", StackSize: 64}
	PolishedBlackstoneBrickWall     = Item{ID: 341, DisplayName: "Polished Blackstone Brick Wall", Name: "polished_blackstone_brick_wall", StackSize: 64}
	CobbledDeepslateWall            = Item{ID: 342, DisplayName: "Cobbled Deepslate Wall", Name: "cobbled_deepslate_wall", StackSize: 64}
	PolishedDeepslateWall           = Item{ID: 343, DisplayName: "Polished Deepslate Wall", Name: "polished_deepslate_wall", StackSize: 64}
	DeepslateBrickWall              = Item{ID: 344, DisplayName: "Deepslate Brick Wall", Name: "deepslate_brick_wall", StackSize: 64}
	DeepslateTileWall               = Item{ID: 345, DisplayName: "Deepslate Tile Wall", Name: "deepslate_tile_wall", StackSize: 64}
	Anvil                           = Item{ID: 346, DisplayName: "Anvil", Name: "anvil", StackSize: 64}
	ChippedAnvil                    = Item{ID: 347, DisplayName: "Chipped Anvil", Name: "chipped_anvil", StackSize: 64}
	DamagedAnvil                    = Item{ID: 348, DisplayName: "Damaged Anvil", Name: "damaged_anvil", StackSize: 64}
	ChiseledQuartzBlock             = Item{ID: 349, DisplayName: "Chiseled Quartz Block", Name: "chiseled_quartz_block", StackSize: 64}
	QuartzBlock                     = Item{ID: 350, DisplayName: "Block of Quartz", Name: "quartz_block", StackSize: 64}
	QuartzBricks                    = Item{ID: 351, DisplayName: "Quartz Bricks", Name: "quartz_bricks", StackSize: 64}
	QuartzPillar                    = Item{ID: 352, DisplayName: "Quartz Pillar", Name: "quartz_pillar", StackSize: 64}
	QuartzStairs                    = Item{ID: 353, DisplayName: "Quartz Stairs", Name: "quartz_stairs", StackSize: 64}
	WhiteTerracotta                 = Item{ID: 354, DisplayName: "White Terracotta", Name: "white_terracotta", StackSize: 64}
	OrangeTerracotta                = Item{ID: 355, DisplayName: "Orange Terracotta", Name: "orange_terracotta", StackSize: 64}
	MagentaTerracotta               = Item{ID: 356, DisplayName: "Magenta Terracotta", Name: "magenta_terracotta", StackSize: 64}
	LightBlueTerracotta             = Item{ID: 357, DisplayName: "Light Blue Terracotta", Name: "light_blue_terracotta", StackSize: 64}
	YellowTerracotta                = Item{ID: 358, DisplayName: "Yellow Terracotta", Name: "yellow_terracotta", StackSize: 64}
	LimeTerracotta                  = Item{ID: 359, DisplayName: "Lime Terracotta", Name: "lime_terracotta", StackSize: 64}
	PinkTerracotta                  = Item{ID: 360, DisplayName: "Pink Terracotta", Name: "pink_terracotta", StackSize: 64}
	GrayTerracotta                  = Item{ID: 361, DisplayName: "Gray Terracotta", Name: "gray_terracotta", StackSize: 64}
	LightGrayTerracotta             = Item{ID: 362, DisplayName: "Light Gray Terracotta", Name: "light_gray_terracotta", StackSize: 64}
	CyanTerracotta                  = Item{ID: 363, DisplayName: "Cyan Terracotta", Name: "cyan_terracotta", StackSize: 64}
	PurpleTerracotta                = Item{ID: 364, DisplayName: "Purple Terracotta", Name: "purple_terracotta", StackSize: 64}
	BlueTerracotta                  = Item{ID: 365, DisplayName: "Blue Terracotta", Name: "blue_terracotta", StackSize: 64}
	BrownTerracotta                 = Item{ID: 366, DisplayName: "Brown Terracotta", Name: "brown_terracotta", StackSize: 64}
	GreenTerracotta                 = Item{ID: 367, DisplayName: "Green Terracotta", Name: "green_terracotta", StackSize: 64}
	RedTerracotta                   = Item{ID: 368, DisplayName: "Red Terracotta", Name: "red_terracotta", StackSize: 64}
	BlackTerracotta                 = Item{ID: 369, DisplayName: "Black Terracotta", Name: "black_terracotta", StackSize: 64}
	Barrier                         = Item{ID: 370, DisplayName: "Barrier", Name: "barrier", StackSize: 64}
	Light                           = Item{ID: 371, DisplayName: "Light", Name: "light", StackSize: 64}
	HayBlock                        = Item{ID: 372, DisplayName: "Hay Bale", Name: "hay_block", StackSize: 64}
	WhiteCarpet                     = Item{ID: 373, DisplayName: "White Carpet", Name: "white_carpet", StackSize: 64}
	OrangeCarpet                    = Item{ID: 374, DisplayName: "Orange Carpet", Name: "orange_carpet", StackSize: 64}
	MagentaCarpet                   = Item{ID: 375, DisplayName: "Magenta Carpet", Name: "magenta_carpet", StackSize: 64}
	LightBlueCarpet                 = Item{ID: 376, DisplayName: "Light Blue Carpet", Name: "light_blue_carpet", StackSize: 64}
	YellowCarpet                    = Item{ID: 377, DisplayName: "Yellow Carpet", Name: "yellow_carpet", StackSize: 64}
	LimeCarpet                      = Item{ID: 378, DisplayName: "Lime Carpet", Name: "lime_carpet", StackSize: 64}
	PinkCarpet                      = Item{ID: 379, DisplayName: "Pink Carpet", Name: "pink_carpet", StackSize: 64}
	GrayCarpet                      = Item{ID: 380, DisplayName: "Gray Carpet", Name: "gray_carpet", StackSize: 64}
	LightGrayCarpet                 = Item{ID: 381, DisplayName: "Light Gray Carpet", Name: "light_gray_carpet", StackSize: 64}
	CyanCarpet                      = Item{ID: 382, DisplayName: "Cyan Carpet", Name: "cyan_carpet", StackSize: 64}
	PurpleCarpet                    = Item{ID: 383, DisplayName: "Purple Carpet", Name: "purple_carpet", StackSize: 64}
	BlueCarpet                      = Item{ID: 384, DisplayName: "Blue Carpet", Name: "blue_carpet", StackSize: 64}
	BrownCarpet                     = Item{ID: 385, DisplayName: "Brown Carpet", Name: "brown_carpet", StackSize: 64}
	GreenCarpet                     = Item{ID: 386, DisplayName: "Green Carpet", Name: "green_carpet", StackSize: 64}
	RedCarpet                       = Item{ID: 387, DisplayName: "Red Carpet", Name: "red_carpet", StackSize: 64}
	BlackCarpet                     = Item{ID: 388, DisplayName: "Black Carpet", Name: "black_carpet", StackSize: 64}
	Terracotta                      = Item{ID: 389, DisplayName: "Terracotta", Name: "terracotta", StackSize: 64}
	PackedIce                       = Item{ID: 390, DisplayName: "Packed Ice", Name: "packed_ice", StackSize: 64}
	AcaciaStairs                    = Item{ID: 391, DisplayName: "Acacia Stairs", Name: "acacia_stairs", StackSize: 64}
	DarkOakStairs                   = Item{ID: 392, DisplayName: "Dark Oak Stairs", Name: "dark_oak_stairs", StackSize: 64}
	DirtPath                        = Item{ID: 393, DisplayName: "Dirt Path", Name: "dirt_path", StackSize: 64}
	Sunflower                       = Item{ID: 394, DisplayName: "Sunflower", Name: "sunflower", StackSize: 64}
	Lilac                           = Item{ID: 395, DisplayName: "Lilac", Name: "lilac", StackSize: 64}
	RoseBush                        = Item{ID: 396, DisplayName: "Rose Bush", Name: "rose_bush", StackSize: 64}
	Peony                           = Item{ID: 397, DisplayName: "Peony", Name: "peony", StackSize: 64}
	TallGrass                       = Item{ID: 398, DisplayName: "Tall Grass", Name: "tall_grass", StackSize: 64}
	LargeFern                       = Item{ID: 399, DisplayName: "Large Fern", Name: "large_fern", StackSize: 64}
	WhiteStainedGlass               = Item{ID: 400, DisplayName: "White Stained Glass", Name: "white_stained_glass", StackSize: 64}
	OrangeStainedGlass              = Item{ID: 401, DisplayName: "Orange Stained Glass", Name: "orange_stained_glass", StackSize: 64}
	MagentaStainedGlass             = Item{ID: 402, DisplayName: "Magenta Stained Glass", Name: "magenta_stained_glass", StackSize: 64}
	LightBlueStainedGlass           = Item{ID: 403, DisplayName: "Light Blue Stained Glass", Name: "light_blue_stained_glass", StackSize: 64}
	YellowStainedGlass              = Item{ID: 404, DisplayName: "Yellow Stained Glass", Name: "yellow_stained_glass", StackSize: 64}
	LimeStainedGlass                = Item{ID: 405, DisplayName: "Lime Stained Glass", Name: "lime_stained_glass", StackSize: 64}
	PinkStainedGlass                = Item{ID: 406, DisplayName: "Pink Stained Glass", Name: "pink_stained_glass", StackSize: 64}
	GrayStainedGlass                = Item{ID: 407, DisplayName: "Gray Stained Glass", Name: "gray_stained_glass", StackSize: 64}
	LightGrayStainedGlass           = Item{ID: 408, DisplayName: "Light Gray Stained Glass", Name: "light_gray_stained_glass", StackSize: 64}
	CyanStainedGlass                = Item{ID: 409, DisplayName: "Cyan Stained Glass", Name: "cyan_stained_glass", StackSize: 64}
	PurpleStainedGlass              = Item{ID: 410, DisplayName: "Purple Stained Glass", Name: "purple_stained_glass", StackSize: 64}
	BlueStainedGlass                = Item{ID: 411, DisplayName: "Blue Stained Glass", Name: "blue_stained_glass", StackSize: 64}
	BrownStainedGlass               = Item{ID: 412, DisplayName: "Brown Stained Glass", Name: "brown_stained_glass", StackSize: 64}
	GreenStainedGlass               = Item{ID: 413, DisplayName: "Green Stained Glass", Name: "green_stained_glass", StackSize: 64}
	RedStainedGlass                 = Item{ID: 414, DisplayName: "Red Stained Glass", Name: "red_stained_glass", StackSize: 64}
	BlackStainedGlass               = Item{ID: 415, DisplayName: "Black Stained Glass", Name: "black_stained_glass", StackSize: 64}
	WhiteStainedGlassPane           = Item{ID: 416, DisplayName: "White Stained Glass Pane", Name: "white_stained_glass_pane", StackSize: 64}
	OrangeStainedGlassPane          = Item{ID: 417, DisplayName: "Orange Stained Glass Pane", Name: "orange_stained_glass_pane", StackSize: 64}
	MagentaStainedGlassPane         = Item{ID: 418, DisplayName: "Magenta Stained Glass Pane", Name: "magenta_stained_glass_pane", StackSize: 64}
	LightBlueStainedGlassPane       = Item{ID: 419, DisplayName: "Light Blue Stained Glass Pane", Name: "light_blue_stained_glass_pane", StackSize: 64}
	YellowStainedGlassPane          = Item{ID: 420, DisplayName: "Yellow Stained Glass Pane", Name: "yellow_stained_glass_pane", StackSize: 64}
	LimeStainedGlassPane            = Item{ID: 421, DisplayName: "Lime Stained Glass Pane", Name: "lime_stained_glass_pane", StackSize: 64}
	PinkStainedGlassPane            = Item{ID: 422, DisplayName: "Pink Stained Glass Pane", Name: "pink_stained_glass_pane", StackSize: 64}
	GrayStainedGlassPane            = Item{ID: 423, DisplayName: "Gray Stained Glass Pane", Name: "gray_stained_glass_pane", StackSize: 64}
	LightGrayStainedGlassPane       = Item{ID: 424, DisplayName: "Light Gray Stained Glass Pane", Name: "light_gray_stained_glass_pane", StackSize: 64}
	CyanStainedGlassPane            = Item{ID: 425, DisplayName: "Cyan Stained Glass Pane", Name: "cyan_stained_glass_pane", StackSize: 64}
	PurpleStainedGlassPane          = Item{ID: 426, DisplayName: "Purple Stained Glass Pane", Name: "purple_stained_glass_pane", StackSize: 64}
	BlueStainedGlassPane            = Item{ID: 427, DisplayName: "Blue Stained Glass Pane", Name: "blue_stained_glass_pane", StackSize: 64}
	BrownStainedGlassPane           = Item{ID: 428, DisplayName: "Brown Stained Glass Pane", Name: "brown_stained_glass_pane", StackSize: 64}
	GreenStainedGlassPane           = Item{ID: 429, DisplayName: "Green Stained Glass Pane", Name: "green_stained_glass_pane", StackSize: 64}
	RedStainedGlassPane             = Item{ID: 430, DisplayName: "Red Stained Glass Pane", Name: "red_stained_glass_pane", StackSize: 64}
	BlackStainedGlassPane           = Item{ID: 431, DisplayName: "Black Stained Glass Pane", Name: "black_stained_glass_pane", StackSize: 64}
	Prismarine                      = Item{ID: 432, DisplayName: "Prismarine", Name: "prismarine", StackSize: 64}
	PrismarineBricks                = Item{ID: 433, DisplayName: "Prismarine Bricks", Name: "prismarine_bricks", StackSize: 64}
	DarkPrismarine                  = Item{ID: 434, DisplayName: "Dark Prismarine", Name: "dark_prismarine", StackSize: 64}
	PrismarineStairs                = Item{ID: 435, DisplayName: "Prismarine Stairs", Name: "prismarine_stairs", StackSize: 64}
	PrismarineBrickStairs           = Item{ID: 436, DisplayName: "Prismarine Brick Stairs", Name: "prismarine_brick_stairs", StackSize: 64}
	DarkPrismarineStairs            = Item{ID: 437, DisplayName: "Dark Prismarine Stairs", Name: "dark_prismarine_stairs", StackSize: 64}
	SeaLantern                      = Item{ID: 438, DisplayName: "Sea Lantern", Name: "sea_lantern", StackSize: 64}
	RedSandstone                    = Item{ID: 439, DisplayName: "Red Sandstone", Name: "red_sandstone", StackSize: 64}
	ChiseledRedSandstone            = Item{ID: 440, DisplayName: "Chiseled Red Sandstone", Name: "chiseled_red_sandstone", StackSize: 64}
	CutRedSandstone                 = Item{ID: 441, DisplayName: "Cut Red Sandstone", Name: "cut_red_sandstone", StackSize: 64}
	RedSandstoneStairs              = Item{ID: 442, DisplayName: "Red Sandstone Stairs", Name: "red_sandstone_stairs", StackSize: 64}
	RepeatingCommandBlock           = Item{ID: 443, DisplayName: "Repeating Command Block", Name: "repeating_command_block", StackSize: 64}
	ChainCommandBlock               = Item{ID: 444, DisplayName: "Chain Command Block", Name: "chain_command_block", StackSize: 64}
	MagmaBlock                      = Item{ID: 445, DisplayName: "Magma Block", Name: "magma_block", StackSize: 64}
	NetherWartBlock                 = Item{ID: 446, DisplayName: "Nether Wart Block", Name: "nether_wart_block", StackSize: 64}
	WarpedWartBlock                 = Item{ID: 447, DisplayName: "Warped Wart Block", Name: "warped_wart_block", StackSize: 64}
	RedNetherBricks                 = Item{ID: 448, DisplayName: "Red Nether Bricks", Name: "red_nether_bricks", StackSize: 64}
	BoneBlock                       = Item{ID: 449, DisplayName: "Bone Block", Name: "bone_block", StackSize: 64}
	StructureVoid                   = Item{ID: 450, DisplayName: "Structure Void", Name: "structure_void", StackSize: 64}
	ShulkerBox                      = Item{ID: 451, DisplayName: "Shulker Box", Name: "shulker_box", StackSize: 1}
	WhiteShulkerBox                 = Item{ID: 452, DisplayName: "White Shulker Box", Name: "white_shulker_box", StackSize: 1}
	OrangeShulkerBox                = Item{ID: 453, DisplayName: "Orange Shulker Box", Name: "orange_shulker_box", StackSize: 1}
	MagentaShulkerBox               = Item{ID: 454, DisplayName: "Magenta Shulker Box", Name: "magenta_shulker_box", StackSize: 1}
	LightBlueShulkerBox             = Item{ID: 455, DisplayName: "Light Blue Shulker Box", Name: "light_blue_shulker_box", StackSize: 1}
	YellowShulkerBox                = Item{ID: 456, DisplayName: "Yellow Shulker Box", Name: "yellow_shulker_box", StackSize: 1}
	LimeShulkerBox                  = Item{ID: 457, DisplayName: "Lime Shulker Box", Name: "lime_shulker_box", StackSize: 1}
	PinkShulkerBox                  = Item{ID: 458, DisplayName: "Pink Shulker Box", Name: "pink_shulker_box", StackSize: 1}
	GrayShulkerBox                  = Item{ID: 459, DisplayName: "Gray Shulker Box", Name: "gray_shulker_box", StackSize: 1}
	LightGrayShulkerBox             = Item{ID: 460, DisplayName: "Light Gray Shulker Box", Name: "light_gray_shulker_box", StackSize: 1}
	CyanShulkerBox                  = Item{ID: 461, DisplayName: "Cyan Shulker Box", Name: "cyan_shulker_box", StackSize: 1}
	PurpleShulkerBox                = Item{ID: 462, DisplayName: "Purple Shulker Box", Name: "purple_shulker_box", StackSize: 1}
	BlueShulkerBox                  = Item{ID: 463, DisplayName: "Blue Shulker Box", Name: "blue_shulker_box", StackSize: 1}
	BrownShulkerBox                 = Item{ID: 464, DisplayName: "Brown Shulker Box", Name: "brown_shulker_box", StackSize: 1}
	GreenShulkerBox                 = Item{ID: 465, DisplayName: "Green Shulker Box", Name: "green_shulker_box", StackSize: 1}
	RedShulkerBox                   = Item{ID: 466, DisplayName: "Red Shulker Box", Name: "red_shulker_box", StackSize: 1}
	BlackShulkerBox                 = Item{ID: 467, DisplayName: "Black Shulker Box", Name: "black_shulker_box", StackSize: 1}
	WhiteGlazedTerracotta           = Item{ID: 468, DisplayName: "White Glazed Terracotta", Name: "white_glazed_terracotta", StackSize: 64}
	OrangeGlazedTerracotta          = Item{ID: 469, DisplayName: "Orange Glazed Terracotta", Name: "orange_glazed_terracotta", StackSize: 64}
	MagentaGlazedTerracotta         = Item{ID: 470, DisplayName: "Magenta Glazed Terracotta", Name: "magenta_glazed_terracotta", StackSize: 64}
	LightBlueGlazedTerracotta       = Item{ID: 471, DisplayName: "Light Blue Glazed Terracotta", Name: "light_blue_glazed_terracotta", StackSize: 64}
	YellowGlazedTerracotta          = Item{ID: 472, DisplayName: "Yellow Glazed Terracotta", Name: "yellow_glazed_terracotta", StackSize: 64}
	LimeGlazedTerracotta            = Item{ID: 473, DisplayName: "Lime Glazed Terracotta", Name: "lime_glazed_terracotta", StackSize: 64}
	PinkGlazedTerracotta            = Item{ID: 474, DisplayName: "Pink Glazed Terracotta", Name: "pink_glazed_terracotta", StackSize: 64}
	GrayGlazedTerracotta            = Item{ID: 475, DisplayName: "Gray Glazed Terracotta", Name: "gray_glazed_terracotta", StackSize: 64}
	LightGrayGlazedTerracotta       = Item{ID: 476, DisplayName: "Light Gray Glazed Terracotta", Name: "light_gray_glazed_terracotta", StackSize: 64}
	CyanGlazedTerracotta            = Item{ID: 477, DisplayName: "Cyan Glazed Terracotta", Name: "cyan_glazed_terracotta", StackSize: 64}
	PurpleGlazedTerracotta          = Item{ID: 478, DisplayName: "Purple Glazed Terracotta", Name: "purple_glazed_terracotta", StackSize: 64}
	BlueGlazedTerracotta            = Item{ID: 479, DisplayName: "Blue Glazed Terracotta", Name: "blue_glazed_terracotta", StackSize: 64}
	BrownGlazedTerracotta           = Item{ID: 480, DisplayName: "Brown Glazed Terracotta", Name: "brown_glazed_terracotta", StackSize: 64}
	GreenGlazedTerracotta           = Item{ID: 481, DisplayName: "Green Glazed Terracotta", Name: "green_glazed_terracotta", StackSize: 64}
	RedGlazedTerracotta             = Item{ID: 482, DisplayName: "Red Glazed Terracotta", Name: "red_glazed_terracotta", StackSize: 64}
	BlackGlazedTerracotta           = Item{ID: 483, DisplayName: "Black Glazed Terracotta", Name: "black_glazed_terracotta", StackSize: 64}
	WhiteConcrete                   = Item{ID: 484, DisplayName: "White Concrete", Name: "white_concrete", StackSize: 64}
	OrangeConcrete                  = Item{ID: 485, DisplayName: "Orange Concrete", Name: "orange_concrete", StackSize: 64}
	MagentaConcrete                 = Item{ID: 486, DisplayName: "Magenta Concrete", Name: "magenta_concrete", StackSize: 64}
	LightBlueConcrete               = Item{ID: 487, DisplayName: "Light Blue Concrete", Name: "light_blue_concrete", StackSize: 64}
	YellowConcrete                  = Item{ID: 488, DisplayName: "Yellow Concrete", Name: "yellow_concrete", StackSize: 64}
	LimeConcrete                    = Item{ID: 489, DisplayName: "Lime Concrete", Name: "lime_concrete", StackSize: 64}
	PinkConcrete                    = Item{ID: 490, DisplayName: "Pink Concrete", Name: "pink_concrete", StackSize: 64}
	GrayConcrete                    = Item{ID: 491, DisplayName: "Gray Concrete", Name: "gray_concrete", StackSize: 64}
	LightGrayConcrete               = Item{ID: 492, DisplayName: "Light Gray Concrete", Name: "light_gray_concrete", StackSize: 64}
	CyanConcrete                    = Item{ID: 493, DisplayName: "Cyan Concrete", Name: "cyan_concrete", StackSize: 64}
	PurpleConcrete                  = Item{ID: 494, DisplayName: "Purple Concrete", Name: "purple_concrete", StackSize: 64}
	BlueConcrete                    = Item{ID: 495, DisplayName: "Blue Concrete", Name: "blue_concrete", StackSize: 64}
	BrownConcrete                   = Item{ID: 496, DisplayName: "Brown Concrete", Name: "brown_concrete", StackSize: 64}
	GreenConcrete                   = Item{ID: 497, DisplayName: "Green Concrete", Name: "green_concrete", StackSize: 64}
	RedConcrete                     = Item{ID: 498, DisplayName: "Red Concrete", Name: "red_concrete", StackSize: 64}
	BlackConcrete                   = Item{ID: 499, DisplayName: "Black Concrete", Name: "black_concrete", StackSize: 64}
	WhiteConcretePowder             = Item{ID: 500, DisplayName: "White Concrete Powder", Name: "white_concrete_powder", StackSize: 64}
	OrangeConcretePowder            = Item{ID: 501, DisplayName: "Orange Concrete Powder", Name: "orange_concrete_powder", StackSize: 64}
	MagentaConcretePowder           = Item{ID: 502, DisplayName: "Magenta Concrete Powder", Name: "magenta_concrete_powder", StackSize: 64}
	LightBlueConcretePowder         = Item{ID: 503, DisplayName: "Light Blue Concrete Powder", Name: "light_blue_concrete_powder", StackSize: 64}
	YellowConcretePowder            = Item{ID: 504, DisplayName: "Yellow Concrete Powder", Name: "yellow_concrete_powder", StackSize: 64}
	LimeConcretePowder              = Item{ID: 505, DisplayName: "Lime Concrete Powder", Name: "lime_concrete_powder", StackSize: 64}
	PinkConcretePowder              = Item{ID: 506, DisplayName: "Pink Concrete Powder", Name: "pink_concrete_powder", StackSize: 64}
	GrayConcretePowder              = Item{ID: 507, DisplayName: "Gray Concrete Powder", Name: "gray_concrete_powder", StackSize: 64}
	LightGrayConcretePowder         = Item{ID: 508, DisplayName: "Light Gray Concrete Powder", Name: "light_gray_concrete_powder", StackSize: 64}
	CyanConcretePowder              = Item{ID: 509, DisplayName: "Cyan Concrete Powder", Name: "cyan_concrete_powder", StackSize: 64}
	PurpleConcretePowder            = Item{ID: 510, DisplayName: "Purple Concrete Powder", Name: "purple_concrete_powder", StackSize: 64}
	BlueConcretePowder              = Item{ID: 511, DisplayName: "Blue Concrete Powder", Name: "blue_concrete_powder", StackSize: 64}
	BrownConcretePowder             = Item{ID: 512, DisplayName: "Brown Concrete Powder", Name: "brown_concrete_powder", StackSize: 64}
	GreenConcretePowder             = Item{ID: 513, DisplayName: "Green Concrete Powder", Name: "green_concrete_powder", StackSize: 64}
	RedConcretePowder               = Item{ID: 514, DisplayName: "Red Concrete Powder", Name: "red_concrete_powder", StackSize: 64}
	BlackConcretePowder             = Item{ID: 515, DisplayName: "Black Concrete Powder", Name: "black_concrete_powder", StackSize: 64}
	TurtleEgg                       = Item{ID: 516, DisplayName: "Turtle Egg", Name: "turtle_egg", StackSize: 64}
	DeadTubeCoralBlock              = Item{ID: 517, DisplayName: "Dead Tube Coral Block", Name: "dead_tube_coral_block", StackSize: 64}
	DeadBrainCoralBlock             = Item{ID: 518, DisplayName: "Dead Brain Coral Block", Name: "dead_brain_coral_block", StackSize: 64}
	DeadBubbleCoralBlock            = Item{ID: 519, DisplayName: "Dead Bubble Coral Block", Name: "dead_bubble_coral_block", StackSize: 64}
	DeadFireCoralBlock              = Item{ID: 520, DisplayName: "Dead Fire Coral Block", Name: "dead_fire_coral_block", StackSize: 64}
	DeadHornCoralBlock              = Item{ID: 521, DisplayName: "Dead Horn Coral Block", Name: "dead_horn_coral_block", StackSize: 64}
	TubeCoralBlock                  = Item{ID: 522, DisplayName: "Tube Coral Block", Name: "tube_coral_block", StackSize: 64}
	BrainCoralBlock                 = Item{ID: 523, DisplayName: "Brain Coral Block", Name: "brain_coral_block", StackSize: 64}
	BubbleCoralBlock                = Item{ID: 524, DisplayName: "Bubble Coral Block", Name: "bubble_coral_block", StackSize: 64}
	FireCoralBlock                  = Item{ID: 525, DisplayName: "Fire Coral Block", Name: "fire_coral_block", StackSize: 64}
	HornCoralBlock                  = Item{ID: 526, DisplayName: "Horn Coral Block", Name: "horn_coral_block", StackSize: 64}
	TubeCoral                       = Item{ID: 527, DisplayName: "Tube Coral", Name: "tube_coral", StackSize: 64}
	BrainCoral                      = Item{ID: 528, DisplayName: "Brain Coral", Name: "brain_coral", StackSize: 64}
	BubbleCoral                     = Item{ID: 529, DisplayName: "Bubble Coral", Name: "bubble_coral", StackSize: 64}
	FireCoral                       = Item{ID: 530, DisplayName: "Fire Coral", Name: "fire_coral", StackSize: 64}
	HornCoral                       = Item{ID: 531, DisplayName: "Horn Coral", Name: "horn_coral", StackSize: 64}
	DeadBrainCoral                  = Item{ID: 532, DisplayName: "Dead Brain Coral", Name: "dead_brain_coral", StackSize: 64}
	DeadBubbleCoral                 = Item{ID: 533, DisplayName: "Dead Bubble Coral", Name: "dead_bubble_coral", StackSize: 64}
	DeadFireCoral                   = Item{ID: 534, DisplayName: "Dead Fire Coral", Name: "dead_fire_coral", StackSize: 64}
	DeadHornCoral                   = Item{ID: 535, DisplayName: "Dead Horn Coral", Name: "dead_horn_coral", StackSize: 64}
	DeadTubeCoral                   = Item{ID: 536, DisplayName: "Dead Tube Coral", Name: "dead_tube_coral", StackSize: 64}
	TubeCoralFan                    = Item{ID: 537, DisplayName: "Tube Coral Fan", Name: "tube_coral_fan", StackSize: 64}
	BrainCoralFan                   = Item{ID: 538, DisplayName: "Brain Coral Fan", Name: "brain_coral_fan", StackSize: 64}
	BubbleCoralFan                  = Item{ID: 539, DisplayName: "Bubble Coral Fan", Name: "bubble_coral_fan", StackSize: 64}
	FireCoralFan                    = Item{ID: 540, DisplayName: "Fire Coral Fan", Name: "fire_coral_fan", StackSize: 64}
	HornCoralFan                    = Item{ID: 541, DisplayName: "Horn Coral Fan", Name: "horn_coral_fan", StackSize: 64}
	DeadTubeCoralFan                = Item{ID: 542, DisplayName: "Dead Tube Coral Fan", Name: "dead_tube_coral_fan", StackSize: 64}
	DeadBrainCoralFan               = Item{ID: 543, DisplayName: "Dead Brain Coral Fan", Name: "dead_brain_coral_fan", StackSize: 64}
	DeadBubbleCoralFan              = Item{ID: 544, DisplayName: "Dead Bubble Coral Fan", Name: "dead_bubble_coral_fan", StackSize: 64}
	DeadFireCoralFan                = Item{ID: 545, DisplayName: "Dead Fire Coral Fan", Name: "dead_fire_coral_fan", StackSize: 64}
	DeadHornCoralFan                = Item{ID: 546, DisplayName: "Dead Horn Coral Fan", Name: "dead_horn_coral_fan", StackSize: 64}
	BlueIce                         = Item{ID: 547, DisplayName: "Blue Ice", Name: "blue_ice", StackSize: 64}
	Conduit                         = Item{ID: 548, DisplayName: "Conduit", Name: "conduit", StackSize: 64}
	PolishedGraniteStairs           = Item{ID: 549, DisplayName: "Polished Granite Stairs", Name: "polished_granite_stairs", StackSize: 64}
	SmoothRedSandstoneStairs        = Item{ID: 550, DisplayName: "Smooth Red Sandstone Stairs", Name: "smooth_red_sandstone_stairs", StackSize: 64}
	MossyStoneBrickStairs           = Item{ID: 551, DisplayName: "Mossy Stone Brick Stairs", Name: "mossy_stone_brick_stairs", StackSize: 64}
	PolishedDioriteStairs           = Item{ID: 552, DisplayName: "Polished Diorite Stairs", Name: "polished_diorite_stairs", StackSize: 64}
	MossyCobblestoneStairs          = Item{ID: 553, DisplayName: "Mossy Cobblestone Stairs", Name: "mossy_cobblestone_stairs", StackSize: 64}
	EndStoneBrickStairs             = Item{ID: 554, DisplayName: "End Stone Brick Stairs", Name: "end_stone_brick_stairs", StackSize: 64}
	StoneStairs                     = Item{ID: 555, DisplayName: "Stone Stairs", Name: "stone_stairs", StackSize: 64}
	SmoothSandstoneStairs           = Item{ID: 556, DisplayName: "Smooth Sandstone Stairs", Name: "smooth_sandstone_stairs", StackSize: 64}
	SmoothQuartzStairs              = Item{ID: 557, DisplayName: "Smooth Quartz Stairs", Name: "smooth_quartz_stairs", StackSize: 64}
	GraniteStairs                   = Item{ID: 558, DisplayName: "Granite Stairs", Name: "granite_stairs", StackSize: 64}
	AndesiteStairs                  = Item{ID: 559, DisplayName: "Andesite Stairs", Name: "andesite_stairs", StackSize: 64}
	RedNetherBrickStairs            = Item{ID: 560, DisplayName: "Red Nether Brick Stairs", Name: "red_nether_brick_stairs", StackSize: 64}
	PolishedAndesiteStairs          = Item{ID: 561, DisplayName: "Polished Andesite Stairs", Name: "polished_andesite_stairs", StackSize: 64}
	DioriteStairs                   = Item{ID: 562, DisplayName: "Diorite Stairs", Name: "diorite_stairs", StackSize: 64}
	CobbledDeepslateStairs          = Item{ID: 563, DisplayName: "Cobbled Deepslate Stairs", Name: "cobbled_deepslate_stairs", StackSize: 64}
	PolishedDeepslateStairs         = Item{ID: 564, DisplayName: "Polished Deepslate Stairs", Name: "polished_deepslate_stairs", StackSize: 64}
	DeepslateBrickStairs            = Item{ID: 565, DisplayName: "Deepslate Brick Stairs", Name: "deepslate_brick_stairs", StackSize: 64}
	DeepslateTileStairs             = Item{ID: 566, DisplayName: "Deepslate Tile Stairs", Name: "deepslate_tile_stairs", StackSize: 64}
	PolishedGraniteSlab             = Item{ID: 567, DisplayName: "Polished Granite Slab", Name: "polished_granite_slab", StackSize: 64}
	SmoothRedSandstoneSlab          = Item{ID: 568, DisplayName: "Smooth Red Sandstone Slab", Name: "smooth_red_sandstone_slab", StackSize: 64}
	MossyStoneBrickSlab             = Item{ID: 569, DisplayName: "Mossy Stone Brick Slab", Name: "mossy_stone_brick_slab", StackSize: 64}
	PolishedDioriteSlab             = Item{ID: 570, DisplayName: "Polished Diorite Slab", Name: "polished_diorite_slab", StackSize: 64}
	MossyCobblestoneSlab            = Item{ID: 571, DisplayName: "Mossy Cobblestone Slab", Name: "mossy_cobblestone_slab", StackSize: 64}
	EndStoneBrickSlab               = Item{ID: 572, DisplayName: "End Stone Brick Slab", Name: "end_stone_brick_slab", StackSize: 64}
	SmoothSandstoneSlab             = Item{ID: 573, DisplayName: "Smooth Sandstone Slab", Name: "smooth_sandstone_slab", StackSize: 64}
	SmoothQuartzSlab                = Item{ID: 574, DisplayName: "Smooth Quartz Slab", Name: "smooth_quartz_slab", StackSize: 64}
	GraniteSlab                     = Item{ID: 575, DisplayName: "Granite Slab", Name: "granite_slab", StackSize: 64}
	AndesiteSlab                    = Item{ID: 576, DisplayName: "Andesite Slab", Name: "andesite_slab", StackSize: 64}
	RedNetherBrickSlab              = Item{ID: 577, DisplayName: "Red Nether Brick Slab", Name: "red_nether_brick_slab", StackSize: 64}
	PolishedAndesiteSlab            = Item{ID: 578, DisplayName: "Polished Andesite Slab", Name: "polished_andesite_slab", StackSize: 64}
	DioriteSlab                     = Item{ID: 579, DisplayName: "Diorite Slab", Name: "diorite_slab", StackSize: 64}
	CobbledDeepslateSlab            = Item{ID: 580, DisplayName: "Cobbled Deepslate Slab", Name: "cobbled_deepslate_slab", StackSize: 64}
	PolishedDeepslateSlab           = Item{ID: 581, DisplayName: "Polished Deepslate Slab", Name: "polished_deepslate_slab", StackSize: 64}
	DeepslateBrickSlab              = Item{ID: 582, DisplayName: "Deepslate Brick Slab", Name: "deepslate_brick_slab", StackSize: 64}
	DeepslateTileSlab               = Item{ID: 583, DisplayName: "Deepslate Tile Slab", Name: "deepslate_tile_slab", StackSize: 64}
	Scaffolding                     = Item{ID: 584, DisplayName: "Scaffolding", Name: "scaffolding", StackSize: 64}
	Redstone                        = Item{ID: 585, DisplayName: "Redstone Dust", Name: "redstone", StackSize: 64}
	RedstoneTorch                   = Item{ID: 586, DisplayName: "Redstone Torch", Name: "redstone_torch", StackSize: 64}
	RedstoneBlock                   = Item{ID: 587, DisplayName: "Block of Redstone", Name: "redstone_block", StackSize: 64}
	Repeater                        = Item{ID: 588, DisplayName: "Redstone Repeater", Name: "repeater", StackSize: 64}
	Comparator                      = Item{ID: 589, DisplayName: "Redstone Comparator", Name: "comparator", StackSize: 64}
	Piston                          = Item{ID: 590, DisplayName: "Piston", Name: "piston", StackSize: 64}
	StickyPiston                    = Item{ID: 591, DisplayName: "Sticky Piston", Name: "sticky_piston", StackSize: 64}
	SlimeBlock                      = Item{ID: 592, DisplayName: "Slime Block", Name: "slime_block", StackSize: 64}
	HoneyBlock                      = Item{ID: 593, DisplayName: "Honey Block", Name: "honey_block", StackSize: 64}
	Observer                        = Item{ID: 594, DisplayName: "Observer", Name: "observer", StackSize: 64}
	Hopper                          = Item{ID: 595, DisplayName: "Hopper", Name: "hopper", StackSize: 64}
	Dispenser                       = Item{ID: 596, DisplayName: "Dispenser", Name: "dispenser", StackSize: 64}
	Dropper                         = Item{ID: 597, DisplayName: "Dropper", Name: "dropper", StackSize: 64}
	Lectern                         = Item{ID: 598, DisplayName: "Lectern", Name: "lectern", StackSize: 64}
	Target                          = Item{ID: 599, DisplayName: "Target", Name: "target", StackSize: 64}
	Lever                           = Item{ID: 600, DisplayName: "Lever", Name: "lever", StackSize: 64}
	LightningRod                    = Item{ID: 601, DisplayName: "Lightning Rod", Name: "lightning_rod", StackSize: 64}
	DaylightDetector                = Item{ID: 602, DisplayName: "Daylight Detector", Name: "daylight_detector", StackSize: 64}
	SculkSensor                     = Item{ID: 603, DisplayName: "Sculk Sensor", Name: "sculk_sensor", StackSize: 64}
	TripwireHook                    = Item{ID: 604, DisplayName: "Tripwire Hook", Name: "tripwire_hook", StackSize: 64}
	TrappedChest                    = Item{ID: 605, DisplayName: "Trapped Chest", Name: "trapped_chest", StackSize: 64}
	Tnt                             = Item{ID: 606, DisplayName: "TNT", Name: "tnt", StackSize: 64}
	RedstoneLamp                    = Item{ID: 607, DisplayName: "Redstone Lamp", Name: "redstone_lamp", StackSize: 64}
	NoteBlock                       = Item{ID: 608, DisplayName: "Note Block", Name: "note_block", StackSize: 64}
	StoneButton                     = Item{ID: 609, DisplayName: "Stone Button", Name: "stone_button", StackSize: 64}
	PolishedBlackstoneButton        = Item{ID: 610, DisplayName: "Polished Blackstone Button", Name: "polished_blackstone_button", StackSize: 64}
	OakButton                       = Item{ID: 611, DisplayName: "Oak Button", Name: "oak_button", StackSize: 64}
	SpruceButton                    = Item{ID: 612, DisplayName: "Spruce Button", Name: "spruce_button", StackSize: 64}
	BirchButton                     = Item{ID: 613, DisplayName: "Birch Button", Name: "birch_button", StackSize: 64}
	JungleButton                    = Item{ID: 614, DisplayName: "Jungle Button", Name: "jungle_button", StackSize: 64}
	AcaciaButton                    = Item{ID: 615, DisplayName: "Acacia Button", Name: "acacia_button", StackSize: 64}
	DarkOakButton                   = Item{ID: 616, DisplayName: "Dark Oak Button", Name: "dark_oak_button", StackSize: 64}
	CrimsonButton                   = Item{ID: 617, DisplayName: "Crimson Button", Name: "crimson_button", StackSize: 64}
	WarpedButton                    = Item{ID: 618, DisplayName: "Warped Button", Name: "warped_button", StackSize: 64}
	StonePressurePlate              = Item{ID: 619, DisplayName: "Stone Pressure Plate", Name: "stone_pressure_plate", StackSize: 64}
	PolishedBlackstonePressurePlate = Item{ID: 620, DisplayName: "Polished Blackstone Pressure Plate", Name: "polished_blackstone_pressure_plate", StackSize: 64}
	LightWeightedPressurePlate      = Item{ID: 621, DisplayName: "Light Weighted Pressure Plate", Name: "light_weighted_pressure_plate", StackSize: 64}
	HeavyWeightedPressurePlate      = Item{ID: 622, DisplayName: "Heavy Weighted Pressure Plate", Name: "heavy_weighted_pressure_plate", StackSize: 64}
	OakPressurePlate                = Item{ID: 623, DisplayName: "Oak Pressure Plate", Name: "oak_pressure_plate", StackSize: 64}
	SprucePressurePlate             = Item{ID: 624, DisplayName: "Spruce Pressure Plate", Name: "spruce_pressure_plate", StackSize: 64}
	BirchPressurePlate              = Item{ID: 625, DisplayName: "Birch Pressure Plate", Name: "birch_pressure_plate", StackSize: 64}
	JunglePressurePlate             = Item{ID: 626, DisplayName: "Jungle Pressure Plate", Name: "jungle_pressure_plate", StackSize: 64}
	AcaciaPressurePlate             = Item{ID: 627, DisplayName: "Acacia Pressure Plate", Name: "acacia_pressure_plate", StackSize: 64}
	DarkOakPressurePlate            = Item{ID: 628, DisplayName: "Dark Oak Pressure Plate", Name: "dark_oak_pressure_plate", StackSize: 64}
	CrimsonPressurePlate            = Item{ID: 629, DisplayName: "Crimson Pressure Plate", Name: "crimson_pressure_plate", StackSize: 64}
	WarpedPressurePlate             = Item{ID: 630, DisplayName: "Warped Pressure Plate", Name: "warped_pressure_plate", StackSize: 64}
	IronDoor                        = Item{ID: 631, DisplayName: "Iron Door", Name: "iron_door", StackSize: 64}
	OakDoor                         = Item{ID: 632, DisplayName: "Oak Door", Name: "oak_door", StackSize: 64}
	SpruceDoor                      = Item{ID: 633, DisplayName: "Spruce Door", Name: "spruce_door", StackSize: 64}
	BirchDoor                       = Item{ID: 634, DisplayName: "Birch Door", Name: "birch_door", StackSize: 64}
	JungleDoor                      = Item{ID: 635, DisplayName: "Jungle Door", Name: "jungle_door", StackSize: 64}
	AcaciaDoor                      = Item{ID: 636, DisplayName: "Acacia Door", Name: "acacia_door", StackSize: 64}
	DarkOakDoor                     = Item{ID: 637, DisplayName: "Dark Oak Door", Name: "dark_oak_door", StackSize: 64}
	CrimsonDoor                     = Item{ID: 638, DisplayName: "Crimson Door", Name: "crimson_door", StackSize: 64}
	WarpedDoor                      = Item{ID: 639, DisplayName: "Warped Door", Name: "warped_door", StackSize: 64}
	IronTrapdoor                    = Item{ID: 640, DisplayName: "Iron Trapdoor", Name: "iron_trapdoor", StackSize: 64}
	OakTrapdoor                     = Item{ID: 641, DisplayName: "Oak Trapdoor", Name: "oak_trapdoor", StackSize: 64}
	SpruceTrapdoor                  = Item{ID: 642, DisplayName: "Spruce Trapdoor", Name: "spruce_trapdoor", StackSize: 64}
	BirchTrapdoor                   = Item{ID: 643, DisplayName: "Birch Trapdoor", Name: "birch_trapdoor", StackSize: 64}
	JungleTrapdoor                  = Item{ID: 644, DisplayName: "Jungle Trapdoor", Name: "jungle_trapdoor", StackSize: 64}
	AcaciaTrapdoor                  = Item{ID: 645, DisplayName: "Acacia Trapdoor", Name: "acacia_trapdoor", StackSize: 64}
	DarkOakTrapdoor                 = Item{ID: 646, DisplayName: "Dark Oak Trapdoor", Name: "dark_oak_trapdoor", StackSize: 64}
	CrimsonTrapdoor                 = Item{ID: 647, DisplayName: "Crimson Trapdoor", Name: "crimson_trapdoor", StackSize: 64}
	WarpedTrapdoor                  = Item{ID: 648, DisplayName: "Warped Trapdoor", Name: "warped_trapdoor", StackSize: 64}
	OakFenceGate                    = Item{ID: 649, DisplayName: "Oak Fence Gate", Name: "oak_fence_gate", StackSize: 64}
	SpruceFenceGate                 = Item{ID: 650, DisplayName: "Spruce Fence Gate", Name: "spruce_fence_gate", StackSize: 64}
	BirchFenceGate                  = Item{ID: 651, DisplayName: "Birch Fence Gate", Name: "birch_fence_gate", StackSize: 64}
	JungleFenceGate                 = Item{ID: 652, DisplayName: "Jungle Fence Gate", Name: "jungle_fence_gate", StackSize: 64}
	AcaciaFenceGate                 = Item{ID: 653, DisplayName: "Acacia Fence Gate", Name: "acacia_fence_gate", StackSize: 64}
	DarkOakFenceGate                = Item{ID: 654, DisplayName: "Dark Oak Fence Gate", Name: "dark_oak_fence_gate", StackSize: 64}
	CrimsonFenceGate                = Item{ID: 655, DisplayName: "Crimson Fence Gate", Name: "crimson_fence_gate", StackSize: 64}
	WarpedFenceGate                 = Item{ID: 656, DisplayName: "Warped Fence Gate", Name: "warped_fence_gate", StackSize: 64}
	PoweredRail                     = Item{ID: 657, DisplayName: "Powered Rail", Name: "powered_rail", StackSize: 64}
	DetectorRail                    = Item{ID: 658, DisplayName: "Detector Rail", Name: "detector_rail", StackSize: 64}
	Rail                            = Item{ID: 659, DisplayName: "Rail", Name: "rail", StackSize: 64}
	ActivatorRail                   = Item{ID: 660, DisplayName: "Activator Rail", Name: "activator_rail", StackSize: 64}
	Saddle                          = Item{ID: 661, DisplayName: "Saddle", Name: "saddle", StackSize: 1}
	Minecart                        = Item{ID: 662, DisplayName: "Minecart", Name: "minecart", StackSize: 1}
	ChestMinecart                   = Item{ID: 663, DisplayName: "Minecart with Chest", Name: "chest_minecart", StackSize: 1}
	FurnaceMinecart                 = Item{ID: 664, DisplayName: "Minecart with Furnace", Name: "furnace_minecart", StackSize: 1}
	TntMinecart                     = Item{ID: 665, DisplayName: "Minecart with TNT", Name: "tnt_minecart", StackSize: 1}
	HopperMinecart                  = Item{ID: 666, DisplayName: "Minecart with Hopper", Name: "hopper_minecart", StackSize: 1}
	CarrotOnAStick                  = Item{ID: 667, DisplayName: "Carrot on a Stick", Name: "carrot_on_a_stick", StackSize: 1}
	WarpedFungusOnAStick            = Item{ID: 668, DisplayName: "Warped Fungus on a Stick", Name: "warped_fungus_on_a_stick", StackSize: 64}
	Elytra                          = Item{ID: 669, DisplayName: "Elytra", Name: "elytra", StackSize: 1}
	OakBoat                         = Item{ID: 670, DisplayName: "Oak Boat", Name: "oak_boat", StackSize: 1}
	SpruceBoat                      = Item{ID: 671, DisplayName: "Spruce Boat", Name: "spruce_boat", StackSize: 1}
	BirchBoat                       = Item{ID: 672, DisplayName: "Birch Boat", Name: "birch_boat", StackSize: 1}
	JungleBoat                      = Item{ID: 673, DisplayName: "Jungle Boat", Name: "jungle_boat", StackSize: 1}
	AcaciaBoat                      = Item{ID: 674, DisplayName: "Acacia Boat", Name: "acacia_boat", StackSize: 1}
	DarkOakBoat                     = Item{ID: 675, DisplayName: "Dark Oak Boat", Name: "dark_oak_boat", StackSize: 1}
	StructureBlock                  = Item{ID: 676, DisplayName: "Structure Block", Name: "structure_block", StackSize: 64}
	Jigsaw                          = Item{ID: 677, DisplayName: "Jigsaw Block", Name: "jigsaw", StackSize: 64}
	TurtleHelmet                    = Item{ID: 678, DisplayName: "Turtle Shell", Name: "turtle_helmet", StackSize: 1}
	Scute                           = Item{ID: 679, DisplayName: "Scute", Name: "scute", StackSize: 64}
	FlintAndSteel                   = Item{ID: 680, DisplayName: "Flint and Steel", Name: "flint_and_steel", StackSize: 1}
	Apple                           = Item{ID: 681, DisplayName: "Apple", Name: "apple", StackSize: 64}
	Bow                             = Item{ID: 682, DisplayName: "Bow", Name: "bow", StackSize: 1}
	Arrow                           = Item{ID: 683, DisplayName: "Arrow", Name: "arrow", StackSize: 64}
	Coal                            = Item{ID: 684, DisplayName: "Coal", Name: "coal", StackSize: 64}
	Charcoal                        = Item{ID: 685, DisplayName: "Charcoal", Name: "charcoal", StackSize: 64}
	Diamond                         = Item{ID: 686, DisplayName: "Diamond", Name: "diamond", StackSize: 64}
	Emerald                         = Item{ID: 687, DisplayName: "Emerald", Name: "emerald", StackSize: 64}
	LapisLazuli                     = Item{ID: 688, DisplayName: "Lapis Lazuli", Name: "lapis_lazuli", StackSize: 64}
	Quartz                          = Item{ID: 689, DisplayName: "Nether Quartz", Name: "quartz", StackSize: 64}
	AmethystShard                   = Item{ID: 690, DisplayName: "Amethyst Shard", Name: "amethyst_shard", StackSize: 64}
	RawIron                         = Item{ID: 691, DisplayName: "Raw Iron", Name: "raw_iron", StackSize: 64}
	IronIngot                       = Item{ID: 692, DisplayName: "Iron Ingot", Name: "iron_ingot", StackSize: 64}
	RawCopper                       = Item{ID: 693, DisplayName: "Raw Copper", Name: "raw_copper", StackSize: 64}
	CopperIngot                     = Item{ID: 694, DisplayName: "Copper Ingot", Name: "copper_ingot", StackSize: 64}
	RawGold                         = Item{ID: 695, DisplayName: "Raw Gold", Name: "raw_gold", StackSize: 64}
	GoldIngot                       = Item{ID: 696, DisplayName: "Gold Ingot", Name: "gold_ingot", StackSize: 64}
	NetheriteIngot                  = Item{ID: 697, DisplayName: "Netherite Ingot", Name: "netherite_ingot", StackSize: 64}
	NetheriteScrap                  = Item{ID: 698, DisplayName: "Netherite Scrap", Name: "netherite_scrap", StackSize: 64}
	WoodenSword                     = Item{ID: 699, DisplayName: "Wooden Sword", Name: "wooden_sword", StackSize: 1}
	WoodenShovel                    = Item{ID: 700, DisplayName: "Wooden Shovel", Name: "wooden_shovel", StackSize: 1}
	WoodenPickaxe                   = Item{ID: 701, DisplayName: "Wooden Pickaxe", Name: "wooden_pickaxe", StackSize: 1}
	WoodenAxe                       = Item{ID: 702, DisplayName: "Wooden Axe", Name: "wooden_axe", StackSize: 1}
	WoodenHoe                       = Item{ID: 703, DisplayName: "Wooden Hoe", Name: "wooden_hoe", StackSize: 1}
	StoneSword                      = Item{ID: 704, DisplayName: "Stone Sword", Name: "stone_sword", StackSize: 1}
	StoneShovel                     = Item{ID: 705, DisplayName: "Stone Shovel", Name: "stone_shovel", StackSize: 1}
	StonePickaxe                    = Item{ID: 706, DisplayName: "Stone Pickaxe", Name: "stone_pickaxe", StackSize: 1}
	StoneAxe                        = Item{ID: 707, DisplayName: "Stone Axe", Name: "stone_axe", StackSize: 1}
	StoneHoe                        = Item{ID: 708, DisplayName: "Stone Hoe", Name: "stone_hoe", StackSize: 1}
	GoldenSword                     = Item{ID: 709, DisplayName: "Golden Sword", Name: "golden_sword", StackSize: 1}
	GoldenShovel                    = Item{ID: 710, DisplayName: "Golden Shovel", Name: "golden_shovel", StackSize: 1}
	GoldenPickaxe                   = Item{ID: 711, DisplayName: "Golden Pickaxe", Name: "golden_pickaxe", StackSize: 1}
	GoldenAxe                       = Item{ID: 712, DisplayName: "Golden Axe", Name: "golden_axe", StackSize: 1}
	GoldenHoe                       = Item{ID: 713, DisplayName: "Golden Hoe", Name: "golden_hoe", StackSize: 1}
	IronSword                       = Item{ID: 714, DisplayName: "Iron Sword", Name: "iron_sword", StackSize: 1}
	IronShovel                      = Item{ID: 715, DisplayName: "Iron Shovel", Name: "iron_shovel", StackSize: 1}
	IronPickaxe                     = Item{ID: 716, DisplayName: "Iron Pickaxe", Name: "iron_pickaxe", StackSize: 1}
	IronAxe                         = Item{ID: 717, DisplayName: "Iron Axe", Name: "iron_axe", StackSize: 1}
	IronHoe                         = Item{ID: 718, DisplayName: "Iron Hoe", Name: "iron_hoe", StackSize: 1}
	DiamondSword                    = Item{ID: 719, DisplayName: "Diamond Sword", Name: "diamond_sword", StackSize: 1}
	DiamondShovel                   = Item{ID: 720, DisplayName: "Diamond Shovel", Name: "diamond_shovel", StackSize: 1}
	DiamondPickaxe                  = Item{ID: 721, DisplayName: "Diamond Pickaxe", Name: "diamond_pickaxe", StackSize: 1}
	DiamondAxe                      = Item{ID: 722, DisplayName: "Diamond Axe", Name: "diamond_axe", StackSize: 1}
	DiamondHoe                      = Item{ID: 723, DisplayName: "Diamond Hoe", Name: "diamond_hoe", StackSize: 1}
	NetheriteSword                  = Item{ID: 724, DisplayName: "Netherite Sword", Name: "netherite_sword", StackSize: 1}
	NetheriteShovel                 = Item{ID: 725, DisplayName: "Netherite Shovel", Name: "netherite_shovel", StackSize: 1}
	NetheritePickaxe                = Item{ID: 726, DisplayName: "Netherite Pickaxe", Name: "netherite_pickaxe", StackSize: 1}
	NetheriteAxe                    = Item{ID: 727, DisplayName: "Netherite Axe", Name: "netherite_axe", StackSize: 1}
	NetheriteHoe                    = Item{ID: 728, DisplayName: "Netherite Hoe", Name: "netherite_hoe", StackSize: 1}
	Stick                           = Item{ID: 729, DisplayName: "Stick", Name: "stick", StackSize: 64}
	Bowl                            = Item{ID: 730, DisplayName: "Bowl", Name: "bowl", StackSize: 64}
	MushroomStew                    = Item{ID: 731, DisplayName: "Mushroom Stew", Name: "mushroom_stew", StackSize: 1}
	String                          = Item{ID: 732, DisplayName: "String", Name: "string", StackSize: 64}
	Feather                         = Item{ID: 733, DisplayName: "Feather", Name: "feather", StackSize: 64}
	Gunpowder                       = Item{ID: 734, DisplayName: "Gunpowder", Name: "gunpowder", StackSize: 64}
	WheatSeeds                      = Item{ID: 735, DisplayName: "Wheat Seeds", Name: "wheat_seeds", StackSize: 64}
	Wheat                           = Item{ID: 736, DisplayName: "Wheat", Name: "wheat", StackSize: 64}
	Bread                           = Item{ID: 737, DisplayName: "Bread", Name: "bread", StackSize: 64}
	LeatherHelmet                   = Item{ID: 738, DisplayName: "Leather Cap", Name: "leather_helmet", StackSize: 1}
	LeatherChestplate               = Item{ID: 739, DisplayName: "Leather Tunic", Name: "leather_chestplate", StackSize: 1}
	LeatherLeggings                 = Item{ID: 740, DisplayName: "Leather Pants", Name: "leather_leggings", StackSize: 1}
	LeatherBoots                    = Item{ID: 741, DisplayName: "Leather Boots", Name: "leather_boots", StackSize: 1}
	ChainmailHelmet                 = Item{ID: 742, DisplayName: "Chainmail Helmet", Name: "chainmail_helmet", StackSize: 1}
	ChainmailChestplate             = Item{ID: 743, DisplayName: "Chainmail Chestplate", Name: "chainmail_chestplate", StackSize: 1}
	ChainmailLeggings               = Item{ID: 744, DisplayName: "Chainmail Leggings", Name: "chainmail_leggings", StackSize: 1}
	ChainmailBoots                  = Item{ID: 745, DisplayName: "Chainmail Boots", Name: "chainmail_boots", StackSize: 1}
	IronHelmet                      = Item{ID: 746, DisplayName: "Iron Helmet", Name: "iron_helmet", StackSize: 1}
	IronChestplate                  = Item{ID: 747, DisplayName: "Iron Chestplate", Name: "iron_chestplate", StackSize: 1}
	IronLeggings                    = Item{ID: 748, DisplayName: "Iron Leggings", Name: "iron_leggings", StackSize: 1}
	IronBoots                       = Item{ID: 749, DisplayName: "Iron Boots", Name: "iron_boots", StackSize: 1}
	DiamondHelmet                   = Item{ID: 750, DisplayName: "Diamond Helmet", Name: "diamond_helmet", StackSize: 1}
	DiamondChestplate               = Item{ID: 751, DisplayName: "Diamond Chestplate", Name: "diamond_chestplate", StackSize: 1}
	DiamondLeggings                 = Item{ID: 752, DisplayName: "Diamond Leggings", Name: "diamond_leggings", StackSize: 1}
	DiamondBoots                    = Item{ID: 753, DisplayName: "Diamond Boots", Name: "diamond_boots", StackSize: 1}
	GoldenHelmet                    = Item{ID: 754, DisplayName: "Golden Helmet", Name: "golden_helmet", StackSize: 1}
	GoldenChestplate                = Item{ID: 755, DisplayName: "Golden Chestplate", Name: "golden_chestplate", StackSize: 1}
	GoldenLeggings                  = Item{ID: 756, DisplayName: "Golden Leggings", Name: "golden_leggings", StackSize: 1}
	GoldenBoots                     = Item{ID: 757, DisplayName: "Golden Boots", Name: "golden_boots", StackSize: 1}
	NetheriteHelmet                 = Item{ID: 758, DisplayName: "Netherite Helmet", Name: "netherite_helmet", StackSize: 1}
	NetheriteChestplate             = Item{ID: 759, DisplayName: "Netherite Chestplate", Name: "netherite_chestplate", StackSize: 1}
	NetheriteLeggings               = Item{ID: 760, DisplayName: "Netherite Leggings", Name: "netherite_leggings", StackSize: 1}
	NetheriteBoots                  = Item{ID: 761, DisplayName: "Netherite Boots", Name: "netherite_boots", StackSize: 1}
	Flint                           = Item{ID: 762, DisplayName: "Flint", Name: "flint", StackSize: 64}
	Porkchop                        = Item{ID: 763, DisplayName: "Raw Porkchop", Name: "porkchop", StackSize: 64}
	CookedPorkchop                  = Item{ID: 764, DisplayName: "Cooked Porkchop", Name: "cooked_porkchop", StackSize: 64}
	Painting                        = Item{ID: 765, DisplayName: "Painting", Name: "painting", StackSize: 64}
	GoldenApple                     = Item{ID: 766, DisplayName: "Golden Apple", Name: "golden_apple", StackSize: 64}
	EnchantedGoldenApple            = Item{ID: 767, DisplayName: "Enchanted Golden Apple", Name: "enchanted_golden_apple", StackSize: 64}
	OakSign                         = Item{ID: 768, DisplayName: "Oak Sign", Name: "oak_sign", StackSize: 16}
	SpruceSign                      = Item{ID: 769, DisplayName: "Spruce Sign", Name: "spruce_sign", StackSize: 16}
	BirchSign                       = Item{ID: 770, DisplayName: "Birch Sign", Name: "birch_sign", StackSize: 16}
	JungleSign                      = Item{ID: 771, DisplayName: "Jungle Sign", Name: "jungle_sign", StackSize: 16}
	AcaciaSign                      = Item{ID: 772, DisplayName: "Acacia Sign", Name: "acacia_sign", StackSize: 16}
	DarkOakSign                     = Item{ID: 773, DisplayName: "Dark Oak Sign", Name: "dark_oak_sign", StackSize: 16}
	CrimsonSign                     = Item{ID: 774, DisplayName: "Crimson Sign", Name: "crimson_sign", StackSize: 16}
	WarpedSign                      = Item{ID: 775, DisplayName: "Warped Sign", Name: "warped_sign", StackSize: 16}
	Bucket                          = Item{ID: 776, DisplayName: "Bucket", Name: "bucket", StackSize: 16}
	WaterBucket                     = Item{ID: 777, DisplayName: "Water Bucket", Name: "water_bucket", StackSize: 1}
	LavaBucket                      = Item{ID: 778, DisplayName: "Lava Bucket", Name: "lava_bucket", StackSize: 1}
	PowderSnowBucket                = Item{ID: 779, DisplayName: "Powder Snow Bucket", Name: "powder_snow_bucket", StackSize: 1}
	Snowball                        = Item{ID: 780, DisplayName: "Snowball", Name: "snowball", StackSize: 16}
	Leather                         = Item{ID: 781, DisplayName: "Leather", Name: "leather", StackSize: 64}
	MilkBucket                      = Item{ID: 782, DisplayName: "Milk Bucket", Name: "milk_bucket", StackSize: 1}
	PufferfishBucket                = Item{ID: 783, DisplayName: "Bucket of Pufferfish", Name: "pufferfish_bucket", StackSize: 1}
	SalmonBucket                    = Item{ID: 784, DisplayName: "Bucket of Salmon", Name: "salmon_bucket", StackSize: 1}
	CodBucket                       = Item{ID: 785, DisplayName: "Bucket of Cod", Name: "cod_bucket", StackSize: 1}
	TropicalFishBucket              = Item{ID: 786, DisplayName: "Bucket of Tropical Fish", Name: "tropical_fish_bucket", StackSize: 1}
	AxolotlBucket                   = Item{ID: 787, DisplayName: "Bucket of Axolotl", Name: "axolotl_bucket", StackSize: 1}
	Brick                           = Item{ID: 788, DisplayName: "Brick", Name: "brick", StackSize: 64}
	ClayBall                        = Item{ID: 789, DisplayName: "Clay Ball", Name: "clay_ball", StackSize: 64}
	DriedKelpBlock                  = Item{ID: 790, DisplayName: "Dried Kelp Block", Name: "dried_kelp_block", StackSize: 64}
	Paper                           = Item{ID: 791, DisplayName: "Paper", Name: "paper", StackSize: 64}
	Book                            = Item{ID: 792, DisplayName: "Book", Name: "book", StackSize: 64}
	SlimeBall                       = Item{ID: 793, DisplayName: "Slimeball", Name: "slime_ball", StackSize: 64}
	Egg                             = Item{ID: 794, DisplayName: "Egg", Name: "egg", StackSize: 16}
	Compass                         = Item{ID: 795, DisplayName: "Compass", Name: "compass", StackSize: 64}
	Bundle                          = Item{ID: 796, DisplayName: "Bundle", Name: "bundle", StackSize: 1}
	FishingRod                      = Item{ID: 797, DisplayName: "Fishing Rod", Name: "fishing_rod", StackSize: 1}
	Clock                           = Item{ID: 798, DisplayName: "Clock", Name: "clock", StackSize: 64}
	Spyglass                        = Item{ID: 799, DisplayName: "Spyglass", Name: "spyglass", StackSize: 1}
	GlowstoneDust                   = Item{ID: 800, DisplayName: "Glowstone Dust", Name: "glowstone_dust", StackSize: 64}
	Cod                             = Item{ID: 801, DisplayName: "Raw Cod", Name: "cod", StackSize: 64}
	Salmon                          = Item{ID: 802, DisplayName: "Raw Salmon", Name: "salmon", StackSize: 64}
	TropicalFish                    = Item{ID: 803, DisplayName: "Tropical Fish", Name: "tropical_fish", StackSize: 64}
	Pufferfish                      = Item{ID: 804, DisplayName: "Pufferfish", Name: "pufferfish", StackSize: 64}
	CookedCod                       = Item{ID: 805, DisplayName: "Cooked Cod", Name: "cooked_cod", StackSize: 64}
	CookedSalmon                    = Item{ID: 806, DisplayName: "Cooked Salmon", Name: "cooked_salmon", StackSize: 64}
	InkSac                          = Item{ID: 807, DisplayName: "Ink Sac", Name: "ink_sac", StackSize: 64}
	GlowInkSac                      = Item{ID: 808, DisplayName: "Glow Ink Sac", Name: "glow_ink_sac", StackSize: 64}
	CocoaBeans                      = Item{ID: 809, DisplayName: "Cocoa Beans", Name: "cocoa_beans", StackSize: 64}
	WhiteDye                        = Item{ID: 810, DisplayName: "White Dye", Name: "white_dye", StackSize: 64}
	OrangeDye                       = Item{ID: 811, DisplayName: "Orange Dye", Name: "orange_dye", StackSize: 64}
	MagentaDye                      = Item{ID: 812, DisplayName: "Magenta Dye", Name: "magenta_dye", StackSize: 64}
	LightBlueDye                    = Item{ID: 813, DisplayName: "Light Blue Dye", Name: "light_blue_dye", StackSize: 64}
	YellowDye                       = Item{ID: 814, DisplayName: "Yellow Dye", Name: "yellow_dye", StackSize: 64}
	LimeDye                         = Item{ID: 815, DisplayName: "Lime Dye", Name: "lime_dye", StackSize: 64}
	PinkDye                         = Item{ID: 816, DisplayName: "Pink Dye", Name: "pink_dye", StackSize: 64}
	GrayDye                         = Item{ID: 817, DisplayName: "Gray Dye", Name: "gray_dye", StackSize: 64}
	LightGrayDye                    = Item{ID: 818, DisplayName: "Light Gray Dye", Name: "light_gray_dye", StackSize: 64}
	CyanDye                         = Item{ID: 819, DisplayName: "Cyan Dye", Name: "cyan_dye", StackSize: 64}
	PurpleDye                       = Item{ID: 820, DisplayName: "Purple Dye", Name: "purple_dye", StackSize: 64}
	BlueDye                         = Item{ID: 821, DisplayName: "Blue Dye", Name: "blue_dye", StackSize: 64}
	BrownDye                        = Item{ID: 822, DisplayName: "Brown Dye", Name: "brown_dye", StackSize: 64}
	GreenDye                        = Item{ID: 823, DisplayName: "Green Dye", Name: "green_dye", StackSize: 64}
	RedDye                          = Item{ID: 824, DisplayName: "Red Dye", Name: "red_dye", StackSize: 64}
	BlackDye                        = Item{ID: 825, DisplayName: "Black Dye", Name: "black_dye", StackSize: 64}
	BoneMeal                        = Item{ID: 826, DisplayName: "Bone Meal", Name: "bone_meal", StackSize: 64}
	Bone                            = Item{ID: 827, DisplayName: "Bone", Name: "bone", StackSize: 64}
	Sugar                           = Item{ID: 828, DisplayName: "Sugar", Name: "sugar", StackSize: 64}
	Cake                            = Item{ID: 829, DisplayName: "Cake", Name: "cake", StackSize: 1}
	WhiteBed                        = Item{ID: 830, DisplayName: "White Bed", Name: "white_bed", StackSize: 1}
	OrangeBed                       = Item{ID: 831, DisplayName: "Orange Bed", Name: "orange_bed", StackSize: 1}
	MagentaBed                      = Item{ID: 832, DisplayName: "Magenta Bed", Name: "magenta_bed", StackSize: 1}
	LightBlueBed                    = Item{ID: 833, DisplayName: "Light Blue Bed", Name: "light_blue_bed", StackSize: 1}
	YellowBed                       = Item{ID: 834, DisplayName: "Yellow Bed", Name: "yellow_bed", StackSize: 1}
	LimeBed                         = Item{ID: 835, DisplayName: "Lime Bed", Name: "lime_bed", StackSize: 1}
	PinkBed                         = Item{ID: 836, DisplayName: "Pink Bed", Name: "pink_bed", StackSize: 1}
	GrayBed                         = Item{ID: 837, DisplayName: "Gray Bed", Name: "gray_bed", StackSize: 1}
	LightGrayBed                    = Item{ID: 838, DisplayName: "Light Gray Bed", Name: "light_gray_bed", StackSize: 1}
	CyanBed                         = Item{ID: 839, DisplayName: "Cyan Bed", Name: "cyan_bed", StackSize: 1}
	PurpleBed                       = Item{ID: 840, DisplayName: "Purple Bed", Name: "purple_bed", StackSize: 1}
	BlueBed                         = Item{ID: 841, DisplayName: "Blue Bed", Name: "blue_bed", StackSize: 1}
	BrownBed                        = Item{ID: 842, DisplayName: "Brown Bed", Name: "brown_bed", StackSize: 1}
	GreenBed                        = Item{ID: 843, DisplayName: "Green Bed", Name: "green_bed", StackSize: 1}
	RedBed                          = Item{ID: 844, DisplayName: "Red Bed", Name: "red_bed", StackSize: 1}
	BlackBed                        = Item{ID: 845, DisplayName: "Black Bed", Name: "black_bed", StackSize: 1}
	Cookie                          = Item{ID: 846, DisplayName: "Cookie", Name: "cookie", StackSize: 64}
	FilledMap                       = Item{ID: 847, DisplayName: "Map", Name: "filled_map", StackSize: 64}
	Shears                          = Item{ID: 848, DisplayName: "Shears", Name: "shears", StackSize: 1}
	MelonSlice                      = Item{ID: 849, DisplayName: "Melon Slice", Name: "melon_slice", StackSize: 64}
	DriedKelp                       = Item{ID: 850, DisplayName: "Dried Kelp", Name: "dried_kelp", StackSize: 64}
	PumpkinSeeds                    = Item{ID: 851, DisplayName: "Pumpkin Seeds", Name: "pumpkin_seeds", StackSize: 64}
	MelonSeeds                      = Item{ID: 852, DisplayName: "Melon Seeds", Name: "melon_seeds", StackSize: 64}
	Beef                            = Item{ID: 853, DisplayName: "Raw Beef", Name: "beef", StackSize: 64}
	CookedBeef                      = Item{ID: 854, DisplayName: "Steak", Name: "cooked_beef", StackSize: 64}
	Chicken                         = Item{ID: 855, DisplayName: "Raw Chicken", Name: "chicken", StackSize: 64}
	CookedChicken                   = Item{ID: 856, DisplayName: "Cooked Chicken", Name: "cooked_chicken", StackSize: 64}
	RottenFlesh                     = Item{ID: 857, DisplayName: "Rotten Flesh", Name: "rotten_flesh", StackSize: 64}
	EnderPearl                      = Item{ID: 858, DisplayName: "Ender Pearl", Name: "ender_pearl", StackSize: 16}
	BlazeRod                        = Item{ID: 859, DisplayName: "Blaze Rod", Name: "blaze_rod", StackSize: 64}
	GhastTear                       = Item{ID: 860, DisplayName: "Ghast Tear", Name: "ghast_tear", StackSize: 64}
	GoldNugget                      = Item{ID: 861, DisplayName: "Gold Nugget", Name: "gold_nugget", StackSize: 64}
	NetherWart                      = Item{ID: 862, DisplayName: "Nether Wart", Name: "nether_wart", StackSize: 64}
	Potion                          = Item{ID: 863, DisplayName: "Potion", Name: "potion", StackSize: 1}
	GlassBottle                     = Item{ID: 864, DisplayName: "Glass Bottle", Name: "glass_bottle", StackSize: 64}
	SpiderEye                       = Item{ID: 865, DisplayName: "Spider Eye", Name: "spider_eye", StackSize: 64}
	FermentedSpiderEye              = Item{ID: 866, DisplayName: "Fermented Spider Eye", Name: "fermented_spider_eye", StackSize: 64}
	BlazePowder                     = Item{ID: 867, DisplayName: "Blaze Powder", Name: "blaze_powder", StackSize: 64}
	MagmaCream                      = Item{ID: 868, DisplayName: "Magma Cream", Name: "magma_cream", StackSize: 64}
	BrewingStand                    = Item{ID: 869, DisplayName: "Brewing Stand", Name: "brewing_stand", StackSize: 64}
	Cauldron                        = Item{ID: 870, DisplayName: "Cauldron", Name: "cauldron", StackSize: 64}
	EnderEye                        = Item{ID: 871, DisplayName: "Eye of Ender", Name: "ender_eye", StackSize: 64}
	GlisteringMelonSlice            = Item{ID: 872, DisplayName: "Glistering Melon Slice", Name: "glistering_melon_slice", StackSize: 64}
	AxolotlSpawnEgg                 = Item{ID: 873, DisplayName: "Axolotl Spawn Egg", Name: "axolotl_spawn_egg", StackSize: 64}
	BatSpawnEgg                     = Item{ID: 874, DisplayName: "Bat Spawn Egg", Name: "bat_spawn_egg", StackSize: 64}
	BeeSpawnEgg                     = Item{ID: 875, DisplayName: "Bee Spawn Egg", Name: "bee_spawn_egg", StackSize: 64}
	BlazeSpawnEgg                   = Item{ID: 876, DisplayName: "Blaze Spawn Egg", Name: "blaze_spawn_egg", StackSize: 64}
	CatSpawnEgg                     = Item{ID: 877, DisplayName: "Cat Spawn Egg", Name: "cat_spawn_egg", StackSize: 64}
	CaveSpiderSpawnEgg              = Item{ID: 878, DisplayName: "Cave Spider Spawn Egg", Name: "cave_spider_spawn_egg", StackSize: 64}
	ChickenSpawnEgg                 = Item{ID: 879, DisplayName: "Chicken Spawn Egg", Name: "chicken_spawn_egg", StackSize: 64}
	CodSpawnEgg                     = Item{ID: 880, DisplayName: "Cod Spawn Egg", Name: "cod_spawn_egg", StackSize: 64}
	CowSpawnEgg                     = Item{ID: 881, DisplayName: "Cow Spawn Egg", Name: "cow_spawn_egg", StackSize: 64}
	CreeperSpawnEgg                 = Item{ID: 882, DisplayName: "Creeper Spawn Egg", Name: "creeper_spawn_egg", StackSize: 64}
	DolphinSpawnEgg                 = Item{ID: 883, DisplayName: "Dolphin Spawn Egg", Name: "dolphin_spawn_egg", StackSize: 64}
	DonkeySpawnEgg                  = Item{ID: 884, DisplayName: "Donkey Spawn Egg", Name: "donkey_spawn_egg", StackSize: 64}
	DrownedSpawnEgg                 = Item{ID: 885, DisplayName: "Drowned Spawn Egg", Name: "drowned_spawn_egg", StackSize: 64}
	ElderGuardianSpawnEgg           = Item{ID: 886, DisplayName: "Elder Guardian Spawn Egg", Name: "elder_guardian_spawn_egg", StackSize: 64}
	EndermanSpawnEgg                = Item{ID: 887, DisplayName: "Enderman Spawn Egg", Name: "enderman_spawn_egg", StackSize: 64}
	EndermiteSpawnEgg               = Item{ID: 888, DisplayName: "Endermite Spawn Egg", Name: "endermite_spawn_egg", StackSize: 64}
	EvokerSpawnEgg                  = Item{ID: 889, DisplayName: "Evoker Spawn Egg", Name: "evoker_spawn_egg", StackSize: 64}
	FoxSpawnEgg                     = Item{ID: 890, DisplayName: "Fox Spawn Egg", Name: "fox_spawn_egg", StackSize: 64}
	GhastSpawnEgg                   = Item{ID: 891, DisplayName: "Ghast Spawn Egg", Name: "ghast_spawn_egg", StackSize: 64}
	GlowSquidSpawnEgg               = Item{ID: 892, DisplayName: "Glow Squid Spawn Egg", Name: "glow_squid_spawn_egg", StackSize: 64}
	GoatSpawnEgg                    = Item{ID: 893, DisplayName: "Goat Spawn Egg", Name: "goat_spawn_egg", StackSize: 64}
	GuardianSpawnEgg                = Item{ID: 894, DisplayName: "Guardian Spawn Egg", Name: "guardian_spawn_egg", StackSize: 64}
	HoglinSpawnEgg                  = Item{ID: 895, DisplayName: "Hoglin Spawn Egg", Name: "hoglin_spawn_egg", StackSize: 64}
	HorseSpawnEgg                   = Item{ID: 896, DisplayName: "Horse Spawn Egg", Name: "horse_spawn_egg", StackSize: 64}
	HuskSpawnEgg                    = Item{ID: 897, DisplayName: "Husk Spawn Egg", Name: "husk_spawn_egg", StackSize: 64}
	LlamaSpawnEgg                   = Item{ID: 898, DisplayName: "Llama Spawn Egg", Name: "llama_spawn_egg", StackSize: 64}
	MagmaCubeSpawnEgg               = Item{ID: 899, DisplayName: "Magma Cube Spawn Egg", Name: "magma_cube_spawn_egg", StackSize: 64}
	MooshroomSpawnEgg               = Item{ID: 900, DisplayName: "Mooshroom Spawn Egg", Name: "mooshroom_spawn_egg", StackSize: 64}
	MuleSpawnEgg                    = Item{ID: 901, DisplayName: "Mule Spawn Egg", Name: "mule_spawn_egg", StackSize: 64}
	OcelotSpawnEgg                  = Item{ID: 902, DisplayName: "Ocelot Spawn Egg", Name: "ocelot_spawn_egg", StackSize: 64}
	PandaSpawnEgg                   = Item{ID: 903, DisplayName: "Panda Spawn Egg", Name: "panda_spawn_egg", StackSize: 64}
	ParrotSpawnEgg                  = Item{ID: 904, DisplayName: "Parrot Spawn Egg", Name: "parrot_spawn_egg", StackSize: 64}
	PhantomSpawnEgg                 = Item{ID: 905, DisplayName: "Phantom Spawn Egg", Name: "phantom_spawn_egg", StackSize: 64}
	PigSpawnEgg                     = Item{ID: 906, DisplayName: "Pig Spawn Egg", Name: "pig_spawn_egg", StackSize: 64}
	PiglinSpawnEgg                  = Item{ID: 907, DisplayName: "Piglin Spawn Egg", Name: "piglin_spawn_egg", StackSize: 64}
	PiglinBruteSpawnEgg             = Item{ID: 908, DisplayName: "Piglin Brute Spawn Egg", Name: "piglin_brute_spawn_egg", StackSize: 64}
	PillagerSpawnEgg                = Item{ID: 909, DisplayName: "Pillager Spawn Egg", Name: "pillager_spawn_egg", StackSize: 64}
	PolarBearSpawnEgg               = Item{ID: 910, DisplayName: "Polar Bear Spawn Egg", Name: "polar_bear_spawn_egg", StackSize: 64}
	PufferfishSpawnEgg              = Item{ID: 911, DisplayName: "Pufferfish Spawn Egg", Name: "pufferfish_spawn_egg", StackSize: 64}
	RabbitSpawnEgg                  = Item{ID: 912, DisplayName: "Rabbit Spawn Egg", Name: "rabbit_spawn_egg", StackSize: 64}
	RavagerSpawnEgg                 = Item{ID: 913, DisplayName: "Ravager Spawn Egg", Name: "ravager_spawn_egg", StackSize: 64}
	SalmonSpawnEgg                  = Item{ID: 914, DisplayName: "Salmon Spawn Egg", Name: "salmon_spawn_egg", StackSize: 64}
	SheepSpawnEgg                   = Item{ID: 915, DisplayName: "Sheep Spawn Egg", Name: "sheep_spawn_egg", StackSize: 64}
	ShulkerSpawnEgg                 = Item{ID: 916, DisplayName: "Shulker Spawn Egg", Name: "shulker_spawn_egg", StackSize: 64}
	SilverfishSpawnEgg              = Item{ID: 917, DisplayName: "Silverfish Spawn Egg", Name: "silverfish_spawn_egg", StackSize: 64}
	SkeletonSpawnEgg                = Item{ID: 918, DisplayName: "Skeleton Spawn Egg", Name: "skeleton_spawn_egg", StackSize: 64}
	SkeletonHorseSpawnEgg           = Item{ID: 919, DisplayName: "Skeleton Horse Spawn Egg", Name: "skeleton_horse_spawn_egg", StackSize: 64}
	SlimeSpawnEgg                   = Item{ID: 920, DisplayName: "Slime Spawn Egg", Name: "slime_spawn_egg", StackSize: 64}
	SpiderSpawnEgg                  = Item{ID: 921, DisplayName: "Spider Spawn Egg", Name: "spider_spawn_egg", StackSize: 64}
	SquidSpawnEgg                   = Item{ID: 922, DisplayName: "Squid Spawn Egg", Name: "squid_spawn_egg", StackSize: 64}
	StraySpawnEgg                   = Item{ID: 923, DisplayName: "Stray Spawn Egg", Name: "stray_spawn_egg", StackSize: 64}
	StriderSpawnEgg                 = Item{ID: 924, DisplayName: "Strider Spawn Egg", Name: "strider_spawn_egg", StackSize: 64}
	TraderLlamaSpawnEgg             = Item{ID: 925, DisplayName: "Trader Llama Spawn Egg", Name: "trader_llama_spawn_egg", StackSize: 64}
	TropicalFishSpawnEgg            = Item{ID: 926, DisplayName: "Tropical Fish Spawn Egg", Name: "tropical_fish_spawn_egg", StackSize: 64}
	TurtleSpawnEgg                  = Item{ID: 927, DisplayName: "Turtle Spawn Egg", Name: "turtle_spawn_egg", StackSize: 64}
	VexSpawnEgg                     = Item{ID: 928, DisplayName: "Vex Spawn Egg", Name: "vex_spawn_egg", StackSize: 64}
	VillagerSpawnEgg                = Item{ID: 929, DisplayName: "Villager Spawn Egg", Name: "villager_spawn_egg", StackSize: 64}
	VindicatorSpawnEgg              = Item{ID: 930, DisplayName: "Vindicator Spawn Egg", Name: "vindicator_spawn_egg", StackSize: 64}
	WanderingTraderSpawnEgg         = Item{ID: 931, DisplayName: "Wandering Trader Spawn Egg", Name: "wandering_trader_spawn_egg", StackSize: 64}
	WitchSpawnEgg                   = Item{ID: 932, DisplayName: "Witch Spawn Egg", Name: "witch_spawn_egg", StackSize: 64}
	WitherSkeletonSpawnEgg          = Item{ID: 933, DisplayName: "Wither Skeleton Spawn Egg", Name: "wither_skeleton_spawn_egg", StackSize: 64}
	WolfSpawnEgg                    = Item{ID: 934, DisplayName: "Wolf Spawn Egg", Name: "wolf_spawn_egg", StackSize: 64}
	ZoglinSpawnEgg                  = Item{ID: 935, DisplayName: "Zoglin Spawn Egg", Name: "zoglin_spawn_egg", StackSize: 64}
	ZombieSpawnEgg                  = Item{ID: 936, DisplayName: "Zombie Spawn Egg", Name: "zombie_spawn_egg", StackSize: 64}
	ZombieHorseSpawnEgg             = Item{ID: 937, DisplayName: "Zombie Horse Spawn Egg", Name: "zombie_horse_spawn_egg", StackSize: 64}
	ZombieVillagerSpawnEgg          = Item{ID: 938, DisplayName: "Zombie Villager Spawn Egg", Name: "zombie_villager_spawn_egg", StackSize: 64}
	ZombifiedPiglinSpawnEgg         = Item{ID: 939, DisplayName: "Zombified Piglin Spawn Egg", Name: "zombified_piglin_spawn_egg", StackSize: 64}
	ExperienceBottle                = Item{ID: 940, DisplayName: "Bottle o' Enchanting", Name: "experience_bottle", StackSize: 64}
	FireCharge                      = Item{ID: 941, DisplayName: "Fire Charge", Name: "fire_charge", StackSize: 64}
	WritableBook                    = Item{ID: 942, DisplayName: "Book and Quill", Name: "writable_book", StackSize: 1}
	WrittenBook                     = Item{ID: 943, DisplayName: "Written Book", Name: "written_book", StackSize: 16}
	ItemFrame                       = Item{ID: 944, DisplayName: "Item Frame", Name: "item_frame", StackSize: 64}
	GlowItemFrame                   = Item{ID: 945, DisplayName: "Glow Item Frame", Name: "glow_item_frame", StackSize: 64}
	FlowerPot                       = Item{ID: 946, DisplayName: "Flower Pot", Name: "flower_pot", StackSize: 64}
	Carrot                          = Item{ID: 947, DisplayName: "Carrot", Name: "carrot", StackSize: 64}
	Potato                          = Item{ID: 948, DisplayName: "Potato", Name: "potato", StackSize: 64}
	BakedPotato                     = Item{ID: 949, DisplayName: "Baked Potato", Name: "baked_potato", StackSize: 64}
	PoisonousPotato                 = Item{ID: 950, DisplayName: "Poisonous Potato", Name: "poisonous_potato", StackSize: 64}
	Map                             = Item{ID: 951, DisplayName: "Empty Map", Name: "map", StackSize: 64}
	GoldenCarrot                    = Item{ID: 952, DisplayName: "Golden Carrot", Name: "golden_carrot", StackSize: 64}
	SkeletonSkull                   = Item{ID: 953, DisplayName: "Skeleton Skull", Name: "skeleton_skull", StackSize: 64}
	WitherSkeletonSkull             = Item{ID: 954, DisplayName: "Wither Skeleton Skull", Name: "wither_skeleton_skull", StackSize: 64}
	PlayerHead                      = Item{ID: 955, DisplayName: "Player Head", Name: "player_head", StackSize: 64}
	ZombieHead                      = Item{ID: 956, DisplayName: "Zombie Head", Name: "zombie_head", StackSize: 64}
	CreeperHead                     = Item{ID: 957, DisplayName: "Creeper Head", Name: "creeper_head", StackSize: 64}
	DragonHead                      = Item{ID: 958, DisplayName: "Dragon Head", Name: "dragon_head", StackSize: 64}
	NetherStar                      = Item{ID: 959, DisplayName: "Nether Star", Name: "nether_star", StackSize: 64}
	PumpkinPie                      = Item{ID: 960, DisplayName: "Pumpkin Pie", Name: "pumpkin_pie", StackSize: 64}
	FireworkRocket                  = Item{ID: 961, DisplayName: "Firework Rocket", Name: "firework_rocket", StackSize: 64}
	FireworkStar                    = Item{ID: 962, DisplayName: "Firework Star", Name: "firework_star", StackSize: 64}
	EnchantedBook                   = Item{ID: 963, DisplayName: "Enchanted Book", Name: "enchanted_book", StackSize: 1}
	NetherBrick                     = Item{ID: 964, DisplayName: "Nether Brick", Name: "nether_brick", StackSize: 64}
	PrismarineShard                 = Item{ID: 965, DisplayName: "Prismarine Shard", Name: "prismarine_shard", StackSize: 64}
	PrismarineCrystals              = Item{ID: 966, DisplayName: "Prismarine Crystals", Name: "prismarine_crystals", StackSize: 64}
	Rabbit                          = Item{ID: 967, DisplayName: "Raw Rabbit", Name: "rabbit", StackSize: 64}
	CookedRabbit                    = Item{ID: 968, DisplayName: "Cooked Rabbit", Name: "cooked_rabbit", StackSize: 64}
	RabbitStew                      = Item{ID: 969, DisplayName: "Rabbit Stew", Name: "rabbit_stew", StackSize: 1}
	RabbitFoot                      = Item{ID: 970, DisplayName: "Rabbit's Foot", Name: "rabbit_foot", StackSize: 64}
	RabbitHide                      = Item{ID: 971, DisplayName: "Rabbit Hide", Name: "rabbit_hide", StackSize: 64}
	ArmorStand                      = Item{ID: 972, DisplayName: "Armor Stand", Name: "armor_stand", StackSize: 16}
	IronHorseArmor                  = Item{ID: 973, DisplayName: "Iron Horse Armor", Name: "iron_horse_armor", StackSize: 1}
	GoldenHorseArmor                = Item{ID: 974, DisplayName: "Golden Horse Armor", Name: "golden_horse_armor", StackSize: 1}
	DiamondHorseArmor               = Item{ID: 975, DisplayName: "Diamond Horse Armor", Name: "diamond_horse_armor", StackSize: 1}
	LeatherHorseArmor               = Item{ID: 976, DisplayName: "Leather Horse Armor", Name: "leather_horse_armor", StackSize: 1}
	Lead                            = Item{ID: 977, DisplayName: "Lead", Name: "lead", StackSize: 64}
	NameTag                         = Item{ID: 978, DisplayName: "Name Tag", Name: "name_tag", StackSize: 64}
	CommandBlockMinecart            = Item{ID: 979, DisplayName: "Minecart with Command Block", Name: "command_block_minecart", StackSize: 1}
	Mutton                          = Item{ID: 980, DisplayName: "Raw Mutton", Name: "mutton", StackSize: 64}
	CookedMutton                    = Item{ID: 981, DisplayName: "Cooked Mutton", Name: "cooked_mutton", StackSize: 64}
	WhiteBanner                     = Item{ID: 982, DisplayName: "White Banner", Name: "white_banner", StackSize: 16}
	OrangeBanner                    = Item{ID: 983, DisplayName: "Orange Banner", Name: "orange_banner", StackSize: 16}
	MagentaBanner                   = Item{ID: 984, DisplayName: "Magenta Banner", Name: "magenta_banner", StackSize: 16}
	LightBlueBanner                 = Item{ID: 985, DisplayName: "Light Blue Banner", Name: "light_blue_banner", StackSize: 16}
	YellowBanner                    = Item{ID: 986, DisplayName: "Yellow Banner", Name: "yellow_banner", StackSize: 16}
	LimeBanner                      = Item{ID: 987, DisplayName: "Lime Banner", Name: "lime_banner", StackSize: 16}
	PinkBanner                      = Item{ID: 988, DisplayName: "Pink Banner", Name: "pink_banner", StackSize: 16}
	GrayBanner                      = Item{ID: 989, DisplayName: "Gray Banner", Name: "gray_banner", StackSize: 16}
	LightGrayBanner                 = Item{ID: 990, DisplayName: "Light Gray Banner", Name: "light_gray_banner", StackSize: 16}
	CyanBanner                      = Item{ID: 991, DisplayName: "Cyan Banner", Name: "cyan_banner", StackSize: 16}
	PurpleBanner                    = Item{ID: 992, DisplayName: "Purple Banner", Name: "purple_banner", StackSize: 16}
	BlueBanner                      = Item{ID: 993, DisplayName: "Blue Banner", Name: "blue_banner", StackSize: 16}
	BrownBanner                     = Item{ID: 994, DisplayName: "Brown Banner", Name: "brown_banner", StackSize: 16}
	GreenBanner                     = Item{ID: 995, DisplayName: "Green Banner", Name: "green_banner", StackSize: 16}
	RedBanner                       = Item{ID: 996, DisplayName: "Red Banner", Name: "red_banner", StackSize: 16}
	BlackBanner                     = Item{ID: 997, DisplayName: "Black Banner", Name: "black_banner", StackSize: 16}
	EndCrystal                      = Item{ID: 998, DisplayName: "End Crystal", Name: "end_crystal", StackSize: 64}
	ChorusFruit                     = Item{ID: 999, DisplayName: "Chorus Fruit", Name: "chorus_fruit", StackSize: 64}
	PoppedChorusFruit               = Item{ID: 1000, DisplayName: "Popped Chorus Fruit", Name: "popped_chorus_fruit", StackSize: 64}
	Beetroot                        = Item{ID: 1001, DisplayName: "Beetroot", Name: "beetroot", StackSize: 64}
	BeetrootSeeds                   = Item{ID: 1002, DisplayName: "Beetroot Seeds", Name: "beetroot_seeds", StackSize: 64}
	BeetrootSoup                    = Item{ID: 1003, DisplayName: "Beetroot Soup", Name: "beetroot_soup", StackSize: 1}
	DragonBreath                    = Item{ID: 1004, DisplayName: "Dragon's Breath", Name: "dragon_breath", StackSize: 64}
	SplashPotion                    = Item{ID: 1005, DisplayName: "Splash Potion", Name: "splash_potion", StackSize: 1}
	SpectralArrow                   = Item{ID: 1006, DisplayName: "Spectral Arrow", Name: "spectral_arrow", StackSize: 64}
	TippedArrow                     = Item{ID: 1007, DisplayName: "Tipped Arrow", Name: "tipped_arrow", StackSize: 64}
	LingeringPotion                 = Item{ID: 1008, DisplayName: "Lingering Potion", Name: "lingering_potion", StackSize: 1}
	Shield                          = Item{ID: 1009, DisplayName: "Shield", Name: "shield", StackSize: 1}
	TotemOfUndying                  = Item{ID: 1010, DisplayName: "Totem of Undying", Name: "totem_of_undying", StackSize: 1}
	ShulkerShell                    = Item{ID: 1011, DisplayName: "Shulker Shell", Name: "shulker_shell", StackSize: 64}
	IronNugget                      = Item{ID: 1012, DisplayName: "Iron Nugget", Name: "iron_nugget", StackSize: 64}
	KnowledgeBook                   = Item{ID: 1013, DisplayName: "Knowledge Book", Name: "knowledge_book", StackSize: 1}
	DebugStick                      = Item{ID: 1014, DisplayName: "Debug Stick", Name: "debug_stick", StackSize: 1}
	MusicDisc13                     = Item{ID: 1015, DisplayName: "13 Disc", Name: "music_disc_13", StackSize: 1}
	MusicDiscCat                    = Item{ID: 1016, DisplayName: "Cat Disc", Name: "music_disc_cat", StackSize: 1}
	MusicDiscBlocks                 = Item{ID: 1017, DisplayName: "Blocks Disc", Name: "music_disc_blocks", StackSize: 1}
	MusicDiscChirp                  = Item{ID: 1018, DisplayName: "Chirp Disc", Name: "music_disc_chirp", StackSize: 1}
	MusicDiscFar                    = Item{ID: 1019, DisplayName: "Far Disc", Name: "music_disc_far", StackSize: 1}
	MusicDiscMall                   = Item{ID: 1020, DisplayName: "Mall Disc", Name: "music_disc_mall", StackSize: 1}
	MusicDiscMellohi                = Item{ID: 1021, DisplayName: "Mellohi Disc", Name: "music_disc_mellohi", StackSize: 1}
	MusicDiscStal                   = Item{ID: 1022, DisplayName: "Stal Disc", Name: "music_disc_stal", StackSize: 1}
	MusicDiscStrad                  = Item{ID: 1023, DisplayName: "Strad Disc", Name: "music_disc_strad", StackSize: 1}
	MusicDiscWard                   = Item{ID: 1024, DisplayName: "Ward Disc", Name: "music_disc_ward", StackSize: 1}
	MusicDisc11                     = Item{ID: 1025, DisplayName: "11 Disc", Name: "music_disc_11", StackSize: 1}
	MusicDiscWait                   = Item{ID: 1026, DisplayName: "Wait Disc", Name: "music_disc_wait", StackSize: 1}
	MusicDiscPigstep                = Item{ID: 1027, DisplayName: "Music Disc", Name: "music_disc_pigstep", StackSize: 1}
	Trident                         = Item{ID: 1028, DisplayName: "Trident", Name: "trident", StackSize: 1}
	PhantomMembrane                 = Item{ID: 1029, DisplayName: "Phantom Membrane", Name: "phantom_membrane", StackSize: 64}
	NautilusShell                   = Item{ID: 1030, DisplayName: "Nautilus Shell", Name: "nautilus_shell", StackSize: 64}
	HeartOfTheSea                   = Item{ID: 1031, DisplayName: "Heart of the Sea", Name: "heart_of_the_sea", StackSize: 64}
	Crossbow                        = Item{ID: 1032, DisplayName: "Crossbow", Name: "crossbow", StackSize: 1}
	SuspiciousStew                  = Item{ID: 1033, DisplayName: "Suspicious Stew", Name: "suspicious_stew", StackSize: 1}
	Loom                            = Item{ID: 1034, DisplayName: "Loom", Name: "loom", StackSize: 64}
	FlowerBannerPattern             = Item{ID: 1035, DisplayName: "Banner Pattern", Name: "flower_banner_pattern", StackSize: 1}
	CreeperBannerPattern            = Item{ID: 1036, DisplayName: "Banner Pattern", Name: "creeper_banner_pattern", StackSize: 1}
	SkullBannerPattern              = Item{ID: 1037, DisplayName: "Banner Pattern", Name: "skull_banner_pattern", StackSize: 1}
	MojangBannerPattern             = Item{ID: 1038, DisplayName: "Banner Pattern", Name: "mojang_banner_pattern", StackSize: 1}
	GlobeBannerPattern              = Item{ID: 1039, DisplayName: "Banner Pattern", Name: "globe_banner_pattern", StackSize: 1}
	PiglinBannerPattern             = Item{ID: 1040, DisplayName: "Banner Pattern", Name: "piglin_banner_pattern", StackSize: 1}
	Composter                       = Item{ID: 1041, DisplayName: "Composter", Name: "composter", StackSize: 64}
	Barrel                          = Item{ID: 1042, DisplayName: "Barrel", Name: "barrel", StackSize: 64}
	Smoker                          = Item{ID: 1043, DisplayName: "Smoker", Name: "smoker", StackSize: 64}
	BlastFurnace                    = Item{ID: 1044, DisplayName: "Blast Furnace", Name: "blast_furnace", StackSize: 64}
	CartographyTable                = Item{ID: 1045, DisplayName: "Cartography Table", Name: "cartography_table", StackSize: 64}
	FletchingTable                  = Item{ID: 1046, DisplayName: "Fletching Table", Name: "fletching_table", StackSize: 64}
	Grindstone                      = Item{ID: 1047, DisplayName: "Grindstone", Name: "grindstone", StackSize: 64}
	SmithingTable                   = Item{ID: 1048, DisplayName: "Smithing Table", Name: "smithing_table", StackSize: 64}
	Stonecutter                     = Item{ID: 1049, DisplayName: "Stonecutter", Name: "stonecutter", StackSize: 64}
	Bell                            = Item{ID: 1050, DisplayName: "Bell", Name: "bell", StackSize: 64}
	Lantern                         = Item{ID: 1051, DisplayName: "Lantern", Name: "lantern", StackSize: 64}
	SoulLantern                     = Item{ID: 1052, DisplayName: "Soul Lantern", Name: "soul_lantern", StackSize: 64}
	SweetBerries                    = Item{ID: 1053, DisplayName: "Sweet Berries", Name: "sweet_berries", StackSize: 64}
	GlowBerries                     = Item{ID: 1054, DisplayName: "Glow Berries", Name: "glow_berries", StackSize: 64}
	Campfire                        = Item{ID: 1055, DisplayName: "Campfire", Name: "campfire", StackSize: 64}
	SoulCampfire                    = Item{ID: 1056, DisplayName: "Soul Campfire", Name: "soul_campfire", StackSize: 64}
	Shroomlight                     = Item{ID: 1057, DisplayName: "Shroomlight", Name: "shroomlight", StackSize: 64}
	Honeycomb                       = Item{ID: 1058, DisplayName: "Honeycomb", Name: "honeycomb", StackSize: 64}
	BeeNest                         = Item{ID: 1059, DisplayName: "Bee Nest", Name: "bee_nest", StackSize: 64}
	Beehive                         = Item{ID: 1060, DisplayName: "Beehive", Name: "beehive", StackSize: 64}
	HoneyBottle                     = Item{ID: 1061, DisplayName: "Honey Bottle", Name: "honey_bottle", StackSize: 16}
	HoneycombBlock                  = Item{ID: 1062, DisplayName: "Honeycomb Block", Name: "honeycomb_block", StackSize: 64}
	Lodestone                       = Item{ID: 1063, DisplayName: "Lodestone", Name: "lodestone", StackSize: 64}
	CryingObsidian                  = Item{ID: 1064, DisplayName: "Crying Obsidian", Name: "crying_obsidian", StackSize: 64}
	Blackstone                      = Item{ID: 1065, DisplayName: "Blackstone", Name: "blackstone", StackSize: 64}
	BlackstoneSlab                  = Item{ID: 1066, DisplayName: "Blackstone Slab", Name: "blackstone_slab", StackSize: 64}
	BlackstoneStairs                = Item{ID: 1067, DisplayName: "Blackstone Stairs", Name: "blackstone_stairs", StackSize: 64}
	GildedBlackstone                = Item{ID: 1068, DisplayName: "Gilded Blackstone", Name: "gilded_blackstone", StackSize: 64}
	PolishedBlackstone              = Item{ID: 1069, DisplayName: "Polished Blackstone", Name: "polished_blackstone", StackSize: 64}
	PolishedBlackstoneSlab          = Item{ID: 1070, DisplayName: "Polished Blackstone Slab", Name: "polished_blackstone_slab", StackSize: 64}
	PolishedBlackstoneStairs        = Item{ID: 1071, DisplayName: "Polished Blackstone Stairs", Name: "polished_blackstone_stairs", StackSize: 64}
	ChiseledPolishedBlackstone      = Item{ID: 1072, DisplayName: "Chiseled Polished Blackstone", Name: "chiseled_polished_blackstone", StackSize: 64}
	PolishedBlackstoneBricks        = Item{ID: 1073, DisplayName: "Polished Blackstone Bricks", Name: "polished_blackstone_bricks", StackSize: 64}
	PolishedBlackstoneBrickSlab     = Item{ID: 1074, DisplayName: "Polished Blackstone Brick Slab", Name: "polished_blackstone_brick_slab", StackSize: 64}
	PolishedBlackstoneBrickStairs   = Item{ID: 1075, DisplayName: "Polished Blackstone Brick Stairs", Name: "polished_blackstone_brick_stairs", StackSize: 64}
	CrackedPolishedBlackstoneBricks = Item{ID: 1076, DisplayName: "Cracked Polished Blackstone Bricks", Name: "cracked_polished_blackstone_bricks", StackSize: 64}
	RespawnAnchor                   = Item{ID: 1077, DisplayName: "Respawn Anchor", Name: "respawn_anchor", StackSize: 64}
	Candle                          = Item{ID: 1078, DisplayName: "Candle", Name: "candle", StackSize: 64}
	WhiteCandle                     = Item{ID: 1079, DisplayName: "White Candle", Name: "white_candle", StackSize: 64}
	OrangeCandle                    = Item{ID: 1080, DisplayName: "Orange Candle", Name: "orange_candle", StackSize: 64}
	MagentaCandle                   = Item{ID: 1081, DisplayName: "Magenta Candle", Name: "magenta_candle", StackSize: 64}
	LightBlueCandle                 = Item{ID: 1082, DisplayName: "Light Blue Candle", Name: "light_blue_candle", StackSize: 64}
	YellowCandle                    = Item{ID: 1083, DisplayName: "Yellow Candle", Name: "yellow_candle", StackSize: 64}
	LimeCandle                      = Item{ID: 1084, DisplayName: "Lime Candle", Name: "lime_candle", StackSize: 64}
	PinkCandle                      = Item{ID: 1085, DisplayName: "Pink Candle", Name: "pink_candle", StackSize: 64}
	GrayCandle                      = Item{ID: 1086, DisplayName: "Gray Candle", Name: "gray_candle", StackSize: 64}
	LightGrayCandle                 = Item{ID: 1087, DisplayName: "Light Gray Candle", Name: "light_gray_candle", StackSize: 64}
	CyanCandle                      = Item{ID: 1088, DisplayName: "Cyan Candle", Name: "cyan_candle", StackSize: 64}
	PurpleCandle                    = Item{ID: 1089, DisplayName: "Purple Candle", Name: "purple_candle", StackSize: 64}
	BlueCandle                      = Item{ID: 1090, DisplayName: "Blue Candle", Name: "blue_candle", StackSize: 64}
	BrownCandle                     = Item{ID: 1091, DisplayName: "Brown Candle", Name: "brown_candle", StackSize: 64}
	GreenCandle                     = Item{ID: 1092, DisplayName: "Green Candle", Name: "green_candle", StackSize: 64}
	RedCandle                       = Item{ID: 1093, DisplayName: "Red Candle", Name: "red_candle", StackSize: 64}
	BlackCandle                     = Item{ID: 1094, DisplayName: "Black Candle", Name: "black_candle", StackSize: 64}
	SmallAmethystBud                = Item{ID: 1095, DisplayName: "Small Amethyst Bud", Name: "small_amethyst_bud", StackSize: 64}
	MediumAmethystBud               = Item{ID: 1096, DisplayName: "Medium Amethyst Bud", Name: "medium_amethyst_bud", StackSize: 64}
	LargeAmethystBud                = Item{ID: 1097, DisplayName: "Large Amethyst Bud", Name: "large_amethyst_bud", StackSize: 64}
	AmethystCluster                 = Item{ID: 1098, DisplayName: "Amethyst Cluster", Name: "amethyst_cluster", StackSize: 64}
	PointedDripstone                = Item{ID: 1099, DisplayName: "Pointed Dripstone", Name: "pointed_dripstone", StackSize: 64}
)

// ByID is an index of minecraft items by their ID.
var ByID = map[ID]*Item{
	1:    &Stone,
	2:    &Granite,
	3:    &PolishedGranite,
	4:    &Diorite,
	5:    &PolishedDiorite,
	6:    &Andesite,
	7:    &PolishedAndesite,
	8:    &Deepslate,
	9:    &CobbledDeepslate,
	10:   &PolishedDeepslate,
	11:   &Calcite,
	12:   &Tuff,
	13:   &DripstoneBlock,
	14:   &GrassBlock,
	15:   &Dirt,
	16:   &CoarseDirt,
	17:   &Podzol,
	18:   &RootedDirt,
	19:   &CrimsonNylium,
	20:   &WarpedNylium,
	21:   &Cobblestone,
	22:   &OakPlanks,
	23:   &SprucePlanks,
	24:   &BirchPlanks,
	25:   &JunglePlanks,
	26:   &AcaciaPlanks,
	27:   &DarkOakPlanks,
	28:   &CrimsonPlanks,
	29:   &WarpedPlanks,
	30:   &OakSapling,
	31:   &SpruceSapling,
	32:   &BirchSapling,
	33:   &JungleSapling,
	34:   &AcaciaSapling,
	35:   &DarkOakSapling,
	36:   &Bedrock,
	37:   &Sand,
	38:   &RedSand,
	39:   &Gravel,
	40:   &CoalOre,
	41:   &DeepslateCoalOre,
	42:   &IronOre,
	43:   &DeepslateIronOre,
	44:   &CopperOre,
	45:   &DeepslateCopperOre,
	46:   &GoldOre,
	47:   &DeepslateGoldOre,
	48:   &RedstoneOre,
	49:   &DeepslateRedstoneOre,
	50:   &EmeraldOre,
	51:   &DeepslateEmeraldOre,
	52:   &LapisOre,
	53:   &DeepslateLapisOre,
	54:   &DiamondOre,
	55:   &DeepslateDiamondOre,
	56:   &NetherGoldOre,
	57:   &NetherQuartzOre,
	58:   &AncientDebris,
	59:   &CoalBlock,
	60:   &RawIronBlock,
	61:   &RawCopperBlock,
	62:   &RawGoldBlock,
	63:   &AmethystBlock,
	64:   &BuddingAmethyst,
	65:   &IronBlock,
	66:   &CopperBlock,
	67:   &GoldBlock,
	68:   &DiamondBlock,
	69:   &NetheriteBlock,
	70:   &ExposedCopper,
	71:   &WeatheredCopper,
	72:   &OxidizedCopper,
	73:   &CutCopper,
	74:   &ExposedCutCopper,
	75:   &WeatheredCutCopper,
	76:   &OxidizedCutCopper,
	77:   &CutCopperStairs,
	78:   &ExposedCutCopperStairs,
	79:   &WeatheredCutCopperStairs,
	80:   &OxidizedCutCopperStairs,
	81:   &CutCopperSlab,
	82:   &ExposedCutCopperSlab,
	83:   &WeatheredCutCopperSlab,
	84:   &OxidizedCutCopperSlab,
	85:   &WaxedCopperBlock,
	86:   &WaxedExposedCopper,
	87:   &WaxedWeatheredCopper,
	88:   &WaxedOxidizedCopper,
	89:   &WaxedCutCopper,
	90:   &WaxedExposedCutCopper,
	91:   &WaxedWeatheredCutCopper,
	92:   &WaxedOxidizedCutCopper,
	93:   &WaxedCutCopperStairs,
	94:   &WaxedExposedCutCopperStairs,
	95:   &WaxedWeatheredCutCopperStairs,
	96:   &WaxedOxidizedCutCopperStairs,
	97:   &WaxedCutCopperSlab,
	98:   &WaxedExposedCutCopperSlab,
	99:   &WaxedWeatheredCutCopperSlab,
	100:  &WaxedOxidizedCutCopperSlab,
	101:  &OakLog,
	102:  &SpruceLog,
	103:  &BirchLog,
	104:  &JungleLog,
	105:  &AcaciaLog,
	106:  &DarkOakLog,
	107:  &CrimsonStem,
	108:  &WarpedStem,
	109:  &StrippedOakLog,
	110:  &StrippedSpruceLog,
	111:  &StrippedBirchLog,
	112:  &StrippedJungleLog,
	113:  &StrippedAcaciaLog,
	114:  &StrippedDarkOakLog,
	115:  &StrippedCrimsonStem,
	116:  &StrippedWarpedStem,
	117:  &StrippedOakWood,
	118:  &StrippedSpruceWood,
	119:  &StrippedBirchWood,
	120:  &StrippedJungleWood,
	121:  &StrippedAcaciaWood,
	122:  &StrippedDarkOakWood,
	123:  &StrippedCrimsonHyphae,
	124:  &StrippedWarpedHyphae,
	125:  &OakWood,
	126:  &SpruceWood,
	127:  &BirchWood,
	128:  &JungleWood,
	129:  &AcaciaWood,
	130:  &DarkOakWood,
	131:  &CrimsonHyphae,
	132:  &WarpedHyphae,
	133:  &OakLeaves,
	134:  &SpruceLeaves,
	135:  &BirchLeaves,
	136:  &JungleLeaves,
	137:  &AcaciaLeaves,
	138:  &DarkOakLeaves,
	139:  &AzaleaLeaves,
	140:  &FloweringAzaleaLeaves,
	141:  &Sponge,
	142:  &WetSponge,
	143:  &Glass,
	144:  &TintedGlass,
	145:  &LapisBlock,
	146:  &Sandstone,
	147:  &ChiseledSandstone,
	148:  &CutSandstone,
	149:  &Cobweb,
	150:  &Grass,
	151:  &Fern,
	152:  &Azalea,
	153:  &FloweringAzalea,
	154:  &DeadBush,
	155:  &Seagrass,
	156:  &SeaPickle,
	157:  &WhiteWool,
	158:  &OrangeWool,
	159:  &MagentaWool,
	160:  &LightBlueWool,
	161:  &YellowWool,
	162:  &LimeWool,
	163:  &PinkWool,
	164:  &GrayWool,
	165:  &LightGrayWool,
	166:  &CyanWool,
	167:  &PurpleWool,
	168:  &BlueWool,
	169:  &BrownWool,
	170:  &GreenWool,
	171:  &RedWool,
	172:  &BlackWool,
	173:  &Dandelion,
	174:  &Poppy,
	175:  &BlueOrchid,
	176:  &Allium,
	177:  &AzureBluet,
	178:  &RedTulip,
	179:  &OrangeTulip,
	180:  &WhiteTulip,
	181:  &PinkTulip,
	182:  &OxeyeDaisy,
	183:  &Cornflower,
	184:  &LilyOfTheValley,
	185:  &WitherRose,
	186:  &SporeBlossom,
	187:  &BrownMushroom,
	188:  &RedMushroom,
	189:  &CrimsonFungus,
	190:  &WarpedFungus,
	191:  &CrimsonRoots,
	192:  &WarpedRoots,
	193:  &NetherSprouts,
	194:  &WeepingVines,
	195:  &TwistingVines,
	196:  &SugarCane,
	197:  &Kelp,
	198:  &MossCarpet,
	199:  &MossBlock,
	200:  &HangingRoots,
	201:  &BigDripleaf,
	202:  &SmallDripleaf,
	203:  &Bamboo,
	204:  &OakSlab,
	205:  &SpruceSlab,
	206:  &BirchSlab,
	207:  &JungleSlab,
	208:  &AcaciaSlab,
	209:  &DarkOakSlab,
	210:  &CrimsonSlab,
	211:  &WarpedSlab,
	212:  &StoneSlab,
	213:  &SmoothStoneSlab,
	214:  &SandstoneSlab,
	215:  &CutSandstoneSlab,
	216:  &PetrifiedOakSlab,
	217:  &CobblestoneSlab,
	218:  &BrickSlab,
	219:  &StoneBrickSlab,
	220:  &NetherBrickSlab,
	221:  &QuartzSlab,
	222:  &RedSandstoneSlab,
	223:  &CutRedSandstoneSlab,
	224:  &PurpurSlab,
	225:  &PrismarineSlab,
	226:  &PrismarineBrickSlab,
	227:  &DarkPrismarineSlab,
	228:  &SmoothQuartz,
	229:  &SmoothRedSandstone,
	230:  &SmoothSandstone,
	231:  &SmoothStone,
	232:  &Bricks,
	233:  &Bookshelf,
	234:  &MossyCobblestone,
	235:  &Obsidian,
	236:  &Torch,
	237:  &EndRod,
	238:  &ChorusPlant,
	239:  &ChorusFlower,
	240:  &PurpurBlock,
	241:  &PurpurPillar,
	242:  &PurpurStairs,
	243:  &Spawner,
	244:  &OakStairs,
	245:  &Chest,
	246:  &CraftingTable,
	247:  &Farmland,
	248:  &Furnace,
	249:  &Ladder,
	250:  &CobblestoneStairs,
	251:  &Snow,
	252:  &Ice,
	253:  &SnowBlock,
	254:  &Cactus,
	255:  &Clay,
	256:  &Jukebox,
	257:  &OakFence,
	258:  &SpruceFence,
	259:  &BirchFence,
	260:  &JungleFence,
	261:  &AcaciaFence,
	262:  &DarkOakFence,
	263:  &CrimsonFence,
	264:  &WarpedFence,
	265:  &Pumpkin,
	266:  &CarvedPumpkin,
	267:  &JackOLantern,
	268:  &Netherrack,
	269:  &SoulSand,
	270:  &SoulSoil,
	271:  &Basalt,
	272:  &PolishedBasalt,
	273:  &SmoothBasalt,
	274:  &SoulTorch,
	275:  &Glowstone,
	276:  &InfestedStone,
	277:  &InfestedCobblestone,
	278:  &InfestedStoneBricks,
	279:  &InfestedMossyStoneBricks,
	280:  &InfestedCrackedStoneBricks,
	281:  &InfestedChiseledStoneBricks,
	282:  &InfestedDeepslate,
	283:  &StoneBricks,
	284:  &MossyStoneBricks,
	285:  &CrackedStoneBricks,
	286:  &ChiseledStoneBricks,
	287:  &DeepslateBricks,
	288:  &CrackedDeepslateBricks,
	289:  &DeepslateTiles,
	290:  &CrackedDeepslateTiles,
	291:  &ChiseledDeepslate,
	292:  &BrownMushroomBlock,
	293:  &RedMushroomBlock,
	294:  &MushroomStem,
	295:  &IronBars,
	296:  &Chain,
	297:  &GlassPane,
	298:  &Melon,
	299:  &Vine,
	300:  &GlowLichen,
	301:  &BrickStairs,
	302:  &StoneBrickStairs,
	303:  &Mycelium,
	304:  &LilyPad,
	305:  &NetherBricks,
	306:  &CrackedNetherBricks,
	307:  &ChiseledNetherBricks,
	308:  &NetherBrickFence,
	309:  &NetherBrickStairs,
	310:  &EnchantingTable,
	311:  &EndPortalFrame,
	312:  &EndStone,
	313:  &EndStoneBricks,
	314:  &DragonEgg,
	315:  &SandstoneStairs,
	316:  &EnderChest,
	317:  &EmeraldBlock,
	318:  &SpruceStairs,
	319:  &BirchStairs,
	320:  &JungleStairs,
	321:  &CrimsonStairs,
	322:  &WarpedStairs,
	323:  &CommandBlock,
	324:  &Beacon,
	325:  &CobblestoneWall,
	326:  &MossyCobblestoneWall,
	327:  &BrickWall,
	328:  &PrismarineWall,
	329:  &RedSandstoneWall,
	330:  &MossyStoneBrickWall,
	331:  &GraniteWall,
	332:  &StoneBrickWall,
	333:  &NetherBrickWall,
	334:  &AndesiteWall,
	335:  &RedNetherBrickWall,
	336:  &SandstoneWall,
	337:  &EndStoneBrickWall,
	338:  &DioriteWall,
	339:  &BlackstoneWall,
	340:  &PolishedBlackstoneWall,
	341:  &PolishedBlackstoneBrickWall,
	342:  &CobbledDeepslateWall,
	343:  &PolishedDeepslateWall,
	344:  &DeepslateBrickWall,
	345:  &DeepslateTileWall,
	346:  &Anvil,
	347:  &ChippedAnvil,
	348:  &DamagedAnvil,
	349:  &ChiseledQuartzBlock,
	350:  &QuartzBlock,
	351:  &QuartzBricks,
	352:  &QuartzPillar,
	353:  &QuartzStairs,
	354:  &WhiteTerracotta,
	355:  &OrangeTerracotta,
	356:  &MagentaTerracotta,
	357:  &LightBlueTerracotta,
	358:  &YellowTerracotta,
	359:  &LimeTerracotta,
	360:  &PinkTerracotta,
	361:  &GrayTerracotta,
	362:  &LightGrayTerracotta,
	363:  &CyanTerracotta,
	364:  &PurpleTerracotta,
	365:  &BlueTerracotta,
	366:  &BrownTerracotta,
	367:  &GreenTerracotta,
	368:  &RedTerracotta,
	369:  &BlackTerracotta,
	370:  &Barrier,
	371:  &Light,
	372:  &HayBlock,
	373:  &WhiteCarpet,
	374:  &OrangeCarpet,
	375:  &MagentaCarpet,
	376:  &LightBlueCarpet,
	377:  &YellowCarpet,
	378:  &LimeCarpet,
	379:  &PinkCarpet,
	380:  &GrayCarpet,
	381:  &LightGrayCarpet,
	382:  &CyanCarpet,
	383:  &PurpleCarpet,
	384:  &BlueCarpet,
	385:  &BrownCarpet,
	386:  &GreenCarpet,
	387:  &RedCarpet,
	388:  &BlackCarpet,
	389:  &Terracotta,
	390:  &PackedIce,
	391:  &AcaciaStairs,
	392:  &DarkOakStairs,
	393:  &DirtPath,
	394:  &Sunflower,
	395:  &Lilac,
	396:  &RoseBush,
	397:  &Peony,
	398:  &TallGrass,
	399:  &LargeFern,
	400:  &WhiteStainedGlass,
	401:  &OrangeStainedGlass,
	402:  &MagentaStainedGlass,
	403:  &LightBlueStainedGlass,
	404:  &YellowStainedGlass,
	405:  &LimeStainedGlass,
	406:  &PinkStainedGlass,
	407:  &GrayStainedGlass,
	408:  &LightGrayStainedGlass,
	409:  &CyanStainedGlass,
	410:  &PurpleStainedGlass,
	411:  &BlueStainedGlass,
	412:  &BrownStainedGlass,
	413:  &GreenStainedGlass,
	414:  &RedStainedGlass,
	415:  &BlackStainedGlass,
	416:  &WhiteStainedGlassPane,
	417:  &OrangeStainedGlassPane,
	418:  &MagentaStainedGlassPane,
	419:  &LightBlueStainedGlassPane,
	420:  &YellowStainedGlassPane,
	421:  &LimeStainedGlassPane,
	422:  &PinkStainedGlassPane,
	423:  &GrayStainedGlassPane,
	424:  &LightGrayStainedGlassPane,
	425:  &CyanStainedGlassPane,
	426:  &PurpleStainedGlassPane,
	427:  &BlueStainedGlassPane,
	428:  &BrownStainedGlassPane,
	429:  &GreenStainedGlassPane,
	430:  &RedStainedGlassPane,
	431:  &BlackStainedGlassPane,
	432:  &Prismarine,
	433:  &PrismarineBricks,
	434:  &DarkPrismarine,
	435:  &PrismarineStairs,
	436:  &PrismarineBrickStairs,
	437:  &DarkPrismarineStairs,
	438:  &SeaLantern,
	439:  &RedSandstone,
	440:  &ChiseledRedSandstone,
	441:  &CutRedSandstone,
	442:  &RedSandstoneStairs,
	443:  &RepeatingCommandBlock,
	444:  &ChainCommandBlock,
	445:  &MagmaBlock,
	446:  &NetherWartBlock,
	447:  &WarpedWartBlock,
	448:  &RedNetherBricks,
	449:  &BoneBlock,
	450:  &StructureVoid,
	451:  &ShulkerBox,
	452:  &WhiteShulkerBox,
	453:  &OrangeShulkerBox,
	454:  &MagentaShulkerBox,
	455:  &LightBlueShulkerBox,
	456:  &YellowShulkerBox,
	457:  &LimeShulkerBox,
	458:  &PinkShulkerBox,
	459:  &GrayShulkerBox,
	460:  &LightGrayShulkerBox,
	461:  &CyanShulkerBox,
	462:  &PurpleShulkerBox,
	463:  &BlueShulkerBox,
	464:  &BrownShulkerBox,
	465:  &GreenShulkerBox,
	466:  &RedShulkerBox,
	467:  &BlackShulkerBox,
	468:  &WhiteGlazedTerracotta,
	469:  &OrangeGlazedTerracotta,
	470:  &MagentaGlazedTerracotta,
	471:  &LightBlueGlazedTerracotta,
	472:  &YellowGlazedTerracotta,
	473:  &LimeGlazedTerracotta,
	474:  &PinkGlazedTerracotta,
	475:  &GrayGlazedTerracotta,
	476:  &LightGrayGlazedTerracotta,
	477:  &CyanGlazedTerracotta,
	478:  &PurpleGlazedTerracotta,
	479:  &BlueGlazedTerracotta,
	480:  &BrownGlazedTerracotta,
	481:  &GreenGlazedTerracotta,
	482:  &RedGlazedTerracotta,
	483:  &BlackGlazedTerracotta,
	484:  &WhiteConcrete,
	485:  &OrangeConcrete,
	486:  &MagentaConcrete,
	487:  &LightBlueConcrete,
	488:  &YellowConcrete,
	489:  &LimeConcrete,
	490:  &PinkConcrete,
	491:  &GrayConcrete,
	492:  &LightGrayConcrete,
	493:  &CyanConcrete,
	494:  &PurpleConcrete,
	495:  &BlueConcrete,
	496:  &BrownConcrete,
	497:  &GreenConcrete,
	498:  &RedConcrete,
	499:  &BlackConcrete,
	500:  &WhiteConcretePowder,
	501:  &OrangeConcretePowder,
	502:  &MagentaConcretePowder,
	503:  &LightBlueConcretePowder,
	504:  &YellowConcretePowder,
	505:  &LimeConcretePowder,
	506:  &PinkConcretePowder,
	507:  &GrayConcretePowder,
	508:  &LightGrayConcretePowder,
	509:  &CyanConcretePowder,
	510:  &PurpleConcretePowder,
	511:  &BlueConcretePowder,
	512:  &BrownConcretePowder,
	513:  &GreenConcretePowder,
	514:  &RedConcretePowder,
	515:  &BlackConcretePowder,
	516:  &TurtleEgg,
	517:  &DeadTubeCoralBlock,
	518:  &DeadBrainCoralBlock,
	519:  &DeadBubbleCoralBlock,
	520:  &DeadFireCoralBlock,
	521:  &DeadHornCoralBlock,
	522:  &TubeCoralBlock,
	523:  &BrainCoralBlock,
	524:  &BubbleCoralBlock,
	525:  &FireCoralBlock,
	526:  &HornCoralBlock,
	527:  &TubeCoral,
	528:  &BrainCoral,
	529:  &BubbleCoral,
	530:  &FireCoral,
	531:  &HornCoral,
	532:  &DeadBrainCoral,
	533:  &DeadBubbleCoral,
	534:  &DeadFireCoral,
	535:  &DeadHornCoral,
	536:  &DeadTubeCoral,
	537:  &TubeCoralFan,
	538:  &BrainCoralFan,
	539:  &BubbleCoralFan,
	540:  &FireCoralFan,
	541:  &HornCoralFan,
	542:  &DeadTubeCoralFan,
	543:  &DeadBrainCoralFan,
	544:  &DeadBubbleCoralFan,
	545:  &DeadFireCoralFan,
	546:  &DeadHornCoralFan,
	547:  &BlueIce,
	548:  &Conduit,
	549:  &PolishedGraniteStairs,
	550:  &SmoothRedSandstoneStairs,
	551:  &MossyStoneBrickStairs,
	552:  &PolishedDioriteStairs,
	553:  &MossyCobblestoneStairs,
	554:  &EndStoneBrickStairs,
	555:  &StoneStairs,
	556:  &SmoothSandstoneStairs,
	557:  &SmoothQuartzStairs,
	558:  &GraniteStairs,
	559:  &AndesiteStairs,
	560:  &RedNetherBrickStairs,
	561:  &PolishedAndesiteStairs,
	562:  &DioriteStairs,
	563:  &CobbledDeepslateStairs,
	564:  &PolishedDeepslateStairs,
	565:  &DeepslateBrickStairs,
	566:  &DeepslateTileStairs,
	567:  &PolishedGraniteSlab,
	568:  &SmoothRedSandstoneSlab,
	569:  &MossyStoneBrickSlab,
	570:  &PolishedDioriteSlab,
	571:  &MossyCobblestoneSlab,
	572:  &EndStoneBrickSlab,
	573:  &SmoothSandstoneSlab,
	574:  &SmoothQuartzSlab,
	575:  &GraniteSlab,
	576:  &AndesiteSlab,
	577:  &RedNetherBrickSlab,
	578:  &PolishedAndesiteSlab,
	579:  &DioriteSlab,
	580:  &CobbledDeepslateSlab,
	581:  &PolishedDeepslateSlab,
	582:  &DeepslateBrickSlab,
	583:  &DeepslateTileSlab,
	584:  &Scaffolding,
	585:  &Redstone,
	586:  &RedstoneTorch,
	587:  &RedstoneBlock,
	588:  &Repeater,
	589:  &Comparator,
	590:  &Piston,
	591:  &StickyPiston,
	592:  &SlimeBlock,
	593:  &HoneyBlock,
	594:  &Observer,
	595:  &Hopper,
	596:  &Dispenser,
	597:  &Dropper,
	598:  &Lectern,
	599:  &Target,
	600:  &Lever,
	601:  &LightningRod,
	602:  &DaylightDetector,
	603:  &SculkSensor,
	604:  &TripwireHook,
	605:  &TrappedChest,
	606:  &Tnt,
	607:  &RedstoneLamp,
	608:  &NoteBlock,
	609:  &StoneButton,
	610:  &PolishedBlackstoneButton,
	611:  &OakButton,
	612:  &SpruceButton,
	613:  &BirchButton,
	614:  &JungleButton,
	615:  &AcaciaButton,
	616:  &DarkOakButton,
	617:  &CrimsonButton,
	618:  &WarpedButton,
	619:  &StonePressurePlate,
	620:  &PolishedBlackstonePressurePlate,
	621:  &LightWeightedPressurePlate,
	622:  &HeavyWeightedPressurePlate,
	623:  &OakPressurePlate,
	624:  &SprucePressurePlate,
	625:  &BirchPressurePlate,
	626:  &JunglePressurePlate,
	627:  &AcaciaPressurePlate,
	628:  &DarkOakPressurePlate,
	629:  &CrimsonPressurePlate,
	630:  &WarpedPressurePlate,
	631:  &IronDoor,
	632:  &OakDoor,
	633:  &SpruceDoor,
	634:  &BirchDoor,
	635:  &JungleDoor,
	636:  &AcaciaDoor,
	637:  &DarkOakDoor,
	638:  &CrimsonDoor,
	639:  &WarpedDoor,
	640:  &IronTrapdoor,
	641:  &OakTrapdoor,
	642:  &SpruceTrapdoor,
	643:  &BirchTrapdoor,
	644:  &JungleTrapdoor,
	645:  &AcaciaTrapdoor,
	646:  &DarkOakTrapdoor,
	647:  &CrimsonTrapdoor,
	648:  &WarpedTrapdoor,
	649:  &OakFenceGate,
	650:  &SpruceFenceGate,
	651:  &BirchFenceGate,
	652:  &JungleFenceGate,
	653:  &AcaciaFenceGate,
	654:  &DarkOakFenceGate,
	655:  &CrimsonFenceGate,
	656:  &WarpedFenceGate,
	657:  &PoweredRail,
	658:  &DetectorRail,
	659:  &Rail,
	660:  &ActivatorRail,
	661:  &Saddle,
	662:  &Minecart,
	663:  &ChestMinecart,
	664:  &FurnaceMinecart,
	665:  &TntMinecart,
	666:  &HopperMinecart,
	667:  &CarrotOnAStick,
	668:  &WarpedFungusOnAStick,
	669:  &Elytra,
	670:  &OakBoat,
	671:  &SpruceBoat,
	672:  &BirchBoat,
	673:  &JungleBoat,
	674:  &AcaciaBoat,
	675:  &DarkOakBoat,
	676:  &StructureBlock,
	677:  &Jigsaw,
	678:  &TurtleHelmet,
	679:  &Scute,
	680:  &FlintAndSteel,
	681:  &Apple,
	682:  &Bow,
	683:  &Arrow,
	684:  &Coal,
	685:  &Charcoal,
	686:  &Diamond,
	687:  &Emerald,
	688:  &LapisLazuli,
	689:  &Quartz,
	690:  &AmethystShard,
	691:  &RawIron,
	692:  &IronIngot,
	693:  &RawCopper,
	694:  &CopperIngot,
	695:  &RawGold,
	696:  &GoldIngot,
	697:  &NetheriteIngot,
	698:  &NetheriteScrap,
	699:  &WoodenSword,
	700:  &WoodenShovel,
	701:  &WoodenPickaxe,
	702:  &WoodenAxe,
	703:  &WoodenHoe,
	704:  &StoneSword,
	705:  &StoneShovel,
	706:  &StonePickaxe,
	707:  &StoneAxe,
	708:  &StoneHoe,
	709:  &GoldenSword,
	710:  &GoldenShovel,
	711:  &GoldenPickaxe,
	712:  &GoldenAxe,
	713:  &GoldenHoe,
	714:  &IronSword,
	715:  &IronShovel,
	716:  &IronPickaxe,
	717:  &IronAxe,
	718:  &IronHoe,
	719:  &DiamondSword,
	720:  &DiamondShovel,
	721:  &DiamondPickaxe,
	722:  &DiamondAxe,
	723:  &DiamondHoe,
	724:  &NetheriteSword,
	725:  &NetheriteShovel,
	726:  &NetheritePickaxe,
	727:  &NetheriteAxe,
	728:  &NetheriteHoe,
	729:  &Stick,
	730:  &Bowl,
	731:  &MushroomStew,
	732:  &String,
	733:  &Feather,
	734:  &Gunpowder,
	735:  &WheatSeeds,
	736:  &Wheat,
	737:  &Bread,
	738:  &LeatherHelmet,
	739:  &LeatherChestplate,
	740:  &LeatherLeggings,
	741:  &LeatherBoots,
	742:  &ChainmailHelmet,
	743:  &ChainmailChestplate,
	744:  &ChainmailLeggings,
	745:  &ChainmailBoots,
	746:  &IronHelmet,
	747:  &IronChestplate,
	748:  &IronLeggings,
	749:  &IronBoots,
	750:  &DiamondHelmet,
	751:  &DiamondChestplate,
	752:  &DiamondLeggings,
	753:  &DiamondBoots,
	754:  &GoldenHelmet,
	755:  &GoldenChestplate,
	756:  &GoldenLeggings,
	757:  &GoldenBoots,
	758:  &NetheriteHelmet,
	759:  &NetheriteChestplate,
	760:  &NetheriteLeggings,
	761:  &NetheriteBoots,
	762:  &Flint,
	763:  &Porkchop,
	764:  &CookedPorkchop,
	765:  &Painting,
	766:  &GoldenApple,
	767:  &EnchantedGoldenApple,
	768:  &OakSign,
	769:  &SpruceSign,
	770:  &BirchSign,
	771:  &JungleSign,
	772:  &AcaciaSign,
	773:  &DarkOakSign,
	774:  &CrimsonSign,
	775:  &WarpedSign,
	776:  &Bucket,
	777:  &WaterBucket,
	778:  &LavaBucket,
	779:  &PowderSnowBucket,
	780:  &Snowball,
	781:  &Leather,
	782:  &MilkBucket,
	783:  &PufferfishBucket,
	784:  &SalmonBucket,
	785:  &CodBucket,
	786:  &TropicalFishBucket,
	787:  &AxolotlBucket,
	788:  &Brick,
	789:  &ClayBall,
	790:  &DriedKelpBlock,
	791:  &Paper,
	792:  &Book,
	793:  &SlimeBall,
	794:  &Egg,
	795:  &Compass,
	796:  &Bundle,
	797:  &FishingRod,
	798:  &Clock,
	799:  &Spyglass,
	800:  &GlowstoneDust,
	801:  &Cod,
	802:  &Salmon,
	803:  &TropicalFish,
	804:  &Pufferfish,
	805:  &CookedCod,
	806:  &CookedSalmon,
	807:  &InkSac,
	808:  &GlowInkSac,
	809:  &CocoaBeans,
	810:  &WhiteDye,
	811:  &OrangeDye,
	812:  &MagentaDye,
	813:  &LightBlueDye,
	814:  &YellowDye,
	815:  &LimeDye,
	816:  &PinkDye,
	817:  &GrayDye,
	818:  &LightGrayDye,
	819:  &CyanDye,
	820:  &PurpleDye,
	821:  &BlueDye,
	822:  &BrownDye,
	823:  &GreenDye,
	824:  &RedDye,
	825:  &BlackDye,
	826:  &BoneMeal,
	827:  &Bone,
	828:  &Sugar,
	829:  &Cake,
	830:  &WhiteBed,
	831:  &OrangeBed,
	832:  &MagentaBed,
	833:  &LightBlueBed,
	834:  &YellowBed,
	835:  &LimeBed,
	836:  &PinkBed,
	837:  &GrayBed,
	838:  &LightGrayBed,
	839:  &CyanBed,
	840:  &PurpleBed,
	841:  &BlueBed,
	842:  &BrownBed,
	843:  &GreenBed,
	844:  &RedBed,
	845:  &BlackBed,
	846:  &Cookie,
	847:  &FilledMap,
	848:  &Shears,
	849:  &MelonSlice,
	850:  &DriedKelp,
	851:  &PumpkinSeeds,
	852:  &MelonSeeds,
	853:  &Beef,
	854:  &CookedBeef,
	855:  &Chicken,
	856:  &CookedChicken,
	857:  &RottenFlesh,
	858:  &EnderPearl,
	859:  &BlazeRod,
	860:  &GhastTear,
	861:  &GoldNugget,
	862:  &NetherWart,
	863:  &Potion,
	864:  &GlassBottle,
	865:  &SpiderEye,
	866:  &FermentedSpiderEye,
	867:  &BlazePowder,
	868:  &MagmaCream,
	869:  &BrewingStand,
	870:  &Cauldron,
	871:  &EnderEye,
	872:  &GlisteringMelonSlice,
	873:  &AxolotlSpawnEgg,
	874:  &BatSpawnEgg,
	875:  &BeeSpawnEgg,
	876:  &BlazeSpawnEgg,
	877:  &CatSpawnEgg,
	878:  &CaveSpiderSpawnEgg,
	879:  &ChickenSpawnEgg,
	880:  &CodSpawnEgg,
	881:  &CowSpawnEgg,
	882:  &CreeperSpawnEgg,
	883:  &DolphinSpawnEgg,
	884:  &DonkeySpawnEgg,
	885:  &DrownedSpawnEgg,
	886:  &ElderGuardianSpawnEgg,
	887:  &EndermanSpawnEgg,
	888:  &EndermiteSpawnEgg,
	889:  &EvokerSpawnEgg,
	890:  &FoxSpawnEgg,
	891:  &GhastSpawnEgg,
	892:  &GlowSquidSpawnEgg,
	893:  &GoatSpawnEgg,
	894:  &GuardianSpawnEgg,
	895:  &HoglinSpawnEgg,
	896:  &HorseSpawnEgg,
	897:  &HuskSpawnEgg,
	898:  &LlamaSpawnEgg,
	899:  &MagmaCubeSpawnEgg,
	900:  &MooshroomSpawnEgg,
	901:  &MuleSpawnEgg,
	902:  &OcelotSpawnEgg,
	903:  &PandaSpawnEgg,
	904:  &ParrotSpawnEgg,
	905:  &PhantomSpawnEgg,
	906:  &PigSpawnEgg,
	907:  &PiglinSpawnEgg,
	908:  &PiglinBruteSpawnEgg,
	909:  &PillagerSpawnEgg,
	910:  &PolarBearSpawnEgg,
	911:  &PufferfishSpawnEgg,
	912:  &RabbitSpawnEgg,
	913:  &RavagerSpawnEgg,
	914:  &SalmonSpawnEgg,
	915:  &SheepSpawnEgg,
	916:  &ShulkerSpawnEgg,
	917:  &SilverfishSpawnEgg,
	918:  &SkeletonSpawnEgg,
	919:  &SkeletonHorseSpawnEgg,
	920:  &SlimeSpawnEgg,
	921:  &SpiderSpawnEgg,
	922:  &SquidSpawnEgg,
	923:  &StraySpawnEgg,
	924:  &StriderSpawnEgg,
	925:  &TraderLlamaSpawnEgg,
	926:  &TropicalFishSpawnEgg,
	927:  &TurtleSpawnEgg,
	928:  &VexSpawnEgg,
	929:  &VillagerSpawnEgg,
	930:  &VindicatorSpawnEgg,
	931:  &WanderingTraderSpawnEgg,
	932:  &WitchSpawnEgg,
	933:  &WitherSkeletonSpawnEgg,
	934:  &WolfSpawnEgg,
	935:  &ZoglinSpawnEgg,
	936:  &ZombieSpawnEgg,
	937:  &ZombieHorseSpawnEgg,
	938:  &ZombieVillagerSpawnEgg,
	939:  &ZombifiedPiglinSpawnEgg,
	940:  &ExperienceBottle,
	941:  &FireCharge,
	942:  &WritableBook,
	943:  &WrittenBook,
	944:  &ItemFrame,
	945:  &GlowItemFrame,
	946:  &FlowerPot,
	947:  &Carrot,
	948:  &Potato,
	949:  &BakedPotato,
	950:  &PoisonousPotato,
	951:  &Map,
	952:  &GoldenCarrot,
	953:  &SkeletonSkull,
	954:  &WitherSkeletonSkull,
	955:  &PlayerHead,
	956:  &ZombieHead,
	957:  &CreeperHead,
	958:  &DragonHead,
	959:  &NetherStar,
	960:  &PumpkinPie,
	961:  &FireworkRocket,
	962:  &FireworkStar,
	963:  &EnchantedBook,
	964:  &NetherBrick,
	965:  &PrismarineShard,
	966:  &PrismarineCrystals,
	967:  &Rabbit,
	968:  &CookedRabbit,
	969:  &RabbitStew,
	970:  &RabbitFoot,
	971:  &RabbitHide,
	972:  &ArmorStand,
	973:  &IronHorseArmor,
	974:  &GoldenHorseArmor,
	975:  &DiamondHorseArmor,
	976:  &LeatherHorseArmor,
	977:  &Lead,
	978:  &NameTag,
	979:  &CommandBlockMinecart,
	980:  &Mutton,
	981:  &CookedMutton,
	982:  &WhiteBanner,
	983:  &OrangeBanner,
	984:  &MagentaBanner,
	985:  &LightBlueBanner,
	986:  &YellowBanner,
	987:  &LimeBanner,
	988:  &PinkBanner,
	989:  &GrayBanner,
	990:  &LightGrayBanner,
	991:  &CyanBanner,
	992:  &PurpleBanner,
	993:  &BlueBanner,
	994:  &BrownBanner,
	995:  &GreenBanner,
	996:  &RedBanner,
	997:  &BlackBanner,
	998:  &EndCrystal,
	999:  &ChorusFruit,
	1000: &PoppedChorusFruit,
	1001: &Beetroot,
	1002: &BeetrootSeeds,
	1003: &BeetrootSoup,
	1004: &DragonBreath,
	1005: &SplashPotion,
	1006: &SpectralArrow,
	1007: &TippedArrow,
	1008: &LingeringPotion,
	1009: &Shield,
	1010: &TotemOfUndying,
	1011: &ShulkerShell,
	1012: &IronNugget,
	1013: &KnowledgeBook,
	1014: &DebugStick,
	1015: &MusicDisc13,
	1016: &MusicDiscCat,
	1017: &MusicDiscBlocks,
	1018: &MusicDiscChirp,
	1019: &MusicDiscFar,
	1020: &MusicDiscMall,
	1021: &MusicDiscMellohi,
	1022: &MusicDiscStal,
	1023: &MusicDiscStrad,
	1024: &MusicDiscWard,
	1025: &MusicDisc11,
	1026: &MusicDiscWait,
	1027: &MusicDiscPigstep,
	1028: &Trident,
	1029: &PhantomMembrane,
	1030: &NautilusShell,
	1031: &HeartOfTheSea,
	1032: &Crossbow,
	1033: &SuspiciousStew,
	1034: &Loom,
	1035: &FlowerBannerPattern,
	1036: &CreeperBannerPattern,
	1037: &SkullBannerPattern,
	1038: &MojangBannerPattern,
	1039: &GlobeBannerPattern,
	1040: &PiglinBannerPattern,
	1041: &Composter,
	1042: &Barrel,
	1043: &Smoker,
	1044: &BlastFurnace,
	1045: &CartographyTable,
	1046: &FletchingTable,
	1047: &Grindstone,
	1048: &SmithingTable,
	1049: &Stonecutter,
	1050: &Bell,
	1051: &Lantern,
	1052: &SoulLantern,
	1053: &SweetBerries,
	1054: &GlowBerries,
	1055: &Campfire,
	1056: &SoulCampfire,
	1057: &Shroomlight,
	1058: &Honeycomb,
	1059: &BeeNest,
	1060: &Beehive,
	1061: &HoneyBottle,
	1062: &HoneycombBlock,
	1063: &Lodestone,
	1064: &CryingObsidian,
	1065: &Blackstone,
	1066: &BlackstoneSlab,
	1067: &BlackstoneStairs,
	1068: &GildedBlackstone,
	1069: &PolishedBlackstone,
	1070: &PolishedBlackstoneSlab,
	1071: &PolishedBlackstoneStairs,
	1072: &ChiseledPolishedBlackstone,
	1073: &PolishedBlackstoneBricks,
	1074: &PolishedBlackstoneBrickSlab,
	1075: &PolishedBlackstoneBrickStairs,
	1076: &CrackedPolishedBlackstoneBricks,
	1077: &RespawnAnchor,
	1078: &Candle,
	1079: &WhiteCandle,
	1080: &OrangeCandle,
	1081: &MagentaCandle,
	1082: &LightBlueCandle,
	1083: &YellowCandle,
	1084: &LimeCandle,
	1085: &PinkCandle,
	1086: &GrayCandle,
	1087: &LightGrayCandle,
	1088: &CyanCandle,
	1089: &PurpleCandle,
	1090: &BlueCandle,
	1091: &BrownCandle,
	1092: &GreenCandle,
	1093: &RedCandle,
	1094: &BlackCandle,
	1095: &SmallAmethystBud,
	1096: &MediumAmethystBud,
	1097: &LargeAmethystBud,
	1098: &AmethystCluster,
	1099: &PointedDripstone,
}

// ByName is an index of minecraft items by their name.
var ByName = map[string]*Item{
	"stone":                              &Stone,
	"granite":                            &Granite,
	"polished_granite":                   &PolishedGranite,
	"diorite":                            &Diorite,
	"polished_diorite":                   &PolishedDiorite,
	"andesite":                           &Andesite,
	"polished_andesite":                  &PolishedAndesite,
	"deepslate":                          &Deepslate,
	"cobbled_deepslate":                  &CobbledDeepslate,
	"polished_deepslate":                 &PolishedDeepslate,
	"calcite":                            &Calcite,
	"tuff":                               &Tuff,
	"dripstone_block":                    &DripstoneBlock,
	"grass_block":                        &GrassBlock,
	"dirt":                               &Dirt,
	"coarse_dirt":                        &CoarseDirt,
	"podzol":                             &Podzol,
	"rooted_dirt":                        &RootedDirt,
	"crimson_nylium":                     &CrimsonNylium,
	"warped_nylium":                      &WarpedNylium,
	"cobblestone":                        &Cobblestone,
	"oak_planks":                         &OakPlanks,
	"spruce_planks":                      &SprucePlanks,
	"birch_planks":                       &BirchPlanks,
	"jungle_planks":                      &JunglePlanks,
	"acacia_planks":                      &AcaciaPlanks,
	"dark_oak_planks":                    &DarkOakPlanks,
	"crimson_planks":                     &CrimsonPlanks,
	"warped_planks":                      &WarpedPlanks,
	"oak_sapling":                        &OakSapling,
	"spruce_sapling":                     &SpruceSapling,
	"birch_sapling":                      &BirchSapling,
	"jungle_sapling":                     &JungleSapling,
	"acacia_sapling":                     &AcaciaSapling,
	"dark_oak_sapling":                   &DarkOakSapling,
	"bedrock":                            &Bedrock,
	"sand":                               &Sand,
	"red_sand":                           &RedSand,
	"gravel":                             &Gravel,
	"coal_ore":                           &CoalOre,
	"deepslate_coal_ore":                 &DeepslateCoalOre,
	"iron_ore":                           &IronOre,
	"deepslate_iron_ore":                 &DeepslateIronOre,
	"copper_ore":                         &CopperOre,
	"deepslate_copper_ore":               &DeepslateCopperOre,
	"gold_ore":                           &GoldOre,
	"deepslate_gold_ore":                 &DeepslateGoldOre,
	"redstone_ore":                       &RedstoneOre,
	"deepslate_redstone_ore":             &DeepslateRedstoneOre,
	"emerald_ore":                        &EmeraldOre,
	"deepslate_emerald_ore":              &DeepslateEmeraldOre,
	"lapis_ore":                          &LapisOre,
	"deepslate_lapis_ore":                &DeepslateLapisOre,
	"diamond_ore":                        &DiamondOre,
	"deepslate_diamond_ore":              &DeepslateDiamondOre,
	"nether_gold_ore":                    &NetherGoldOre,
	"nether_quartz_ore":                  &NetherQuartzOre,
	"ancient_debris":                     &AncientDebris,
	"coal_block":                         &CoalBlock,
	"raw_iron_block":                     &RawIronBlock,
	"raw_copper_block":                   &RawCopperBlock,
	"raw_gold_block":                     &RawGoldBlock,
	"amethyst_block":                     &AmethystBlock,
	"budding_amethyst":                   &BuddingAmethyst,
	"iron_block":                         &IronBlock,
	"copper_block":                       &CopperBlock,
	"gold_block":                         &GoldBlock,
	"diamond_block":                      &DiamondBlock,
	"netherite_block":                    &NetheriteBlock,
	"exposed_copper":                     &ExposedCopper,
	"weathered_copper":                   &WeatheredCopper,
	"oxidized_copper":                    &OxidizedCopper,
	"cut_copper":                         &CutCopper,
	"exposed_cut_copper":                 &ExposedCutCopper,
	"weathered_cut_copper":               &WeatheredCutCopper,
	"oxidized_cut_copper":                &OxidizedCutCopper,
	"cut_copper_stairs":                  &CutCopperStairs,
	"exposed_cut_copper_stairs":          &ExposedCutCopperStairs,
	"weathered_cut_copper_stairs":        &WeatheredCutCopperStairs,
	"oxidized_cut_copper_stairs":         &OxidizedCutCopperStairs,
	"cut_copper_slab":                    &CutCopperSlab,
	"exposed_cut_copper_slab":            &ExposedCutCopperSlab,
	"weathered_cut_copper_slab":          &WeatheredCutCopperSlab,
	"oxidized_cut_copper_slab":           &OxidizedCutCopperSlab,
	"waxed_copper_block":                 &WaxedCopperBlock,
	"waxed_exposed_copper":               &WaxedExposedCopper,
	"waxed_weathered_copper":             &WaxedWeatheredCopper,
	"waxed_oxidized_copper":              &WaxedOxidizedCopper,
	"waxed_cut_copper":                   &WaxedCutCopper,
	"waxed_exposed_cut_copper":           &WaxedExposedCutCopper,
	"waxed_weathered_cut_copper":         &WaxedWeatheredCutCopper,
	"waxed_oxidized_cut_copper":          &WaxedOxidizedCutCopper,
	"waxed_cut_copper_stairs":            &WaxedCutCopperStairs,
	"waxed_exposed_cut_copper_stairs":    &WaxedExposedCutCopperStairs,
	"waxed_weathered_cut_copper_stairs":  &WaxedWeatheredCutCopperStairs,
	"waxed_oxidized_cut_copper_stairs":   &WaxedOxidizedCutCopperStairs,
	"waxed_cut_copper_slab":              &WaxedCutCopperSlab,
	"waxed_exposed_cut_copper_slab":      &WaxedExposedCutCopperSlab,
	"waxed_weathered_cut_copper_slab":    &WaxedWeatheredCutCopperSlab,
	"waxed_oxidized_cut_copper_slab":     &WaxedOxidizedCutCopperSlab,
	"oak_log":                            &OakLog,
	"spruce_log":                         &SpruceLog,
	"birch_log":                          &BirchLog,
	"jungle_log":                         &JungleLog,
	"acacia_log":                         &AcaciaLog,
	"dark_oak_log":                       &DarkOakLog,
	"crimson_stem":                       &CrimsonStem,
	"warped_stem":                        &WarpedStem,
	"stripped_oak_log":                   &StrippedOakLog,
	"stripped_spruce_log":                &StrippedSpruceLog,
	"stripped_birch_log":                 &StrippedBirchLog,
	"stripped_jungle_log":                &StrippedJungleLog,
	"stripped_acacia_log":                &StrippedAcaciaLog,
	"stripped_dark_oak_log":              &StrippedDarkOakLog,
	"stripped_crimson_stem":              &StrippedCrimsonStem,
	"stripped_warped_stem":               &StrippedWarpedStem,
	"stripped_oak_wood":                  &StrippedOakWood,
	"stripped_spruce_wood":               &StrippedSpruceWood,
	"stripped_birch_wood":                &StrippedBirchWood,
	"stripped_jungle_wood":               &StrippedJungleWood,
	"stripped_acacia_wood":               &StrippedAcaciaWood,
	"stripped_dark_oak_wood":             &StrippedDarkOakWood,
	"stripped_crimson_hyphae":            &StrippedCrimsonHyphae,
	"stripped_warped_hyphae":             &StrippedWarpedHyphae,
	"oak_wood":                           &OakWood,
	"spruce_wood":                        &SpruceWood,
	"birch_wood":                         &BirchWood,
	"jungle_wood":                        &JungleWood,
	"acacia_wood":                        &AcaciaWood,
	"dark_oak_wood":                      &DarkOakWood,
	"crimson_hyphae":                     &CrimsonHyphae,
	"warped_hyphae":                      &WarpedHyphae,
	"oak_leaves":                         &OakLeaves,
	"spruce_leaves":                      &SpruceLeaves,
	"birch_leaves":                       &BirchLeaves,
	"jungle_leaves":                      &JungleLeaves,
	"acacia_leaves":                      &AcaciaLeaves,
	"dark_oak_leaves":                    &DarkOakLeaves,
	"azalea_leaves":                      &AzaleaLeaves,
	"flowering_azalea_leaves":            &FloweringAzaleaLeaves,
	"sponge":                             &Sponge,
	"wet_sponge":                         &WetSponge,
	"glass":                              &Glass,
	"tinted_glass":                       &TintedGlass,
	"lapis_block":                        &LapisBlock,
	"sandstone":                          &Sandstone,
	"chiseled_sandstone":                 &ChiseledSandstone,
	"cut_sandstone":                      &CutSandstone,
	"cobweb":                             &Cobweb,
	"grass":                              &Grass,
	"fern":                               &Fern,
	"azalea":                             &Azalea,
	"flowering_azalea":                   &FloweringAzalea,
	"dead_bush":                          &DeadBush,
	"seagrass":                           &Seagrass,
	"sea_pickle":                         &SeaPickle,
	"white_wool":                         &WhiteWool,
	"orange_wool":                        &OrangeWool,
	"magenta_wool":                       &MagentaWool,
	"light_blue_wool":                    &LightBlueWool,
	"yellow_wool":                        &YellowWool,
	"lime_wool":                          &LimeWool,
	"pink_wool":                          &PinkWool,
	"gray_wool":                          &GrayWool,
	"light_gray_wool":                    &LightGrayWool,
	"cyan_wool":                          &CyanWool,
	"purple_wool":                        &PurpleWool,
	"blue_wool":                          &BlueWool,
	"brown_wool":                         &BrownWool,
	"green_wool":                         &GreenWool,
	"red_wool":                           &RedWool,
	"black_wool":                         &BlackWool,
	"dandelion":                          &Dandelion,
	"poppy":                              &Poppy,
	"blue_orchid":                        &BlueOrchid,
	"allium":                             &Allium,
	"azure_bluet":                        &AzureBluet,
	"red_tulip":                          &RedTulip,
	"orange_tulip":                       &OrangeTulip,
	"white_tulip":                        &WhiteTulip,
	"pink_tulip":                         &PinkTulip,
	"oxeye_daisy":                        &OxeyeDaisy,
	"cornflower":                         &Cornflower,
	"lily_of_the_valley":                 &LilyOfTheValley,
	"wither_rose":                        &WitherRose,
	"spore_blossom":                      &SporeBlossom,
	"brown_mushroom":                     &BrownMushroom,
	"red_mushroom":                       &RedMushroom,
	"crimson_fungus":                     &CrimsonFungus,
	"warped_fungus":                      &WarpedFungus,
	"crimson_roots":                      &CrimsonRoots,
	"warped_roots":                       &WarpedRoots,
	"nether_sprouts":                     &NetherSprouts,
	"weeping_vines":                      &WeepingVines,
	"twisting_vines":                     &TwistingVines,
	"sugar_cane":                         &SugarCane,
	"kelp":                               &Kelp,
	"moss_carpet":                        &MossCarpet,
	"moss_block":                         &MossBlock,
	"hanging_roots":                      &HangingRoots,
	"big_dripleaf":                       &BigDripleaf,
	"small_dripleaf":                     &SmallDripleaf,
	"bamboo":                             &Bamboo,
	"oak_slab":                           &OakSlab,
	"spruce_slab":                        &SpruceSlab,
	"birch_slab":                         &BirchSlab,
	"jungle_slab":                        &JungleSlab,
	"acacia_slab":                        &AcaciaSlab,
	"dark_oak_slab":                      &DarkOakSlab,
	"crimson_slab":                       &CrimsonSlab,
	"warped_slab":                        &WarpedSlab,
	"stone_slab":                         &StoneSlab,
	"smooth_stone_slab":                  &SmoothStoneSlab,
	"sandstone_slab":                     &SandstoneSlab,
	"cut_sandstone_slab":                 &CutSandstoneSlab,
	"petrified_oak_slab":                 &PetrifiedOakSlab,
	"cobblestone_slab":                   &CobblestoneSlab,
	"brick_slab":                         &BrickSlab,
	"stone_brick_slab":                   &StoneBrickSlab,
	"nether_brick_slab":                  &NetherBrickSlab,
	"quartz_slab":                        &QuartzSlab,
	"red_sandstone_slab":                 &RedSandstoneSlab,
	"cut_red_sandstone_slab":             &CutRedSandstoneSlab,
	"purpur_slab":                        &PurpurSlab,
	"prismarine_slab":                    &PrismarineSlab,
	"prismarine_brick_slab":              &PrismarineBrickSlab,
	"dark_prismarine_slab":               &DarkPrismarineSlab,
	"smooth_quartz":                      &SmoothQuartz,
	"smooth_red_sandstone":               &SmoothRedSandstone,
	"smooth_sandstone":                   &SmoothSandstone,
	"smooth_stone":                       &SmoothStone,
	"bricks":                             &Bricks,
	"bookshelf":                          &Bookshelf,
	"mossy_cobblestone":                  &MossyCobblestone,
	"obsidian":                           &Obsidian,
	"torch":                              &Torch,
	"end_rod":                            &EndRod,
	"chorus_plant":                       &ChorusPlant,
	"chorus_flower":                      &ChorusFlower,
	"purpur_block":                       &PurpurBlock,
	"purpur_pillar":                      &PurpurPillar,
	"purpur_stairs":                      &PurpurStairs,
	"spawner":                            &Spawner,
	"oak_stairs":                         &OakStairs,
	"chest":                              &Chest,
	"crafting_table":                     &CraftingTable,
	"farmland":                           &Farmland,
	"furnace":                            &Furnace,
	"ladder":                             &Ladder,
	"cobblestone_stairs":                 &CobblestoneStairs,
	"snow":                               &Snow,
	"ice":                                &Ice,
	"snow_block":                         &SnowBlock,
	"cactus":                             &Cactus,
	"clay":                               &Clay,
	"jukebox":                            &Jukebox,
	"oak_fence":                          &OakFence,
	"spruce_fence":                       &SpruceFence,
	"birch_fence":                        &BirchFence,
	"jungle_fence":                       &JungleFence,
	"acacia_fence":                       &AcaciaFence,
	"dark_oak_fence":                     &DarkOakFence,
	"crimson_fence":                      &CrimsonFence,
	"warped_fence":                       &WarpedFence,
	"pumpkin":                            &Pumpkin,
	"carved_pumpkin":                     &CarvedPumpkin,
	"jack_o_lantern":                     &JackOLantern,
	"netherrack":                         &Netherrack,
	"soul_sand":                          &SoulSand,
	"soul_soil":                          &SoulSoil,
	"basalt":                             &Basalt,
	"polished_basalt":                    &PolishedBasalt,
	"smooth_basalt":                      &SmoothBasalt,
	"soul_torch":                         &SoulTorch,
	"glowstone":                          &Glowstone,
	"infested_stone":                     &InfestedStone,
	"infested_cobblestone":               &InfestedCobblestone,
	"infested_stone_bricks":              &InfestedStoneBricks,
	"infested_mossy_stone_bricks":        &InfestedMossyStoneBricks,
	"infested_cracked_stone_bricks":      &InfestedCrackedStoneBricks,
	"infested_chiseled_stone_bricks":     &InfestedChiseledStoneBricks,
	"infested_deepslate":                 &InfestedDeepslate,
	"stone_bricks":                       &StoneBricks,
	"mossy_stone_bricks":                 &MossyStoneBricks,
	"cracked_stone_bricks":               &CrackedStoneBricks,
	"chiseled_stone_bricks":              &ChiseledStoneBricks,
	"deepslate_bricks":                   &DeepslateBricks,
	"cracked_deepslate_bricks":           &CrackedDeepslateBricks,
	"deepslate_tiles":                    &DeepslateTiles,
	"cracked_deepslate_tiles":            &CrackedDeepslateTiles,
	"chiseled_deepslate":                 &ChiseledDeepslate,
	"brown_mushroom_block":               &BrownMushroomBlock,
	"red_mushroom_block":                 &RedMushroomBlock,
	"mushroom_stem":                      &MushroomStem,
	"iron_bars":                          &IronBars,
	"chain":                              &Chain,
	"glass_pane":                         &GlassPane,
	"melon":                              &Melon,
	"vine":                               &Vine,
	"glow_lichen":                        &GlowLichen,
	"brick_stairs":                       &BrickStairs,
	"stone_brick_stairs":                 &StoneBrickStairs,
	"mycelium":                           &Mycelium,
	"lily_pad":                           &LilyPad,
	"nether_bricks":                      &NetherBricks,
	"cracked_nether_bricks":              &CrackedNetherBricks,
	"chiseled_nether_bricks":             &ChiseledNetherBricks,
	"nether_brick_fence":                 &NetherBrickFence,
	"nether_brick_stairs":                &NetherBrickStairs,
	"enchanting_table":                   &EnchantingTable,
	"end_portal_frame":                   &EndPortalFrame,
	"end_stone":                          &EndStone,
	"end_stone_bricks":                   &EndStoneBricks,
	"dragon_egg":                         &DragonEgg,
	"sandstone_stairs":                   &SandstoneStairs,
	"ender_chest":                        &EnderChest,
	"emerald_block":                      &EmeraldBlock,
	"spruce_stairs":                      &SpruceStairs,
	"birch_stairs":                       &BirchStairs,
	"jungle_stairs":                      &JungleStairs,
	"crimson_stairs":                     &CrimsonStairs,
	"warped_stairs":                      &WarpedStairs,
	"command_block":                      &CommandBlock,
	"beacon":                             &Beacon,
	"cobblestone_wall":                   &CobblestoneWall,
	"mossy_cobblestone_wall":             &MossyCobblestoneWall,
	"brick_wall":                         &BrickWall,
	"prismarine_wall":                    &PrismarineWall,
	"red_sandstone_wall":                 &RedSandstoneWall,
	"mossy_stone_brick_wall":             &MossyStoneBrickWall,
	"granite_wall":                       &GraniteWall,
	"stone_brick_wall":                   &StoneBrickWall,
	"nether_brick_wall":                  &NetherBrickWall,
	"andesite_wall":                      &AndesiteWall,
	"red_nether_brick_wall":              &RedNetherBrickWall,
	"sandstone_wall":                     &SandstoneWall,
	"end_stone_brick_wall":               &EndStoneBrickWall,
	"diorite_wall":                       &DioriteWall,
	"blackstone_wall":                    &BlackstoneWall,
	"polished_blackstone_wall":           &PolishedBlackstoneWall,
	"polished_blackstone_brick_wall":     &PolishedBlackstoneBrickWall,
	"cobbled_deepslate_wall":             &CobbledDeepslateWall,
	"polished_deepslate_wall":            &PolishedDeepslateWall,
	"deepslate_brick_wall":               &DeepslateBrickWall,
	"deepslate_tile_wall":                &DeepslateTileWall,
	"anvil":                              &Anvil,
	"chipped_anvil":                      &ChippedAnvil,
	"damaged_anvil":                      &DamagedAnvil,
	"chiseled_quartz_block":              &ChiseledQuartzBlock,
	"quartz_block":                       &QuartzBlock,
	"quartz_bricks":                      &QuartzBricks,
	"quartz_pillar":                      &QuartzPillar,
	"quartz_stairs":                      &QuartzStairs,
	"white_terracotta":                   &WhiteTerracotta,
	"orange_terracotta":                  &OrangeTerracotta,
	"magenta_terracotta":                 &MagentaTerracotta,
	"light_blue_terracotta":              &LightBlueTerracotta,
	"yellow_terracotta":                  &YellowTerracotta,
	"lime_terracotta":                    &LimeTerracotta,
	"pink_terracotta":                    &PinkTerracotta,
	"gray_terracotta":                    &GrayTerracotta,
	"light_gray_terracotta":              &LightGrayTerracotta,
	"cyan_terracotta":                    &CyanTerracotta,
	"purple_terracotta":                  &PurpleTerracotta,
	"blue_terracotta":                    &BlueTerracotta,
	"brown_terracotta":                   &BrownTerracotta,
	"green_terracotta":                   &GreenTerracotta,
	"red_terracotta":                     &RedTerracotta,
	"black_terracotta":                   &BlackTerracotta,
	"barrier":                            &Barrier,
	"light":                              &Light,
	"hay_block":                          &HayBlock,
	"white_carpet":                       &WhiteCarpet,
	"orange_carpet":                      &OrangeCarpet,
	"magenta_carpet":                     &MagentaCarpet,
	"light_blue_carpet":                  &LightBlueCarpet,
	"yellow_carpet":                      &YellowCarpet,
	"lime_carpet":                        &LimeCarpet,
	"pink_carpet":                        &PinkCarpet,
	"gray_carpet":                        &GrayCarpet,
	"light_gray_carpet":                  &LightGrayCarpet,
	"cyan_carpet":                        &CyanCarpet,
	"purple_carpet":                      &PurpleCarpet,
	"blue_carpet":                        &BlueCarpet,
	"brown_carpet":                       &BrownCarpet,
	"green_carpet":                       &GreenCarpet,
	"red_carpet":                         &RedCarpet,
	"black_carpet":                       &BlackCarpet,
	"terracotta":                         &Terracotta,
	"packed_ice":                         &PackedIce,
	"acacia_stairs":                      &AcaciaStairs,
	"dark_oak_stairs":                    &DarkOakStairs,
	"dirt_path":                          &DirtPath,
	"sunflower":                          &Sunflower,
	"lilac":                              &Lilac,
	"rose_bush":                          &RoseBush,
	"peony":                              &Peony,
	"tall_grass":                         &TallGrass,
	"large_fern":                         &LargeFern,
	"white_stained_glass":                &WhiteStainedGlass,
	"orange_stained_glass":               &OrangeStainedGlass,
	"magenta_stained_glass":              &MagentaStainedGlass,
	"light_blue_stained_glass":           &LightBlueStainedGlass,
	"yellow_stained_glass":               &YellowStainedGlass,
	"lime_stained_glass":                 &LimeStainedGlass,
	"pink_stained_glass":                 &PinkStainedGlass,
	"gray_stained_glass":                 &GrayStainedGlass,
	"light_gray_stained_glass":           &LightGrayStainedGlass,
	"cyan_stained_glass":                 &CyanStainedGlass,
	"purple_stained_glass":               &PurpleStainedGlass,
	"blue_stained_glass":                 &BlueStainedGlass,
	"brown_stained_glass":                &BrownStainedGlass,
	"green_stained_glass":                &GreenStainedGlass,
	"red_stained_glass":                  &RedStainedGlass,
	"black_stained_glass":                &BlackStainedGlass,
	"white_stained_glass_pane":           &WhiteStainedGlassPane,
	"orange_stained_glass_pane":          &OrangeStainedGlassPane,
	"magenta_stained_glass_pane":         &MagentaStainedGlassPane,
	"light_blue_stained_glass_pane":      &LightBlueStainedGlassPane,
	"yellow_stained_glass_pane":          &YellowStainedGlassPane,
	"lime_stained_glass_pane":            &LimeStainedGlassPane,
	"pink_stained_glass_pane":            &PinkStainedGlassPane,
	"gray_stained_glass_pane":            &GrayStainedGlassPane,
	"light_gray_stained_glass_pane":      &LightGrayStainedGlassPane,
	"cyan_stained_glass_pane":            &CyanStainedGlassPane,
	"purple_stained_glass_pane":          &PurpleStainedGlassPane,
	"blue_stained_glass_pane":            &BlueStainedGlassPane,
	"brown_stained_glass_pane":           &BrownStainedGlassPane,
	"green_stained_glass_pane":           &GreenStainedGlassPane,
	"red_stained_glass_pane":             &RedStainedGlassPane,
	"black_stained_glass_pane":           &BlackStainedGlassPane,
	"prismarine":                         &Prismarine,
	"prismarine_bricks":                  &PrismarineBricks,
	"dark_prismarine":                    &DarkPrismarine,
	"prismarine_stairs":                  &PrismarineStairs,
	"prismarine_brick_stairs":            &PrismarineBrickStairs,
	"dark_prismarine_stairs":             &DarkPrismarineStairs,
	"sea_lantern":                        &SeaLantern,
	"red_sandstone":                      &RedSandstone,
	"chiseled_red_sandstone":             &ChiseledRedSandstone,
	"cut_red_sandstone":                  &CutRedSandstone,
	"red_sandstone_stairs":               &RedSandstoneStairs,
	"repeating_command_block":            &RepeatingCommandBlock,
	"chain_command_block":                &ChainCommandBlock,
	"magma_block":                        &MagmaBlock,
	"nether_wart_block":                  &NetherWartBlock,
	"warped_wart_block":                  &WarpedWartBlock,
	"red_nether_bricks":                  &RedNetherBricks,
	"bone_block":                         &BoneBlock,
	"structure_void":                     &StructureVoid,
	"shulker_box":                        &ShulkerBox,
	"white_shulker_box":                  &WhiteShulkerBox,
	"orange_shulker_box":                 &OrangeShulkerBox,
	"magenta_shulker_box":                &MagentaShulkerBox,
	"light_blue_shulker_box":             &LightBlueShulkerBox,
	"yellow_shulker_box":                 &YellowShulkerBox,
	"lime_shulker_box":                   &LimeShulkerBox,
	"pink_shulker_box":                   &PinkShulkerBox,
	"gray_shulker_box":                   &GrayShulkerBox,
	"light_gray_shulker_box":             &LightGrayShulkerBox,
	"cyan_shulker_box":                   &CyanShulkerBox,
	"purple_shulker_box":                 &PurpleShulkerBox,
	"blue_shulker_box":                   &BlueShulkerBox,
	"brown_shulker_box":                  &BrownShulkerBox,
	"green_shulker_box":                  &GreenShulkerBox,
	"red_shulker_box":                    &RedShulkerBox,
	"black_shulker_box":                  &BlackShulkerBox,
	"white_glazed_terracotta":            &WhiteGlazedTerracotta,
	"orange_glazed_terracotta":           &OrangeGlazedTerracotta,
	"magenta_glazed_terracotta":          &MagentaGlazedTerracotta,
	"light_blue_glazed_terracotta":       &LightBlueGlazedTerracotta,
	"yellow_glazed_terracotta":           &YellowGlazedTerracotta,
	"lime_glazed_terracotta":             &LimeGlazedTerracotta,
	"pink_glazed_terracotta":             &PinkGlazedTerracotta,
	"gray_glazed_terracotta":             &GrayGlazedTerracotta,
	"light_gray_glazed_terracotta":       &LightGrayGlazedTerracotta,
	"cyan_glazed_terracotta":             &CyanGlazedTerracotta,
	"purple_glazed_terracotta":           &PurpleGlazedTerracotta,
	"blue_glazed_terracotta":             &BlueGlazedTerracotta,
	"brown_glazed_terracotta":            &BrownGlazedTerracotta,
	"green_glazed_terracotta":            &GreenGlazedTerracotta,
	"red_glazed_terracotta":              &RedGlazedTerracotta,
	"black_glazed_terracotta":            &BlackGlazedTerracotta,
	"white_concrete":                     &WhiteConcrete,
	"orange_concrete":                    &OrangeConcrete,
	"magenta_concrete":                   &MagentaConcrete,
	"light_blue_concrete":                &LightBlueConcrete,
	"yellow_concrete":                    &YellowConcrete,
	"lime_concrete":                      &LimeConcrete,
	"pink_concrete":                      &PinkConcrete,
	"gray_concrete":                      &GrayConcrete,
	"light_gray_concrete":                &LightGrayConcrete,
	"cyan_concrete":                      &CyanConcrete,
	"purple_concrete":                    &PurpleConcrete,
	"blue_concrete":                      &BlueConcrete,
	"brown_concrete":                     &BrownConcrete,
	"green_concrete":                     &GreenConcrete,
	"red_concrete":                       &RedConcrete,
	"black_concrete":                     &BlackConcrete,
	"white_concrete_powder":              &WhiteConcretePowder,
	"orange_concrete_powder":             &OrangeConcretePowder,
	"magenta_concrete_powder":            &MagentaConcretePowder,
	"light_blue_concrete_powder":         &LightBlueConcretePowder,
	"yellow_concrete_powder":             &YellowConcretePowder,
	"lime_concrete_powder":               &LimeConcretePowder,
	"pink_concrete_powder":               &PinkConcretePowder,
	"gray_concrete_powder":               &GrayConcretePowder,
	"light_gray_concrete_powder":         &LightGrayConcretePowder,
	"cyan_concrete_powder":               &CyanConcretePowder,
	"purple_concrete_powder":             &PurpleConcretePowder,
	"blue_concrete_powder":               &BlueConcretePowder,
	"brown_concrete_powder":              &BrownConcretePowder,
	"green_concrete_powder":              &GreenConcretePowder,
	"red_concrete_powder":                &RedConcretePowder,
	"black_concrete_powder":              &BlackConcretePowder,
	"turtle_egg":                         &TurtleEgg,
	"dead_tube_coral_block":              &DeadTubeCoralBlock,
	"dead_brain_coral_block":             &DeadBrainCoralBlock,
	"dead_bubble_coral_block":            &DeadBubbleCoralBlock,
	"dead_fire_coral_block":              &DeadFireCoralBlock,
	"dead_horn_coral_block":              &DeadHornCoralBlock,
	"tube_coral_block":                   &TubeCoralBlock,
	"brain_coral_block":                  &BrainCoralBlock,
	"bubble_coral_block":                 &BubbleCoralBlock,
	"fire_coral_block":                   &FireCoralBlock,
	"horn_coral_block":                   &HornCoralBlock,
	"tube_coral":                         &TubeCoral,
	"brain_coral":                        &BrainCoral,
	"bubble_coral":                       &BubbleCoral,
	"fire_coral":                         &FireCoral,
	"horn_coral":                         &HornCoral,
	"dead_brain_coral":                   &DeadBrainCoral,
	"dead_bubble_coral":                  &DeadBubbleCoral,
	"dead_fire_coral":                    &DeadFireCoral,
	"dead_horn_coral":                    &DeadHornCoral,
	"dead_tube_coral":                    &DeadTubeCoral,
	"tube_coral_fan":                     &TubeCoralFan,
	"brain_coral_fan":                    &BrainCoralFan,
	"bubble_coral_fan":                   &BubbleCoralFan,
	"fire_coral_fan":                     &FireCoralFan,
	"horn_coral_fan":                     &HornCoralFan,
	"dead_tube_coral_fan":                &DeadTubeCoralFan,
	"dead_brain_coral_fan":               &DeadBrainCoralFan,
	"dead_bubble_coral_fan":              &DeadBubbleCoralFan,
	"dead_fire_coral_fan":                &DeadFireCoralFan,
	"dead_horn_coral_fan":                &DeadHornCoralFan,
	"blue_ice":                           &BlueIce,
	"conduit":                            &Conduit,
	"polished_granite_stairs":            &PolishedGraniteStairs,
	"smooth_red_sandstone_stairs":        &SmoothRedSandstoneStairs,
	"mossy_stone_brick_stairs":           &MossyStoneBrickStairs,
	"polished_diorite_stairs":            &PolishedDioriteStairs,
	"mossy_cobblestone_stairs":           &MossyCobblestoneStairs,
	"end_stone_brick_stairs":             &EndStoneBrickStairs,
	"stone_stairs":                       &StoneStairs,
	"smooth_sandstone_stairs":            &SmoothSandstoneStairs,
	"smooth_quartz_stairs":               &SmoothQuartzStairs,
	"granite_stairs":                     &GraniteStairs,
	"andesite_stairs":                    &AndesiteStairs,
	"red_nether_brick_stairs":            &RedNetherBrickStairs,
	"polished_andesite_stairs":           &PolishedAndesiteStairs,
	"diorite_stairs":                     &DioriteStairs,
	"cobbled_deepslate_stairs":           &CobbledDeepslateStairs,
	"polished_deepslate_stairs":          &PolishedDeepslateStairs,
	"deepslate_brick_stairs":             &DeepslateBrickStairs,
	"deepslate_tile_stairs":              &DeepslateTileStairs,
	"polished_granite_slab":              &PolishedGraniteSlab,
	"smooth_red_sandstone_slab":          &SmoothRedSandstoneSlab,
	"mossy_stone_brick_slab":             &MossyStoneBrickSlab,
	"polished_diorite_slab":              &PolishedDioriteSlab,
	"mossy_cobblestone_slab":             &MossyCobblestoneSlab,
	"end_stone_brick_slab":               &EndStoneBrickSlab,
	"smooth_sandstone_slab":              &SmoothSandstoneSlab,
	"smooth_quartz_slab":                 &SmoothQuartzSlab,
	"granite_slab":                       &GraniteSlab,
	"andesite_slab":                      &AndesiteSlab,
	"red_nether_brick_slab":              &RedNetherBrickSlab,
	"polished_andesite_slab":             &PolishedAndesiteSlab,
	"diorite_slab":                       &DioriteSlab,
	"cobbled_deepslate_slab":             &CobbledDeepslateSlab,
	"polished_deepslate_slab":            &PolishedDeepslateSlab,
	"deepslate_brick_slab":               &DeepslateBrickSlab,
	"deepslate_tile_slab":                &DeepslateTileSlab,
	"scaffolding":                        &Scaffolding,
	"redstone":                           &Redstone,
	"redstone_torch":                     &RedstoneTorch,
	"redstone_block":                     &RedstoneBlock,
	"repeater":                           &Repeater,
	"comparator":                         &Comparator,
	"piston":                             &Piston,
	"sticky_piston":                      &StickyPiston,
	"slime_block":                        &SlimeBlock,
	"honey_block":                        &HoneyBlock,
	"observer":                           &Observer,
	"hopper":                             &Hopper,
	"dispenser":                          &Dispenser,
	"dropper":                            &Dropper,
	"lectern":                            &Lectern,
	"target":                             &Target,
	"lever":                              &Lever,
	"lightning_rod":                      &LightningRod,
	"daylight_detector":                  &DaylightDetector,
	"sculk_sensor":                       &SculkSensor,
	"tripwire_hook":                      &TripwireHook,
	"trapped_chest":                      &TrappedChest,
	"tnt":                                &Tnt,
	"redstone_lamp":                      &RedstoneLamp,
	"note_block":                         &NoteBlock,
	"stone_button":                       &StoneButton,
	"polished_blackstone_button":         &PolishedBlackstoneButton,
	"oak_button":                         &OakButton,
	"spruce_button":                      &SpruceButton,
	"birch_button":                       &BirchButton,
	"jungle_button":                      &JungleButton,
	"acacia_button":                      &AcaciaButton,
	"dark_oak_button":                    &DarkOakButton,
	"crimson_button":                     &CrimsonButton,
	"warped_button":                      &WarpedButton,
	"stone_pressure_plate":               &StonePressurePlate,
	"polished_blackstone_pressure_plate": &PolishedBlackstonePressurePlate,
	"light_weighted_pressure_plate":      &LightWeightedPressurePlate,
	"heavy_weighted_pressure_plate":      &HeavyWeightedPressurePlate,
	"oak_pressure_plate":                 &OakPressurePlate,
	"spruce_pressure_plate":              &SprucePressurePlate,
	"birch_pressure_plate":               &BirchPressurePlate,
	"jungle_pressure_plate":              &JunglePressurePlate,
	"acacia_pressure_plate":              &AcaciaPressurePlate,
	"dark_oak_pressure_plate":            &DarkOakPressurePlate,
	"crimson_pressure_plate":             &CrimsonPressurePlate,
	"warped_pressure_plate":              &WarpedPressurePlate,
	"iron_door":                          &IronDoor,
	"oak_door":                           &OakDoor,
	"spruce_door":                        &SpruceDoor,
	"birch_door":                         &BirchDoor,
	"jungle_door":                        &JungleDoor,
	"acacia_door":                        &AcaciaDoor,
	"dark_oak_door":                      &DarkOakDoor,
	"crimson_door":                       &CrimsonDoor,
	"warped_door":                        &WarpedDoor,
	"iron_trapdoor":                      &IronTrapdoor,
	"oak_trapdoor":                       &OakTrapdoor,
	"spruce_trapdoor":                    &SpruceTrapdoor,
	"birch_trapdoor":                     &BirchTrapdoor,
	"jungle_trapdoor":                    &JungleTrapdoor,
	"acacia_trapdoor":                    &AcaciaTrapdoor,
	"dark_oak_trapdoor":                  &DarkOakTrapdoor,
	"crimson_trapdoor":                   &CrimsonTrapdoor,
	"warped_trapdoor":                    &WarpedTrapdoor,
	"oak_fence_gate":                     &OakFenceGate,
	"spruce_fence_gate":                  &SpruceFenceGate,
	"birch_fence_gate":                   &BirchFenceGate,
	"jungle_fence_gate":                  &JungleFenceGate,
	"acacia_fence_gate":                  &AcaciaFenceGate,
	"dark_oak_fence_gate":                &DarkOakFenceGate,
	"crimson_fence_gate":                 &CrimsonFenceGate,
	"warped_fence_gate":                  &WarpedFenceGate,
	"powered_rail":                       &PoweredRail,
	"detector_rail":                      &DetectorRail,
	"rail":                               &Rail,
	"activator_rail":                     &ActivatorRail,
	"saddle":                             &Saddle,
	"minecart":                           &Minecart,
	"chest_minecart":                     &ChestMinecart,
	"furnace_minecart":                   &FurnaceMinecart,
	"tnt_minecart":                       &TntMinecart,
	"hopper_minecart":                    &HopperMinecart,
	"carrot_on_a_stick":                  &CarrotOnAStick,
	"warped_fungus_on_a_stick":           &WarpedFungusOnAStick,
	"elytra":                             &Elytra,
	"oak_boat":                           &OakBoat,
	"spruce_boat":                        &SpruceBoat,
	"birch_boat":                         &BirchBoat,
	"jungle_boat":                        &JungleBoat,
	"acacia_boat":                        &AcaciaBoat,
	"dark_oak_boat":                      &DarkOakBoat,
	"structure_block":                    &StructureBlock,
	"jigsaw":                             &Jigsaw,
	"turtle_helmet":                      &TurtleHelmet,
	"scute":                              &Scute,
	"flint_and_steel":                    &FlintAndSteel,
	"apple":                              &Apple,
	"bow":                                &Bow,
	"arrow":                              &Arrow,
	"coal":                               &Coal,
	"charcoal":                           &Charcoal,
	"diamond":                            &Diamond,
	"emerald":                            &Emerald,
	"lapis_lazuli":                       &LapisLazuli,
	"quartz":                             &Quartz,
	"amethyst_shard":                     &AmethystShard,
	"raw_iron":                           &RawIron,
	"iron_ingot":                         &IronIngot,
	"raw_copper":                         &RawCopper,
	"copper_ingot":                       &CopperIngot,
	"raw_gold":                           &RawGold,
	"gold_ingot":                         &GoldIngot,
	"netherite_ingot":                    &NetheriteIngot,
	"netherite_scrap":                    &NetheriteScrap,
	"wooden_sword":                       &WoodenSword,
	"wooden_shovel":                      &WoodenShovel,
	"wooden_pickaxe":                     &WoodenPickaxe,
	"wooden_axe":                         &WoodenAxe,
	"wooden_hoe":                         &WoodenHoe,
	"stone_sword":                        &StoneSword,
	"stone_shovel":                       &StoneShovel,
	"stone_pickaxe":                      &StonePickaxe,
	"stone_axe":                          &StoneAxe,
	"stone_hoe":                          &StoneHoe,
	"golden_sword":                       &GoldenSword,
	"golden_shovel":                      &GoldenShovel,
	"golden_pickaxe":                     &GoldenPickaxe,
	"golden_axe":                         &GoldenAxe,
	"golden_hoe":                         &GoldenHoe,
	"iron_sword":                         &IronSword,
	"iron_shovel":                        &IronShovel,
	"iron_pickaxe":                       &IronPickaxe,
	"iron_axe":                           &IronAxe,
	"iron_hoe":                           &IronHoe,
	"diamond_sword":                      &DiamondSword,
	"diamond_shovel":                     &DiamondShovel,
	"diamond_pickaxe":                    &DiamondPickaxe,
	"diamond_axe":                        &DiamondAxe,
	"diamond_hoe":                        &DiamondHoe,
	"netherite_sword":                    &NetheriteSword,
	"netherite_shovel":                   &NetheriteShovel,
	"netherite_pickaxe":                  &NetheritePickaxe,
	"netherite_axe":                      &NetheriteAxe,
	"netherite_hoe":                      &NetheriteHoe,
	"stick":                              &Stick,
	"bowl":                               &Bowl,
	"mushroom_stew":                      &MushroomStew,
	"string":                             &String,
	"feather":                            &Feather,
	"gunpowder":                          &Gunpowder,
	"wheat_seeds":                        &WheatSeeds,
	"wheat":                              &Wheat,
	"bread":                              &Bread,
	"leather_helmet":                     &LeatherHelmet,
	"leather_chestplate":                 &LeatherChestplate,
	"leather_leggings":                   &LeatherLeggings,
	"leather_boots":                      &LeatherBoots,
	"chainmail_helmet":                   &ChainmailHelmet,
	"chainmail_chestplate":               &ChainmailChestplate,
	"chainmail_leggings":                 &ChainmailLeggings,
	"chainmail_boots":                    &ChainmailBoots,
	"iron_helmet":                        &IronHelmet,
	"iron_chestplate":                    &IronChestplate,
	"iron_leggings":                      &IronLeggings,
	"iron_boots":                         &IronBoots,
	"diamond_helmet":                     &DiamondHelmet,
	"diamond_chestplate":                 &DiamondChestplate,
	"diamond_leggings":                   &DiamondLeggings,
	"diamond_boots":                      &DiamondBoots,
	"golden_helmet":                      &GoldenHelmet,
	"golden_chestplate":                  &GoldenChestplate,
	"golden_leggings":                    &GoldenLeggings,
	"golden_boots":                       &GoldenBoots,
	"netherite_helmet":                   &NetheriteHelmet,
	"netherite_chestplate":               &NetheriteChestplate,
	"netherite_leggings":                 &NetheriteLeggings,
	"netherite_boots":                    &NetheriteBoots,
	"flint":                              &Flint,
	"porkchop":                           &Porkchop,
	"cooked_porkchop":                    &CookedPorkchop,
	"painting":                           &Painting,
	"golden_apple":                       &GoldenApple,
	"enchanted_golden_apple":             &EnchantedGoldenApple,
	"oak_sign":                           &OakSign,
	"spruce_sign":                        &SpruceSign,
	"birch_sign":                         &BirchSign,
	"jungle_sign":                        &JungleSign,
	"acacia_sign":                        &AcaciaSign,
	"dark_oak_sign":                      &DarkOakSign,
	"crimson_sign":                       &CrimsonSign,
	"warped_sign":                        &WarpedSign,
	"bucket":                             &Bucket,
	"water_bucket":                       &WaterBucket,
	"lava_bucket":                        &LavaBucket,
	"powder_snow_bucket":                 &PowderSnowBucket,
	"snowball":                           &Snowball,
	"leather":                            &Leather,
	"milk_bucket":                        &MilkBucket,
	"pufferfish_bucket":                  &PufferfishBucket,
	"salmon_bucket":                      &SalmonBucket,
	"cod_bucket":                         &CodBucket,
	"tropical_fish_bucket":               &TropicalFishBucket,
	"axolotl_bucket":                     &AxolotlBucket,
	"brick":                              &Brick,
	"clay_ball":                          &ClayBall,
	"dried_kelp_block":                   &DriedKelpBlock,
	"paper":                              &Paper,
	"book":                               &Book,
	"slime_ball":                         &SlimeBall,
	"egg":                                &Egg,
	"compass":                            &Compass,
	"bundle":                             &Bundle,
	"fishing_rod":                        &FishingRod,
	"clock":                              &Clock,
	"spyglass":                           &Spyglass,
	"glowstone_dust":                     &GlowstoneDust,
	"cod":                                &Cod,
	"salmon":                             &Salmon,
	"tropical_fish":                      &TropicalFish,
	"pufferfish":                         &Pufferfish,
	"cooked_cod":                         &CookedCod,
	"cooked_salmon":                      &CookedSalmon,
	"ink_sac":                            &InkSac,
	"glow_ink_sac":                       &GlowInkSac,
	"cocoa_beans":                        &CocoaBeans,
	"white_dye":                          &WhiteDye,
	"orange_dye":                         &OrangeDye,
	"magenta_dye":                        &MagentaDye,
	"light_blue_dye":                     &LightBlueDye,
	"yellow_dye":                         &YellowDye,
	"lime_dye":                           &LimeDye,
	"pink_dye":                           &PinkDye,
	"gray_dye":                           &GrayDye,
	"light_gray_dye":                     &LightGrayDye,
	"cyan_dye":                           &CyanDye,
	"purple_dye":                         &PurpleDye,
	"blue_dye":                           &BlueDye,
	"brown_dye":                          &BrownDye,
	"green_dye":                          &GreenDye,
	"red_dye":                            &RedDye,
	"black_dye":                          &BlackDye,
	"bone_meal":                          &BoneMeal,
	"bone":                               &Bone,
	"sugar":                              &Sugar,
	"cake":                               &Cake,
	"white_bed":                          &WhiteBed,
	"orange_bed":                         &OrangeBed,
	"magenta_bed":                        &MagentaBed,
	"light_blue_bed":                     &LightBlueBed,
	"yellow_bed":                         &YellowBed,
	"lime_bed":                           &LimeBed,
	"pink_bed":                           &PinkBed,
	"gray_bed":                           &GrayBed,
	"light_gray_bed":                     &LightGrayBed,
	"cyan_bed":                           &CyanBed,
	"purple_bed":                         &PurpleBed,
	"blue_bed":                           &BlueBed,
	"brown_bed":                          &BrownBed,
	"green_bed":                          &GreenBed,
	"red_bed":                            &RedBed,
	"black_bed":                          &BlackBed,
	"cookie":                             &Cookie,
	"filled_map":                         &FilledMap,
	"shears":                             &Shears,
	"melon_slice":                        &MelonSlice,
	"dried_kelp":                         &DriedKelp,
	"pumpkin_seeds":                      &PumpkinSeeds,
	"melon_seeds":                        &MelonSeeds,
	"beef":                               &Beef,
	"cooked_beef":                        &CookedBeef,
	"chicken":                            &Chicken,
	"cooked_chicken":                     &CookedChicken,
	"rotten_flesh":                       &RottenFlesh,
	"ender_pearl":                        &EnderPearl,
	"blaze_rod":                          &BlazeRod,
	"ghast_tear":                         &GhastTear,
	"gold_nugget":                        &GoldNugget,
	"nether_wart":                        &NetherWart,
	"potion":                             &Potion,
	"glass_bottle":                       &GlassBottle,
	"spider_eye":                         &SpiderEye,
	"fermented_spider_eye":               &FermentedSpiderEye,
	"blaze_powder":                       &BlazePowder,
	"magma_cream":                        &MagmaCream,
	"brewing_stand":                      &BrewingStand,
	"cauldron":                           &Cauldron,
	"ender_eye":                          &EnderEye,
	"glistering_melon_slice":             &GlisteringMelonSlice,
	"axolotl_spawn_egg":                  &AxolotlSpawnEgg,
	"bat_spawn_egg":                      &BatSpawnEgg,
	"bee_spawn_egg":                      &BeeSpawnEgg,
	"blaze_spawn_egg":                    &BlazeSpawnEgg,
	"cat_spawn_egg":                      &CatSpawnEgg,
	"cave_spider_spawn_egg":              &CaveSpiderSpawnEgg,
	"chicken_spawn_egg":                  &ChickenSpawnEgg,
	"cod_spawn_egg":                      &CodSpawnEgg,
	"cow_spawn_egg":                      &CowSpawnEgg,
	"creeper_spawn_egg":                  &CreeperSpawnEgg,
	"dolphin_spawn_egg":                  &DolphinSpawnEgg,
	"donkey_spawn_egg":                   &DonkeySpawnEgg,
	"drowned_spawn_egg":                  &DrownedSpawnEgg,
	"elder_guardian_spawn_egg":           &ElderGuardianSpawnEgg,
	"enderman_spawn_egg":                 &EndermanSpawnEgg,
	"endermite_spawn_egg":                &EndermiteSpawnEgg,
	"evoker_spawn_egg":                   &EvokerSpawnEgg,
	"fox_spawn_egg":                      &FoxSpawnEgg,
	"ghast_spawn_egg":                    &GhastSpawnEgg,
	"glow_squid_spawn_egg":               &GlowSquidSpawnEgg,
	"goat_spawn_egg":                     &GoatSpawnEgg,
	"guardian_spawn_egg":                 &GuardianSpawnEgg,
	"hoglin_spawn_egg":                   &HoglinSpawnEgg,
	"horse_spawn_egg":                    &HorseSpawnEgg,
	"husk_spawn_egg":                     &HuskSpawnEgg,
	"llama_spawn_egg":                    &LlamaSpawnEgg,
	"magma_cube_spawn_egg":               &MagmaCubeSpawnEgg,
	"mooshroom_spawn_egg":                &MooshroomSpawnEgg,
	"mule_spawn_egg":                     &MuleSpawnEgg,
	"ocelot_spawn_egg":                   &OcelotSpawnEgg,
	"panda_spawn_egg":                    &PandaSpawnEgg,
	"parrot_spawn_egg":                   &ParrotSpawnEgg,
	"phantom_spawn_egg":                  &PhantomSpawnEgg,
	"pig_spawn_egg":                      &PigSpawnEgg,
	"piglin_spawn_egg":                   &PiglinSpawnEgg,
	"piglin_brute_spawn_egg":             &PiglinBruteSpawnEgg,
	"pillager_spawn_egg":                 &PillagerSpawnEgg,
	"polar_bear_spawn_egg":               &PolarBearSpawnEgg,
	"pufferfish_spawn_egg":               &PufferfishSpawnEgg,
	"rabbit_spawn_egg":                   &RabbitSpawnEgg,
	"ravager_spawn_egg":                  &RavagerSpawnEgg,
	"salmon_spawn_egg":                   &SalmonSpawnEgg,
	"sheep_spawn_egg":                    &SheepSpawnEgg,
	"shulker_spawn_egg":                  &ShulkerSpawnEgg,
	"silverfish_spawn_egg":               &SilverfishSpawnEgg,
	"skeleton_spawn_egg":                 &SkeletonSpawnEgg,
	"skeleton_horse_spawn_egg":           &SkeletonHorseSpawnEgg,
	"slime_spawn_egg":                    &SlimeSpawnEgg,
	"spider_spawn_egg":                   &SpiderSpawnEgg,
	"squid_spawn_egg":                    &SquidSpawnEgg,
	"stray_spawn_egg":                    &StraySpawnEgg,
	"strider_spawn_egg":                  &StriderSpawnEgg,
	"trader_llama_spawn_egg":             &TraderLlamaSpawnEgg,
	"tropical_fish_spawn_egg":            &TropicalFishSpawnEgg,
	"turtle_spawn_egg":                   &TurtleSpawnEgg,
	"vex_spawn_egg":                      &VexSpawnEgg,
	"villager_spawn_egg":                 &VillagerSpawnEgg,
	"vindicator_spawn_egg":               &VindicatorSpawnEgg,
	"wandering_trader_spawn_egg":         &WanderingTraderSpawnEgg,
	"witch_spawn_egg":                    &WitchSpawnEgg,
	"wither_skeleton_spawn_egg":          &WitherSkeletonSpawnEgg,
	"wolf_spawn_egg":                     &WolfSpawnEgg,
	"zoglin_spawn_egg":                   &ZoglinSpawnEgg,
	"zombie_spawn_egg":                   &ZombieSpawnEgg,
	"zombie_horse_spawn_egg":             &ZombieHorseSpawnEgg,
	"zombie_villager_spawn_egg":          &ZombieVillagerSpawnEgg,
	"zombified_piglin_spawn_egg":         &ZombifiedPiglinSpawnEgg,
	"experience_bottle":                  &ExperienceBottle,
	"fire_charge":                        &FireCharge,
	"writable_book":                      &WritableBook,
	"written_book":                       &WrittenBook,
	"item_frame":                         &ItemFrame,
	"glow_item_frame":                    &GlowItemFrame,
	"flower_pot":                         &FlowerPot,
	"carrot":                             &Carrot,
	"potato":                             &Potato,
	"baked_potato":                       &BakedPotato,
	"poisonous_potato":                   &PoisonousPotato,
	"map":                                &Map,
	"golden_carrot":                      &GoldenCarrot,
	"skeleton_skull":                     &SkeletonSkull,
	"wither_skeleton_skull":              &WitherSkeletonSkull,
	"player_head":                        &PlayerHead,
	"zombie_head":                        &ZombieHead,
	"creeper_head":                       &CreeperHead,
	"dragon_head":                        &DragonHead,
	"nether_star":                        &NetherStar,
	"pumpkin_pie":                        &PumpkinPie,
	"firework_rocket":                    &FireworkRocket,
	"firework_star":                      &FireworkStar,
	"enchanted_book":                     &EnchantedBook,
	"nether_brick":                       &NetherBrick,
	"prismarine_shard":                   &PrismarineShard,
	"prismarine_crystals":                &PrismarineCrystals,
	"rabbit":                             &Rabbit,
	"cooked_rabbit":                      &CookedRabbit,
	"rabbit_stew":                        &RabbitStew,
	"rabbit_foot":                        &RabbitFoot,
	"rabbit_hide":                        &RabbitHide,
	"armor_stand":                        &ArmorStand,
	"iron_horse_armor":                   &IronHorseArmor,
	"golden_horse_armor":                 &GoldenHorseArmor,
	"diamond_horse_armor":                &DiamondHorseArmor,
	"leather_horse_armor":                &LeatherHorseArmor,
	"lead":                               &Lead,
	"name_tag":                           &NameTag,
	"command_block_minecart":             &CommandBlockMinecart,
	"mutton":                             &Mutton,
	"cooked_mutton":                      &CookedMutton,
	"white_banner":                       &WhiteBanner,
	"orange_banner":                      &OrangeBanner,
	"magenta_banner":                     &MagentaBanner,
	"light_blue_banner":                  &LightBlueBanner,
	"yellow_banner":                      &YellowBanner,
	"lime_banner":                        &LimeBanner,
	"pink_banner":                        &PinkBanner,
	"gray_banner":                        &GrayBanner,
	"light_gray_banner":                  &LightGrayBanner,
	"cyan_banner":                        &CyanBanner,
	"purple_banner":                      &PurpleBanner,
	"blue_banner":                        &BlueBanner,
	"brown_banner":                       &BrownBanner,
	"green_banner":                       &GreenBanner,
	"red_banner":                         &RedBanner,
	"black_banner":                       &BlackBanner,
	"end_crystal":                        &EndCrystal,
	"chorus_fruit":                       &ChorusFruit,
	"popped_chorus_fruit":                &PoppedChorusFruit,
	"beetroot":                           &Beetroot,
	"beetroot_seeds":                     &BeetrootSeeds,
	"beetroot_soup":                      &BeetrootSoup,
	"dragon_breath":                      &DragonBreath,
	"splash_potion":                      &SplashPotion,
	"spectral_arrow":                     &SpectralArrow,
	"tipped_arrow":                       &TippedArrow,
	"lingering_potion":                   &LingeringPotion,
	"shield":                             &Shield,
	"totem_of_undying":                   &TotemOfUndying,
	"shulker_shell":                      &ShulkerShell,
	"iron_nugget":                        &IronNugget,
	"knowledge_book":                     &KnowledgeBook,
	"debug_stick":                        &DebugStick,
	"music_disc_13":                      &MusicDisc13,
	"music_disc_cat":                     &MusicDiscCat,
	"music_disc_blocks":                  &MusicDiscBlocks,
	"music_disc_chirp":                   &MusicDiscChirp,
	"music_disc_far":                     &MusicDiscFar,
	"music_disc_mall":                    &MusicDiscMall,
	"music_disc_mellohi":                 &MusicDiscMellohi,
	"music_disc_stal":                    &MusicDiscStal,
	"music_disc_strad":                   &MusicDiscStrad,
	"music_disc_ward":                    &MusicDiscWard,
	"music_disc_11":                      &MusicDisc11,
	"music_disc_wait":                    &MusicDiscWait,
	"music_disc_pigstep":                 &MusicDiscPigstep,
	"trident":                            &Trident,
	"phantom_membrane":                   &PhantomMembrane,
	"nautilus_shell":                     &NautilusShell,
	"heart_of_the_sea":                   &HeartOfTheSea,
	"crossbow":                           &Crossbow,
	"suspicious_stew":                    &SuspiciousStew,
	"loom":                               &Loom,
	"flower_banner_pattern":              &FlowerBannerPattern,
	"creeper_banner_pattern":             &CreeperBannerPattern,
	"skull_banner_pattern":               &SkullBannerPattern,
	"mojang_banner_pattern":              &MojangBannerPattern,
	"globe_banner_pattern":               &GlobeBannerPattern,
	"piglin_banner_pattern":              &PiglinBannerPattern,
	"composter":                          &Composter,
	"barrel":                             &Barrel,
	"smoker":                             &Smoker,
	"blast_furnace":                      &BlastFurnace,
	"cartography_table":                  &CartographyTable,
	"fletching_table":                    &FletchingTable,
	"grindstone":                         &Grindstone,
	"smithing_table":                     &SmithingTable,
	"stonecutter":                        &Stonecutter,
	"bell":                               &Bell,
	"lantern":                            &Lantern,
	"soul_lantern":                       &SoulLantern,
	"sweet_berries":                      &SweetBerries,
	"glow_berries":                       &GlowBerries,
	"campfire":                           &Campfire,
	"soul_campfire":                      &SoulCampfire,
	"shroomlight":                        &Shroomlight,
	"honeycomb":                          &Honeycomb,
	"bee_nest":                           &BeeNest,
	"beehive":                            &Beehive,
	"honey_bottle":                       &HoneyBottle,
	"honeycomb_block":                    &HoneycombBlock,
	"lodestone":                          &Lodestone,
	"crying_obsidian":                    &CryingObsidian,
	"blackstone":                         &Blackstone,
	"blackstone_slab":                    &BlackstoneSlab,
	"blackstone_stairs":                  &BlackstoneStairs,
	"gilded_blackstone":                  &GildedBlackstone,
	"polished_blackstone":                &PolishedBlackstone,
	"polished_blackstone_slab":           &PolishedBlackstoneSlab,
	"polished_blackstone_stairs":         &PolishedBlackstoneStairs,
	"chiseled_polished_blackstone":       &ChiseledPolishedBlackstone,
	"polished_blackstone_bricks":         &PolishedBlackstoneBricks,
	"polished_blackstone_brick_slab":     &PolishedBlackstoneBrickSlab,
	"polished_blackstone_brick_stairs":   &PolishedBlackstoneBrickStairs,
	"cracked_polished_blackstone_bricks": &CrackedPolishedBlackstoneBricks,
	"respawn_anchor":                     &RespawnAnchor,
	"candle":                             &Candle,
	"white_candle":                       &WhiteCandle,
	"orange_candle":                      &OrangeCandle,
	"magenta_candle":                     &MagentaCandle,
	"light_blue_candle":                  &LightBlueCandle,
	"yellow_candle":                      &YellowCandle,
	"lime_candle":                        &LimeCandle,
	"pink_candle":                        &PinkCandle,
	"gray_candle":                        &GrayCandle,
	"light_gray_candle":                  &LightGrayCandle,
	"cyan_candle":                        &CyanCandle,
	"purple_candle":                      &PurpleCandle,
	"blue_candle":                        &BlueCandle,
	"brown_candle":                       &BrownCandle,
	"green_candle":                       &GreenCandle,
	"red_candle":                         &RedCandle,
	"black_candle":                       &BlackCandle,
	"small_amethyst_bud":                 &SmallAmethystBud,
	"medium_amethyst_bud":                &MediumAmethystBud,
	"large_amethyst_bud":                 &LargeAmethystBud,
	"amethyst_cluster":                   &AmethystCluster,
	"pointed_dripstone":                  &PointedDripstone,
}
