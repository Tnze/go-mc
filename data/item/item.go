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
	Air                             = Item{ID: 0, DisplayName: "Air", Name: "air", StackSize: 0}
	Stone                           = Item{ID: 1, DisplayName: "Stone", Name: "stone", StackSize: 64}
	Granite                         = Item{ID: 2, DisplayName: "Granite", Name: "granite", StackSize: 64}
	PolishedGranite                 = Item{ID: 3, DisplayName: "Polished Granite", Name: "polished_granite", StackSize: 64}
	Diorite                         = Item{ID: 4, DisplayName: "Diorite", Name: "diorite", StackSize: 64}
	PolishedDiorite                 = Item{ID: 5, DisplayName: "Polished Diorite", Name: "polished_diorite", StackSize: 64}
	Andesite                        = Item{ID: 6, DisplayName: "Andesite", Name: "andesite", StackSize: 64}
	PolishedAndesite                = Item{ID: 7, DisplayName: "Polished Andesite", Name: "polished_andesite", StackSize: 64}
	GrassBlock                      = Item{ID: 8, DisplayName: "Grass Block", Name: "grass_block", StackSize: 64}
	Dirt                            = Item{ID: 9, DisplayName: "Dirt", Name: "dirt", StackSize: 64}
	CoarseDirt                      = Item{ID: 10, DisplayName: "Coarse Dirt", Name: "coarse_dirt", StackSize: 64}
	Podzol                          = Item{ID: 11, DisplayName: "Podzol", Name: "podzol", StackSize: 64}
	CrimsonNylium                   = Item{ID: 12, DisplayName: "Crimson Nylium", Name: "crimson_nylium", StackSize: 64}
	WarpedNylium                    = Item{ID: 13, DisplayName: "Warped Nylium", Name: "warped_nylium", StackSize: 64}
	Cobblestone                     = Item{ID: 14, DisplayName: "Cobblestone", Name: "cobblestone", StackSize: 64}
	OakPlanks                       = Item{ID: 15, DisplayName: "Oak Planks", Name: "oak_planks", StackSize: 64}
	SprucePlanks                    = Item{ID: 16, DisplayName: "Spruce Planks", Name: "spruce_planks", StackSize: 64}
	BirchPlanks                     = Item{ID: 17, DisplayName: "Birch Planks", Name: "birch_planks", StackSize: 64}
	JunglePlanks                    = Item{ID: 18, DisplayName: "Jungle Planks", Name: "jungle_planks", StackSize: 64}
	AcaciaPlanks                    = Item{ID: 19, DisplayName: "Acacia Planks", Name: "acacia_planks", StackSize: 64}
	DarkOakPlanks                   = Item{ID: 20, DisplayName: "Dark Oak Planks", Name: "dark_oak_planks", StackSize: 64}
	CrimsonPlanks                   = Item{ID: 21, DisplayName: "Crimson Planks", Name: "crimson_planks", StackSize: 64}
	WarpedPlanks                    = Item{ID: 22, DisplayName: "Warped Planks", Name: "warped_planks", StackSize: 64}
	OakSapling                      = Item{ID: 23, DisplayName: "Oak Sapling", Name: "oak_sapling", StackSize: 64}
	SpruceSapling                   = Item{ID: 24, DisplayName: "Spruce Sapling", Name: "spruce_sapling", StackSize: 64}
	BirchSapling                    = Item{ID: 25, DisplayName: "Birch Sapling", Name: "birch_sapling", StackSize: 64}
	JungleSapling                   = Item{ID: 26, DisplayName: "Jungle Sapling", Name: "jungle_sapling", StackSize: 64}
	AcaciaSapling                   = Item{ID: 27, DisplayName: "Acacia Sapling", Name: "acacia_sapling", StackSize: 64}
	DarkOakSapling                  = Item{ID: 28, DisplayName: "Dark Oak Sapling", Name: "dark_oak_sapling", StackSize: 64}
	Bedrock                         = Item{ID: 29, DisplayName: "Bedrock", Name: "bedrock", StackSize: 64}
	Sand                            = Item{ID: 30, DisplayName: "Sand", Name: "sand", StackSize: 64}
	RedSand                         = Item{ID: 31, DisplayName: "Red Sand", Name: "red_sand", StackSize: 64}
	Gravel                          = Item{ID: 32, DisplayName: "Gravel", Name: "gravel", StackSize: 64}
	GoldOre                         = Item{ID: 33, DisplayName: "Gold Ore", Name: "gold_ore", StackSize: 64}
	IronOre                         = Item{ID: 34, DisplayName: "Iron Ore", Name: "iron_ore", StackSize: 64}
	CoalOre                         = Item{ID: 35, DisplayName: "Coal Ore", Name: "coal_ore", StackSize: 64}
	NetherGoldOre                   = Item{ID: 36, DisplayName: "Nether Gold Ore", Name: "nether_gold_ore", StackSize: 64}
	OakLog                          = Item{ID: 37, DisplayName: "Oak Log", Name: "oak_log", StackSize: 64}
	SpruceLog                       = Item{ID: 38, DisplayName: "Spruce Log", Name: "spruce_log", StackSize: 64}
	BirchLog                        = Item{ID: 39, DisplayName: "Birch Log", Name: "birch_log", StackSize: 64}
	JungleLog                       = Item{ID: 40, DisplayName: "Jungle Log", Name: "jungle_log", StackSize: 64}
	AcaciaLog                       = Item{ID: 41, DisplayName: "Acacia Log", Name: "acacia_log", StackSize: 64}
	DarkOakLog                      = Item{ID: 42, DisplayName: "Dark Oak Log", Name: "dark_oak_log", StackSize: 64}
	CrimsonStem                     = Item{ID: 43, DisplayName: "Crimson Stem", Name: "crimson_stem", StackSize: 64}
	WarpedStem                      = Item{ID: 44, DisplayName: "Warped Stem", Name: "warped_stem", StackSize: 64}
	StrippedOakLog                  = Item{ID: 45, DisplayName: "Stripped Oak Log", Name: "stripped_oak_log", StackSize: 64}
	StrippedSpruceLog               = Item{ID: 46, DisplayName: "Stripped Spruce Log", Name: "stripped_spruce_log", StackSize: 64}
	StrippedBirchLog                = Item{ID: 47, DisplayName: "Stripped Birch Log", Name: "stripped_birch_log", StackSize: 64}
	StrippedJungleLog               = Item{ID: 48, DisplayName: "Stripped Jungle Log", Name: "stripped_jungle_log", StackSize: 64}
	StrippedAcaciaLog               = Item{ID: 49, DisplayName: "Stripped Acacia Log", Name: "stripped_acacia_log", StackSize: 64}
	StrippedDarkOakLog              = Item{ID: 50, DisplayName: "Stripped Dark Oak Log", Name: "stripped_dark_oak_log", StackSize: 64}
	StrippedCrimsonStem             = Item{ID: 51, DisplayName: "Stripped Crimson Stem", Name: "stripped_crimson_stem", StackSize: 64}
	StrippedWarpedStem              = Item{ID: 52, DisplayName: "Stripped Warped Stem", Name: "stripped_warped_stem", StackSize: 64}
	StrippedOakWood                 = Item{ID: 53, DisplayName: "Stripped Oak Wood", Name: "stripped_oak_wood", StackSize: 64}
	StrippedSpruceWood              = Item{ID: 54, DisplayName: "Stripped Spruce Wood", Name: "stripped_spruce_wood", StackSize: 64}
	StrippedBirchWood               = Item{ID: 55, DisplayName: "Stripped Birch Wood", Name: "stripped_birch_wood", StackSize: 64}
	StrippedJungleWood              = Item{ID: 56, DisplayName: "Stripped Jungle Wood", Name: "stripped_jungle_wood", StackSize: 64}
	StrippedAcaciaWood              = Item{ID: 57, DisplayName: "Stripped Acacia Wood", Name: "stripped_acacia_wood", StackSize: 64}
	StrippedDarkOakWood             = Item{ID: 58, DisplayName: "Stripped Dark Oak Wood", Name: "stripped_dark_oak_wood", StackSize: 64}
	StrippedCrimsonHyphae           = Item{ID: 59, DisplayName: "Stripped Crimson Hyphae", Name: "stripped_crimson_hyphae", StackSize: 64}
	StrippedWarpedHyphae            = Item{ID: 60, DisplayName: "Stripped Warped Hyphae", Name: "stripped_warped_hyphae", StackSize: 64}
	OakWood                         = Item{ID: 61, DisplayName: "Oak Wood", Name: "oak_wood", StackSize: 64}
	SpruceWood                      = Item{ID: 62, DisplayName: "Spruce Wood", Name: "spruce_wood", StackSize: 64}
	BirchWood                       = Item{ID: 63, DisplayName: "Birch Wood", Name: "birch_wood", StackSize: 64}
	JungleWood                      = Item{ID: 64, DisplayName: "Jungle Wood", Name: "jungle_wood", StackSize: 64}
	AcaciaWood                      = Item{ID: 65, DisplayName: "Acacia Wood", Name: "acacia_wood", StackSize: 64}
	DarkOakWood                     = Item{ID: 66, DisplayName: "Dark Oak Wood", Name: "dark_oak_wood", StackSize: 64}
	CrimsonHyphae                   = Item{ID: 67, DisplayName: "Crimson Hyphae", Name: "crimson_hyphae", StackSize: 64}
	WarpedHyphae                    = Item{ID: 68, DisplayName: "Warped Hyphae", Name: "warped_hyphae", StackSize: 64}
	OakLeaves                       = Item{ID: 69, DisplayName: "Oak Leaves", Name: "oak_leaves", StackSize: 64}
	SpruceLeaves                    = Item{ID: 70, DisplayName: "Spruce Leaves", Name: "spruce_leaves", StackSize: 64}
	BirchLeaves                     = Item{ID: 71, DisplayName: "Birch Leaves", Name: "birch_leaves", StackSize: 64}
	JungleLeaves                    = Item{ID: 72, DisplayName: "Jungle Leaves", Name: "jungle_leaves", StackSize: 64}
	AcaciaLeaves                    = Item{ID: 73, DisplayName: "Acacia Leaves", Name: "acacia_leaves", StackSize: 64}
	DarkOakLeaves                   = Item{ID: 74, DisplayName: "Dark Oak Leaves", Name: "dark_oak_leaves", StackSize: 64}
	Sponge                          = Item{ID: 75, DisplayName: "Sponge", Name: "sponge", StackSize: 64}
	WetSponge                       = Item{ID: 76, DisplayName: "Wet Sponge", Name: "wet_sponge", StackSize: 64}
	Glass                           = Item{ID: 77, DisplayName: "Glass", Name: "glass", StackSize: 64}
	LapisOre                        = Item{ID: 78, DisplayName: "Lapis Lazuli Ore", Name: "lapis_ore", StackSize: 64}
	LapisBlock                      = Item{ID: 79, DisplayName: "Lapis Lazuli Block", Name: "lapis_block", StackSize: 64}
	Dispenser                       = Item{ID: 80, DisplayName: "Dispenser", Name: "dispenser", StackSize: 64}
	Sandstone                       = Item{ID: 81, DisplayName: "Sandstone", Name: "sandstone", StackSize: 64}
	ChiseledSandstone               = Item{ID: 82, DisplayName: "Chiseled Sandstone", Name: "chiseled_sandstone", StackSize: 64}
	CutSandstone                    = Item{ID: 83, DisplayName: "Cut Sandstone", Name: "cut_sandstone", StackSize: 64}
	NoteBlock                       = Item{ID: 84, DisplayName: "Note Block", Name: "note_block", StackSize: 64}
	PoweredRail                     = Item{ID: 85, DisplayName: "Powered Rail", Name: "powered_rail", StackSize: 64}
	DetectorRail                    = Item{ID: 86, DisplayName: "Detector Rail", Name: "detector_rail", StackSize: 64}
	StickyPiston                    = Item{ID: 87, DisplayName: "Sticky Piston", Name: "sticky_piston", StackSize: 64}
	Cobweb                          = Item{ID: 88, DisplayName: "Cobweb", Name: "cobweb", StackSize: 64}
	Grass                           = Item{ID: 89, DisplayName: "Grass", Name: "grass", StackSize: 64}
	Fern                            = Item{ID: 90, DisplayName: "Fern", Name: "fern", StackSize: 64}
	DeadBush                        = Item{ID: 91, DisplayName: "Dead Bush", Name: "dead_bush", StackSize: 64}
	Seagrass                        = Item{ID: 92, DisplayName: "Seagrass", Name: "seagrass", StackSize: 64}
	SeaPickle                       = Item{ID: 93, DisplayName: "Sea Pickle", Name: "sea_pickle", StackSize: 64}
	Piston                          = Item{ID: 94, DisplayName: "Piston", Name: "piston", StackSize: 64}
	WhiteWool                       = Item{ID: 95, DisplayName: "White Wool", Name: "white_wool", StackSize: 64}
	OrangeWool                      = Item{ID: 96, DisplayName: "Orange Wool", Name: "orange_wool", StackSize: 64}
	MagentaWool                     = Item{ID: 97, DisplayName: "Magenta Wool", Name: "magenta_wool", StackSize: 64}
	LightBlueWool                   = Item{ID: 98, DisplayName: "Light Blue Wool", Name: "light_blue_wool", StackSize: 64}
	YellowWool                      = Item{ID: 99, DisplayName: "Yellow Wool", Name: "yellow_wool", StackSize: 64}
	LimeWool                        = Item{ID: 100, DisplayName: "Lime Wool", Name: "lime_wool", StackSize: 64}
	PinkWool                        = Item{ID: 101, DisplayName: "Pink Wool", Name: "pink_wool", StackSize: 64}
	GrayWool                        = Item{ID: 102, DisplayName: "Gray Wool", Name: "gray_wool", StackSize: 64}
	LightGrayWool                   = Item{ID: 103, DisplayName: "Light Gray Wool", Name: "light_gray_wool", StackSize: 64}
	CyanWool                        = Item{ID: 104, DisplayName: "Cyan Wool", Name: "cyan_wool", StackSize: 64}
	PurpleWool                      = Item{ID: 105, DisplayName: "Purple Wool", Name: "purple_wool", StackSize: 64}
	BlueWool                        = Item{ID: 106, DisplayName: "Blue Wool", Name: "blue_wool", StackSize: 64}
	BrownWool                       = Item{ID: 107, DisplayName: "Brown Wool", Name: "brown_wool", StackSize: 64}
	GreenWool                       = Item{ID: 108, DisplayName: "Green Wool", Name: "green_wool", StackSize: 64}
	RedWool                         = Item{ID: 109, DisplayName: "Red Wool", Name: "red_wool", StackSize: 64}
	BlackWool                       = Item{ID: 110, DisplayName: "Black Wool", Name: "black_wool", StackSize: 64}
	Dandelion                       = Item{ID: 111, DisplayName: "Dandelion", Name: "dandelion", StackSize: 64}
	Poppy                           = Item{ID: 112, DisplayName: "Poppy", Name: "poppy", StackSize: 64}
	BlueOrchid                      = Item{ID: 113, DisplayName: "Blue Orchid", Name: "blue_orchid", StackSize: 64}
	Allium                          = Item{ID: 114, DisplayName: "Allium", Name: "allium", StackSize: 64}
	AzureBluet                      = Item{ID: 115, DisplayName: "Azure Bluet", Name: "azure_bluet", StackSize: 64}
	RedTulip                        = Item{ID: 116, DisplayName: "Red Tulip", Name: "red_tulip", StackSize: 64}
	OrangeTulip                     = Item{ID: 117, DisplayName: "Orange Tulip", Name: "orange_tulip", StackSize: 64}
	WhiteTulip                      = Item{ID: 118, DisplayName: "White Tulip", Name: "white_tulip", StackSize: 64}
	PinkTulip                       = Item{ID: 119, DisplayName: "Pink Tulip", Name: "pink_tulip", StackSize: 64}
	OxeyeDaisy                      = Item{ID: 120, DisplayName: "Oxeye Daisy", Name: "oxeye_daisy", StackSize: 64}
	Cornflower                      = Item{ID: 121, DisplayName: "Cornflower", Name: "cornflower", StackSize: 64}
	LilyOfTheValley                 = Item{ID: 122, DisplayName: "Lily of the Valley", Name: "lily_of_the_valley", StackSize: 64}
	WitherRose                      = Item{ID: 123, DisplayName: "Wither Rose", Name: "wither_rose", StackSize: 64}
	BrownMushroom                   = Item{ID: 124, DisplayName: "Brown Mushroom", Name: "brown_mushroom", StackSize: 64}
	RedMushroom                     = Item{ID: 125, DisplayName: "Red Mushroom", Name: "red_mushroom", StackSize: 64}
	CrimsonFungus                   = Item{ID: 126, DisplayName: "Crimson Fungus", Name: "crimson_fungus", StackSize: 64}
	WarpedFungus                    = Item{ID: 127, DisplayName: "Warped Fungus", Name: "warped_fungus", StackSize: 64}
	CrimsonRoots                    = Item{ID: 128, DisplayName: "Crimson Roots", Name: "crimson_roots", StackSize: 64}
	WarpedRoots                     = Item{ID: 129, DisplayName: "Warped Roots", Name: "warped_roots", StackSize: 64}
	NetherSprouts                   = Item{ID: 130, DisplayName: "Nether Sprouts", Name: "nether_sprouts", StackSize: 64}
	WeepingVines                    = Item{ID: 131, DisplayName: "Weeping Vines", Name: "weeping_vines", StackSize: 64}
	TwistingVines                   = Item{ID: 132, DisplayName: "Twisting Vines", Name: "twisting_vines", StackSize: 64}
	SugarCane                       = Item{ID: 133, DisplayName: "Sugar Cane", Name: "sugar_cane", StackSize: 64}
	Kelp                            = Item{ID: 134, DisplayName: "Kelp", Name: "kelp", StackSize: 64}
	Bamboo                          = Item{ID: 135, DisplayName: "Bamboo", Name: "bamboo", StackSize: 64}
	GoldBlock                       = Item{ID: 136, DisplayName: "Block of Gold", Name: "gold_block", StackSize: 64}
	IronBlock                       = Item{ID: 137, DisplayName: "Block of Iron", Name: "iron_block", StackSize: 64}
	OakSlab                         = Item{ID: 138, DisplayName: "Oak Slab", Name: "oak_slab", StackSize: 64}
	SpruceSlab                      = Item{ID: 139, DisplayName: "Spruce Slab", Name: "spruce_slab", StackSize: 64}
	BirchSlab                       = Item{ID: 140, DisplayName: "Birch Slab", Name: "birch_slab", StackSize: 64}
	JungleSlab                      = Item{ID: 141, DisplayName: "Jungle Slab", Name: "jungle_slab", StackSize: 64}
	AcaciaSlab                      = Item{ID: 142, DisplayName: "Acacia Slab", Name: "acacia_slab", StackSize: 64}
	DarkOakSlab                     = Item{ID: 143, DisplayName: "Dark Oak Slab", Name: "dark_oak_slab", StackSize: 64}
	CrimsonSlab                     = Item{ID: 144, DisplayName: "Crimson Slab", Name: "crimson_slab", StackSize: 64}
	WarpedSlab                      = Item{ID: 145, DisplayName: "Warped Slab", Name: "warped_slab", StackSize: 64}
	StoneSlab                       = Item{ID: 146, DisplayName: "Stone Slab", Name: "stone_slab", StackSize: 64}
	SmoothStoneSlab                 = Item{ID: 147, DisplayName: "Smooth Stone Slab", Name: "smooth_stone_slab", StackSize: 64}
	SandstoneSlab                   = Item{ID: 148, DisplayName: "Sandstone Slab", Name: "sandstone_slab", StackSize: 64}
	CutSandstoneSlab                = Item{ID: 149, DisplayName: "Cut Sandstone Slab", Name: "cut_sandstone_slab", StackSize: 64}
	PetrifiedOakSlab                = Item{ID: 150, DisplayName: "Petrified Oak Slab", Name: "petrified_oak_slab", StackSize: 64}
	CobblestoneSlab                 = Item{ID: 151, DisplayName: "Cobblestone Slab", Name: "cobblestone_slab", StackSize: 64}
	BrickSlab                       = Item{ID: 152, DisplayName: "Brick Slab", Name: "brick_slab", StackSize: 64}
	StoneBrickSlab                  = Item{ID: 153, DisplayName: "Stone Brick Slab", Name: "stone_brick_slab", StackSize: 64}
	NetherBrickSlab                 = Item{ID: 154, DisplayName: "Nether Brick Slab", Name: "nether_brick_slab", StackSize: 64}
	QuartzSlab                      = Item{ID: 155, DisplayName: "Quartz Slab", Name: "quartz_slab", StackSize: 64}
	RedSandstoneSlab                = Item{ID: 156, DisplayName: "Red Sandstone Slab", Name: "red_sandstone_slab", StackSize: 64}
	CutRedSandstoneSlab             = Item{ID: 157, DisplayName: "Cut Red Sandstone Slab", Name: "cut_red_sandstone_slab", StackSize: 64}
	PurpurSlab                      = Item{ID: 158, DisplayName: "Purpur Slab", Name: "purpur_slab", StackSize: 64}
	PrismarineSlab                  = Item{ID: 159, DisplayName: "Prismarine Slab", Name: "prismarine_slab", StackSize: 64}
	PrismarineBrickSlab             = Item{ID: 160, DisplayName: "Prismarine Brick Slab", Name: "prismarine_brick_slab", StackSize: 64}
	DarkPrismarineSlab              = Item{ID: 161, DisplayName: "Dark Prismarine Slab", Name: "dark_prismarine_slab", StackSize: 64}
	SmoothQuartz                    = Item{ID: 162, DisplayName: "Smooth Quartz Block", Name: "smooth_quartz", StackSize: 64}
	SmoothRedSandstone              = Item{ID: 163, DisplayName: "Smooth Red Sandstone", Name: "smooth_red_sandstone", StackSize: 64}
	SmoothSandstone                 = Item{ID: 164, DisplayName: "Smooth Sandstone", Name: "smooth_sandstone", StackSize: 64}
	SmoothStone                     = Item{ID: 165, DisplayName: "Smooth Stone", Name: "smooth_stone", StackSize: 64}
	Bricks                          = Item{ID: 166, DisplayName: "Bricks", Name: "bricks", StackSize: 64}
	Tnt                             = Item{ID: 167, DisplayName: "TNT", Name: "tnt", StackSize: 64}
	Bookshelf                       = Item{ID: 168, DisplayName: "Bookshelf", Name: "bookshelf", StackSize: 64}
	MossyCobblestone                = Item{ID: 169, DisplayName: "Mossy Cobblestone", Name: "mossy_cobblestone", StackSize: 64}
	Obsidian                        = Item{ID: 170, DisplayName: "Obsidian", Name: "obsidian", StackSize: 64}
	Torch                           = Item{ID: 171, DisplayName: "Torch", Name: "torch", StackSize: 64}
	EndRod                          = Item{ID: 172, DisplayName: "End Rod", Name: "end_rod", StackSize: 64}
	ChorusPlant                     = Item{ID: 173, DisplayName: "Chorus Plant", Name: "chorus_plant", StackSize: 64}
	ChorusFlower                    = Item{ID: 174, DisplayName: "Chorus Flower", Name: "chorus_flower", StackSize: 64}
	PurpurBlock                     = Item{ID: 175, DisplayName: "Purpur Block", Name: "purpur_block", StackSize: 64}
	PurpurPillar                    = Item{ID: 176, DisplayName: "Purpur Pillar", Name: "purpur_pillar", StackSize: 64}
	PurpurStairs                    = Item{ID: 177, DisplayName: "Purpur Stairs", Name: "purpur_stairs", StackSize: 64}
	Spawner                         = Item{ID: 178, DisplayName: "Spawner", Name: "spawner", StackSize: 64}
	OakStairs                       = Item{ID: 179, DisplayName: "Oak Stairs", Name: "oak_stairs", StackSize: 64}
	Chest                           = Item{ID: 180, DisplayName: "Chest", Name: "chest", StackSize: 64}
	DiamondOre                      = Item{ID: 181, DisplayName: "Diamond Ore", Name: "diamond_ore", StackSize: 64}
	DiamondBlock                    = Item{ID: 182, DisplayName: "Block of Diamond", Name: "diamond_block", StackSize: 64}
	CraftingTable                   = Item{ID: 183, DisplayName: "Crafting Table", Name: "crafting_table", StackSize: 64}
	Farmland                        = Item{ID: 184, DisplayName: "Farmland", Name: "farmland", StackSize: 64}
	Furnace                         = Item{ID: 185, DisplayName: "Furnace", Name: "furnace", StackSize: 64}
	Ladder                          = Item{ID: 186, DisplayName: "Ladder", Name: "ladder", StackSize: 64}
	Rail                            = Item{ID: 187, DisplayName: "Rail", Name: "rail", StackSize: 64}
	CobblestoneStairs               = Item{ID: 188, DisplayName: "Cobblestone Stairs", Name: "cobblestone_stairs", StackSize: 64}
	Lever                           = Item{ID: 189, DisplayName: "Lever", Name: "lever", StackSize: 64}
	StonePressurePlate              = Item{ID: 190, DisplayName: "Stone Pressure Plate", Name: "stone_pressure_plate", StackSize: 64}
	OakPressurePlate                = Item{ID: 191, DisplayName: "Oak Pressure Plate", Name: "oak_pressure_plate", StackSize: 64}
	SprucePressurePlate             = Item{ID: 192, DisplayName: "Spruce Pressure Plate", Name: "spruce_pressure_plate", StackSize: 64}
	BirchPressurePlate              = Item{ID: 193, DisplayName: "Birch Pressure Plate", Name: "birch_pressure_plate", StackSize: 64}
	JunglePressurePlate             = Item{ID: 194, DisplayName: "Jungle Pressure Plate", Name: "jungle_pressure_plate", StackSize: 64}
	AcaciaPressurePlate             = Item{ID: 195, DisplayName: "Acacia Pressure Plate", Name: "acacia_pressure_plate", StackSize: 64}
	DarkOakPressurePlate            = Item{ID: 196, DisplayName: "Dark Oak Pressure Plate", Name: "dark_oak_pressure_plate", StackSize: 64}
	CrimsonPressurePlate            = Item{ID: 197, DisplayName: "Crimson Pressure Plate", Name: "crimson_pressure_plate", StackSize: 64}
	WarpedPressurePlate             = Item{ID: 198, DisplayName: "Warped Pressure Plate", Name: "warped_pressure_plate", StackSize: 64}
	PolishedBlackstonePressurePlate = Item{ID: 199, DisplayName: "Polished Blackstone Pressure Plate", Name: "polished_blackstone_pressure_plate", StackSize: 64}
	RedstoneOre                     = Item{ID: 200, DisplayName: "Redstone Ore", Name: "redstone_ore", StackSize: 64}
	RedstoneTorch                   = Item{ID: 201, DisplayName: "Redstone Torch", Name: "redstone_torch", StackSize: 64}
	Snow                            = Item{ID: 202, DisplayName: "Snow", Name: "snow", StackSize: 64}
	Ice                             = Item{ID: 203, DisplayName: "Ice", Name: "ice", StackSize: 64}
	SnowBlock                       = Item{ID: 204, DisplayName: "Snow Block", Name: "snow_block", StackSize: 64}
	Cactus                          = Item{ID: 205, DisplayName: "Cactus", Name: "cactus", StackSize: 64}
	Clay                            = Item{ID: 206, DisplayName: "Clay", Name: "clay", StackSize: 64}
	Jukebox                         = Item{ID: 207, DisplayName: "Jukebox", Name: "jukebox", StackSize: 64}
	OakFence                        = Item{ID: 208, DisplayName: "Oak Fence", Name: "oak_fence", StackSize: 64}
	SpruceFence                     = Item{ID: 209, DisplayName: "Spruce Fence", Name: "spruce_fence", StackSize: 64}
	BirchFence                      = Item{ID: 210, DisplayName: "Birch Fence", Name: "birch_fence", StackSize: 64}
	JungleFence                     = Item{ID: 211, DisplayName: "Jungle Fence", Name: "jungle_fence", StackSize: 64}
	AcaciaFence                     = Item{ID: 212, DisplayName: "Acacia Fence", Name: "acacia_fence", StackSize: 64}
	DarkOakFence                    = Item{ID: 213, DisplayName: "Dark Oak Fence", Name: "dark_oak_fence", StackSize: 64}
	CrimsonFence                    = Item{ID: 214, DisplayName: "Crimson Fence", Name: "crimson_fence", StackSize: 64}
	WarpedFence                     = Item{ID: 215, DisplayName: "Warped Fence", Name: "warped_fence", StackSize: 64}
	Pumpkin                         = Item{ID: 216, DisplayName: "Pumpkin", Name: "pumpkin", StackSize: 64}
	CarvedPumpkin                   = Item{ID: 217, DisplayName: "Carved Pumpkin", Name: "carved_pumpkin", StackSize: 64}
	Netherrack                      = Item{ID: 218, DisplayName: "Netherrack", Name: "netherrack", StackSize: 64}
	SoulSand                        = Item{ID: 219, DisplayName: "Soul Sand", Name: "soul_sand", StackSize: 64}
	SoulSoil                        = Item{ID: 220, DisplayName: "Soul Soil", Name: "soul_soil", StackSize: 64}
	Basalt                          = Item{ID: 221, DisplayName: "Basalt", Name: "basalt", StackSize: 64}
	PolishedBasalt                  = Item{ID: 222, DisplayName: "Polished Basalt", Name: "polished_basalt", StackSize: 64}
	SoulTorch                       = Item{ID: 223, DisplayName: "Soul Torch", Name: "soul_torch", StackSize: 64}
	Glowstone                       = Item{ID: 224, DisplayName: "Glowstone", Name: "glowstone", StackSize: 64}
	JackOLantern                    = Item{ID: 225, DisplayName: "Jack o'Lantern", Name: "jack_o_lantern", StackSize: 64}
	OakTrapdoor                     = Item{ID: 226, DisplayName: "Oak Trapdoor", Name: "oak_trapdoor", StackSize: 64}
	SpruceTrapdoor                  = Item{ID: 227, DisplayName: "Spruce Trapdoor", Name: "spruce_trapdoor", StackSize: 64}
	BirchTrapdoor                   = Item{ID: 228, DisplayName: "Birch Trapdoor", Name: "birch_trapdoor", StackSize: 64}
	JungleTrapdoor                  = Item{ID: 229, DisplayName: "Jungle Trapdoor", Name: "jungle_trapdoor", StackSize: 64}
	AcaciaTrapdoor                  = Item{ID: 230, DisplayName: "Acacia Trapdoor", Name: "acacia_trapdoor", StackSize: 64}
	DarkOakTrapdoor                 = Item{ID: 231, DisplayName: "Dark Oak Trapdoor", Name: "dark_oak_trapdoor", StackSize: 64}
	CrimsonTrapdoor                 = Item{ID: 232, DisplayName: "Crimson Trapdoor", Name: "crimson_trapdoor", StackSize: 64}
	WarpedTrapdoor                  = Item{ID: 233, DisplayName: "Warped Trapdoor", Name: "warped_trapdoor", StackSize: 64}
	InfestedStone                   = Item{ID: 234, DisplayName: "Infested Stone", Name: "infested_stone", StackSize: 64}
	InfestedCobblestone             = Item{ID: 235, DisplayName: "Infested Cobblestone", Name: "infested_cobblestone", StackSize: 64}
	InfestedStoneBricks             = Item{ID: 236, DisplayName: "Infested Stone Bricks", Name: "infested_stone_bricks", StackSize: 64}
	InfestedMossyStoneBricks        = Item{ID: 237, DisplayName: "Infested Mossy Stone Bricks", Name: "infested_mossy_stone_bricks", StackSize: 64}
	InfestedCrackedStoneBricks      = Item{ID: 238, DisplayName: "Infested Cracked Stone Bricks", Name: "infested_cracked_stone_bricks", StackSize: 64}
	InfestedChiseledStoneBricks     = Item{ID: 239, DisplayName: "Infested Chiseled Stone Bricks", Name: "infested_chiseled_stone_bricks", StackSize: 64}
	StoneBricks                     = Item{ID: 240, DisplayName: "Stone Bricks", Name: "stone_bricks", StackSize: 64}
	MossyStoneBricks                = Item{ID: 241, DisplayName: "Mossy Stone Bricks", Name: "mossy_stone_bricks", StackSize: 64}
	CrackedStoneBricks              = Item{ID: 242, DisplayName: "Cracked Stone Bricks", Name: "cracked_stone_bricks", StackSize: 64}
	ChiseledStoneBricks             = Item{ID: 243, DisplayName: "Chiseled Stone Bricks", Name: "chiseled_stone_bricks", StackSize: 64}
	BrownMushroomBlock              = Item{ID: 244, DisplayName: "Brown Mushroom Block", Name: "brown_mushroom_block", StackSize: 64}
	RedMushroomBlock                = Item{ID: 245, DisplayName: "Red Mushroom Block", Name: "red_mushroom_block", StackSize: 64}
	MushroomStem                    = Item{ID: 246, DisplayName: "Mushroom Stem", Name: "mushroom_stem", StackSize: 64}
	IronBars                        = Item{ID: 247, DisplayName: "Iron Bars", Name: "iron_bars", StackSize: 64}
	Chain                           = Item{ID: 248, DisplayName: "Chain", Name: "chain", StackSize: 64}
	GlassPane                       = Item{ID: 249, DisplayName: "Glass Pane", Name: "glass_pane", StackSize: 64}
	Melon                           = Item{ID: 250, DisplayName: "Melon", Name: "melon", StackSize: 64}
	Vine                            = Item{ID: 251, DisplayName: "Vines", Name: "vine", StackSize: 64}
	OakFenceGate                    = Item{ID: 252, DisplayName: "Oak Fence Gate", Name: "oak_fence_gate", StackSize: 64}
	SpruceFenceGate                 = Item{ID: 253, DisplayName: "Spruce Fence Gate", Name: "spruce_fence_gate", StackSize: 64}
	BirchFenceGate                  = Item{ID: 254, DisplayName: "Birch Fence Gate", Name: "birch_fence_gate", StackSize: 64}
	JungleFenceGate                 = Item{ID: 255, DisplayName: "Jungle Fence Gate", Name: "jungle_fence_gate", StackSize: 64}
	AcaciaFenceGate                 = Item{ID: 256, DisplayName: "Acacia Fence Gate", Name: "acacia_fence_gate", StackSize: 64}
	DarkOakFenceGate                = Item{ID: 257, DisplayName: "Dark Oak Fence Gate", Name: "dark_oak_fence_gate", StackSize: 64}
	CrimsonFenceGate                = Item{ID: 258, DisplayName: "Crimson Fence Gate", Name: "crimson_fence_gate", StackSize: 64}
	WarpedFenceGate                 = Item{ID: 259, DisplayName: "Warped Fence Gate", Name: "warped_fence_gate", StackSize: 64}
	BrickStairs                     = Item{ID: 260, DisplayName: "Brick Stairs", Name: "brick_stairs", StackSize: 64}
	StoneBrickStairs                = Item{ID: 261, DisplayName: "Stone Brick Stairs", Name: "stone_brick_stairs", StackSize: 64}
	Mycelium                        = Item{ID: 262, DisplayName: "Mycelium", Name: "mycelium", StackSize: 64}
	LilyPad                         = Item{ID: 263, DisplayName: "Lily Pad", Name: "lily_pad", StackSize: 64}
	NetherBricks                    = Item{ID: 264, DisplayName: "Nether Bricks", Name: "nether_bricks", StackSize: 64}
	CrackedNetherBricks             = Item{ID: 265, DisplayName: "Cracked Nether Bricks", Name: "cracked_nether_bricks", StackSize: 64}
	ChiseledNetherBricks            = Item{ID: 266, DisplayName: "Chiseled Nether Bricks", Name: "chiseled_nether_bricks", StackSize: 64}
	NetherBrickFence                = Item{ID: 267, DisplayName: "Nether Brick Fence", Name: "nether_brick_fence", StackSize: 64}
	NetherBrickStairs               = Item{ID: 268, DisplayName: "Nether Brick Stairs", Name: "nether_brick_stairs", StackSize: 64}
	EnchantingTable                 = Item{ID: 269, DisplayName: "Enchanting Table", Name: "enchanting_table", StackSize: 64}
	EndPortalFrame                  = Item{ID: 270, DisplayName: "End Portal Frame", Name: "end_portal_frame", StackSize: 64}
	EndStone                        = Item{ID: 271, DisplayName: "End Stone", Name: "end_stone", StackSize: 64}
	EndStoneBricks                  = Item{ID: 272, DisplayName: "End Stone Bricks", Name: "end_stone_bricks", StackSize: 64}
	DragonEgg                       = Item{ID: 273, DisplayName: "Dragon Egg", Name: "dragon_egg", StackSize: 64}
	RedstoneLamp                    = Item{ID: 274, DisplayName: "Redstone Lamp", Name: "redstone_lamp", StackSize: 64}
	SandstoneStairs                 = Item{ID: 275, DisplayName: "Sandstone Stairs", Name: "sandstone_stairs", StackSize: 64}
	EmeraldOre                      = Item{ID: 276, DisplayName: "Emerald Ore", Name: "emerald_ore", StackSize: 64}
	EnderChest                      = Item{ID: 277, DisplayName: "Ender Chest", Name: "ender_chest", StackSize: 64}
	TripwireHook                    = Item{ID: 278, DisplayName: "Tripwire Hook", Name: "tripwire_hook", StackSize: 64}
	EmeraldBlock                    = Item{ID: 279, DisplayName: "Block of Emerald", Name: "emerald_block", StackSize: 64}
	SpruceStairs                    = Item{ID: 280, DisplayName: "Spruce Stairs", Name: "spruce_stairs", StackSize: 64}
	BirchStairs                     = Item{ID: 281, DisplayName: "Birch Stairs", Name: "birch_stairs", StackSize: 64}
	JungleStairs                    = Item{ID: 282, DisplayName: "Jungle Stairs", Name: "jungle_stairs", StackSize: 64}
	CrimsonStairs                   = Item{ID: 283, DisplayName: "Crimson Stairs", Name: "crimson_stairs", StackSize: 64}
	WarpedStairs                    = Item{ID: 284, DisplayName: "Warped Stairs", Name: "warped_stairs", StackSize: 64}
	CommandBlock                    = Item{ID: 285, DisplayName: "Command Block", Name: "command_block", StackSize: 64}
	Beacon                          = Item{ID: 286, DisplayName: "Beacon", Name: "beacon", StackSize: 64}
	CobblestoneWall                 = Item{ID: 287, DisplayName: "Cobblestone Wall", Name: "cobblestone_wall", StackSize: 64}
	MossyCobblestoneWall            = Item{ID: 288, DisplayName: "Mossy Cobblestone Wall", Name: "mossy_cobblestone_wall", StackSize: 64}
	BrickWall                       = Item{ID: 289, DisplayName: "Brick Wall", Name: "brick_wall", StackSize: 64}
	PrismarineWall                  = Item{ID: 290, DisplayName: "Prismarine Wall", Name: "prismarine_wall", StackSize: 64}
	RedSandstoneWall                = Item{ID: 291, DisplayName: "Red Sandstone Wall", Name: "red_sandstone_wall", StackSize: 64}
	MossyStoneBrickWall             = Item{ID: 292, DisplayName: "Mossy Stone Brick Wall", Name: "mossy_stone_brick_wall", StackSize: 64}
	GraniteWall                     = Item{ID: 293, DisplayName: "Granite Wall", Name: "granite_wall", StackSize: 64}
	StoneBrickWall                  = Item{ID: 294, DisplayName: "Stone Brick Wall", Name: "stone_brick_wall", StackSize: 64}
	NetherBrickWall                 = Item{ID: 295, DisplayName: "Nether Brick Wall", Name: "nether_brick_wall", StackSize: 64}
	AndesiteWall                    = Item{ID: 296, DisplayName: "Andesite Wall", Name: "andesite_wall", StackSize: 64}
	RedNetherBrickWall              = Item{ID: 297, DisplayName: "Red Nether Brick Wall", Name: "red_nether_brick_wall", StackSize: 64}
	SandstoneWall                   = Item{ID: 298, DisplayName: "Sandstone Wall", Name: "sandstone_wall", StackSize: 64}
	EndStoneBrickWall               = Item{ID: 299, DisplayName: "End Stone Brick Wall", Name: "end_stone_brick_wall", StackSize: 64}
	DioriteWall                     = Item{ID: 300, DisplayName: "Diorite Wall", Name: "diorite_wall", StackSize: 64}
	BlackstoneWall                  = Item{ID: 301, DisplayName: "Blackstone Wall", Name: "blackstone_wall", StackSize: 64}
	PolishedBlackstoneWall          = Item{ID: 302, DisplayName: "Polished Blackstone Wall", Name: "polished_blackstone_wall", StackSize: 64}
	PolishedBlackstoneBrickWall     = Item{ID: 303, DisplayName: "Polished Blackstone Brick Wall", Name: "polished_blackstone_brick_wall", StackSize: 64}
	StoneButton                     = Item{ID: 304, DisplayName: "Stone Button", Name: "stone_button", StackSize: 64}
	OakButton                       = Item{ID: 305, DisplayName: "Oak Button", Name: "oak_button", StackSize: 64}
	SpruceButton                    = Item{ID: 306, DisplayName: "Spruce Button", Name: "spruce_button", StackSize: 64}
	BirchButton                     = Item{ID: 307, DisplayName: "Birch Button", Name: "birch_button", StackSize: 64}
	JungleButton                    = Item{ID: 308, DisplayName: "Jungle Button", Name: "jungle_button", StackSize: 64}
	AcaciaButton                    = Item{ID: 309, DisplayName: "Acacia Button", Name: "acacia_button", StackSize: 64}
	DarkOakButton                   = Item{ID: 310, DisplayName: "Dark Oak Button", Name: "dark_oak_button", StackSize: 64}
	CrimsonButton                   = Item{ID: 311, DisplayName: "Crimson Button", Name: "crimson_button", StackSize: 64}
	WarpedButton                    = Item{ID: 312, DisplayName: "Warped Button", Name: "warped_button", StackSize: 64}
	PolishedBlackstoneButton        = Item{ID: 313, DisplayName: "Polished Blackstone Button", Name: "polished_blackstone_button", StackSize: 64}
	Anvil                           = Item{ID: 314, DisplayName: "Anvil", Name: "anvil", StackSize: 64}
	ChippedAnvil                    = Item{ID: 315, DisplayName: "Chipped Anvil", Name: "chipped_anvil", StackSize: 64}
	DamagedAnvil                    = Item{ID: 316, DisplayName: "Damaged Anvil", Name: "damaged_anvil", StackSize: 64}
	TrappedChest                    = Item{ID: 317, DisplayName: "Trapped Chest", Name: "trapped_chest", StackSize: 64}
	LightWeightedPressurePlate      = Item{ID: 318, DisplayName: "Light Weighted Pressure Plate", Name: "light_weighted_pressure_plate", StackSize: 64}
	HeavyWeightedPressurePlate      = Item{ID: 319, DisplayName: "Heavy Weighted Pressure Plate", Name: "heavy_weighted_pressure_plate", StackSize: 64}
	DaylightDetector                = Item{ID: 320, DisplayName: "Daylight Detector", Name: "daylight_detector", StackSize: 64}
	RedstoneBlock                   = Item{ID: 321, DisplayName: "Block of Redstone", Name: "redstone_block", StackSize: 64}
	NetherQuartzOre                 = Item{ID: 322, DisplayName: "Nether Quartz Ore", Name: "nether_quartz_ore", StackSize: 64}
	Hopper                          = Item{ID: 323, DisplayName: "Hopper", Name: "hopper", StackSize: 64}
	ChiseledQuartzBlock             = Item{ID: 324, DisplayName: "Chiseled Quartz Block", Name: "chiseled_quartz_block", StackSize: 64}
	QuartzBlock                     = Item{ID: 325, DisplayName: "Block of Quartz", Name: "quartz_block", StackSize: 64}
	QuartzBricks                    = Item{ID: 326, DisplayName: "Quartz Bricks", Name: "quartz_bricks", StackSize: 64}
	QuartzPillar                    = Item{ID: 327, DisplayName: "Quartz Pillar", Name: "quartz_pillar", StackSize: 64}
	QuartzStairs                    = Item{ID: 328, DisplayName: "Quartz Stairs", Name: "quartz_stairs", StackSize: 64}
	ActivatorRail                   = Item{ID: 329, DisplayName: "Activator Rail", Name: "activator_rail", StackSize: 64}
	Dropper                         = Item{ID: 330, DisplayName: "Dropper", Name: "dropper", StackSize: 64}
	WhiteTerracotta                 = Item{ID: 331, DisplayName: "White Terracotta", Name: "white_terracotta", StackSize: 64}
	OrangeTerracotta                = Item{ID: 332, DisplayName: "Orange Terracotta", Name: "orange_terracotta", StackSize: 64}
	MagentaTerracotta               = Item{ID: 333, DisplayName: "Magenta Terracotta", Name: "magenta_terracotta", StackSize: 64}
	LightBlueTerracotta             = Item{ID: 334, DisplayName: "Light Blue Terracotta", Name: "light_blue_terracotta", StackSize: 64}
	YellowTerracotta                = Item{ID: 335, DisplayName: "Yellow Terracotta", Name: "yellow_terracotta", StackSize: 64}
	LimeTerracotta                  = Item{ID: 336, DisplayName: "Lime Terracotta", Name: "lime_terracotta", StackSize: 64}
	PinkTerracotta                  = Item{ID: 337, DisplayName: "Pink Terracotta", Name: "pink_terracotta", StackSize: 64}
	GrayTerracotta                  = Item{ID: 338, DisplayName: "Gray Terracotta", Name: "gray_terracotta", StackSize: 64}
	LightGrayTerracotta             = Item{ID: 339, DisplayName: "Light Gray Terracotta", Name: "light_gray_terracotta", StackSize: 64}
	CyanTerracotta                  = Item{ID: 340, DisplayName: "Cyan Terracotta", Name: "cyan_terracotta", StackSize: 64}
	PurpleTerracotta                = Item{ID: 341, DisplayName: "Purple Terracotta", Name: "purple_terracotta", StackSize: 64}
	BlueTerracotta                  = Item{ID: 342, DisplayName: "Blue Terracotta", Name: "blue_terracotta", StackSize: 64}
	BrownTerracotta                 = Item{ID: 343, DisplayName: "Brown Terracotta", Name: "brown_terracotta", StackSize: 64}
	GreenTerracotta                 = Item{ID: 344, DisplayName: "Green Terracotta", Name: "green_terracotta", StackSize: 64}
	RedTerracotta                   = Item{ID: 345, DisplayName: "Red Terracotta", Name: "red_terracotta", StackSize: 64}
	BlackTerracotta                 = Item{ID: 346, DisplayName: "Black Terracotta", Name: "black_terracotta", StackSize: 64}
	Barrier                         = Item{ID: 347, DisplayName: "Barrier", Name: "barrier", StackSize: 64}
	IronTrapdoor                    = Item{ID: 348, DisplayName: "Iron Trapdoor", Name: "iron_trapdoor", StackSize: 64}
	HayBlock                        = Item{ID: 349, DisplayName: "Hay Bale", Name: "hay_block", StackSize: 64}
	WhiteCarpet                     = Item{ID: 350, DisplayName: "White Carpet", Name: "white_carpet", StackSize: 64}
	OrangeCarpet                    = Item{ID: 351, DisplayName: "Orange Carpet", Name: "orange_carpet", StackSize: 64}
	MagentaCarpet                   = Item{ID: 352, DisplayName: "Magenta Carpet", Name: "magenta_carpet", StackSize: 64}
	LightBlueCarpet                 = Item{ID: 353, DisplayName: "Light Blue Carpet", Name: "light_blue_carpet", StackSize: 64}
	YellowCarpet                    = Item{ID: 354, DisplayName: "Yellow Carpet", Name: "yellow_carpet", StackSize: 64}
	LimeCarpet                      = Item{ID: 355, DisplayName: "Lime Carpet", Name: "lime_carpet", StackSize: 64}
	PinkCarpet                      = Item{ID: 356, DisplayName: "Pink Carpet", Name: "pink_carpet", StackSize: 64}
	GrayCarpet                      = Item{ID: 357, DisplayName: "Gray Carpet", Name: "gray_carpet", StackSize: 64}
	LightGrayCarpet                 = Item{ID: 358, DisplayName: "Light Gray Carpet", Name: "light_gray_carpet", StackSize: 64}
	CyanCarpet                      = Item{ID: 359, DisplayName: "Cyan Carpet", Name: "cyan_carpet", StackSize: 64}
	PurpleCarpet                    = Item{ID: 360, DisplayName: "Purple Carpet", Name: "purple_carpet", StackSize: 64}
	BlueCarpet                      = Item{ID: 361, DisplayName: "Blue Carpet", Name: "blue_carpet", StackSize: 64}
	BrownCarpet                     = Item{ID: 362, DisplayName: "Brown Carpet", Name: "brown_carpet", StackSize: 64}
	GreenCarpet                     = Item{ID: 363, DisplayName: "Green Carpet", Name: "green_carpet", StackSize: 64}
	RedCarpet                       = Item{ID: 364, DisplayName: "Red Carpet", Name: "red_carpet", StackSize: 64}
	BlackCarpet                     = Item{ID: 365, DisplayName: "Black Carpet", Name: "black_carpet", StackSize: 64}
	Terracotta                      = Item{ID: 366, DisplayName: "Terracotta", Name: "terracotta", StackSize: 64}
	CoalBlock                       = Item{ID: 367, DisplayName: "Block of Coal", Name: "coal_block", StackSize: 64}
	PackedIce                       = Item{ID: 368, DisplayName: "Packed Ice", Name: "packed_ice", StackSize: 64}
	AcaciaStairs                    = Item{ID: 369, DisplayName: "Acacia Stairs", Name: "acacia_stairs", StackSize: 64}
	DarkOakStairs                   = Item{ID: 370, DisplayName: "Dark Oak Stairs", Name: "dark_oak_stairs", StackSize: 64}
	SlimeBlock                      = Item{ID: 371, DisplayName: "Slime Block", Name: "slime_block", StackSize: 64}
	GrassPath                       = Item{ID: 372, DisplayName: "Grass Path", Name: "grass_path", StackSize: 64}
	Sunflower                       = Item{ID: 373, DisplayName: "Sunflower", Name: "sunflower", StackSize: 64}
	Lilac                           = Item{ID: 374, DisplayName: "Lilac", Name: "lilac", StackSize: 64}
	RoseBush                        = Item{ID: 375, DisplayName: "Rose Bush", Name: "rose_bush", StackSize: 64}
	Peony                           = Item{ID: 376, DisplayName: "Peony", Name: "peony", StackSize: 64}
	TallGrass                       = Item{ID: 377, DisplayName: "Tall Grass", Name: "tall_grass", StackSize: 64}
	LargeFern                       = Item{ID: 378, DisplayName: "Large Fern", Name: "large_fern", StackSize: 64}
	WhiteStainedGlass               = Item{ID: 379, DisplayName: "White Stained Glass", Name: "white_stained_glass", StackSize: 64}
	OrangeStainedGlass              = Item{ID: 380, DisplayName: "Orange Stained Glass", Name: "orange_stained_glass", StackSize: 64}
	MagentaStainedGlass             = Item{ID: 381, DisplayName: "Magenta Stained Glass", Name: "magenta_stained_glass", StackSize: 64}
	LightBlueStainedGlass           = Item{ID: 382, DisplayName: "Light Blue Stained Glass", Name: "light_blue_stained_glass", StackSize: 64}
	YellowStainedGlass              = Item{ID: 383, DisplayName: "Yellow Stained Glass", Name: "yellow_stained_glass", StackSize: 64}
	LimeStainedGlass                = Item{ID: 384, DisplayName: "Lime Stained Glass", Name: "lime_stained_glass", StackSize: 64}
	PinkStainedGlass                = Item{ID: 385, DisplayName: "Pink Stained Glass", Name: "pink_stained_glass", StackSize: 64}
	GrayStainedGlass                = Item{ID: 386, DisplayName: "Gray Stained Glass", Name: "gray_stained_glass", StackSize: 64}
	LightGrayStainedGlass           = Item{ID: 387, DisplayName: "Light Gray Stained Glass", Name: "light_gray_stained_glass", StackSize: 64}
	CyanStainedGlass                = Item{ID: 388, DisplayName: "Cyan Stained Glass", Name: "cyan_stained_glass", StackSize: 64}
	PurpleStainedGlass              = Item{ID: 389, DisplayName: "Purple Stained Glass", Name: "purple_stained_glass", StackSize: 64}
	BlueStainedGlass                = Item{ID: 390, DisplayName: "Blue Stained Glass", Name: "blue_stained_glass", StackSize: 64}
	BrownStainedGlass               = Item{ID: 391, DisplayName: "Brown Stained Glass", Name: "brown_stained_glass", StackSize: 64}
	GreenStainedGlass               = Item{ID: 392, DisplayName: "Green Stained Glass", Name: "green_stained_glass", StackSize: 64}
	RedStainedGlass                 = Item{ID: 393, DisplayName: "Red Stained Glass", Name: "red_stained_glass", StackSize: 64}
	BlackStainedGlass               = Item{ID: 394, DisplayName: "Black Stained Glass", Name: "black_stained_glass", StackSize: 64}
	WhiteStainedGlassPane           = Item{ID: 395, DisplayName: "White Stained Glass Pane", Name: "white_stained_glass_pane", StackSize: 64}
	OrangeStainedGlassPane          = Item{ID: 396, DisplayName: "Orange Stained Glass Pane", Name: "orange_stained_glass_pane", StackSize: 64}
	MagentaStainedGlassPane         = Item{ID: 397, DisplayName: "Magenta Stained Glass Pane", Name: "magenta_stained_glass_pane", StackSize: 64}
	LightBlueStainedGlassPane       = Item{ID: 398, DisplayName: "Light Blue Stained Glass Pane", Name: "light_blue_stained_glass_pane", StackSize: 64}
	YellowStainedGlassPane          = Item{ID: 399, DisplayName: "Yellow Stained Glass Pane", Name: "yellow_stained_glass_pane", StackSize: 64}
	LimeStainedGlassPane            = Item{ID: 400, DisplayName: "Lime Stained Glass Pane", Name: "lime_stained_glass_pane", StackSize: 64}
	PinkStainedGlassPane            = Item{ID: 401, DisplayName: "Pink Stained Glass Pane", Name: "pink_stained_glass_pane", StackSize: 64}
	GrayStainedGlassPane            = Item{ID: 402, DisplayName: "Gray Stained Glass Pane", Name: "gray_stained_glass_pane", StackSize: 64}
	LightGrayStainedGlassPane       = Item{ID: 403, DisplayName: "Light Gray Stained Glass Pane", Name: "light_gray_stained_glass_pane", StackSize: 64}
	CyanStainedGlassPane            = Item{ID: 404, DisplayName: "Cyan Stained Glass Pane", Name: "cyan_stained_glass_pane", StackSize: 64}
	PurpleStainedGlassPane          = Item{ID: 405, DisplayName: "Purple Stained Glass Pane", Name: "purple_stained_glass_pane", StackSize: 64}
	BlueStainedGlassPane            = Item{ID: 406, DisplayName: "Blue Stained Glass Pane", Name: "blue_stained_glass_pane", StackSize: 64}
	BrownStainedGlassPane           = Item{ID: 407, DisplayName: "Brown Stained Glass Pane", Name: "brown_stained_glass_pane", StackSize: 64}
	GreenStainedGlassPane           = Item{ID: 408, DisplayName: "Green Stained Glass Pane", Name: "green_stained_glass_pane", StackSize: 64}
	RedStainedGlassPane             = Item{ID: 409, DisplayName: "Red Stained Glass Pane", Name: "red_stained_glass_pane", StackSize: 64}
	BlackStainedGlassPane           = Item{ID: 410, DisplayName: "Black Stained Glass Pane", Name: "black_stained_glass_pane", StackSize: 64}
	Prismarine                      = Item{ID: 411, DisplayName: "Prismarine", Name: "prismarine", StackSize: 64}
	PrismarineBricks                = Item{ID: 412, DisplayName: "Prismarine Bricks", Name: "prismarine_bricks", StackSize: 64}
	DarkPrismarine                  = Item{ID: 413, DisplayName: "Dark Prismarine", Name: "dark_prismarine", StackSize: 64}
	PrismarineStairs                = Item{ID: 414, DisplayName: "Prismarine Stairs", Name: "prismarine_stairs", StackSize: 64}
	PrismarineBrickStairs           = Item{ID: 415, DisplayName: "Prismarine Brick Stairs", Name: "prismarine_brick_stairs", StackSize: 64}
	DarkPrismarineStairs            = Item{ID: 416, DisplayName: "Dark Prismarine Stairs", Name: "dark_prismarine_stairs", StackSize: 64}
	SeaLantern                      = Item{ID: 417, DisplayName: "Sea Lantern", Name: "sea_lantern", StackSize: 64}
	RedSandstone                    = Item{ID: 418, DisplayName: "Red Sandstone", Name: "red_sandstone", StackSize: 64}
	ChiseledRedSandstone            = Item{ID: 419, DisplayName: "Chiseled Red Sandstone", Name: "chiseled_red_sandstone", StackSize: 64}
	CutRedSandstone                 = Item{ID: 420, DisplayName: "Cut Red Sandstone", Name: "cut_red_sandstone", StackSize: 64}
	RedSandstoneStairs              = Item{ID: 421, DisplayName: "Red Sandstone Stairs", Name: "red_sandstone_stairs", StackSize: 64}
	RepeatingCommandBlock           = Item{ID: 422, DisplayName: "Repeating Command Block", Name: "repeating_command_block", StackSize: 64}
	ChainCommandBlock               = Item{ID: 423, DisplayName: "Chain Command Block", Name: "chain_command_block", StackSize: 64}
	MagmaBlock                      = Item{ID: 424, DisplayName: "Magma Block", Name: "magma_block", StackSize: 64}
	NetherWartBlock                 = Item{ID: 425, DisplayName: "Nether Wart Block", Name: "nether_wart_block", StackSize: 64}
	WarpedWartBlock                 = Item{ID: 426, DisplayName: "Warped Wart Block", Name: "warped_wart_block", StackSize: 64}
	RedNetherBricks                 = Item{ID: 427, DisplayName: "Red Nether Bricks", Name: "red_nether_bricks", StackSize: 64}
	BoneBlock                       = Item{ID: 428, DisplayName: "Bone Block", Name: "bone_block", StackSize: 64}
	StructureVoid                   = Item{ID: 429, DisplayName: "Structure Void", Name: "structure_void", StackSize: 64}
	Observer                        = Item{ID: 430, DisplayName: "Observer", Name: "observer", StackSize: 64}
	ShulkerBox                      = Item{ID: 431, DisplayName: "Shulker Box", Name: "shulker_box", StackSize: 1}
	WhiteShulkerBox                 = Item{ID: 432, DisplayName: "White Shulker Box", Name: "white_shulker_box", StackSize: 1}
	OrangeShulkerBox                = Item{ID: 433, DisplayName: "Orange Shulker Box", Name: "orange_shulker_box", StackSize: 1}
	MagentaShulkerBox               = Item{ID: 434, DisplayName: "Magenta Shulker Box", Name: "magenta_shulker_box", StackSize: 1}
	LightBlueShulkerBox             = Item{ID: 435, DisplayName: "Light Blue Shulker Box", Name: "light_blue_shulker_box", StackSize: 1}
	YellowShulkerBox                = Item{ID: 436, DisplayName: "Yellow Shulker Box", Name: "yellow_shulker_box", StackSize: 1}
	LimeShulkerBox                  = Item{ID: 437, DisplayName: "Lime Shulker Box", Name: "lime_shulker_box", StackSize: 1}
	PinkShulkerBox                  = Item{ID: 438, DisplayName: "Pink Shulker Box", Name: "pink_shulker_box", StackSize: 1}
	GrayShulkerBox                  = Item{ID: 439, DisplayName: "Gray Shulker Box", Name: "gray_shulker_box", StackSize: 1}
	LightGrayShulkerBox             = Item{ID: 440, DisplayName: "Light Gray Shulker Box", Name: "light_gray_shulker_box", StackSize: 1}
	CyanShulkerBox                  = Item{ID: 441, DisplayName: "Cyan Shulker Box", Name: "cyan_shulker_box", StackSize: 1}
	PurpleShulkerBox                = Item{ID: 442, DisplayName: "Purple Shulker Box", Name: "purple_shulker_box", StackSize: 1}
	BlueShulkerBox                  = Item{ID: 443, DisplayName: "Blue Shulker Box", Name: "blue_shulker_box", StackSize: 1}
	BrownShulkerBox                 = Item{ID: 444, DisplayName: "Brown Shulker Box", Name: "brown_shulker_box", StackSize: 1}
	GreenShulkerBox                 = Item{ID: 445, DisplayName: "Green Shulker Box", Name: "green_shulker_box", StackSize: 1}
	RedShulkerBox                   = Item{ID: 446, DisplayName: "Red Shulker Box", Name: "red_shulker_box", StackSize: 1}
	BlackShulkerBox                 = Item{ID: 447, DisplayName: "Black Shulker Box", Name: "black_shulker_box", StackSize: 1}
	WhiteGlazedTerracotta           = Item{ID: 448, DisplayName: "White Glazed Terracotta", Name: "white_glazed_terracotta", StackSize: 64}
	OrangeGlazedTerracotta          = Item{ID: 449, DisplayName: "Orange Glazed Terracotta", Name: "orange_glazed_terracotta", StackSize: 64}
	MagentaGlazedTerracotta         = Item{ID: 450, DisplayName: "Magenta Glazed Terracotta", Name: "magenta_glazed_terracotta", StackSize: 64}
	LightBlueGlazedTerracotta       = Item{ID: 451, DisplayName: "Light Blue Glazed Terracotta", Name: "light_blue_glazed_terracotta", StackSize: 64}
	YellowGlazedTerracotta          = Item{ID: 452, DisplayName: "Yellow Glazed Terracotta", Name: "yellow_glazed_terracotta", StackSize: 64}
	LimeGlazedTerracotta            = Item{ID: 453, DisplayName: "Lime Glazed Terracotta", Name: "lime_glazed_terracotta", StackSize: 64}
	PinkGlazedTerracotta            = Item{ID: 454, DisplayName: "Pink Glazed Terracotta", Name: "pink_glazed_terracotta", StackSize: 64}
	GrayGlazedTerracotta            = Item{ID: 455, DisplayName: "Gray Glazed Terracotta", Name: "gray_glazed_terracotta", StackSize: 64}
	LightGrayGlazedTerracotta       = Item{ID: 456, DisplayName: "Light Gray Glazed Terracotta", Name: "light_gray_glazed_terracotta", StackSize: 64}
	CyanGlazedTerracotta            = Item{ID: 457, DisplayName: "Cyan Glazed Terracotta", Name: "cyan_glazed_terracotta", StackSize: 64}
	PurpleGlazedTerracotta          = Item{ID: 458, DisplayName: "Purple Glazed Terracotta", Name: "purple_glazed_terracotta", StackSize: 64}
	BlueGlazedTerracotta            = Item{ID: 459, DisplayName: "Blue Glazed Terracotta", Name: "blue_glazed_terracotta", StackSize: 64}
	BrownGlazedTerracotta           = Item{ID: 460, DisplayName: "Brown Glazed Terracotta", Name: "brown_glazed_terracotta", StackSize: 64}
	GreenGlazedTerracotta           = Item{ID: 461, DisplayName: "Green Glazed Terracotta", Name: "green_glazed_terracotta", StackSize: 64}
	RedGlazedTerracotta             = Item{ID: 462, DisplayName: "Red Glazed Terracotta", Name: "red_glazed_terracotta", StackSize: 64}
	BlackGlazedTerracotta           = Item{ID: 463, DisplayName: "Black Glazed Terracotta", Name: "black_glazed_terracotta", StackSize: 64}
	WhiteConcrete                   = Item{ID: 464, DisplayName: "White Concrete", Name: "white_concrete", StackSize: 64}
	OrangeConcrete                  = Item{ID: 465, DisplayName: "Orange Concrete", Name: "orange_concrete", StackSize: 64}
	MagentaConcrete                 = Item{ID: 466, DisplayName: "Magenta Concrete", Name: "magenta_concrete", StackSize: 64}
	LightBlueConcrete               = Item{ID: 467, DisplayName: "Light Blue Concrete", Name: "light_blue_concrete", StackSize: 64}
	YellowConcrete                  = Item{ID: 468, DisplayName: "Yellow Concrete", Name: "yellow_concrete", StackSize: 64}
	LimeConcrete                    = Item{ID: 469, DisplayName: "Lime Concrete", Name: "lime_concrete", StackSize: 64}
	PinkConcrete                    = Item{ID: 470, DisplayName: "Pink Concrete", Name: "pink_concrete", StackSize: 64}
	GrayConcrete                    = Item{ID: 471, DisplayName: "Gray Concrete", Name: "gray_concrete", StackSize: 64}
	LightGrayConcrete               = Item{ID: 472, DisplayName: "Light Gray Concrete", Name: "light_gray_concrete", StackSize: 64}
	CyanConcrete                    = Item{ID: 473, DisplayName: "Cyan Concrete", Name: "cyan_concrete", StackSize: 64}
	PurpleConcrete                  = Item{ID: 474, DisplayName: "Purple Concrete", Name: "purple_concrete", StackSize: 64}
	BlueConcrete                    = Item{ID: 475, DisplayName: "Blue Concrete", Name: "blue_concrete", StackSize: 64}
	BrownConcrete                   = Item{ID: 476, DisplayName: "Brown Concrete", Name: "brown_concrete", StackSize: 64}
	GreenConcrete                   = Item{ID: 477, DisplayName: "Green Concrete", Name: "green_concrete", StackSize: 64}
	RedConcrete                     = Item{ID: 478, DisplayName: "Red Concrete", Name: "red_concrete", StackSize: 64}
	BlackConcrete                   = Item{ID: 479, DisplayName: "Black Concrete", Name: "black_concrete", StackSize: 64}
	WhiteConcretePowder             = Item{ID: 480, DisplayName: "White Concrete Powder", Name: "white_concrete_powder", StackSize: 64}
	OrangeConcretePowder            = Item{ID: 481, DisplayName: "Orange Concrete Powder", Name: "orange_concrete_powder", StackSize: 64}
	MagentaConcretePowder           = Item{ID: 482, DisplayName: "Magenta Concrete Powder", Name: "magenta_concrete_powder", StackSize: 64}
	LightBlueConcretePowder         = Item{ID: 483, DisplayName: "Light Blue Concrete Powder", Name: "light_blue_concrete_powder", StackSize: 64}
	YellowConcretePowder            = Item{ID: 484, DisplayName: "Yellow Concrete Powder", Name: "yellow_concrete_powder", StackSize: 64}
	LimeConcretePowder              = Item{ID: 485, DisplayName: "Lime Concrete Powder", Name: "lime_concrete_powder", StackSize: 64}
	PinkConcretePowder              = Item{ID: 486, DisplayName: "Pink Concrete Powder", Name: "pink_concrete_powder", StackSize: 64}
	GrayConcretePowder              = Item{ID: 487, DisplayName: "Gray Concrete Powder", Name: "gray_concrete_powder", StackSize: 64}
	LightGrayConcretePowder         = Item{ID: 488, DisplayName: "Light Gray Concrete Powder", Name: "light_gray_concrete_powder", StackSize: 64}
	CyanConcretePowder              = Item{ID: 489, DisplayName: "Cyan Concrete Powder", Name: "cyan_concrete_powder", StackSize: 64}
	PurpleConcretePowder            = Item{ID: 490, DisplayName: "Purple Concrete Powder", Name: "purple_concrete_powder", StackSize: 64}
	BlueConcretePowder              = Item{ID: 491, DisplayName: "Blue Concrete Powder", Name: "blue_concrete_powder", StackSize: 64}
	BrownConcretePowder             = Item{ID: 492, DisplayName: "Brown Concrete Powder", Name: "brown_concrete_powder", StackSize: 64}
	GreenConcretePowder             = Item{ID: 493, DisplayName: "Green Concrete Powder", Name: "green_concrete_powder", StackSize: 64}
	RedConcretePowder               = Item{ID: 494, DisplayName: "Red Concrete Powder", Name: "red_concrete_powder", StackSize: 64}
	BlackConcretePowder             = Item{ID: 495, DisplayName: "Black Concrete Powder", Name: "black_concrete_powder", StackSize: 64}
	TurtleEgg                       = Item{ID: 496, DisplayName: "Turtle Egg", Name: "turtle_egg", StackSize: 64}
	DeadTubeCoralBlock              = Item{ID: 497, DisplayName: "Dead Tube Coral Block", Name: "dead_tube_coral_block", StackSize: 64}
	DeadBrainCoralBlock             = Item{ID: 498, DisplayName: "Dead Brain Coral Block", Name: "dead_brain_coral_block", StackSize: 64}
	DeadBubbleCoralBlock            = Item{ID: 499, DisplayName: "Dead Bubble Coral Block", Name: "dead_bubble_coral_block", StackSize: 64}
	DeadFireCoralBlock              = Item{ID: 500, DisplayName: "Dead Fire Coral Block", Name: "dead_fire_coral_block", StackSize: 64}
	DeadHornCoralBlock              = Item{ID: 501, DisplayName: "Dead Horn Coral Block", Name: "dead_horn_coral_block", StackSize: 64}
	TubeCoralBlock                  = Item{ID: 502, DisplayName: "Tube Coral Block", Name: "tube_coral_block", StackSize: 64}
	BrainCoralBlock                 = Item{ID: 503, DisplayName: "Brain Coral Block", Name: "brain_coral_block", StackSize: 64}
	BubbleCoralBlock                = Item{ID: 504, DisplayName: "Bubble Coral Block", Name: "bubble_coral_block", StackSize: 64}
	FireCoralBlock                  = Item{ID: 505, DisplayName: "Fire Coral Block", Name: "fire_coral_block", StackSize: 64}
	HornCoralBlock                  = Item{ID: 506, DisplayName: "Horn Coral Block", Name: "horn_coral_block", StackSize: 64}
	TubeCoral                       = Item{ID: 507, DisplayName: "Tube Coral", Name: "tube_coral", StackSize: 64}
	BrainCoral                      = Item{ID: 508, DisplayName: "Brain Coral", Name: "brain_coral", StackSize: 64}
	BubbleCoral                     = Item{ID: 509, DisplayName: "Bubble Coral", Name: "bubble_coral", StackSize: 64}
	FireCoral                       = Item{ID: 510, DisplayName: "Fire Coral", Name: "fire_coral", StackSize: 64}
	HornCoral                       = Item{ID: 511, DisplayName: "Horn Coral", Name: "horn_coral", StackSize: 64}
	DeadBrainCoral                  = Item{ID: 512, DisplayName: "Dead Brain Coral", Name: "dead_brain_coral", StackSize: 64}
	DeadBubbleCoral                 = Item{ID: 513, DisplayName: "Dead Bubble Coral", Name: "dead_bubble_coral", StackSize: 64}
	DeadFireCoral                   = Item{ID: 514, DisplayName: "Dead Fire Coral", Name: "dead_fire_coral", StackSize: 64}
	DeadHornCoral                   = Item{ID: 515, DisplayName: "Dead Horn Coral", Name: "dead_horn_coral", StackSize: 64}
	DeadTubeCoral                   = Item{ID: 516, DisplayName: "Dead Tube Coral", Name: "dead_tube_coral", StackSize: 64}
	TubeCoralFan                    = Item{ID: 517, DisplayName: "Tube Coral Fan", Name: "tube_coral_fan", StackSize: 64}
	BrainCoralFan                   = Item{ID: 518, DisplayName: "Brain Coral Fan", Name: "brain_coral_fan", StackSize: 64}
	BubbleCoralFan                  = Item{ID: 519, DisplayName: "Bubble Coral Fan", Name: "bubble_coral_fan", StackSize: 64}
	FireCoralFan                    = Item{ID: 520, DisplayName: "Fire Coral Fan", Name: "fire_coral_fan", StackSize: 64}
	HornCoralFan                    = Item{ID: 521, DisplayName: "Horn Coral Fan", Name: "horn_coral_fan", StackSize: 64}
	DeadTubeCoralFan                = Item{ID: 522, DisplayName: "Dead Tube Coral Fan", Name: "dead_tube_coral_fan", StackSize: 64}
	DeadBrainCoralFan               = Item{ID: 523, DisplayName: "Dead Brain Coral Fan", Name: "dead_brain_coral_fan", StackSize: 64}
	DeadBubbleCoralFan              = Item{ID: 524, DisplayName: "Dead Bubble Coral Fan", Name: "dead_bubble_coral_fan", StackSize: 64}
	DeadFireCoralFan                = Item{ID: 525, DisplayName: "Dead Fire Coral Fan", Name: "dead_fire_coral_fan", StackSize: 64}
	DeadHornCoralFan                = Item{ID: 526, DisplayName: "Dead Horn Coral Fan", Name: "dead_horn_coral_fan", StackSize: 64}
	BlueIce                         = Item{ID: 527, DisplayName: "Blue Ice", Name: "blue_ice", StackSize: 64}
	Conduit                         = Item{ID: 528, DisplayName: "Conduit", Name: "conduit", StackSize: 64}
	PolishedGraniteStairs           = Item{ID: 529, DisplayName: "Polished Granite Stairs", Name: "polished_granite_stairs", StackSize: 64}
	SmoothRedSandstoneStairs        = Item{ID: 530, DisplayName: "Smooth Red Sandstone Stairs", Name: "smooth_red_sandstone_stairs", StackSize: 64}
	MossyStoneBrickStairs           = Item{ID: 531, DisplayName: "Mossy Stone Brick Stairs", Name: "mossy_stone_brick_stairs", StackSize: 64}
	PolishedDioriteStairs           = Item{ID: 532, DisplayName: "Polished Diorite Stairs", Name: "polished_diorite_stairs", StackSize: 64}
	MossyCobblestoneStairs          = Item{ID: 533, DisplayName: "Mossy Cobblestone Stairs", Name: "mossy_cobblestone_stairs", StackSize: 64}
	EndStoneBrickStairs             = Item{ID: 534, DisplayName: "End Stone Brick Stairs", Name: "end_stone_brick_stairs", StackSize: 64}
	StoneStairs                     = Item{ID: 535, DisplayName: "Stone Stairs", Name: "stone_stairs", StackSize: 64}
	SmoothSandstoneStairs           = Item{ID: 536, DisplayName: "Smooth Sandstone Stairs", Name: "smooth_sandstone_stairs", StackSize: 64}
	SmoothQuartzStairs              = Item{ID: 537, DisplayName: "Smooth Quartz Stairs", Name: "smooth_quartz_stairs", StackSize: 64}
	GraniteStairs                   = Item{ID: 538, DisplayName: "Granite Stairs", Name: "granite_stairs", StackSize: 64}
	AndesiteStairs                  = Item{ID: 539, DisplayName: "Andesite Stairs", Name: "andesite_stairs", StackSize: 64}
	RedNetherBrickStairs            = Item{ID: 540, DisplayName: "Red Nether Brick Stairs", Name: "red_nether_brick_stairs", StackSize: 64}
	PolishedAndesiteStairs          = Item{ID: 541, DisplayName: "Polished Andesite Stairs", Name: "polished_andesite_stairs", StackSize: 64}
	DioriteStairs                   = Item{ID: 542, DisplayName: "Diorite Stairs", Name: "diorite_stairs", StackSize: 64}
	PolishedGraniteSlab             = Item{ID: 543, DisplayName: "Polished Granite Slab", Name: "polished_granite_slab", StackSize: 64}
	SmoothRedSandstoneSlab          = Item{ID: 544, DisplayName: "Smooth Red Sandstone Slab", Name: "smooth_red_sandstone_slab", StackSize: 64}
	MossyStoneBrickSlab             = Item{ID: 545, DisplayName: "Mossy Stone Brick Slab", Name: "mossy_stone_brick_slab", StackSize: 64}
	PolishedDioriteSlab             = Item{ID: 546, DisplayName: "Polished Diorite Slab", Name: "polished_diorite_slab", StackSize: 64}
	MossyCobblestoneSlab            = Item{ID: 547, DisplayName: "Mossy Cobblestone Slab", Name: "mossy_cobblestone_slab", StackSize: 64}
	EndStoneBrickSlab               = Item{ID: 548, DisplayName: "End Stone Brick Slab", Name: "end_stone_brick_slab", StackSize: 64}
	SmoothSandstoneSlab             = Item{ID: 549, DisplayName: "Smooth Sandstone Slab", Name: "smooth_sandstone_slab", StackSize: 64}
	SmoothQuartzSlab                = Item{ID: 550, DisplayName: "Smooth Quartz Slab", Name: "smooth_quartz_slab", StackSize: 64}
	GraniteSlab                     = Item{ID: 551, DisplayName: "Granite Slab", Name: "granite_slab", StackSize: 64}
	AndesiteSlab                    = Item{ID: 552, DisplayName: "Andesite Slab", Name: "andesite_slab", StackSize: 64}
	RedNetherBrickSlab              = Item{ID: 553, DisplayName: "Red Nether Brick Slab", Name: "red_nether_brick_slab", StackSize: 64}
	PolishedAndesiteSlab            = Item{ID: 554, DisplayName: "Polished Andesite Slab", Name: "polished_andesite_slab", StackSize: 64}
	DioriteSlab                     = Item{ID: 555, DisplayName: "Diorite Slab", Name: "diorite_slab", StackSize: 64}
	Scaffolding                     = Item{ID: 556, DisplayName: "Scaffolding", Name: "scaffolding", StackSize: 64}
	IronDoor                        = Item{ID: 557, DisplayName: "Iron Door", Name: "iron_door", StackSize: 64}
	OakDoor                         = Item{ID: 558, DisplayName: "Oak Door", Name: "oak_door", StackSize: 64}
	SpruceDoor                      = Item{ID: 559, DisplayName: "Spruce Door", Name: "spruce_door", StackSize: 64}
	BirchDoor                       = Item{ID: 560, DisplayName: "Birch Door", Name: "birch_door", StackSize: 64}
	JungleDoor                      = Item{ID: 561, DisplayName: "Jungle Door", Name: "jungle_door", StackSize: 64}
	AcaciaDoor                      = Item{ID: 562, DisplayName: "Acacia Door", Name: "acacia_door", StackSize: 64}
	DarkOakDoor                     = Item{ID: 563, DisplayName: "Dark Oak Door", Name: "dark_oak_door", StackSize: 64}
	CrimsonDoor                     = Item{ID: 564, DisplayName: "Crimson Door", Name: "crimson_door", StackSize: 64}
	WarpedDoor                      = Item{ID: 565, DisplayName: "Warped Door", Name: "warped_door", StackSize: 64}
	Repeater                        = Item{ID: 566, DisplayName: "Redstone Repeater", Name: "repeater", StackSize: 64}
	Comparator                      = Item{ID: 567, DisplayName: "Redstone Comparator", Name: "comparator", StackSize: 64}
	StructureBlock                  = Item{ID: 568, DisplayName: "Structure Block", Name: "structure_block", StackSize: 64}
	Jigsaw                          = Item{ID: 569, DisplayName: "Jigsaw Block", Name: "jigsaw", StackSize: 64}
	TurtleHelmet                    = Item{ID: 570, DisplayName: "Turtle Shell", Name: "turtle_helmet", StackSize: 1}
	Scute                           = Item{ID: 571, DisplayName: "Scute", Name: "scute", StackSize: 64}
	FlintAndSteel                   = Item{ID: 572, DisplayName: "Flint and Steel", Name: "flint_and_steel", StackSize: 1}
	Apple                           = Item{ID: 573, DisplayName: "Apple", Name: "apple", StackSize: 64}
	Bow                             = Item{ID: 574, DisplayName: "Bow", Name: "bow", StackSize: 1}
	Arrow                           = Item{ID: 575, DisplayName: "Arrow", Name: "arrow", StackSize: 64}
	Coal                            = Item{ID: 576, DisplayName: "Coal", Name: "coal", StackSize: 64}
	Charcoal                        = Item{ID: 577, DisplayName: "Charcoal", Name: "charcoal", StackSize: 64}
	Diamond                         = Item{ID: 578, DisplayName: "Diamond", Name: "diamond", StackSize: 64}
	IronIngot                       = Item{ID: 579, DisplayName: "Iron Ingot", Name: "iron_ingot", StackSize: 64}
	GoldIngot                       = Item{ID: 580, DisplayName: "Gold Ingot", Name: "gold_ingot", StackSize: 64}
	NetheriteIngot                  = Item{ID: 581, DisplayName: "Netherite Ingot", Name: "netherite_ingot", StackSize: 64}
	NetheriteScrap                  = Item{ID: 582, DisplayName: "Netherite Scrap", Name: "netherite_scrap", StackSize: 64}
	WoodenSword                     = Item{ID: 583, DisplayName: "Wooden Sword", Name: "wooden_sword", StackSize: 1}
	WoodenShovel                    = Item{ID: 584, DisplayName: "Wooden Shovel", Name: "wooden_shovel", StackSize: 1}
	WoodenPickaxe                   = Item{ID: 585, DisplayName: "Wooden Pickaxe", Name: "wooden_pickaxe", StackSize: 1}
	WoodenAxe                       = Item{ID: 586, DisplayName: "Wooden Axe", Name: "wooden_axe", StackSize: 1}
	WoodenHoe                       = Item{ID: 587, DisplayName: "Wooden Hoe", Name: "wooden_hoe", StackSize: 1}
	StoneSword                      = Item{ID: 588, DisplayName: "Stone Sword", Name: "stone_sword", StackSize: 1}
	StoneShovel                     = Item{ID: 589, DisplayName: "Stone Shovel", Name: "stone_shovel", StackSize: 1}
	StonePickaxe                    = Item{ID: 590, DisplayName: "Stone Pickaxe", Name: "stone_pickaxe", StackSize: 1}
	StoneAxe                        = Item{ID: 591, DisplayName: "Stone Axe", Name: "stone_axe", StackSize: 1}
	StoneHoe                        = Item{ID: 592, DisplayName: "Stone Hoe", Name: "stone_hoe", StackSize: 1}
	GoldenSword                     = Item{ID: 593, DisplayName: "Golden Sword", Name: "golden_sword", StackSize: 1}
	GoldenShovel                    = Item{ID: 594, DisplayName: "Golden Shovel", Name: "golden_shovel", StackSize: 1}
	GoldenPickaxe                   = Item{ID: 595, DisplayName: "Golden Pickaxe", Name: "golden_pickaxe", StackSize: 1}
	GoldenAxe                       = Item{ID: 596, DisplayName: "Golden Axe", Name: "golden_axe", StackSize: 1}
	GoldenHoe                       = Item{ID: 597, DisplayName: "Golden Hoe", Name: "golden_hoe", StackSize: 1}
	IronSword                       = Item{ID: 598, DisplayName: "Iron Sword", Name: "iron_sword", StackSize: 1}
	IronShovel                      = Item{ID: 599, DisplayName: "Iron Shovel", Name: "iron_shovel", StackSize: 1}
	IronPickaxe                     = Item{ID: 600, DisplayName: "Iron Pickaxe", Name: "iron_pickaxe", StackSize: 1}
	IronAxe                         = Item{ID: 601, DisplayName: "Iron Axe", Name: "iron_axe", StackSize: 1}
	IronHoe                         = Item{ID: 602, DisplayName: "Iron Hoe", Name: "iron_hoe", StackSize: 1}
	DiamondSword                    = Item{ID: 603, DisplayName: "Diamond Sword", Name: "diamond_sword", StackSize: 1}
	DiamondShovel                   = Item{ID: 604, DisplayName: "Diamond Shovel", Name: "diamond_shovel", StackSize: 1}
	DiamondPickaxe                  = Item{ID: 605, DisplayName: "Diamond Pickaxe", Name: "diamond_pickaxe", StackSize: 1}
	DiamondAxe                      = Item{ID: 606, DisplayName: "Diamond Axe", Name: "diamond_axe", StackSize: 1}
	DiamondHoe                      = Item{ID: 607, DisplayName: "Diamond Hoe", Name: "diamond_hoe", StackSize: 1}
	NetheriteSword                  = Item{ID: 608, DisplayName: "Netherite Sword", Name: "netherite_sword", StackSize: 1}
	NetheriteShovel                 = Item{ID: 609, DisplayName: "Netherite Shovel", Name: "netherite_shovel", StackSize: 1}
	NetheritePickaxe                = Item{ID: 610, DisplayName: "Netherite Pickaxe", Name: "netherite_pickaxe", StackSize: 1}
	NetheriteAxe                    = Item{ID: 611, DisplayName: "Netherite Axe", Name: "netherite_axe", StackSize: 1}
	NetheriteHoe                    = Item{ID: 612, DisplayName: "Netherite Hoe", Name: "netherite_hoe", StackSize: 1}
	Stick                           = Item{ID: 613, DisplayName: "Stick", Name: "stick", StackSize: 64}
	Bowl                            = Item{ID: 614, DisplayName: "Bowl", Name: "bowl", StackSize: 64}
	MushroomStew                    = Item{ID: 615, DisplayName: "Mushroom Stew", Name: "mushroom_stew", StackSize: 1}
	String                          = Item{ID: 616, DisplayName: "String", Name: "string", StackSize: 64}
	Feather                         = Item{ID: 617, DisplayName: "Feather", Name: "feather", StackSize: 64}
	Gunpowder                       = Item{ID: 618, DisplayName: "Gunpowder", Name: "gunpowder", StackSize: 64}
	WheatSeeds                      = Item{ID: 619, DisplayName: "Wheat Seeds", Name: "wheat_seeds", StackSize: 64}
	Wheat                           = Item{ID: 620, DisplayName: "Wheat", Name: "wheat", StackSize: 64}
	Bread                           = Item{ID: 621, DisplayName: "Bread", Name: "bread", StackSize: 64}
	LeatherHelmet                   = Item{ID: 622, DisplayName: "Leather Cap", Name: "leather_helmet", StackSize: 1}
	LeatherChestplate               = Item{ID: 623, DisplayName: "Leather Tunic", Name: "leather_chestplate", StackSize: 1}
	LeatherLeggings                 = Item{ID: 624, DisplayName: "Leather Pants", Name: "leather_leggings", StackSize: 1}
	LeatherBoots                    = Item{ID: 625, DisplayName: "Leather Boots", Name: "leather_boots", StackSize: 1}
	ChainmailHelmet                 = Item{ID: 626, DisplayName: "Chainmail Helmet", Name: "chainmail_helmet", StackSize: 1}
	ChainmailChestplate             = Item{ID: 627, DisplayName: "Chainmail Chestplate", Name: "chainmail_chestplate", StackSize: 1}
	ChainmailLeggings               = Item{ID: 628, DisplayName: "Chainmail Leggings", Name: "chainmail_leggings", StackSize: 1}
	ChainmailBoots                  = Item{ID: 629, DisplayName: "Chainmail Boots", Name: "chainmail_boots", StackSize: 1}
	IronHelmet                      = Item{ID: 630, DisplayName: "Iron Helmet", Name: "iron_helmet", StackSize: 1}
	IronChestplate                  = Item{ID: 631, DisplayName: "Iron Chestplate", Name: "iron_chestplate", StackSize: 1}
	IronLeggings                    = Item{ID: 632, DisplayName: "Iron Leggings", Name: "iron_leggings", StackSize: 1}
	IronBoots                       = Item{ID: 633, DisplayName: "Iron Boots", Name: "iron_boots", StackSize: 1}
	DiamondHelmet                   = Item{ID: 634, DisplayName: "Diamond Helmet", Name: "diamond_helmet", StackSize: 1}
	DiamondChestplate               = Item{ID: 635, DisplayName: "Diamond Chestplate", Name: "diamond_chestplate", StackSize: 1}
	DiamondLeggings                 = Item{ID: 636, DisplayName: "Diamond Leggings", Name: "diamond_leggings", StackSize: 1}
	DiamondBoots                    = Item{ID: 637, DisplayName: "Diamond Boots", Name: "diamond_boots", StackSize: 1}
	GoldenHelmet                    = Item{ID: 638, DisplayName: "Golden Helmet", Name: "golden_helmet", StackSize: 1}
	GoldenChestplate                = Item{ID: 639, DisplayName: "Golden Chestplate", Name: "golden_chestplate", StackSize: 1}
	GoldenLeggings                  = Item{ID: 640, DisplayName: "Golden Leggings", Name: "golden_leggings", StackSize: 1}
	GoldenBoots                     = Item{ID: 641, DisplayName: "Golden Boots", Name: "golden_boots", StackSize: 1}
	NetheriteHelmet                 = Item{ID: 642, DisplayName: "Netherite Helmet", Name: "netherite_helmet", StackSize: 1}
	NetheriteChestplate             = Item{ID: 643, DisplayName: "Netherite Chestplate", Name: "netherite_chestplate", StackSize: 1}
	NetheriteLeggings               = Item{ID: 644, DisplayName: "Netherite Leggings", Name: "netherite_leggings", StackSize: 1}
	NetheriteBoots                  = Item{ID: 645, DisplayName: "Netherite Boots", Name: "netherite_boots", StackSize: 1}
	Flint                           = Item{ID: 646, DisplayName: "Flint", Name: "flint", StackSize: 64}
	Porkchop                        = Item{ID: 647, DisplayName: "Raw Porkchop", Name: "porkchop", StackSize: 64}
	CookedPorkchop                  = Item{ID: 648, DisplayName: "Cooked Porkchop", Name: "cooked_porkchop", StackSize: 64}
	Painting                        = Item{ID: 649, DisplayName: "Painting", Name: "painting", StackSize: 64}
	GoldenApple                     = Item{ID: 650, DisplayName: "Golden Apple", Name: "golden_apple", StackSize: 64}
	EnchantedGoldenApple            = Item{ID: 651, DisplayName: "Enchanted Golden Apple", Name: "enchanted_golden_apple", StackSize: 64}
	OakSign                         = Item{ID: 652, DisplayName: "Oak Sign", Name: "oak_sign", StackSize: 16}
	SpruceSign                      = Item{ID: 653, DisplayName: "Spruce Sign", Name: "spruce_sign", StackSize: 16}
	BirchSign                       = Item{ID: 654, DisplayName: "Birch Sign", Name: "birch_sign", StackSize: 16}
	JungleSign                      = Item{ID: 655, DisplayName: "Jungle Sign", Name: "jungle_sign", StackSize: 16}
	AcaciaSign                      = Item{ID: 656, DisplayName: "Acacia Sign", Name: "acacia_sign", StackSize: 16}
	DarkOakSign                     = Item{ID: 657, DisplayName: "Dark Oak Sign", Name: "dark_oak_sign", StackSize: 16}
	CrimsonSign                     = Item{ID: 658, DisplayName: "Crimson Sign", Name: "crimson_sign", StackSize: 16}
	WarpedSign                      = Item{ID: 659, DisplayName: "Warped Sign", Name: "warped_sign", StackSize: 16}
	Bucket                          = Item{ID: 660, DisplayName: "Bucket", Name: "bucket", StackSize: 16}
	WaterBucket                     = Item{ID: 661, DisplayName: "Water Bucket", Name: "water_bucket", StackSize: 1}
	LavaBucket                      = Item{ID: 662, DisplayName: "Lava Bucket", Name: "lava_bucket", StackSize: 1}
	Minecart                        = Item{ID: 663, DisplayName: "Minecart", Name: "minecart", StackSize: 1}
	Saddle                          = Item{ID: 664, DisplayName: "Saddle", Name: "saddle", StackSize: 1}
	Redstone                        = Item{ID: 665, DisplayName: "Redstone Dust", Name: "redstone", StackSize: 64}
	Snowball                        = Item{ID: 666, DisplayName: "Snowball", Name: "snowball", StackSize: 16}
	OakBoat                         = Item{ID: 667, DisplayName: "Oak Boat", Name: "oak_boat", StackSize: 1}
	Leather                         = Item{ID: 668, DisplayName: "Leather", Name: "leather", StackSize: 64}
	MilkBucket                      = Item{ID: 669, DisplayName: "Milk Bucket", Name: "milk_bucket", StackSize: 1}
	PufferfishBucket                = Item{ID: 670, DisplayName: "Bucket of Pufferfish", Name: "pufferfish_bucket", StackSize: 1}
	SalmonBucket                    = Item{ID: 671, DisplayName: "Bucket of Salmon", Name: "salmon_bucket", StackSize: 1}
	CodBucket                       = Item{ID: 672, DisplayName: "Bucket of Cod", Name: "cod_bucket", StackSize: 1}
	TropicalFishBucket              = Item{ID: 673, DisplayName: "Bucket of Tropical Fish", Name: "tropical_fish_bucket", StackSize: 1}
	Brick                           = Item{ID: 674, DisplayName: "Brick", Name: "brick", StackSize: 64}
	ClayBall                        = Item{ID: 675, DisplayName: "Clay Ball", Name: "clay_ball", StackSize: 64}
	DriedKelpBlock                  = Item{ID: 676, DisplayName: "Dried Kelp Block", Name: "dried_kelp_block", StackSize: 64}
	Paper                           = Item{ID: 677, DisplayName: "Paper", Name: "paper", StackSize: 64}
	Book                            = Item{ID: 678, DisplayName: "Book", Name: "book", StackSize: 64}
	SlimeBall                       = Item{ID: 679, DisplayName: "Slimeball", Name: "slime_ball", StackSize: 64}
	ChestMinecart                   = Item{ID: 680, DisplayName: "Minecart with Chest", Name: "chest_minecart", StackSize: 1}
	FurnaceMinecart                 = Item{ID: 681, DisplayName: "Minecart with Furnace", Name: "furnace_minecart", StackSize: 1}
	Egg                             = Item{ID: 682, DisplayName: "Egg", Name: "egg", StackSize: 16}
	Compass                         = Item{ID: 683, DisplayName: "Compass", Name: "compass", StackSize: 64}
	FishingRod                      = Item{ID: 684, DisplayName: "Fishing Rod", Name: "fishing_rod", StackSize: 1}
	Clock                           = Item{ID: 685, DisplayName: "Clock", Name: "clock", StackSize: 64}
	GlowstoneDust                   = Item{ID: 686, DisplayName: "Glowstone Dust", Name: "glowstone_dust", StackSize: 64}
	Cod                             = Item{ID: 687, DisplayName: "Raw Cod", Name: "cod", StackSize: 64}
	Salmon                          = Item{ID: 688, DisplayName: "Raw Salmon", Name: "salmon", StackSize: 64}
	TropicalFish                    = Item{ID: 689, DisplayName: "Tropical Fish", Name: "tropical_fish", StackSize: 64}
	Pufferfish                      = Item{ID: 690, DisplayName: "Pufferfish", Name: "pufferfish", StackSize: 64}
	CookedCod                       = Item{ID: 691, DisplayName: "Cooked Cod", Name: "cooked_cod", StackSize: 64}
	CookedSalmon                    = Item{ID: 692, DisplayName: "Cooked Salmon", Name: "cooked_salmon", StackSize: 64}
	InkSac                          = Item{ID: 693, DisplayName: "Ink Sac", Name: "ink_sac", StackSize: 64}
	CocoaBeans                      = Item{ID: 694, DisplayName: "Cocoa Beans", Name: "cocoa_beans", StackSize: 64}
	LapisLazuli                     = Item{ID: 695, DisplayName: "Lapis Lazuli", Name: "lapis_lazuli", StackSize: 64}
	WhiteDye                        = Item{ID: 696, DisplayName: "White Dye", Name: "white_dye", StackSize: 64}
	OrangeDye                       = Item{ID: 697, DisplayName: "Orange Dye", Name: "orange_dye", StackSize: 64}
	MagentaDye                      = Item{ID: 698, DisplayName: "Magenta Dye", Name: "magenta_dye", StackSize: 64}
	LightBlueDye                    = Item{ID: 699, DisplayName: "Light Blue Dye", Name: "light_blue_dye", StackSize: 64}
	YellowDye                       = Item{ID: 700, DisplayName: "Yellow Dye", Name: "yellow_dye", StackSize: 64}
	LimeDye                         = Item{ID: 701, DisplayName: "Lime Dye", Name: "lime_dye", StackSize: 64}
	PinkDye                         = Item{ID: 702, DisplayName: "Pink Dye", Name: "pink_dye", StackSize: 64}
	GrayDye                         = Item{ID: 703, DisplayName: "Gray Dye", Name: "gray_dye", StackSize: 64}
	LightGrayDye                    = Item{ID: 704, DisplayName: "Light Gray Dye", Name: "light_gray_dye", StackSize: 64}
	CyanDye                         = Item{ID: 705, DisplayName: "Cyan Dye", Name: "cyan_dye", StackSize: 64}
	PurpleDye                       = Item{ID: 706, DisplayName: "Purple Dye", Name: "purple_dye", StackSize: 64}
	BlueDye                         = Item{ID: 707, DisplayName: "Blue Dye", Name: "blue_dye", StackSize: 64}
	BrownDye                        = Item{ID: 708, DisplayName: "Brown Dye", Name: "brown_dye", StackSize: 64}
	GreenDye                        = Item{ID: 709, DisplayName: "Green Dye", Name: "green_dye", StackSize: 64}
	RedDye                          = Item{ID: 710, DisplayName: "Red Dye", Name: "red_dye", StackSize: 64}
	BlackDye                        = Item{ID: 711, DisplayName: "Black Dye", Name: "black_dye", StackSize: 64}
	BoneMeal                        = Item{ID: 712, DisplayName: "Bone Meal", Name: "bone_meal", StackSize: 64}
	Bone                            = Item{ID: 713, DisplayName: "Bone", Name: "bone", StackSize: 64}
	Sugar                           = Item{ID: 714, DisplayName: "Sugar", Name: "sugar", StackSize: 64}
	Cake                            = Item{ID: 715, DisplayName: "Cake", Name: "cake", StackSize: 1}
	WhiteBed                        = Item{ID: 716, DisplayName: "White Bed", Name: "white_bed", StackSize: 1}
	OrangeBed                       = Item{ID: 717, DisplayName: "Orange Bed", Name: "orange_bed", StackSize: 1}
	MagentaBed                      = Item{ID: 718, DisplayName: "Magenta Bed", Name: "magenta_bed", StackSize: 1}
	LightBlueBed                    = Item{ID: 719, DisplayName: "Light Blue Bed", Name: "light_blue_bed", StackSize: 1}
	YellowBed                       = Item{ID: 720, DisplayName: "Yellow Bed", Name: "yellow_bed", StackSize: 1}
	LimeBed                         = Item{ID: 721, DisplayName: "Lime Bed", Name: "lime_bed", StackSize: 1}
	PinkBed                         = Item{ID: 722, DisplayName: "Pink Bed", Name: "pink_bed", StackSize: 1}
	GrayBed                         = Item{ID: 723, DisplayName: "Gray Bed", Name: "gray_bed", StackSize: 1}
	LightGrayBed                    = Item{ID: 724, DisplayName: "Light Gray Bed", Name: "light_gray_bed", StackSize: 1}
	CyanBed                         = Item{ID: 725, DisplayName: "Cyan Bed", Name: "cyan_bed", StackSize: 1}
	PurpleBed                       = Item{ID: 726, DisplayName: "Purple Bed", Name: "purple_bed", StackSize: 1}
	BlueBed                         = Item{ID: 727, DisplayName: "Blue Bed", Name: "blue_bed", StackSize: 1}
	BrownBed                        = Item{ID: 728, DisplayName: "Brown Bed", Name: "brown_bed", StackSize: 1}
	GreenBed                        = Item{ID: 729, DisplayName: "Green Bed", Name: "green_bed", StackSize: 1}
	RedBed                          = Item{ID: 730, DisplayName: "Red Bed", Name: "red_bed", StackSize: 1}
	BlackBed                        = Item{ID: 731, DisplayName: "Black Bed", Name: "black_bed", StackSize: 1}
	Cookie                          = Item{ID: 732, DisplayName: "Cookie", Name: "cookie", StackSize: 64}
	FilledMap                       = Item{ID: 733, DisplayName: "Map", Name: "filled_map", StackSize: 64}
	Shears                          = Item{ID: 734, DisplayName: "Shears", Name: "shears", StackSize: 1}
	MelonSlice                      = Item{ID: 735, DisplayName: "Melon Slice", Name: "melon_slice", StackSize: 64}
	DriedKelp                       = Item{ID: 736, DisplayName: "Dried Kelp", Name: "dried_kelp", StackSize: 64}
	PumpkinSeeds                    = Item{ID: 737, DisplayName: "Pumpkin Seeds", Name: "pumpkin_seeds", StackSize: 64}
	MelonSeeds                      = Item{ID: 738, DisplayName: "Melon Seeds", Name: "melon_seeds", StackSize: 64}
	Beef                            = Item{ID: 739, DisplayName: "Raw Beef", Name: "beef", StackSize: 64}
	CookedBeef                      = Item{ID: 740, DisplayName: "Steak", Name: "cooked_beef", StackSize: 64}
	Chicken                         = Item{ID: 741, DisplayName: "Raw Chicken", Name: "chicken", StackSize: 64}
	CookedChicken                   = Item{ID: 742, DisplayName: "Cooked Chicken", Name: "cooked_chicken", StackSize: 64}
	RottenFlesh                     = Item{ID: 743, DisplayName: "Rotten Flesh", Name: "rotten_flesh", StackSize: 64}
	EnderPearl                      = Item{ID: 744, DisplayName: "Ender Pearl", Name: "ender_pearl", StackSize: 16}
	BlazeRod                        = Item{ID: 745, DisplayName: "Blaze Rod", Name: "blaze_rod", StackSize: 64}
	GhastTear                       = Item{ID: 746, DisplayName: "Ghast Tear", Name: "ghast_tear", StackSize: 64}
	GoldNugget                      = Item{ID: 747, DisplayName: "Gold Nugget", Name: "gold_nugget", StackSize: 64}
	NetherWart                      = Item{ID: 748, DisplayName: "Nether Wart", Name: "nether_wart", StackSize: 64}
	Potion                          = Item{ID: 749, DisplayName: "Potion", Name: "potion", StackSize: 1}
	GlassBottle                     = Item{ID: 750, DisplayName: "Glass Bottle", Name: "glass_bottle", StackSize: 64}
	SpiderEye                       = Item{ID: 751, DisplayName: "Spider Eye", Name: "spider_eye", StackSize: 64}
	FermentedSpiderEye              = Item{ID: 752, DisplayName: "Fermented Spider Eye", Name: "fermented_spider_eye", StackSize: 64}
	BlazePowder                     = Item{ID: 753, DisplayName: "Blaze Powder", Name: "blaze_powder", StackSize: 64}
	MagmaCream                      = Item{ID: 754, DisplayName: "Magma Cream", Name: "magma_cream", StackSize: 64}
	BrewingStand                    = Item{ID: 755, DisplayName: "Brewing Stand", Name: "brewing_stand", StackSize: 64}
	Cauldron                        = Item{ID: 756, DisplayName: "Cauldron", Name: "cauldron", StackSize: 64}
	EnderEye                        = Item{ID: 757, DisplayName: "Eye of Ender", Name: "ender_eye", StackSize: 64}
	GlisteringMelonSlice            = Item{ID: 758, DisplayName: "Glistering Melon Slice", Name: "glistering_melon_slice", StackSize: 64}
	BatSpawnEgg                     = Item{ID: 759, DisplayName: "Bat Spawn Egg", Name: "bat_spawn_egg", StackSize: 64}
	BeeSpawnEgg                     = Item{ID: 760, DisplayName: "Bee Spawn Egg", Name: "bee_spawn_egg", StackSize: 64}
	BlazeSpawnEgg                   = Item{ID: 761, DisplayName: "Blaze Spawn Egg", Name: "blaze_spawn_egg", StackSize: 64}
	CatSpawnEgg                     = Item{ID: 762, DisplayName: "Cat Spawn Egg", Name: "cat_spawn_egg", StackSize: 64}
	CaveSpiderSpawnEgg              = Item{ID: 763, DisplayName: "Cave Spider Spawn Egg", Name: "cave_spider_spawn_egg", StackSize: 64}
	ChickenSpawnEgg                 = Item{ID: 764, DisplayName: "Chicken Spawn Egg", Name: "chicken_spawn_egg", StackSize: 64}
	CodSpawnEgg                     = Item{ID: 765, DisplayName: "Cod Spawn Egg", Name: "cod_spawn_egg", StackSize: 64}
	CowSpawnEgg                     = Item{ID: 766, DisplayName: "Cow Spawn Egg", Name: "cow_spawn_egg", StackSize: 64}
	CreeperSpawnEgg                 = Item{ID: 767, DisplayName: "Creeper Spawn Egg", Name: "creeper_spawn_egg", StackSize: 64}
	DolphinSpawnEgg                 = Item{ID: 768, DisplayName: "Dolphin Spawn Egg", Name: "dolphin_spawn_egg", StackSize: 64}
	DonkeySpawnEgg                  = Item{ID: 769, DisplayName: "Donkey Spawn Egg", Name: "donkey_spawn_egg", StackSize: 64}
	DrownedSpawnEgg                 = Item{ID: 770, DisplayName: "Drowned Spawn Egg", Name: "drowned_spawn_egg", StackSize: 64}
	ElderGuardianSpawnEgg           = Item{ID: 771, DisplayName: "Elder Guardian Spawn Egg", Name: "elder_guardian_spawn_egg", StackSize: 64}
	EndermanSpawnEgg                = Item{ID: 772, DisplayName: "Enderman Spawn Egg", Name: "enderman_spawn_egg", StackSize: 64}
	EndermiteSpawnEgg               = Item{ID: 773, DisplayName: "Endermite Spawn Egg", Name: "endermite_spawn_egg", StackSize: 64}
	EvokerSpawnEgg                  = Item{ID: 774, DisplayName: "Evoker Spawn Egg", Name: "evoker_spawn_egg", StackSize: 64}
	FoxSpawnEgg                     = Item{ID: 775, DisplayName: "Fox Spawn Egg", Name: "fox_spawn_egg", StackSize: 64}
	GhastSpawnEgg                   = Item{ID: 776, DisplayName: "Ghast Spawn Egg", Name: "ghast_spawn_egg", StackSize: 64}
	GuardianSpawnEgg                = Item{ID: 777, DisplayName: "Guardian Spawn Egg", Name: "guardian_spawn_egg", StackSize: 64}
	HoglinSpawnEgg                  = Item{ID: 778, DisplayName: "Hoglin Spawn Egg", Name: "hoglin_spawn_egg", StackSize: 64}
	HorseSpawnEgg                   = Item{ID: 779, DisplayName: "Horse Spawn Egg", Name: "horse_spawn_egg", StackSize: 64}
	HuskSpawnEgg                    = Item{ID: 780, DisplayName: "Husk Spawn Egg", Name: "husk_spawn_egg", StackSize: 64}
	LlamaSpawnEgg                   = Item{ID: 781, DisplayName: "Llama Spawn Egg", Name: "llama_spawn_egg", StackSize: 64}
	MagmaCubeSpawnEgg               = Item{ID: 782, DisplayName: "Magma Cube Spawn Egg", Name: "magma_cube_spawn_egg", StackSize: 64}
	MooshroomSpawnEgg               = Item{ID: 783, DisplayName: "Mooshroom Spawn Egg", Name: "mooshroom_spawn_egg", StackSize: 64}
	MuleSpawnEgg                    = Item{ID: 784, DisplayName: "Mule Spawn Egg", Name: "mule_spawn_egg", StackSize: 64}
	OcelotSpawnEgg                  = Item{ID: 785, DisplayName: "Ocelot Spawn Egg", Name: "ocelot_spawn_egg", StackSize: 64}
	PandaSpawnEgg                   = Item{ID: 786, DisplayName: "Panda Spawn Egg", Name: "panda_spawn_egg", StackSize: 64}
	ParrotSpawnEgg                  = Item{ID: 787, DisplayName: "Parrot Spawn Egg", Name: "parrot_spawn_egg", StackSize: 64}
	PhantomSpawnEgg                 = Item{ID: 788, DisplayName: "Phantom Spawn Egg", Name: "phantom_spawn_egg", StackSize: 64}
	PigSpawnEgg                     = Item{ID: 789, DisplayName: "Pig Spawn Egg", Name: "pig_spawn_egg", StackSize: 64}
	PiglinSpawnEgg                  = Item{ID: 790, DisplayName: "Piglin Spawn Egg", Name: "piglin_spawn_egg", StackSize: 64}
	PiglinBruteSpawnEgg             = Item{ID: 791, DisplayName: "Piglin Brute Spawn Egg", Name: "piglin_brute_spawn_egg", StackSize: 64}
	PillagerSpawnEgg                = Item{ID: 792, DisplayName: "Pillager Spawn Egg", Name: "pillager_spawn_egg", StackSize: 64}
	PolarBearSpawnEgg               = Item{ID: 793, DisplayName: "Polar Bear Spawn Egg", Name: "polar_bear_spawn_egg", StackSize: 64}
	PufferfishSpawnEgg              = Item{ID: 794, DisplayName: "Pufferfish Spawn Egg", Name: "pufferfish_spawn_egg", StackSize: 64}
	RabbitSpawnEgg                  = Item{ID: 795, DisplayName: "Rabbit Spawn Egg", Name: "rabbit_spawn_egg", StackSize: 64}
	RavagerSpawnEgg                 = Item{ID: 796, DisplayName: "Ravager Spawn Egg", Name: "ravager_spawn_egg", StackSize: 64}
	SalmonSpawnEgg                  = Item{ID: 797, DisplayName: "Salmon Spawn Egg", Name: "salmon_spawn_egg", StackSize: 64}
	SheepSpawnEgg                   = Item{ID: 798, DisplayName: "Sheep Spawn Egg", Name: "sheep_spawn_egg", StackSize: 64}
	ShulkerSpawnEgg                 = Item{ID: 799, DisplayName: "Shulker Spawn Egg", Name: "shulker_spawn_egg", StackSize: 64}
	SilverfishSpawnEgg              = Item{ID: 800, DisplayName: "Silverfish Spawn Egg", Name: "silverfish_spawn_egg", StackSize: 64}
	SkeletonSpawnEgg                = Item{ID: 801, DisplayName: "Skeleton Spawn Egg", Name: "skeleton_spawn_egg", StackSize: 64}
	SkeletonHorseSpawnEgg           = Item{ID: 802, DisplayName: "Skeleton Horse Spawn Egg", Name: "skeleton_horse_spawn_egg", StackSize: 64}
	SlimeSpawnEgg                   = Item{ID: 803, DisplayName: "Slime Spawn Egg", Name: "slime_spawn_egg", StackSize: 64}
	SpiderSpawnEgg                  = Item{ID: 804, DisplayName: "Spider Spawn Egg", Name: "spider_spawn_egg", StackSize: 64}
	SquidSpawnEgg                   = Item{ID: 805, DisplayName: "Squid Spawn Egg", Name: "squid_spawn_egg", StackSize: 64}
	StraySpawnEgg                   = Item{ID: 806, DisplayName: "Stray Spawn Egg", Name: "stray_spawn_egg", StackSize: 64}
	StriderSpawnEgg                 = Item{ID: 807, DisplayName: "Strider Spawn Egg", Name: "strider_spawn_egg", StackSize: 64}
	TraderLlamaSpawnEgg             = Item{ID: 808, DisplayName: "Trader Llama Spawn Egg", Name: "trader_llama_spawn_egg", StackSize: 64}
	TropicalFishSpawnEgg            = Item{ID: 809, DisplayName: "Tropical Fish Spawn Egg", Name: "tropical_fish_spawn_egg", StackSize: 64}
	TurtleSpawnEgg                  = Item{ID: 810, DisplayName: "Turtle Spawn Egg", Name: "turtle_spawn_egg", StackSize: 64}
	VexSpawnEgg                     = Item{ID: 811, DisplayName: "Vex Spawn Egg", Name: "vex_spawn_egg", StackSize: 64}
	VillagerSpawnEgg                = Item{ID: 812, DisplayName: "Villager Spawn Egg", Name: "villager_spawn_egg", StackSize: 64}
	VindicatorSpawnEgg              = Item{ID: 813, DisplayName: "Vindicator Spawn Egg", Name: "vindicator_spawn_egg", StackSize: 64}
	WanderingTraderSpawnEgg         = Item{ID: 814, DisplayName: "Wandering Trader Spawn Egg", Name: "wandering_trader_spawn_egg", StackSize: 64}
	WitchSpawnEgg                   = Item{ID: 815, DisplayName: "Witch Spawn Egg", Name: "witch_spawn_egg", StackSize: 64}
	WitherSkeletonSpawnEgg          = Item{ID: 816, DisplayName: "Wither Skeleton Spawn Egg", Name: "wither_skeleton_spawn_egg", StackSize: 64}
	WolfSpawnEgg                    = Item{ID: 817, DisplayName: "Wolf Spawn Egg", Name: "wolf_spawn_egg", StackSize: 64}
	ZoglinSpawnEgg                  = Item{ID: 818, DisplayName: "Zoglin Spawn Egg", Name: "zoglin_spawn_egg", StackSize: 64}
	ZombieSpawnEgg                  = Item{ID: 819, DisplayName: "Zombie Spawn Egg", Name: "zombie_spawn_egg", StackSize: 64}
	ZombieHorseSpawnEgg             = Item{ID: 820, DisplayName: "Zombie Horse Spawn Egg", Name: "zombie_horse_spawn_egg", StackSize: 64}
	ZombieVillagerSpawnEgg          = Item{ID: 821, DisplayName: "Zombie Villager Spawn Egg", Name: "zombie_villager_spawn_egg", StackSize: 64}
	ZombifiedPiglinSpawnEgg         = Item{ID: 822, DisplayName: "Zombified Piglin Spawn Egg", Name: "zombified_piglin_spawn_egg", StackSize: 64}
	ExperienceBottle                = Item{ID: 823, DisplayName: "Bottle o' Enchanting", Name: "experience_bottle", StackSize: 64}
	FireCharge                      = Item{ID: 824, DisplayName: "Fire Charge", Name: "fire_charge", StackSize: 64}
	WritableBook                    = Item{ID: 825, DisplayName: "Book and Quill", Name: "writable_book", StackSize: 1}
	WrittenBook                     = Item{ID: 826, DisplayName: "Written Book", Name: "written_book", StackSize: 16}
	Emerald                         = Item{ID: 827, DisplayName: "Emerald", Name: "emerald", StackSize: 64}
	ItemFrame                       = Item{ID: 828, DisplayName: "Item Frame", Name: "item_frame", StackSize: 64}
	FlowerPot                       = Item{ID: 829, DisplayName: "Flower Pot", Name: "flower_pot", StackSize: 64}
	Carrot                          = Item{ID: 830, DisplayName: "Carrot", Name: "carrot", StackSize: 64}
	Potato                          = Item{ID: 831, DisplayName: "Potato", Name: "potato", StackSize: 64}
	BakedPotato                     = Item{ID: 832, DisplayName: "Baked Potato", Name: "baked_potato", StackSize: 64}
	PoisonousPotato                 = Item{ID: 833, DisplayName: "Poisonous Potato", Name: "poisonous_potato", StackSize: 64}
	Map                             = Item{ID: 834, DisplayName: "Empty Map", Name: "map", StackSize: 64}
	GoldenCarrot                    = Item{ID: 835, DisplayName: "Golden Carrot", Name: "golden_carrot", StackSize: 64}
	SkeletonSkull                   = Item{ID: 836, DisplayName: "Skeleton Skull", Name: "skeleton_skull", StackSize: 64}
	WitherSkeletonSkull             = Item{ID: 837, DisplayName: "Wither Skeleton Skull", Name: "wither_skeleton_skull", StackSize: 64}
	PlayerHead                      = Item{ID: 838, DisplayName: "Player Head", Name: "player_head", StackSize: 64}
	ZombieHead                      = Item{ID: 839, DisplayName: "Zombie Head", Name: "zombie_head", StackSize: 64}
	CreeperHead                     = Item{ID: 840, DisplayName: "Creeper Head", Name: "creeper_head", StackSize: 64}
	DragonHead                      = Item{ID: 841, DisplayName: "Dragon Head", Name: "dragon_head", StackSize: 64}
	CarrotOnAStick                  = Item{ID: 842, DisplayName: "Carrot on a Stick", Name: "carrot_on_a_stick", StackSize: 1}
	WarpedFungusOnAStick            = Item{ID: 843, DisplayName: "Warped Fungus on a Stick", Name: "warped_fungus_on_a_stick", StackSize: 64}
	NetherStar                      = Item{ID: 844, DisplayName: "Nether Star", Name: "nether_star", StackSize: 64}
	PumpkinPie                      = Item{ID: 845, DisplayName: "Pumpkin Pie", Name: "pumpkin_pie", StackSize: 64}
	FireworkRocket                  = Item{ID: 846, DisplayName: "Firework Rocket", Name: "firework_rocket", StackSize: 64}
	FireworkStar                    = Item{ID: 847, DisplayName: "Firework Star", Name: "firework_star", StackSize: 64}
	EnchantedBook                   = Item{ID: 848, DisplayName: "Enchanted Book", Name: "enchanted_book", StackSize: 1}
	NetherBrick                     = Item{ID: 849, DisplayName: "Nether Brick", Name: "nether_brick", StackSize: 64}
	Quartz                          = Item{ID: 850, DisplayName: "Nether Quartz", Name: "quartz", StackSize: 64}
	TntMinecart                     = Item{ID: 851, DisplayName: "Minecart with TNT", Name: "tnt_minecart", StackSize: 1}
	HopperMinecart                  = Item{ID: 852, DisplayName: "Minecart with Hopper", Name: "hopper_minecart", StackSize: 1}
	PrismarineShard                 = Item{ID: 853, DisplayName: "Prismarine Shard", Name: "prismarine_shard", StackSize: 64}
	PrismarineCrystals              = Item{ID: 854, DisplayName: "Prismarine Crystals", Name: "prismarine_crystals", StackSize: 64}
	Rabbit                          = Item{ID: 855, DisplayName: "Raw Rabbit", Name: "rabbit", StackSize: 64}
	CookedRabbit                    = Item{ID: 856, DisplayName: "Cooked Rabbit", Name: "cooked_rabbit", StackSize: 64}
	RabbitStew                      = Item{ID: 857, DisplayName: "Rabbit Stew", Name: "rabbit_stew", StackSize: 1}
	RabbitFoot                      = Item{ID: 858, DisplayName: "Rabbit's Foot", Name: "rabbit_foot", StackSize: 64}
	RabbitHide                      = Item{ID: 859, DisplayName: "Rabbit Hide", Name: "rabbit_hide", StackSize: 64}
	ArmorStand                      = Item{ID: 860, DisplayName: "Armor Stand", Name: "armor_stand", StackSize: 16}
	IronHorseArmor                  = Item{ID: 861, DisplayName: "Iron Horse Armor", Name: "iron_horse_armor", StackSize: 1}
	GoldenHorseArmor                = Item{ID: 862, DisplayName: "Golden Horse Armor", Name: "golden_horse_armor", StackSize: 1}
	DiamondHorseArmor               = Item{ID: 863, DisplayName: "Diamond Horse Armor", Name: "diamond_horse_armor", StackSize: 1}
	LeatherHorseArmor               = Item{ID: 864, DisplayName: "Leather Horse Armor", Name: "leather_horse_armor", StackSize: 1}
	Lead                            = Item{ID: 865, DisplayName: "Lead", Name: "lead", StackSize: 64}
	NameTag                         = Item{ID: 866, DisplayName: "Name Tag", Name: "name_tag", StackSize: 64}
	CommandBlockMinecart            = Item{ID: 867, DisplayName: "Minecart with Command Block", Name: "command_block_minecart", StackSize: 1}
	Mutton                          = Item{ID: 868, DisplayName: "Raw Mutton", Name: "mutton", StackSize: 64}
	CookedMutton                    = Item{ID: 869, DisplayName: "Cooked Mutton", Name: "cooked_mutton", StackSize: 64}
	WhiteBanner                     = Item{ID: 870, DisplayName: "White Banner", Name: "white_banner", StackSize: 16}
	OrangeBanner                    = Item{ID: 871, DisplayName: "Orange Banner", Name: "orange_banner", StackSize: 16}
	MagentaBanner                   = Item{ID: 872, DisplayName: "Magenta Banner", Name: "magenta_banner", StackSize: 16}
	LightBlueBanner                 = Item{ID: 873, DisplayName: "Light Blue Banner", Name: "light_blue_banner", StackSize: 16}
	YellowBanner                    = Item{ID: 874, DisplayName: "Yellow Banner", Name: "yellow_banner", StackSize: 16}
	LimeBanner                      = Item{ID: 875, DisplayName: "Lime Banner", Name: "lime_banner", StackSize: 16}
	PinkBanner                      = Item{ID: 876, DisplayName: "Pink Banner", Name: "pink_banner", StackSize: 16}
	GrayBanner                      = Item{ID: 877, DisplayName: "Gray Banner", Name: "gray_banner", StackSize: 16}
	LightGrayBanner                 = Item{ID: 878, DisplayName: "Light Gray Banner", Name: "light_gray_banner", StackSize: 16}
	CyanBanner                      = Item{ID: 879, DisplayName: "Cyan Banner", Name: "cyan_banner", StackSize: 16}
	PurpleBanner                    = Item{ID: 880, DisplayName: "Purple Banner", Name: "purple_banner", StackSize: 16}
	BlueBanner                      = Item{ID: 881, DisplayName: "Blue Banner", Name: "blue_banner", StackSize: 16}
	BrownBanner                     = Item{ID: 882, DisplayName: "Brown Banner", Name: "brown_banner", StackSize: 16}
	GreenBanner                     = Item{ID: 883, DisplayName: "Green Banner", Name: "green_banner", StackSize: 16}
	RedBanner                       = Item{ID: 884, DisplayName: "Red Banner", Name: "red_banner", StackSize: 16}
	BlackBanner                     = Item{ID: 885, DisplayName: "Black Banner", Name: "black_banner", StackSize: 16}
	EndCrystal                      = Item{ID: 886, DisplayName: "End Crystal", Name: "end_crystal", StackSize: 64}
	ChorusFruit                     = Item{ID: 887, DisplayName: "Chorus Fruit", Name: "chorus_fruit", StackSize: 64}
	PoppedChorusFruit               = Item{ID: 888, DisplayName: "Popped Chorus Fruit", Name: "popped_chorus_fruit", StackSize: 64}
	Beetroot                        = Item{ID: 889, DisplayName: "Beetroot", Name: "beetroot", StackSize: 64}
	BeetrootSeeds                   = Item{ID: 890, DisplayName: "Beetroot Seeds", Name: "beetroot_seeds", StackSize: 64}
	BeetrootSoup                    = Item{ID: 891, DisplayName: "Beetroot Soup", Name: "beetroot_soup", StackSize: 1}
	DragonBreath                    = Item{ID: 892, DisplayName: "Dragon's Breath", Name: "dragon_breath", StackSize: 64}
	SplashPotion                    = Item{ID: 893, DisplayName: "Splash Potion", Name: "splash_potion", StackSize: 1}
	SpectralArrow                   = Item{ID: 894, DisplayName: "Spectral Arrow", Name: "spectral_arrow", StackSize: 64}
	TippedArrow                     = Item{ID: 895, DisplayName: "Tipped Arrow", Name: "tipped_arrow", StackSize: 64}
	LingeringPotion                 = Item{ID: 896, DisplayName: "Lingering Potion", Name: "lingering_potion", StackSize: 1}
	Shield                          = Item{ID: 897, DisplayName: "Shield", Name: "shield", StackSize: 1}
	Elytra                          = Item{ID: 898, DisplayName: "Elytra", Name: "elytra", StackSize: 1}
	SpruceBoat                      = Item{ID: 899, DisplayName: "Spruce Boat", Name: "spruce_boat", StackSize: 1}
	BirchBoat                       = Item{ID: 900, DisplayName: "Birch Boat", Name: "birch_boat", StackSize: 1}
	JungleBoat                      = Item{ID: 901, DisplayName: "Jungle Boat", Name: "jungle_boat", StackSize: 1}
	AcaciaBoat                      = Item{ID: 902, DisplayName: "Acacia Boat", Name: "acacia_boat", StackSize: 1}
	DarkOakBoat                     = Item{ID: 903, DisplayName: "Dark Oak Boat", Name: "dark_oak_boat", StackSize: 1}
	TotemOfUndying                  = Item{ID: 904, DisplayName: "Totem of Undying", Name: "totem_of_undying", StackSize: 1}
	ShulkerShell                    = Item{ID: 905, DisplayName: "Shulker Shell", Name: "shulker_shell", StackSize: 64}
	IronNugget                      = Item{ID: 906, DisplayName: "Iron Nugget", Name: "iron_nugget", StackSize: 64}
	KnowledgeBook                   = Item{ID: 907, DisplayName: "Knowledge Book", Name: "knowledge_book", StackSize: 1}
	DebugStick                      = Item{ID: 908, DisplayName: "Debug Stick", Name: "debug_stick", StackSize: 1}
	MusicDisc13                     = Item{ID: 909, DisplayName: "13 Disc", Name: "music_disc_13", StackSize: 1}
	MusicDiscCat                    = Item{ID: 910, DisplayName: "Cat Disc", Name: "music_disc_cat", StackSize: 1}
	MusicDiscBlocks                 = Item{ID: 911, DisplayName: "Blocks Disc", Name: "music_disc_blocks", StackSize: 1}
	MusicDiscChirp                  = Item{ID: 912, DisplayName: "Chirp Disc", Name: "music_disc_chirp", StackSize: 1}
	MusicDiscFar                    = Item{ID: 913, DisplayName: "Far Disc", Name: "music_disc_far", StackSize: 1}
	MusicDiscMall                   = Item{ID: 914, DisplayName: "Mall Disc", Name: "music_disc_mall", StackSize: 1}
	MusicDiscMellohi                = Item{ID: 915, DisplayName: "Mellohi Disc", Name: "music_disc_mellohi", StackSize: 1}
	MusicDiscStal                   = Item{ID: 916, DisplayName: "Stal Disc", Name: "music_disc_stal", StackSize: 1}
	MusicDiscStrad                  = Item{ID: 917, DisplayName: "Strad Disc", Name: "music_disc_strad", StackSize: 1}
	MusicDiscWard                   = Item{ID: 918, DisplayName: "Ward Disc", Name: "music_disc_ward", StackSize: 1}
	MusicDisc11                     = Item{ID: 919, DisplayName: "11 Disc", Name: "music_disc_11", StackSize: 1}
	MusicDiscWait                   = Item{ID: 920, DisplayName: "Wait Disc", Name: "music_disc_wait", StackSize: 1}
	MusicDiscPigstep                = Item{ID: 921, DisplayName: "Music Disc", Name: "music_disc_pigstep", StackSize: 1}
	Trident                         = Item{ID: 922, DisplayName: "Trident", Name: "trident", StackSize: 1}
	PhantomMembrane                 = Item{ID: 923, DisplayName: "Phantom Membrane", Name: "phantom_membrane", StackSize: 64}
	NautilusShell                   = Item{ID: 924, DisplayName: "Nautilus Shell", Name: "nautilus_shell", StackSize: 64}
	HeartOfTheSea                   = Item{ID: 925, DisplayName: "Heart of the Sea", Name: "heart_of_the_sea", StackSize: 64}
	Crossbow                        = Item{ID: 926, DisplayName: "Crossbow", Name: "crossbow", StackSize: 1}
	SuspiciousStew                  = Item{ID: 927, DisplayName: "Suspicious Stew", Name: "suspicious_stew", StackSize: 1}
	Loom                            = Item{ID: 928, DisplayName: "Loom", Name: "loom", StackSize: 64}
	FlowerBannerPattern             = Item{ID: 929, DisplayName: "Banner Pattern", Name: "flower_banner_pattern", StackSize: 1}
	CreeperBannerPattern            = Item{ID: 930, DisplayName: "Banner Pattern", Name: "creeper_banner_pattern", StackSize: 1}
	SkullBannerPattern              = Item{ID: 931, DisplayName: "Banner Pattern", Name: "skull_banner_pattern", StackSize: 1}
	MojangBannerPattern             = Item{ID: 932, DisplayName: "Banner Pattern", Name: "mojang_banner_pattern", StackSize: 1}
	GlobeBannerPattern              = Item{ID: 933, DisplayName: "Banner Pattern", Name: "globe_banner_pattern", StackSize: 1}
	PiglinBannerPattern             = Item{ID: 934, DisplayName: "Banner Pattern", Name: "piglin_banner_pattern", StackSize: 1}
	Composter                       = Item{ID: 935, DisplayName: "Composter", Name: "composter", StackSize: 64}
	Barrel                          = Item{ID: 936, DisplayName: "Barrel", Name: "barrel", StackSize: 64}
	Smoker                          = Item{ID: 937, DisplayName: "Smoker", Name: "smoker", StackSize: 64}
	BlastFurnace                    = Item{ID: 938, DisplayName: "Blast Furnace", Name: "blast_furnace", StackSize: 64}
	CartographyTable                = Item{ID: 939, DisplayName: "Cartography Table", Name: "cartography_table", StackSize: 64}
	FletchingTable                  = Item{ID: 940, DisplayName: "Fletching Table", Name: "fletching_table", StackSize: 64}
	Grindstone                      = Item{ID: 941, DisplayName: "Grindstone", Name: "grindstone", StackSize: 64}
	Lectern                         = Item{ID: 942, DisplayName: "Lectern", Name: "lectern", StackSize: 64}
	SmithingTable                   = Item{ID: 943, DisplayName: "Smithing Table", Name: "smithing_table", StackSize: 64}
	Stonecutter                     = Item{ID: 944, DisplayName: "Stonecutter", Name: "stonecutter", StackSize: 64}
	Bell                            = Item{ID: 945, DisplayName: "Bell", Name: "bell", StackSize: 64}
	Lantern                         = Item{ID: 946, DisplayName: "Lantern", Name: "lantern", StackSize: 64}
	SoulLantern                     = Item{ID: 947, DisplayName: "Soul Lantern", Name: "soul_lantern", StackSize: 64}
	SweetBerries                    = Item{ID: 948, DisplayName: "Sweet Berries", Name: "sweet_berries", StackSize: 64}
	Campfire                        = Item{ID: 949, DisplayName: "Campfire", Name: "campfire", StackSize: 64}
	SoulCampfire                    = Item{ID: 950, DisplayName: "Soul Campfire", Name: "soul_campfire", StackSize: 64}
	Shroomlight                     = Item{ID: 951, DisplayName: "Shroomlight", Name: "shroomlight", StackSize: 64}
	Honeycomb                       = Item{ID: 952, DisplayName: "Honeycomb", Name: "honeycomb", StackSize: 64}
	BeeNest                         = Item{ID: 953, DisplayName: "Bee Nest", Name: "bee_nest", StackSize: 64}
	Beehive                         = Item{ID: 954, DisplayName: "Beehive", Name: "beehive", StackSize: 64}
	HoneyBottle                     = Item{ID: 955, DisplayName: "Honey Bottle", Name: "honey_bottle", StackSize: 16}
	HoneyBlock                      = Item{ID: 956, DisplayName: "Honey Block", Name: "honey_block", StackSize: 64}
	HoneycombBlock                  = Item{ID: 957, DisplayName: "Honeycomb Block", Name: "honeycomb_block", StackSize: 64}
	Lodestone                       = Item{ID: 958, DisplayName: "Lodestone", Name: "lodestone", StackSize: 64}
	NetheriteBlock                  = Item{ID: 959, DisplayName: "Block of Netherite", Name: "netherite_block", StackSize: 64}
	AncientDebris                   = Item{ID: 960, DisplayName: "Ancient Debris", Name: "ancient_debris", StackSize: 64}
	Target                          = Item{ID: 961, DisplayName: "Target", Name: "target", StackSize: 64}
	CryingObsidian                  = Item{ID: 962, DisplayName: "Crying Obsidian", Name: "crying_obsidian", StackSize: 64}
	Blackstone                      = Item{ID: 963, DisplayName: "Blackstone", Name: "blackstone", StackSize: 64}
	BlackstoneSlab                  = Item{ID: 964, DisplayName: "Blackstone Slab", Name: "blackstone_slab", StackSize: 64}
	BlackstoneStairs                = Item{ID: 965, DisplayName: "Blackstone Stairs", Name: "blackstone_stairs", StackSize: 64}
	GildedBlackstone                = Item{ID: 966, DisplayName: "Gilded Blackstone", Name: "gilded_blackstone", StackSize: 64}
	PolishedBlackstone              = Item{ID: 967, DisplayName: "Polished Blackstone", Name: "polished_blackstone", StackSize: 64}
	PolishedBlackstoneSlab          = Item{ID: 968, DisplayName: "Polished Blackstone Slab", Name: "polished_blackstone_slab", StackSize: 64}
	PolishedBlackstoneStairs        = Item{ID: 969, DisplayName: "Polished Blackstone Stairs", Name: "polished_blackstone_stairs", StackSize: 64}
	ChiseledPolishedBlackstone      = Item{ID: 970, DisplayName: "Chiseled Polished Blackstone", Name: "chiseled_polished_blackstone", StackSize: 64}
	PolishedBlackstoneBricks        = Item{ID: 971, DisplayName: "Polished Blackstone Bricks", Name: "polished_blackstone_bricks", StackSize: 64}
	PolishedBlackstoneBrickSlab     = Item{ID: 972, DisplayName: "Polished Blackstone Brick Slab", Name: "polished_blackstone_brick_slab", StackSize: 64}
	PolishedBlackstoneBrickStairs   = Item{ID: 973, DisplayName: "Polished Blackstone Brick Stairs", Name: "polished_blackstone_brick_stairs", StackSize: 64}
	CrackedPolishedBlackstoneBricks = Item{ID: 974, DisplayName: "Cracked Polished Blackstone Bricks", Name: "cracked_polished_blackstone_bricks", StackSize: 64}
	RespawnAnchor                   = Item{ID: 975, DisplayName: "Respawn Anchor", Name: "respawn_anchor", StackSize: 64}
)

// ByID is an index of minecraft items by their ID.
var ByID = map[ID]*Item{
	0:   &Air,
	1:   &Stone,
	2:   &Granite,
	3:   &PolishedGranite,
	4:   &Diorite,
	5:   &PolishedDiorite,
	6:   &Andesite,
	7:   &PolishedAndesite,
	8:   &GrassBlock,
	9:   &Dirt,
	10:  &CoarseDirt,
	11:  &Podzol,
	12:  &CrimsonNylium,
	13:  &WarpedNylium,
	14:  &Cobblestone,
	15:  &OakPlanks,
	16:  &SprucePlanks,
	17:  &BirchPlanks,
	18:  &JunglePlanks,
	19:  &AcaciaPlanks,
	20:  &DarkOakPlanks,
	21:  &CrimsonPlanks,
	22:  &WarpedPlanks,
	23:  &OakSapling,
	24:  &SpruceSapling,
	25:  &BirchSapling,
	26:  &JungleSapling,
	27:  &AcaciaSapling,
	28:  &DarkOakSapling,
	29:  &Bedrock,
	30:  &Sand,
	31:  &RedSand,
	32:  &Gravel,
	33:  &GoldOre,
	34:  &IronOre,
	35:  &CoalOre,
	36:  &NetherGoldOre,
	37:  &OakLog,
	38:  &SpruceLog,
	39:  &BirchLog,
	40:  &JungleLog,
	41:  &AcaciaLog,
	42:  &DarkOakLog,
	43:  &CrimsonStem,
	44:  &WarpedStem,
	45:  &StrippedOakLog,
	46:  &StrippedSpruceLog,
	47:  &StrippedBirchLog,
	48:  &StrippedJungleLog,
	49:  &StrippedAcaciaLog,
	50:  &StrippedDarkOakLog,
	51:  &StrippedCrimsonStem,
	52:  &StrippedWarpedStem,
	53:  &StrippedOakWood,
	54:  &StrippedSpruceWood,
	55:  &StrippedBirchWood,
	56:  &StrippedJungleWood,
	57:  &StrippedAcaciaWood,
	58:  &StrippedDarkOakWood,
	59:  &StrippedCrimsonHyphae,
	60:  &StrippedWarpedHyphae,
	61:  &OakWood,
	62:  &SpruceWood,
	63:  &BirchWood,
	64:  &JungleWood,
	65:  &AcaciaWood,
	66:  &DarkOakWood,
	67:  &CrimsonHyphae,
	68:  &WarpedHyphae,
	69:  &OakLeaves,
	70:  &SpruceLeaves,
	71:  &BirchLeaves,
	72:  &JungleLeaves,
	73:  &AcaciaLeaves,
	74:  &DarkOakLeaves,
	75:  &Sponge,
	76:  &WetSponge,
	77:  &Glass,
	78:  &LapisOre,
	79:  &LapisBlock,
	80:  &Dispenser,
	81:  &Sandstone,
	82:  &ChiseledSandstone,
	83:  &CutSandstone,
	84:  &NoteBlock,
	85:  &PoweredRail,
	86:  &DetectorRail,
	87:  &StickyPiston,
	88:  &Cobweb,
	89:  &Grass,
	90:  &Fern,
	91:  &DeadBush,
	92:  &Seagrass,
	93:  &SeaPickle,
	94:  &Piston,
	95:  &WhiteWool,
	96:  &OrangeWool,
	97:  &MagentaWool,
	98:  &LightBlueWool,
	99:  &YellowWool,
	100: &LimeWool,
	101: &PinkWool,
	102: &GrayWool,
	103: &LightGrayWool,
	104: &CyanWool,
	105: &PurpleWool,
	106: &BlueWool,
	107: &BrownWool,
	108: &GreenWool,
	109: &RedWool,
	110: &BlackWool,
	111: &Dandelion,
	112: &Poppy,
	113: &BlueOrchid,
	114: &Allium,
	115: &AzureBluet,
	116: &RedTulip,
	117: &OrangeTulip,
	118: &WhiteTulip,
	119: &PinkTulip,
	120: &OxeyeDaisy,
	121: &Cornflower,
	122: &LilyOfTheValley,
	123: &WitherRose,
	124: &BrownMushroom,
	125: &RedMushroom,
	126: &CrimsonFungus,
	127: &WarpedFungus,
	128: &CrimsonRoots,
	129: &WarpedRoots,
	130: &NetherSprouts,
	131: &WeepingVines,
	132: &TwistingVines,
	133: &SugarCane,
	134: &Kelp,
	135: &Bamboo,
	136: &GoldBlock,
	137: &IronBlock,
	138: &OakSlab,
	139: &SpruceSlab,
	140: &BirchSlab,
	141: &JungleSlab,
	142: &AcaciaSlab,
	143: &DarkOakSlab,
	144: &CrimsonSlab,
	145: &WarpedSlab,
	146: &StoneSlab,
	147: &SmoothStoneSlab,
	148: &SandstoneSlab,
	149: &CutSandstoneSlab,
	150: &PetrifiedOakSlab,
	151: &CobblestoneSlab,
	152: &BrickSlab,
	153: &StoneBrickSlab,
	154: &NetherBrickSlab,
	155: &QuartzSlab,
	156: &RedSandstoneSlab,
	157: &CutRedSandstoneSlab,
	158: &PurpurSlab,
	159: &PrismarineSlab,
	160: &PrismarineBrickSlab,
	161: &DarkPrismarineSlab,
	162: &SmoothQuartz,
	163: &SmoothRedSandstone,
	164: &SmoothSandstone,
	165: &SmoothStone,
	166: &Bricks,
	167: &Tnt,
	168: &Bookshelf,
	169: &MossyCobblestone,
	170: &Obsidian,
	171: &Torch,
	172: &EndRod,
	173: &ChorusPlant,
	174: &ChorusFlower,
	175: &PurpurBlock,
	176: &PurpurPillar,
	177: &PurpurStairs,
	178: &Spawner,
	179: &OakStairs,
	180: &Chest,
	181: &DiamondOre,
	182: &DiamondBlock,
	183: &CraftingTable,
	184: &Farmland,
	185: &Furnace,
	186: &Ladder,
	187: &Rail,
	188: &CobblestoneStairs,
	189: &Lever,
	190: &StonePressurePlate,
	191: &OakPressurePlate,
	192: &SprucePressurePlate,
	193: &BirchPressurePlate,
	194: &JunglePressurePlate,
	195: &AcaciaPressurePlate,
	196: &DarkOakPressurePlate,
	197: &CrimsonPressurePlate,
	198: &WarpedPressurePlate,
	199: &PolishedBlackstonePressurePlate,
	200: &RedstoneOre,
	201: &RedstoneTorch,
	202: &Snow,
	203: &Ice,
	204: &SnowBlock,
	205: &Cactus,
	206: &Clay,
	207: &Jukebox,
	208: &OakFence,
	209: &SpruceFence,
	210: &BirchFence,
	211: &JungleFence,
	212: &AcaciaFence,
	213: &DarkOakFence,
	214: &CrimsonFence,
	215: &WarpedFence,
	216: &Pumpkin,
	217: &CarvedPumpkin,
	218: &Netherrack,
	219: &SoulSand,
	220: &SoulSoil,
	221: &Basalt,
	222: &PolishedBasalt,
	223: &SoulTorch,
	224: &Glowstone,
	225: &JackOLantern,
	226: &OakTrapdoor,
	227: &SpruceTrapdoor,
	228: &BirchTrapdoor,
	229: &JungleTrapdoor,
	230: &AcaciaTrapdoor,
	231: &DarkOakTrapdoor,
	232: &CrimsonTrapdoor,
	233: &WarpedTrapdoor,
	234: &InfestedStone,
	235: &InfestedCobblestone,
	236: &InfestedStoneBricks,
	237: &InfestedMossyStoneBricks,
	238: &InfestedCrackedStoneBricks,
	239: &InfestedChiseledStoneBricks,
	240: &StoneBricks,
	241: &MossyStoneBricks,
	242: &CrackedStoneBricks,
	243: &ChiseledStoneBricks,
	244: &BrownMushroomBlock,
	245: &RedMushroomBlock,
	246: &MushroomStem,
	247: &IronBars,
	248: &Chain,
	249: &GlassPane,
	250: &Melon,
	251: &Vine,
	252: &OakFenceGate,
	253: &SpruceFenceGate,
	254: &BirchFenceGate,
	255: &JungleFenceGate,
	256: &AcaciaFenceGate,
	257: &DarkOakFenceGate,
	258: &CrimsonFenceGate,
	259: &WarpedFenceGate,
	260: &BrickStairs,
	261: &StoneBrickStairs,
	262: &Mycelium,
	263: &LilyPad,
	264: &NetherBricks,
	265: &CrackedNetherBricks,
	266: &ChiseledNetherBricks,
	267: &NetherBrickFence,
	268: &NetherBrickStairs,
	269: &EnchantingTable,
	270: &EndPortalFrame,
	271: &EndStone,
	272: &EndStoneBricks,
	273: &DragonEgg,
	274: &RedstoneLamp,
	275: &SandstoneStairs,
	276: &EmeraldOre,
	277: &EnderChest,
	278: &TripwireHook,
	279: &EmeraldBlock,
	280: &SpruceStairs,
	281: &BirchStairs,
	282: &JungleStairs,
	283: &CrimsonStairs,
	284: &WarpedStairs,
	285: &CommandBlock,
	286: &Beacon,
	287: &CobblestoneWall,
	288: &MossyCobblestoneWall,
	289: &BrickWall,
	290: &PrismarineWall,
	291: &RedSandstoneWall,
	292: &MossyStoneBrickWall,
	293: &GraniteWall,
	294: &StoneBrickWall,
	295: &NetherBrickWall,
	296: &AndesiteWall,
	297: &RedNetherBrickWall,
	298: &SandstoneWall,
	299: &EndStoneBrickWall,
	300: &DioriteWall,
	301: &BlackstoneWall,
	302: &PolishedBlackstoneWall,
	303: &PolishedBlackstoneBrickWall,
	304: &StoneButton,
	305: &OakButton,
	306: &SpruceButton,
	307: &BirchButton,
	308: &JungleButton,
	309: &AcaciaButton,
	310: &DarkOakButton,
	311: &CrimsonButton,
	312: &WarpedButton,
	313: &PolishedBlackstoneButton,
	314: &Anvil,
	315: &ChippedAnvil,
	316: &DamagedAnvil,
	317: &TrappedChest,
	318: &LightWeightedPressurePlate,
	319: &HeavyWeightedPressurePlate,
	320: &DaylightDetector,
	321: &RedstoneBlock,
	322: &NetherQuartzOre,
	323: &Hopper,
	324: &ChiseledQuartzBlock,
	325: &QuartzBlock,
	326: &QuartzBricks,
	327: &QuartzPillar,
	328: &QuartzStairs,
	329: &ActivatorRail,
	330: &Dropper,
	331: &WhiteTerracotta,
	332: &OrangeTerracotta,
	333: &MagentaTerracotta,
	334: &LightBlueTerracotta,
	335: &YellowTerracotta,
	336: &LimeTerracotta,
	337: &PinkTerracotta,
	338: &GrayTerracotta,
	339: &LightGrayTerracotta,
	340: &CyanTerracotta,
	341: &PurpleTerracotta,
	342: &BlueTerracotta,
	343: &BrownTerracotta,
	344: &GreenTerracotta,
	345: &RedTerracotta,
	346: &BlackTerracotta,
	347: &Barrier,
	348: &IronTrapdoor,
	349: &HayBlock,
	350: &WhiteCarpet,
	351: &OrangeCarpet,
	352: &MagentaCarpet,
	353: &LightBlueCarpet,
	354: &YellowCarpet,
	355: &LimeCarpet,
	356: &PinkCarpet,
	357: &GrayCarpet,
	358: &LightGrayCarpet,
	359: &CyanCarpet,
	360: &PurpleCarpet,
	361: &BlueCarpet,
	362: &BrownCarpet,
	363: &GreenCarpet,
	364: &RedCarpet,
	365: &BlackCarpet,
	366: &Terracotta,
	367: &CoalBlock,
	368: &PackedIce,
	369: &AcaciaStairs,
	370: &DarkOakStairs,
	371: &SlimeBlock,
	372: &GrassPath,
	373: &Sunflower,
	374: &Lilac,
	375: &RoseBush,
	376: &Peony,
	377: &TallGrass,
	378: &LargeFern,
	379: &WhiteStainedGlass,
	380: &OrangeStainedGlass,
	381: &MagentaStainedGlass,
	382: &LightBlueStainedGlass,
	383: &YellowStainedGlass,
	384: &LimeStainedGlass,
	385: &PinkStainedGlass,
	386: &GrayStainedGlass,
	387: &LightGrayStainedGlass,
	388: &CyanStainedGlass,
	389: &PurpleStainedGlass,
	390: &BlueStainedGlass,
	391: &BrownStainedGlass,
	392: &GreenStainedGlass,
	393: &RedStainedGlass,
	394: &BlackStainedGlass,
	395: &WhiteStainedGlassPane,
	396: &OrangeStainedGlassPane,
	397: &MagentaStainedGlassPane,
	398: &LightBlueStainedGlassPane,
	399: &YellowStainedGlassPane,
	400: &LimeStainedGlassPane,
	401: &PinkStainedGlassPane,
	402: &GrayStainedGlassPane,
	403: &LightGrayStainedGlassPane,
	404: &CyanStainedGlassPane,
	405: &PurpleStainedGlassPane,
	406: &BlueStainedGlassPane,
	407: &BrownStainedGlassPane,
	408: &GreenStainedGlassPane,
	409: &RedStainedGlassPane,
	410: &BlackStainedGlassPane,
	411: &Prismarine,
	412: &PrismarineBricks,
	413: &DarkPrismarine,
	414: &PrismarineStairs,
	415: &PrismarineBrickStairs,
	416: &DarkPrismarineStairs,
	417: &SeaLantern,
	418: &RedSandstone,
	419: &ChiseledRedSandstone,
	420: &CutRedSandstone,
	421: &RedSandstoneStairs,
	422: &RepeatingCommandBlock,
	423: &ChainCommandBlock,
	424: &MagmaBlock,
	425: &NetherWartBlock,
	426: &WarpedWartBlock,
	427: &RedNetherBricks,
	428: &BoneBlock,
	429: &StructureVoid,
	430: &Observer,
	431: &ShulkerBox,
	432: &WhiteShulkerBox,
	433: &OrangeShulkerBox,
	434: &MagentaShulkerBox,
	435: &LightBlueShulkerBox,
	436: &YellowShulkerBox,
	437: &LimeShulkerBox,
	438: &PinkShulkerBox,
	439: &GrayShulkerBox,
	440: &LightGrayShulkerBox,
	441: &CyanShulkerBox,
	442: &PurpleShulkerBox,
	443: &BlueShulkerBox,
	444: &BrownShulkerBox,
	445: &GreenShulkerBox,
	446: &RedShulkerBox,
	447: &BlackShulkerBox,
	448: &WhiteGlazedTerracotta,
	449: &OrangeGlazedTerracotta,
	450: &MagentaGlazedTerracotta,
	451: &LightBlueGlazedTerracotta,
	452: &YellowGlazedTerracotta,
	453: &LimeGlazedTerracotta,
	454: &PinkGlazedTerracotta,
	455: &GrayGlazedTerracotta,
	456: &LightGrayGlazedTerracotta,
	457: &CyanGlazedTerracotta,
	458: &PurpleGlazedTerracotta,
	459: &BlueGlazedTerracotta,
	460: &BrownGlazedTerracotta,
	461: &GreenGlazedTerracotta,
	462: &RedGlazedTerracotta,
	463: &BlackGlazedTerracotta,
	464: &WhiteConcrete,
	465: &OrangeConcrete,
	466: &MagentaConcrete,
	467: &LightBlueConcrete,
	468: &YellowConcrete,
	469: &LimeConcrete,
	470: &PinkConcrete,
	471: &GrayConcrete,
	472: &LightGrayConcrete,
	473: &CyanConcrete,
	474: &PurpleConcrete,
	475: &BlueConcrete,
	476: &BrownConcrete,
	477: &GreenConcrete,
	478: &RedConcrete,
	479: &BlackConcrete,
	480: &WhiteConcretePowder,
	481: &OrangeConcretePowder,
	482: &MagentaConcretePowder,
	483: &LightBlueConcretePowder,
	484: &YellowConcretePowder,
	485: &LimeConcretePowder,
	486: &PinkConcretePowder,
	487: &GrayConcretePowder,
	488: &LightGrayConcretePowder,
	489: &CyanConcretePowder,
	490: &PurpleConcretePowder,
	491: &BlueConcretePowder,
	492: &BrownConcretePowder,
	493: &GreenConcretePowder,
	494: &RedConcretePowder,
	495: &BlackConcretePowder,
	496: &TurtleEgg,
	497: &DeadTubeCoralBlock,
	498: &DeadBrainCoralBlock,
	499: &DeadBubbleCoralBlock,
	500: &DeadFireCoralBlock,
	501: &DeadHornCoralBlock,
	502: &TubeCoralBlock,
	503: &BrainCoralBlock,
	504: &BubbleCoralBlock,
	505: &FireCoralBlock,
	506: &HornCoralBlock,
	507: &TubeCoral,
	508: &BrainCoral,
	509: &BubbleCoral,
	510: &FireCoral,
	511: &HornCoral,
	512: &DeadBrainCoral,
	513: &DeadBubbleCoral,
	514: &DeadFireCoral,
	515: &DeadHornCoral,
	516: &DeadTubeCoral,
	517: &TubeCoralFan,
	518: &BrainCoralFan,
	519: &BubbleCoralFan,
	520: &FireCoralFan,
	521: &HornCoralFan,
	522: &DeadTubeCoralFan,
	523: &DeadBrainCoralFan,
	524: &DeadBubbleCoralFan,
	525: &DeadFireCoralFan,
	526: &DeadHornCoralFan,
	527: &BlueIce,
	528: &Conduit,
	529: &PolishedGraniteStairs,
	530: &SmoothRedSandstoneStairs,
	531: &MossyStoneBrickStairs,
	532: &PolishedDioriteStairs,
	533: &MossyCobblestoneStairs,
	534: &EndStoneBrickStairs,
	535: &StoneStairs,
	536: &SmoothSandstoneStairs,
	537: &SmoothQuartzStairs,
	538: &GraniteStairs,
	539: &AndesiteStairs,
	540: &RedNetherBrickStairs,
	541: &PolishedAndesiteStairs,
	542: &DioriteStairs,
	543: &PolishedGraniteSlab,
	544: &SmoothRedSandstoneSlab,
	545: &MossyStoneBrickSlab,
	546: &PolishedDioriteSlab,
	547: &MossyCobblestoneSlab,
	548: &EndStoneBrickSlab,
	549: &SmoothSandstoneSlab,
	550: &SmoothQuartzSlab,
	551: &GraniteSlab,
	552: &AndesiteSlab,
	553: &RedNetherBrickSlab,
	554: &PolishedAndesiteSlab,
	555: &DioriteSlab,
	556: &Scaffolding,
	557: &IronDoor,
	558: &OakDoor,
	559: &SpruceDoor,
	560: &BirchDoor,
	561: &JungleDoor,
	562: &AcaciaDoor,
	563: &DarkOakDoor,
	564: &CrimsonDoor,
	565: &WarpedDoor,
	566: &Repeater,
	567: &Comparator,
	568: &StructureBlock,
	569: &Jigsaw,
	570: &TurtleHelmet,
	571: &Scute,
	572: &FlintAndSteel,
	573: &Apple,
	574: &Bow,
	575: &Arrow,
	576: &Coal,
	577: &Charcoal,
	578: &Diamond,
	579: &IronIngot,
	580: &GoldIngot,
	581: &NetheriteIngot,
	582: &NetheriteScrap,
	583: &WoodenSword,
	584: &WoodenShovel,
	585: &WoodenPickaxe,
	586: &WoodenAxe,
	587: &WoodenHoe,
	588: &StoneSword,
	589: &StoneShovel,
	590: &StonePickaxe,
	591: &StoneAxe,
	592: &StoneHoe,
	593: &GoldenSword,
	594: &GoldenShovel,
	595: &GoldenPickaxe,
	596: &GoldenAxe,
	597: &GoldenHoe,
	598: &IronSword,
	599: &IronShovel,
	600: &IronPickaxe,
	601: &IronAxe,
	602: &IronHoe,
	603: &DiamondSword,
	604: &DiamondShovel,
	605: &DiamondPickaxe,
	606: &DiamondAxe,
	607: &DiamondHoe,
	608: &NetheriteSword,
	609: &NetheriteShovel,
	610: &NetheritePickaxe,
	611: &NetheriteAxe,
	612: &NetheriteHoe,
	613: &Stick,
	614: &Bowl,
	615: &MushroomStew,
	616: &String,
	617: &Feather,
	618: &Gunpowder,
	619: &WheatSeeds,
	620: &Wheat,
	621: &Bread,
	622: &LeatherHelmet,
	623: &LeatherChestplate,
	624: &LeatherLeggings,
	625: &LeatherBoots,
	626: &ChainmailHelmet,
	627: &ChainmailChestplate,
	628: &ChainmailLeggings,
	629: &ChainmailBoots,
	630: &IronHelmet,
	631: &IronChestplate,
	632: &IronLeggings,
	633: &IronBoots,
	634: &DiamondHelmet,
	635: &DiamondChestplate,
	636: &DiamondLeggings,
	637: &DiamondBoots,
	638: &GoldenHelmet,
	639: &GoldenChestplate,
	640: &GoldenLeggings,
	641: &GoldenBoots,
	642: &NetheriteHelmet,
	643: &NetheriteChestplate,
	644: &NetheriteLeggings,
	645: &NetheriteBoots,
	646: &Flint,
	647: &Porkchop,
	648: &CookedPorkchop,
	649: &Painting,
	650: &GoldenApple,
	651: &EnchantedGoldenApple,
	652: &OakSign,
	653: &SpruceSign,
	654: &BirchSign,
	655: &JungleSign,
	656: &AcaciaSign,
	657: &DarkOakSign,
	658: &CrimsonSign,
	659: &WarpedSign,
	660: &Bucket,
	661: &WaterBucket,
	662: &LavaBucket,
	663: &Minecart,
	664: &Saddle,
	665: &Redstone,
	666: &Snowball,
	667: &OakBoat,
	668: &Leather,
	669: &MilkBucket,
	670: &PufferfishBucket,
	671: &SalmonBucket,
	672: &CodBucket,
	673: &TropicalFishBucket,
	674: &Brick,
	675: &ClayBall,
	676: &DriedKelpBlock,
	677: &Paper,
	678: &Book,
	679: &SlimeBall,
	680: &ChestMinecart,
	681: &FurnaceMinecart,
	682: &Egg,
	683: &Compass,
	684: &FishingRod,
	685: &Clock,
	686: &GlowstoneDust,
	687: &Cod,
	688: &Salmon,
	689: &TropicalFish,
	690: &Pufferfish,
	691: &CookedCod,
	692: &CookedSalmon,
	693: &InkSac,
	694: &CocoaBeans,
	695: &LapisLazuli,
	696: &WhiteDye,
	697: &OrangeDye,
	698: &MagentaDye,
	699: &LightBlueDye,
	700: &YellowDye,
	701: &LimeDye,
	702: &PinkDye,
	703: &GrayDye,
	704: &LightGrayDye,
	705: &CyanDye,
	706: &PurpleDye,
	707: &BlueDye,
	708: &BrownDye,
	709: &GreenDye,
	710: &RedDye,
	711: &BlackDye,
	712: &BoneMeal,
	713: &Bone,
	714: &Sugar,
	715: &Cake,
	716: &WhiteBed,
	717: &OrangeBed,
	718: &MagentaBed,
	719: &LightBlueBed,
	720: &YellowBed,
	721: &LimeBed,
	722: &PinkBed,
	723: &GrayBed,
	724: &LightGrayBed,
	725: &CyanBed,
	726: &PurpleBed,
	727: &BlueBed,
	728: &BrownBed,
	729: &GreenBed,
	730: &RedBed,
	731: &BlackBed,
	732: &Cookie,
	733: &FilledMap,
	734: &Shears,
	735: &MelonSlice,
	736: &DriedKelp,
	737: &PumpkinSeeds,
	738: &MelonSeeds,
	739: &Beef,
	740: &CookedBeef,
	741: &Chicken,
	742: &CookedChicken,
	743: &RottenFlesh,
	744: &EnderPearl,
	745: &BlazeRod,
	746: &GhastTear,
	747: &GoldNugget,
	748: &NetherWart,
	749: &Potion,
	750: &GlassBottle,
	751: &SpiderEye,
	752: &FermentedSpiderEye,
	753: &BlazePowder,
	754: &MagmaCream,
	755: &BrewingStand,
	756: &Cauldron,
	757: &EnderEye,
	758: &GlisteringMelonSlice,
	759: &BatSpawnEgg,
	760: &BeeSpawnEgg,
	761: &BlazeSpawnEgg,
	762: &CatSpawnEgg,
	763: &CaveSpiderSpawnEgg,
	764: &ChickenSpawnEgg,
	765: &CodSpawnEgg,
	766: &CowSpawnEgg,
	767: &CreeperSpawnEgg,
	768: &DolphinSpawnEgg,
	769: &DonkeySpawnEgg,
	770: &DrownedSpawnEgg,
	771: &ElderGuardianSpawnEgg,
	772: &EndermanSpawnEgg,
	773: &EndermiteSpawnEgg,
	774: &EvokerSpawnEgg,
	775: &FoxSpawnEgg,
	776: &GhastSpawnEgg,
	777: &GuardianSpawnEgg,
	778: &HoglinSpawnEgg,
	779: &HorseSpawnEgg,
	780: &HuskSpawnEgg,
	781: &LlamaSpawnEgg,
	782: &MagmaCubeSpawnEgg,
	783: &MooshroomSpawnEgg,
	784: &MuleSpawnEgg,
	785: &OcelotSpawnEgg,
	786: &PandaSpawnEgg,
	787: &ParrotSpawnEgg,
	788: &PhantomSpawnEgg,
	789: &PigSpawnEgg,
	790: &PiglinSpawnEgg,
	791: &PiglinBruteSpawnEgg,
	792: &PillagerSpawnEgg,
	793: &PolarBearSpawnEgg,
	794: &PufferfishSpawnEgg,
	795: &RabbitSpawnEgg,
	796: &RavagerSpawnEgg,
	797: &SalmonSpawnEgg,
	798: &SheepSpawnEgg,
	799: &ShulkerSpawnEgg,
	800: &SilverfishSpawnEgg,
	801: &SkeletonSpawnEgg,
	802: &SkeletonHorseSpawnEgg,
	803: &SlimeSpawnEgg,
	804: &SpiderSpawnEgg,
	805: &SquidSpawnEgg,
	806: &StraySpawnEgg,
	807: &StriderSpawnEgg,
	808: &TraderLlamaSpawnEgg,
	809: &TropicalFishSpawnEgg,
	810: &TurtleSpawnEgg,
	811: &VexSpawnEgg,
	812: &VillagerSpawnEgg,
	813: &VindicatorSpawnEgg,
	814: &WanderingTraderSpawnEgg,
	815: &WitchSpawnEgg,
	816: &WitherSkeletonSpawnEgg,
	817: &WolfSpawnEgg,
	818: &ZoglinSpawnEgg,
	819: &ZombieSpawnEgg,
	820: &ZombieHorseSpawnEgg,
	821: &ZombieVillagerSpawnEgg,
	822: &ZombifiedPiglinSpawnEgg,
	823: &ExperienceBottle,
	824: &FireCharge,
	825: &WritableBook,
	826: &WrittenBook,
	827: &Emerald,
	828: &ItemFrame,
	829: &FlowerPot,
	830: &Carrot,
	831: &Potato,
	832: &BakedPotato,
	833: &PoisonousPotato,
	834: &Map,
	835: &GoldenCarrot,
	836: &SkeletonSkull,
	837: &WitherSkeletonSkull,
	838: &PlayerHead,
	839: &ZombieHead,
	840: &CreeperHead,
	841: &DragonHead,
	842: &CarrotOnAStick,
	843: &WarpedFungusOnAStick,
	844: &NetherStar,
	845: &PumpkinPie,
	846: &FireworkRocket,
	847: &FireworkStar,
	848: &EnchantedBook,
	849: &NetherBrick,
	850: &Quartz,
	851: &TntMinecart,
	852: &HopperMinecart,
	853: &PrismarineShard,
	854: &PrismarineCrystals,
	855: &Rabbit,
	856: &CookedRabbit,
	857: &RabbitStew,
	858: &RabbitFoot,
	859: &RabbitHide,
	860: &ArmorStand,
	861: &IronHorseArmor,
	862: &GoldenHorseArmor,
	863: &DiamondHorseArmor,
	864: &LeatherHorseArmor,
	865: &Lead,
	866: &NameTag,
	867: &CommandBlockMinecart,
	868: &Mutton,
	869: &CookedMutton,
	870: &WhiteBanner,
	871: &OrangeBanner,
	872: &MagentaBanner,
	873: &LightBlueBanner,
	874: &YellowBanner,
	875: &LimeBanner,
	876: &PinkBanner,
	877: &GrayBanner,
	878: &LightGrayBanner,
	879: &CyanBanner,
	880: &PurpleBanner,
	881: &BlueBanner,
	882: &BrownBanner,
	883: &GreenBanner,
	884: &RedBanner,
	885: &BlackBanner,
	886: &EndCrystal,
	887: &ChorusFruit,
	888: &PoppedChorusFruit,
	889: &Beetroot,
	890: &BeetrootSeeds,
	891: &BeetrootSoup,
	892: &DragonBreath,
	893: &SplashPotion,
	894: &SpectralArrow,
	895: &TippedArrow,
	896: &LingeringPotion,
	897: &Shield,
	898: &Elytra,
	899: &SpruceBoat,
	900: &BirchBoat,
	901: &JungleBoat,
	902: &AcaciaBoat,
	903: &DarkOakBoat,
	904: &TotemOfUndying,
	905: &ShulkerShell,
	906: &IronNugget,
	907: &KnowledgeBook,
	908: &DebugStick,
	909: &MusicDisc13,
	910: &MusicDiscCat,
	911: &MusicDiscBlocks,
	912: &MusicDiscChirp,
	913: &MusicDiscFar,
	914: &MusicDiscMall,
	915: &MusicDiscMellohi,
	916: &MusicDiscStal,
	917: &MusicDiscStrad,
	918: &MusicDiscWard,
	919: &MusicDisc11,
	920: &MusicDiscWait,
	921: &MusicDiscPigstep,
	922: &Trident,
	923: &PhantomMembrane,
	924: &NautilusShell,
	925: &HeartOfTheSea,
	926: &Crossbow,
	927: &SuspiciousStew,
	928: &Loom,
	929: &FlowerBannerPattern,
	930: &CreeperBannerPattern,
	931: &SkullBannerPattern,
	932: &MojangBannerPattern,
	933: &GlobeBannerPattern,
	934: &PiglinBannerPattern,
	935: &Composter,
	936: &Barrel,
	937: &Smoker,
	938: &BlastFurnace,
	939: &CartographyTable,
	940: &FletchingTable,
	941: &Grindstone,
	942: &Lectern,
	943: &SmithingTable,
	944: &Stonecutter,
	945: &Bell,
	946: &Lantern,
	947: &SoulLantern,
	948: &SweetBerries,
	949: &Campfire,
	950: &SoulCampfire,
	951: &Shroomlight,
	952: &Honeycomb,
	953: &BeeNest,
	954: &Beehive,
	955: &HoneyBottle,
	956: &HoneyBlock,
	957: &HoneycombBlock,
	958: &Lodestone,
	959: &NetheriteBlock,
	960: &AncientDebris,
	961: &Target,
	962: &CryingObsidian,
	963: &Blackstone,
	964: &BlackstoneSlab,
	965: &BlackstoneStairs,
	966: &GildedBlackstone,
	967: &PolishedBlackstone,
	968: &PolishedBlackstoneSlab,
	969: &PolishedBlackstoneStairs,
	970: &ChiseledPolishedBlackstone,
	971: &PolishedBlackstoneBricks,
	972: &PolishedBlackstoneBrickSlab,
	973: &PolishedBlackstoneBrickStairs,
	974: &CrackedPolishedBlackstoneBricks,
	975: &RespawnAnchor,
}

// ByName is an index of minecraft items by their name.
var ByName = map[string]*Item{
	"air":                                &Air,
	"stone":                              &Stone,
	"granite":                            &Granite,
	"polished_granite":                   &PolishedGranite,
	"diorite":                            &Diorite,
	"polished_diorite":                   &PolishedDiorite,
	"andesite":                           &Andesite,
	"polished_andesite":                  &PolishedAndesite,
	"grass_block":                        &GrassBlock,
	"dirt":                               &Dirt,
	"coarse_dirt":                        &CoarseDirt,
	"podzol":                             &Podzol,
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
	"gold_ore":                           &GoldOre,
	"iron_ore":                           &IronOre,
	"coal_ore":                           &CoalOre,
	"nether_gold_ore":                    &NetherGoldOre,
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
	"sponge":                             &Sponge,
	"wet_sponge":                         &WetSponge,
	"glass":                              &Glass,
	"lapis_ore":                          &LapisOre,
	"lapis_block":                        &LapisBlock,
	"dispenser":                          &Dispenser,
	"sandstone":                          &Sandstone,
	"chiseled_sandstone":                 &ChiseledSandstone,
	"cut_sandstone":                      &CutSandstone,
	"note_block":                         &NoteBlock,
	"powered_rail":                       &PoweredRail,
	"detector_rail":                      &DetectorRail,
	"sticky_piston":                      &StickyPiston,
	"cobweb":                             &Cobweb,
	"grass":                              &Grass,
	"fern":                               &Fern,
	"dead_bush":                          &DeadBush,
	"seagrass":                           &Seagrass,
	"sea_pickle":                         &SeaPickle,
	"piston":                             &Piston,
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
	"bamboo":                             &Bamboo,
	"gold_block":                         &GoldBlock,
	"iron_block":                         &IronBlock,
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
	"tnt":                                &Tnt,
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
	"diamond_ore":                        &DiamondOre,
	"diamond_block":                      &DiamondBlock,
	"crafting_table":                     &CraftingTable,
	"farmland":                           &Farmland,
	"furnace":                            &Furnace,
	"ladder":                             &Ladder,
	"rail":                               &Rail,
	"cobblestone_stairs":                 &CobblestoneStairs,
	"lever":                              &Lever,
	"stone_pressure_plate":               &StonePressurePlate,
	"oak_pressure_plate":                 &OakPressurePlate,
	"spruce_pressure_plate":              &SprucePressurePlate,
	"birch_pressure_plate":               &BirchPressurePlate,
	"jungle_pressure_plate":              &JunglePressurePlate,
	"acacia_pressure_plate":              &AcaciaPressurePlate,
	"dark_oak_pressure_plate":            &DarkOakPressurePlate,
	"crimson_pressure_plate":             &CrimsonPressurePlate,
	"warped_pressure_plate":              &WarpedPressurePlate,
	"polished_blackstone_pressure_plate": &PolishedBlackstonePressurePlate,
	"redstone_ore":                       &RedstoneOre,
	"redstone_torch":                     &RedstoneTorch,
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
	"netherrack":                         &Netherrack,
	"soul_sand":                          &SoulSand,
	"soul_soil":                          &SoulSoil,
	"basalt":                             &Basalt,
	"polished_basalt":                    &PolishedBasalt,
	"soul_torch":                         &SoulTorch,
	"glowstone":                          &Glowstone,
	"jack_o_lantern":                     &JackOLantern,
	"oak_trapdoor":                       &OakTrapdoor,
	"spruce_trapdoor":                    &SpruceTrapdoor,
	"birch_trapdoor":                     &BirchTrapdoor,
	"jungle_trapdoor":                    &JungleTrapdoor,
	"acacia_trapdoor":                    &AcaciaTrapdoor,
	"dark_oak_trapdoor":                  &DarkOakTrapdoor,
	"crimson_trapdoor":                   &CrimsonTrapdoor,
	"warped_trapdoor":                    &WarpedTrapdoor,
	"infested_stone":                     &InfestedStone,
	"infested_cobblestone":               &InfestedCobblestone,
	"infested_stone_bricks":              &InfestedStoneBricks,
	"infested_mossy_stone_bricks":        &InfestedMossyStoneBricks,
	"infested_cracked_stone_bricks":      &InfestedCrackedStoneBricks,
	"infested_chiseled_stone_bricks":     &InfestedChiseledStoneBricks,
	"stone_bricks":                       &StoneBricks,
	"mossy_stone_bricks":                 &MossyStoneBricks,
	"cracked_stone_bricks":               &CrackedStoneBricks,
	"chiseled_stone_bricks":              &ChiseledStoneBricks,
	"brown_mushroom_block":               &BrownMushroomBlock,
	"red_mushroom_block":                 &RedMushroomBlock,
	"mushroom_stem":                      &MushroomStem,
	"iron_bars":                          &IronBars,
	"chain":                              &Chain,
	"glass_pane":                         &GlassPane,
	"melon":                              &Melon,
	"vine":                               &Vine,
	"oak_fence_gate":                     &OakFenceGate,
	"spruce_fence_gate":                  &SpruceFenceGate,
	"birch_fence_gate":                   &BirchFenceGate,
	"jungle_fence_gate":                  &JungleFenceGate,
	"acacia_fence_gate":                  &AcaciaFenceGate,
	"dark_oak_fence_gate":                &DarkOakFenceGate,
	"crimson_fence_gate":                 &CrimsonFenceGate,
	"warped_fence_gate":                  &WarpedFenceGate,
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
	"redstone_lamp":                      &RedstoneLamp,
	"sandstone_stairs":                   &SandstoneStairs,
	"emerald_ore":                        &EmeraldOre,
	"ender_chest":                        &EnderChest,
	"tripwire_hook":                      &TripwireHook,
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
	"stone_button":                       &StoneButton,
	"oak_button":                         &OakButton,
	"spruce_button":                      &SpruceButton,
	"birch_button":                       &BirchButton,
	"jungle_button":                      &JungleButton,
	"acacia_button":                      &AcaciaButton,
	"dark_oak_button":                    &DarkOakButton,
	"crimson_button":                     &CrimsonButton,
	"warped_button":                      &WarpedButton,
	"polished_blackstone_button":         &PolishedBlackstoneButton,
	"anvil":                              &Anvil,
	"chipped_anvil":                      &ChippedAnvil,
	"damaged_anvil":                      &DamagedAnvil,
	"trapped_chest":                      &TrappedChest,
	"light_weighted_pressure_plate":      &LightWeightedPressurePlate,
	"heavy_weighted_pressure_plate":      &HeavyWeightedPressurePlate,
	"daylight_detector":                  &DaylightDetector,
	"redstone_block":                     &RedstoneBlock,
	"nether_quartz_ore":                  &NetherQuartzOre,
	"hopper":                             &Hopper,
	"chiseled_quartz_block":              &ChiseledQuartzBlock,
	"quartz_block":                       &QuartzBlock,
	"quartz_bricks":                      &QuartzBricks,
	"quartz_pillar":                      &QuartzPillar,
	"quartz_stairs":                      &QuartzStairs,
	"activator_rail":                     &ActivatorRail,
	"dropper":                            &Dropper,
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
	"iron_trapdoor":                      &IronTrapdoor,
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
	"coal_block":                         &CoalBlock,
	"packed_ice":                         &PackedIce,
	"acacia_stairs":                      &AcaciaStairs,
	"dark_oak_stairs":                    &DarkOakStairs,
	"slime_block":                        &SlimeBlock,
	"grass_path":                         &GrassPath,
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
	"observer":                           &Observer,
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
	"scaffolding":                        &Scaffolding,
	"iron_door":                          &IronDoor,
	"oak_door":                           &OakDoor,
	"spruce_door":                        &SpruceDoor,
	"birch_door":                         &BirchDoor,
	"jungle_door":                        &JungleDoor,
	"acacia_door":                        &AcaciaDoor,
	"dark_oak_door":                      &DarkOakDoor,
	"crimson_door":                       &CrimsonDoor,
	"warped_door":                        &WarpedDoor,
	"repeater":                           &Repeater,
	"comparator":                         &Comparator,
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
	"iron_ingot":                         &IronIngot,
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
	"minecart":                           &Minecart,
	"saddle":                             &Saddle,
	"redstone":                           &Redstone,
	"snowball":                           &Snowball,
	"oak_boat":                           &OakBoat,
	"leather":                            &Leather,
	"milk_bucket":                        &MilkBucket,
	"pufferfish_bucket":                  &PufferfishBucket,
	"salmon_bucket":                      &SalmonBucket,
	"cod_bucket":                         &CodBucket,
	"tropical_fish_bucket":               &TropicalFishBucket,
	"brick":                              &Brick,
	"clay_ball":                          &ClayBall,
	"dried_kelp_block":                   &DriedKelpBlock,
	"paper":                              &Paper,
	"book":                               &Book,
	"slime_ball":                         &SlimeBall,
	"chest_minecart":                     &ChestMinecart,
	"furnace_minecart":                   &FurnaceMinecart,
	"egg":                                &Egg,
	"compass":                            &Compass,
	"fishing_rod":                        &FishingRod,
	"clock":                              &Clock,
	"glowstone_dust":                     &GlowstoneDust,
	"cod":                                &Cod,
	"salmon":                             &Salmon,
	"tropical_fish":                      &TropicalFish,
	"pufferfish":                         &Pufferfish,
	"cooked_cod":                         &CookedCod,
	"cooked_salmon":                      &CookedSalmon,
	"ink_sac":                            &InkSac,
	"cocoa_beans":                        &CocoaBeans,
	"lapis_lazuli":                       &LapisLazuli,
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
	"emerald":                            &Emerald,
	"item_frame":                         &ItemFrame,
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
	"carrot_on_a_stick":                  &CarrotOnAStick,
	"warped_fungus_on_a_stick":           &WarpedFungusOnAStick,
	"nether_star":                        &NetherStar,
	"pumpkin_pie":                        &PumpkinPie,
	"firework_rocket":                    &FireworkRocket,
	"firework_star":                      &FireworkStar,
	"enchanted_book":                     &EnchantedBook,
	"nether_brick":                       &NetherBrick,
	"quartz":                             &Quartz,
	"tnt_minecart":                       &TntMinecart,
	"hopper_minecart":                    &HopperMinecart,
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
	"elytra":                             &Elytra,
	"spruce_boat":                        &SpruceBoat,
	"birch_boat":                         &BirchBoat,
	"jungle_boat":                        &JungleBoat,
	"acacia_boat":                        &AcaciaBoat,
	"dark_oak_boat":                      &DarkOakBoat,
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
	"lectern":                            &Lectern,
	"smithing_table":                     &SmithingTable,
	"stonecutter":                        &Stonecutter,
	"bell":                               &Bell,
	"lantern":                            &Lantern,
	"soul_lantern":                       &SoulLantern,
	"sweet_berries":                      &SweetBerries,
	"campfire":                           &Campfire,
	"soul_campfire":                      &SoulCampfire,
	"shroomlight":                        &Shroomlight,
	"honeycomb":                          &Honeycomb,
	"bee_nest":                           &BeeNest,
	"beehive":                            &Beehive,
	"honey_bottle":                       &HoneyBottle,
	"honey_block":                        &HoneyBlock,
	"honeycomb_block":                    &HoneycombBlock,
	"lodestone":                          &Lodestone,
	"netherite_block":                    &NetheriteBlock,
	"ancient_debris":                     &AncientDebris,
	"target":                             &Target,
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
}
