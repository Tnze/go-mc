package data

import "encoding/json"

var itemIDs map[string]struct {
	ProtocolID int `json:"protocol_id"`
}

// ItemNameByID store Item's ID and
var ItemNameByID []string

func init() {
	json.Unmarshal([]byte(itemIDsJSON), &itemIDs)
	ItemNameByID = make([]string, 876+1)
	for i, v := range itemIDs {
		ItemNameByID[v.ProtocolID] = i
	}
}

// Generate with follow steps:
// java -cp minecraft_server.1.15.jar net.minecraft.data.Main --all
// {reports/registries.json}.minecraft:block.entries
var itemIDsJSON = `{
    "minecraft:air": {
        "protocol_id": 0
    },
    "minecraft:stone": {
        "protocol_id": 1
    },
    "minecraft:granite": {
        "protocol_id": 2
    },
    "minecraft:polished_granite": {
        "protocol_id": 3
    },
    "minecraft:diorite": {
        "protocol_id": 4
    },
    "minecraft:polished_diorite": {
        "protocol_id": 5
    },
    "minecraft:andesite": {
        "protocol_id": 6
    },
    "minecraft:polished_andesite": {
        "protocol_id": 7
    },
    "minecraft:grass_block": {
        "protocol_id": 8
    },
    "minecraft:dirt": {
        "protocol_id": 9
    },
    "minecraft:coarse_dirt": {
        "protocol_id": 10
    },
    "minecraft:podzol": {
        "protocol_id": 11
    },
    "minecraft:cobblestone": {
        "protocol_id": 12
    },
    "minecraft:oak_planks": {
        "protocol_id": 13
    },
    "minecraft:spruce_planks": {
        "protocol_id": 14
    },
    "minecraft:birch_planks": {
        "protocol_id": 15
    },
    "minecraft:jungle_planks": {
        "protocol_id": 16
    },
    "minecraft:acacia_planks": {
        "protocol_id": 17
    },
    "minecraft:dark_oak_planks": {
        "protocol_id": 18
    },
    "minecraft:oak_sapling": {
        "protocol_id": 19
    },
    "minecraft:spruce_sapling": {
        "protocol_id": 20
    },
    "minecraft:birch_sapling": {
        "protocol_id": 21
    },
    "minecraft:jungle_sapling": {
        "protocol_id": 22
    },
    "minecraft:acacia_sapling": {
        "protocol_id": 23
    },
    "minecraft:dark_oak_sapling": {
        "protocol_id": 24
    },
    "minecraft:bedrock": {
        "protocol_id": 25
    },
    "minecraft:water": {
        "protocol_id": 26
    },
    "minecraft:lava": {
        "protocol_id": 27
    },
    "minecraft:sand": {
        "protocol_id": 28
    },
    "minecraft:red_sand": {
        "protocol_id": 29
    },
    "minecraft:gravel": {
        "protocol_id": 30
    },
    "minecraft:gold_ore": {
        "protocol_id": 31
    },
    "minecraft:iron_ore": {
        "protocol_id": 32
    },
    "minecraft:coal_ore": {
        "protocol_id": 33
    },
    "minecraft:nether_gold_ore": {
        "protocol_id": 34
    },
    "minecraft:oak_log": {
        "protocol_id": 35
    },
    "minecraft:spruce_log": {
        "protocol_id": 36
    },
    "minecraft:birch_log": {
        "protocol_id": 37
    },
    "minecraft:jungle_log": {
        "protocol_id": 38
    },
    "minecraft:acacia_log": {
        "protocol_id": 39
    },
    "minecraft:dark_oak_log": {
        "protocol_id": 40
    },
    "minecraft:stripped_spruce_log": {
        "protocol_id": 41
    },
    "minecraft:stripped_birch_log": {
        "protocol_id": 42
    },
    "minecraft:stripped_jungle_log": {
        "protocol_id": 43
    },
    "minecraft:stripped_acacia_log": {
        "protocol_id": 44
    },
    "minecraft:stripped_dark_oak_log": {
        "protocol_id": 45
    },
    "minecraft:stripped_oak_log": {
        "protocol_id": 46
    },
    "minecraft:oak_wood": {
        "protocol_id": 47
    },
    "minecraft:spruce_wood": {
        "protocol_id": 48
    },
    "minecraft:birch_wood": {
        "protocol_id": 49
    },
    "minecraft:jungle_wood": {
        "protocol_id": 50
    },
    "minecraft:acacia_wood": {
        "protocol_id": 51
    },
    "minecraft:dark_oak_wood": {
        "protocol_id": 52
    },
    "minecraft:stripped_oak_wood": {
        "protocol_id": 53
    },
    "minecraft:stripped_spruce_wood": {
        "protocol_id": 54
    },
    "minecraft:stripped_birch_wood": {
        "protocol_id": 55
    },
    "minecraft:stripped_jungle_wood": {
        "protocol_id": 56
    },
    "minecraft:stripped_acacia_wood": {
        "protocol_id": 57
    },
    "minecraft:stripped_dark_oak_wood": {
        "protocol_id": 58
    },
    "minecraft:oak_leaves": {
        "protocol_id": 59
    },
    "minecraft:spruce_leaves": {
        "protocol_id": 60
    },
    "minecraft:birch_leaves": {
        "protocol_id": 61
    },
    "minecraft:jungle_leaves": {
        "protocol_id": 62
    },
    "minecraft:acacia_leaves": {
        "protocol_id": 63
    },
    "minecraft:dark_oak_leaves": {
        "protocol_id": 64
    },
    "minecraft:sponge": {
        "protocol_id": 65
    },
    "minecraft:wet_sponge": {
        "protocol_id": 66
    },
    "minecraft:glass": {
        "protocol_id": 67
    },
    "minecraft:lapis_ore": {
        "protocol_id": 68
    },
    "minecraft:lapis_block": {
        "protocol_id": 69
    },
    "minecraft:dispenser": {
        "protocol_id": 70
    },
    "minecraft:sandstone": {
        "protocol_id": 71
    },
    "minecraft:chiseled_sandstone": {
        "protocol_id": 72
    },
    "minecraft:cut_sandstone": {
        "protocol_id": 73
    },
    "minecraft:note_block": {
        "protocol_id": 74
    },
    "minecraft:white_bed": {
        "protocol_id": 75
    },
    "minecraft:orange_bed": {
        "protocol_id": 76
    },
    "minecraft:magenta_bed": {
        "protocol_id": 77
    },
    "minecraft:light_blue_bed": {
        "protocol_id": 78
    },
    "minecraft:yellow_bed": {
        "protocol_id": 79
    },
    "minecraft:lime_bed": {
        "protocol_id": 80
    },
    "minecraft:pink_bed": {
        "protocol_id": 81
    },
    "minecraft:gray_bed": {
        "protocol_id": 82
    },
    "minecraft:light_gray_bed": {
        "protocol_id": 83
    },
    "minecraft:cyan_bed": {
        "protocol_id": 84
    },
    "minecraft:purple_bed": {
        "protocol_id": 85
    },
    "minecraft:blue_bed": {
        "protocol_id": 86
    },
    "minecraft:brown_bed": {
        "protocol_id": 87
    },
    "minecraft:green_bed": {
        "protocol_id": 88
    },
    "minecraft:red_bed": {
        "protocol_id": 89
    },
    "minecraft:black_bed": {
        "protocol_id": 90
    },
    "minecraft:powered_rail": {
        "protocol_id": 91
    },
    "minecraft:detector_rail": {
        "protocol_id": 92
    },
    "minecraft:sticky_piston": {
        "protocol_id": 93
    },
    "minecraft:cobweb": {
        "protocol_id": 94
    },
    "minecraft:grass": {
        "protocol_id": 95
    },
    "minecraft:fern": {
        "protocol_id": 96
    },
    "minecraft:dead_bush": {
        "protocol_id": 97
    },
    "minecraft:seagrass": {
        "protocol_id": 98
    },
    "minecraft:tall_seagrass": {
        "protocol_id": 99
    },
    "minecraft:piston": {
        "protocol_id": 100
    },
    "minecraft:piston_head": {
        "protocol_id": 101
    },
    "minecraft:white_wool": {
        "protocol_id": 102
    },
    "minecraft:orange_wool": {
        "protocol_id": 103
    },
    "minecraft:magenta_wool": {
        "protocol_id": 104
    },
    "minecraft:light_blue_wool": {
        "protocol_id": 105
    },
    "minecraft:yellow_wool": {
        "protocol_id": 106
    },
    "minecraft:lime_wool": {
        "protocol_id": 107
    },
    "minecraft:pink_wool": {
        "protocol_id": 108
    },
    "minecraft:gray_wool": {
        "protocol_id": 109
    },
    "minecraft:light_gray_wool": {
        "protocol_id": 110
    },
    "minecraft:cyan_wool": {
        "protocol_id": 111
    },
    "minecraft:purple_wool": {
        "protocol_id": 112
    },
    "minecraft:blue_wool": {
        "protocol_id": 113
    },
    "minecraft:brown_wool": {
        "protocol_id": 114
    },
    "minecraft:green_wool": {
        "protocol_id": 115
    },
    "minecraft:red_wool": {
        "protocol_id": 116
    },
    "minecraft:black_wool": {
        "protocol_id": 117
    },
    "minecraft:moving_piston": {
        "protocol_id": 118
    },
    "minecraft:dandelion": {
        "protocol_id": 119
    },
    "minecraft:poppy": {
        "protocol_id": 120
    },
    "minecraft:blue_orchid": {
        "protocol_id": 121
    },
    "minecraft:allium": {
        "protocol_id": 122
    },
    "minecraft:azure_bluet": {
        "protocol_id": 123
    },
    "minecraft:red_tulip": {
        "protocol_id": 124
    },
    "minecraft:orange_tulip": {
        "protocol_id": 125
    },
    "minecraft:white_tulip": {
        "protocol_id": 126
    },
    "minecraft:pink_tulip": {
        "protocol_id": 127
    },
    "minecraft:oxeye_daisy": {
        "protocol_id": 128
    },
    "minecraft:cornflower": {
        "protocol_id": 129
    },
    "minecraft:wither_rose": {
        "protocol_id": 130
    },
    "minecraft:lily_of_the_valley": {
        "protocol_id": 131
    },
    "minecraft:brown_mushroom": {
        "protocol_id": 132
    },
    "minecraft:red_mushroom": {
        "protocol_id": 133
    },
    "minecraft:gold_block": {
        "protocol_id": 134
    },
    "minecraft:iron_block": {
        "protocol_id": 135
    },
    "minecraft:bricks": {
        "protocol_id": 136
    },
    "minecraft:tnt": {
        "protocol_id": 137
    },
    "minecraft:bookshelf": {
        "protocol_id": 138
    },
    "minecraft:mossy_cobblestone": {
        "protocol_id": 139
    },
    "minecraft:obsidian": {
        "protocol_id": 140
    },
    "minecraft:torch": {
        "protocol_id": 141
    },
    "minecraft:wall_torch": {
        "protocol_id": 142
    },
    "minecraft:fire": {
        "protocol_id": 143
    },
    "minecraft:soul_fire": {
        "protocol_id": 144
    },
    "minecraft:spawner": {
        "protocol_id": 145
    },
    "minecraft:oak_stairs": {
        "protocol_id": 146
    },
    "minecraft:chest": {
        "protocol_id": 147
    },
    "minecraft:redstone_wire": {
        "protocol_id": 148
    },
    "minecraft:diamond_ore": {
        "protocol_id": 149
    },
    "minecraft:diamond_block": {
        "protocol_id": 150
    },
    "minecraft:crafting_table": {
        "protocol_id": 151
    },
    "minecraft:wheat": {
        "protocol_id": 152
    },
    "minecraft:farmland": {
        "protocol_id": 153
    },
    "minecraft:furnace": {
        "protocol_id": 154
    },
    "minecraft:oak_sign": {
        "protocol_id": 155
    },
    "minecraft:spruce_sign": {
        "protocol_id": 156
    },
    "minecraft:birch_sign": {
        "protocol_id": 157
    },
    "minecraft:acacia_sign": {
        "protocol_id": 158
    },
    "minecraft:jungle_sign": {
        "protocol_id": 159
    },
    "minecraft:dark_oak_sign": {
        "protocol_id": 160
    },
    "minecraft:oak_door": {
        "protocol_id": 161
    },
    "minecraft:ladder": {
        "protocol_id": 162
    },
    "minecraft:rail": {
        "protocol_id": 163
    },
    "minecraft:cobblestone_stairs": {
        "protocol_id": 164
    },
    "minecraft:oak_wall_sign": {
        "protocol_id": 165
    },
    "minecraft:spruce_wall_sign": {
        "protocol_id": 166
    },
    "minecraft:birch_wall_sign": {
        "protocol_id": 167
    },
    "minecraft:acacia_wall_sign": {
        "protocol_id": 168
    },
    "minecraft:jungle_wall_sign": {
        "protocol_id": 169
    },
    "minecraft:dark_oak_wall_sign": {
        "protocol_id": 170
    },
    "minecraft:lever": {
        "protocol_id": 171
    },
    "minecraft:stone_pressure_plate": {
        "protocol_id": 172
    },
    "minecraft:iron_door": {
        "protocol_id": 173
    },
    "minecraft:oak_pressure_plate": {
        "protocol_id": 174
    },
    "minecraft:spruce_pressure_plate": {
        "protocol_id": 175
    },
    "minecraft:birch_pressure_plate": {
        "protocol_id": 176
    },
    "minecraft:jungle_pressure_plate": {
        "protocol_id": 177
    },
    "minecraft:acacia_pressure_plate": {
        "protocol_id": 178
    },
    "minecraft:dark_oak_pressure_plate": {
        "protocol_id": 179
    },
    "minecraft:redstone_ore": {
        "protocol_id": 180
    },
    "minecraft:redstone_torch": {
        "protocol_id": 181
    },
    "minecraft:redstone_wall_torch": {
        "protocol_id": 182
    },
    "minecraft:stone_button": {
        "protocol_id": 183
    },
    "minecraft:snow": {
        "protocol_id": 184
    },
    "minecraft:ice": {
        "protocol_id": 185
    },
    "minecraft:snow_block": {
        "protocol_id": 186
    },
    "minecraft:cactus": {
        "protocol_id": 187
    },
    "minecraft:clay": {
        "protocol_id": 188
    },
    "minecraft:sugar_cane": {
        "protocol_id": 189
    },
    "minecraft:jukebox": {
        "protocol_id": 190
    },
    "minecraft:oak_fence": {
        "protocol_id": 191
    },
    "minecraft:pumpkin": {
        "protocol_id": 192
    },
    "minecraft:netherrack": {
        "protocol_id": 193
    },
    "minecraft:soul_sand": {
        "protocol_id": 194
    },
    "minecraft:soul_soil": {
        "protocol_id": 195
    },
    "minecraft:basalt": {
        "protocol_id": 196
    },
    "minecraft:polished_basalt": {
        "protocol_id": 197
    },
    "minecraft:soul_torch": {
        "protocol_id": 198
    },
    "minecraft:soul_wall_torch": {
        "protocol_id": 199
    },
    "minecraft:glowstone": {
        "protocol_id": 200
    },
    "minecraft:nether_portal": {
        "protocol_id": 201
    },
    "minecraft:carved_pumpkin": {
        "protocol_id": 202
    },
    "minecraft:jack_o_lantern": {
        "protocol_id": 203
    },
    "minecraft:cake": {
        "protocol_id": 204
    },
    "minecraft:repeater": {
        "protocol_id": 205
    },
    "minecraft:white_stained_glass": {
        "protocol_id": 206
    },
    "minecraft:orange_stained_glass": {
        "protocol_id": 207
    },
    "minecraft:magenta_stained_glass": {
        "protocol_id": 208
    },
    "minecraft:light_blue_stained_glass": {
        "protocol_id": 209
    },
    "minecraft:yellow_stained_glass": {
        "protocol_id": 210
    },
    "minecraft:lime_stained_glass": {
        "protocol_id": 211
    },
    "minecraft:pink_stained_glass": {
        "protocol_id": 212
    },
    "minecraft:gray_stained_glass": {
        "protocol_id": 213
    },
    "minecraft:light_gray_stained_glass": {
        "protocol_id": 214
    },
    "minecraft:cyan_stained_glass": {
        "protocol_id": 215
    },
    "minecraft:purple_stained_glass": {
        "protocol_id": 216
    },
    "minecraft:blue_stained_glass": {
        "protocol_id": 217
    },
    "minecraft:brown_stained_glass": {
        "protocol_id": 218
    },
    "minecraft:green_stained_glass": {
        "protocol_id": 219
    },
    "minecraft:red_stained_glass": {
        "protocol_id": 220
    },
    "minecraft:black_stained_glass": {
        "protocol_id": 221
    },
    "minecraft:oak_trapdoor": {
        "protocol_id": 222
    },
    "minecraft:spruce_trapdoor": {
        "protocol_id": 223
    },
    "minecraft:birch_trapdoor": {
        "protocol_id": 224
    },
    "minecraft:jungle_trapdoor": {
        "protocol_id": 225
    },
    "minecraft:acacia_trapdoor": {
        "protocol_id": 226
    },
    "minecraft:dark_oak_trapdoor": {
        "protocol_id": 227
    },
    "minecraft:stone_bricks": {
        "protocol_id": 228
    },
    "minecraft:mossy_stone_bricks": {
        "protocol_id": 229
    },
    "minecraft:cracked_stone_bricks": {
        "protocol_id": 230
    },
    "minecraft:chiseled_stone_bricks": {
        "protocol_id": 231
    },
    "minecraft:infested_stone": {
        "protocol_id": 232
    },
    "minecraft:infested_cobblestone": {
        "protocol_id": 233
    },
    "minecraft:infested_stone_bricks": {
        "protocol_id": 234
    },
    "minecraft:infested_mossy_stone_bricks": {
        "protocol_id": 235
    },
    "minecraft:infested_cracked_stone_bricks": {
        "protocol_id": 236
    },
    "minecraft:infested_chiseled_stone_bricks": {
        "protocol_id": 237
    },
    "minecraft:brown_mushroom_block": {
        "protocol_id": 238
    },
    "minecraft:red_mushroom_block": {
        "protocol_id": 239
    },
    "minecraft:mushroom_stem": {
        "protocol_id": 240
    },
    "minecraft:iron_bars": {
        "protocol_id": 241
    },
    "minecraft:chain": {
        "protocol_id": 242
    },
    "minecraft:glass_pane": {
        "protocol_id": 243
    },
    "minecraft:melon": {
        "protocol_id": 244
    },
    "minecraft:attached_pumpkin_stem": {
        "protocol_id": 245
    },
    "minecraft:attached_melon_stem": {
        "protocol_id": 246
    },
    "minecraft:pumpkin_stem": {
        "protocol_id": 247
    },
    "minecraft:melon_stem": {
        "protocol_id": 248
    },
    "minecraft:vine": {
        "protocol_id": 249
    },
    "minecraft:oak_fence_gate": {
        "protocol_id": 250
    },
    "minecraft:brick_stairs": {
        "protocol_id": 251
    },
    "minecraft:stone_brick_stairs": {
        "protocol_id": 252
    },
    "minecraft:mycelium": {
        "protocol_id": 253
    },
    "minecraft:lily_pad": {
        "protocol_id": 254
    },
    "minecraft:nether_bricks": {
        "protocol_id": 255
    },
    "minecraft:nether_brick_fence": {
        "protocol_id": 256
    },
    "minecraft:nether_brick_stairs": {
        "protocol_id": 257
    },
    "minecraft:nether_wart": {
        "protocol_id": 258
    },
    "minecraft:enchanting_table": {
        "protocol_id": 259
    },
    "minecraft:brewing_stand": {
        "protocol_id": 260
    },
    "minecraft:cauldron": {
        "protocol_id": 261
    },
    "minecraft:end_portal": {
        "protocol_id": 262
    },
    "minecraft:end_portal_frame": {
        "protocol_id": 263
    },
    "minecraft:end_stone": {
        "protocol_id": 264
    },
    "minecraft:dragon_egg": {
        "protocol_id": 265
    },
    "minecraft:redstone_lamp": {
        "protocol_id": 266
    },
    "minecraft:cocoa": {
        "protocol_id": 267
    },
    "minecraft:sandstone_stairs": {
        "protocol_id": 268
    },
    "minecraft:emerald_ore": {
        "protocol_id": 269
    },
    "minecraft:ender_chest": {
        "protocol_id": 270
    },
    "minecraft:tripwire_hook": {
        "protocol_id": 271
    },
    "minecraft:tripwire": {
        "protocol_id": 272
    },
    "minecraft:emerald_block": {
        "protocol_id": 273
    },
    "minecraft:spruce_stairs": {
        "protocol_id": 274
    },
    "minecraft:birch_stairs": {
        "protocol_id": 275
    },
    "minecraft:jungle_stairs": {
        "protocol_id": 276
    },
    "minecraft:command_block": {
        "protocol_id": 277
    },
    "minecraft:beacon": {
        "protocol_id": 278
    },
    "minecraft:cobblestone_wall": {
        "protocol_id": 279
    },
    "minecraft:mossy_cobblestone_wall": {
        "protocol_id": 280
    },
    "minecraft:flower_pot": {
        "protocol_id": 281
    },
    "minecraft:potted_oak_sapling": {
        "protocol_id": 282
    },
    "minecraft:potted_spruce_sapling": {
        "protocol_id": 283
    },
    "minecraft:potted_birch_sapling": {
        "protocol_id": 284
    },
    "minecraft:potted_jungle_sapling": {
        "protocol_id": 285
    },
    "minecraft:potted_acacia_sapling": {
        "protocol_id": 286
    },
    "minecraft:potted_dark_oak_sapling": {
        "protocol_id": 287
    },
    "minecraft:potted_fern": {
        "protocol_id": 288
    },
    "minecraft:potted_dandelion": {
        "protocol_id": 289
    },
    "minecraft:potted_poppy": {
        "protocol_id": 290
    },
    "minecraft:potted_blue_orchid": {
        "protocol_id": 291
    },
    "minecraft:potted_allium": {
        "protocol_id": 292
    },
    "minecraft:potted_azure_bluet": {
        "protocol_id": 293
    },
    "minecraft:potted_red_tulip": {
        "protocol_id": 294
    },
    "minecraft:potted_orange_tulip": {
        "protocol_id": 295
    },
    "minecraft:potted_white_tulip": {
        "protocol_id": 296
    },
    "minecraft:potted_pink_tulip": {
        "protocol_id": 297
    },
    "minecraft:potted_oxeye_daisy": {
        "protocol_id": 298
    },
    "minecraft:potted_cornflower": {
        "protocol_id": 299
    },
    "minecraft:potted_lily_of_the_valley": {
        "protocol_id": 300
    },
    "minecraft:potted_wither_rose": {
        "protocol_id": 301
    },
    "minecraft:potted_red_mushroom": {
        "protocol_id": 302
    },
    "minecraft:potted_brown_mushroom": {
        "protocol_id": 303
    },
    "minecraft:potted_dead_bush": {
        "protocol_id": 304
    },
    "minecraft:potted_cactus": {
        "protocol_id": 305
    },
    "minecraft:carrots": {
        "protocol_id": 306
    },
    "minecraft:potatoes": {
        "protocol_id": 307
    },
    "minecraft:oak_button": {
        "protocol_id": 308
    },
    "minecraft:spruce_button": {
        "protocol_id": 309
    },
    "minecraft:birch_button": {
        "protocol_id": 310
    },
    "minecraft:jungle_button": {
        "protocol_id": 311
    },
    "minecraft:acacia_button": {
        "protocol_id": 312
    },
    "minecraft:dark_oak_button": {
        "protocol_id": 313
    },
    "minecraft:skeleton_skull": {
        "protocol_id": 314
    },
    "minecraft:skeleton_wall_skull": {
        "protocol_id": 315
    },
    "minecraft:wither_skeleton_skull": {
        "protocol_id": 316
    },
    "minecraft:wither_skeleton_wall_skull": {
        "protocol_id": 317
    },
    "minecraft:zombie_head": {
        "protocol_id": 318
    },
    "minecraft:zombie_wall_head": {
        "protocol_id": 319
    },
    "minecraft:player_head": {
        "protocol_id": 320
    },
    "minecraft:player_wall_head": {
        "protocol_id": 321
    },
    "minecraft:creeper_head": {
        "protocol_id": 322
    },
    "minecraft:creeper_wall_head": {
        "protocol_id": 323
    },
    "minecraft:dragon_head": {
        "protocol_id": 324
    },
    "minecraft:dragon_wall_head": {
        "protocol_id": 325
    },
    "minecraft:anvil": {
        "protocol_id": 326
    },
    "minecraft:chipped_anvil": {
        "protocol_id": 327
    },
    "minecraft:damaged_anvil": {
        "protocol_id": 328
    },
    "minecraft:trapped_chest": {
        "protocol_id": 329
    },
    "minecraft:light_weighted_pressure_plate": {
        "protocol_id": 330
    },
    "minecraft:heavy_weighted_pressure_plate": {
        "protocol_id": 331
    },
    "minecraft:comparator": {
        "protocol_id": 332
    },
    "minecraft:daylight_detector": {
        "protocol_id": 333
    },
    "minecraft:redstone_block": {
        "protocol_id": 334
    },
    "minecraft:nether_quartz_ore": {
        "protocol_id": 335
    },
    "minecraft:hopper": {
        "protocol_id": 336
    },
    "minecraft:quartz_block": {
        "protocol_id": 337
    },
    "minecraft:chiseled_quartz_block": {
        "protocol_id": 338
    },
    "minecraft:quartz_pillar": {
        "protocol_id": 339
    },
    "minecraft:quartz_stairs": {
        "protocol_id": 340
    },
    "minecraft:activator_rail": {
        "protocol_id": 341
    },
    "minecraft:dropper": {
        "protocol_id": 342
    },
    "minecraft:white_terracotta": {
        "protocol_id": 343
    },
    "minecraft:orange_terracotta": {
        "protocol_id": 344
    },
    "minecraft:magenta_terracotta": {
        "protocol_id": 345
    },
    "minecraft:light_blue_terracotta": {
        "protocol_id": 346
    },
    "minecraft:yellow_terracotta": {
        "protocol_id": 347
    },
    "minecraft:lime_terracotta": {
        "protocol_id": 348
    },
    "minecraft:pink_terracotta": {
        "protocol_id": 349
    },
    "minecraft:gray_terracotta": {
        "protocol_id": 350
    },
    "minecraft:light_gray_terracotta": {
        "protocol_id": 351
    },
    "minecraft:cyan_terracotta": {
        "protocol_id": 352
    },
    "minecraft:purple_terracotta": {
        "protocol_id": 353
    },
    "minecraft:blue_terracotta": {
        "protocol_id": 354
    },
    "minecraft:brown_terracotta": {
        "protocol_id": 355
    },
    "minecraft:green_terracotta": {
        "protocol_id": 356
    },
    "minecraft:red_terracotta": {
        "protocol_id": 357
    },
    "minecraft:black_terracotta": {
        "protocol_id": 358
    },
    "minecraft:white_stained_glass_pane": {
        "protocol_id": 359
    },
    "minecraft:orange_stained_glass_pane": {
        "protocol_id": 360
    },
    "minecraft:magenta_stained_glass_pane": {
        "protocol_id": 361
    },
    "minecraft:light_blue_stained_glass_pane": {
        "protocol_id": 362
    },
    "minecraft:yellow_stained_glass_pane": {
        "protocol_id": 363
    },
    "minecraft:lime_stained_glass_pane": {
        "protocol_id": 364
    },
    "minecraft:pink_stained_glass_pane": {
        "protocol_id": 365
    },
    "minecraft:gray_stained_glass_pane": {
        "protocol_id": 366
    },
    "minecraft:light_gray_stained_glass_pane": {
        "protocol_id": 367
    },
    "minecraft:cyan_stained_glass_pane": {
        "protocol_id": 368
    },
    "minecraft:purple_stained_glass_pane": {
        "protocol_id": 369
    },
    "minecraft:blue_stained_glass_pane": {
        "protocol_id": 370
    },
    "minecraft:brown_stained_glass_pane": {
        "protocol_id": 371
    },
    "minecraft:green_stained_glass_pane": {
        "protocol_id": 372
    },
    "minecraft:red_stained_glass_pane": {
        "protocol_id": 373
    },
    "minecraft:black_stained_glass_pane": {
        "protocol_id": 374
    },
    "minecraft:acacia_stairs": {
        "protocol_id": 375
    },
    "minecraft:dark_oak_stairs": {
        "protocol_id": 376
    },
    "minecraft:slime_block": {
        "protocol_id": 377
    },
    "minecraft:barrier": {
        "protocol_id": 378
    },
    "minecraft:iron_trapdoor": {
        "protocol_id": 379
    },
    "minecraft:prismarine": {
        "protocol_id": 380
    },
    "minecraft:prismarine_bricks": {
        "protocol_id": 381
    },
    "minecraft:dark_prismarine": {
        "protocol_id": 382
    },
    "minecraft:prismarine_stairs": {
        "protocol_id": 383
    },
    "minecraft:prismarine_brick_stairs": {
        "protocol_id": 384
    },
    "minecraft:dark_prismarine_stairs": {
        "protocol_id": 385
    },
    "minecraft:prismarine_slab": {
        "protocol_id": 386
    },
    "minecraft:prismarine_brick_slab": {
        "protocol_id": 387
    },
    "minecraft:dark_prismarine_slab": {
        "protocol_id": 388
    },
    "minecraft:sea_lantern": {
        "protocol_id": 389
    },
    "minecraft:hay_block": {
        "protocol_id": 390
    },
    "minecraft:white_carpet": {
        "protocol_id": 391
    },
    "minecraft:orange_carpet": {
        "protocol_id": 392
    },
    "minecraft:magenta_carpet": {
        "protocol_id": 393
    },
    "minecraft:light_blue_carpet": {
        "protocol_id": 394
    },
    "minecraft:yellow_carpet": {
        "protocol_id": 395
    },
    "minecraft:lime_carpet": {
        "protocol_id": 396
    },
    "minecraft:pink_carpet": {
        "protocol_id": 397
    },
    "minecraft:gray_carpet": {
        "protocol_id": 398
    },
    "minecraft:light_gray_carpet": {
        "protocol_id": 399
    },
    "minecraft:cyan_carpet": {
        "protocol_id": 400
    },
    "minecraft:purple_carpet": {
        "protocol_id": 401
    },
    "minecraft:blue_carpet": {
        "protocol_id": 402
    },
    "minecraft:brown_carpet": {
        "protocol_id": 403
    },
    "minecraft:green_carpet": {
        "protocol_id": 404
    },
    "minecraft:red_carpet": {
        "protocol_id": 405
    },
    "minecraft:black_carpet": {
        "protocol_id": 406
    },
    "minecraft:terracotta": {
        "protocol_id": 407
    },
    "minecraft:coal_block": {
        "protocol_id": 408
    },
    "minecraft:packed_ice": {
        "protocol_id": 409
    },
    "minecraft:sunflower": {
        "protocol_id": 410
    },
    "minecraft:lilac": {
        "protocol_id": 411
    },
    "minecraft:rose_bush": {
        "protocol_id": 412
    },
    "minecraft:peony": {
        "protocol_id": 413
    },
    "minecraft:tall_grass": {
        "protocol_id": 414
    },
    "minecraft:large_fern": {
        "protocol_id": 415
    },
    "minecraft:white_banner": {
        "protocol_id": 416
    },
    "minecraft:orange_banner": {
        "protocol_id": 417
    },
    "minecraft:magenta_banner": {
        "protocol_id": 418
    },
    "minecraft:light_blue_banner": {
        "protocol_id": 419
    },
    "minecraft:yellow_banner": {
        "protocol_id": 420
    },
    "minecraft:lime_banner": {
        "protocol_id": 421
    },
    "minecraft:pink_banner": {
        "protocol_id": 422
    },
    "minecraft:gray_banner": {
        "protocol_id": 423
    },
    "minecraft:light_gray_banner": {
        "protocol_id": 424
    },
    "minecraft:cyan_banner": {
        "protocol_id": 425
    },
    "minecraft:purple_banner": {
        "protocol_id": 426
    },
    "minecraft:blue_banner": {
        "protocol_id": 427
    },
    "minecraft:brown_banner": {
        "protocol_id": 428
    },
    "minecraft:green_banner": {
        "protocol_id": 429
    },
    "minecraft:red_banner": {
        "protocol_id": 430
    },
    "minecraft:black_banner": {
        "protocol_id": 431
    },
    "minecraft:white_wall_banner": {
        "protocol_id": 432
    },
    "minecraft:orange_wall_banner": {
        "protocol_id": 433
    },
    "minecraft:magenta_wall_banner": {
        "protocol_id": 434
    },
    "minecraft:light_blue_wall_banner": {
        "protocol_id": 435
    },
    "minecraft:yellow_wall_banner": {
        "protocol_id": 436
    },
    "minecraft:lime_wall_banner": {
        "protocol_id": 437
    },
    "minecraft:pink_wall_banner": {
        "protocol_id": 438
    },
    "minecraft:gray_wall_banner": {
        "protocol_id": 439
    },
    "minecraft:light_gray_wall_banner": {
        "protocol_id": 440
    },
    "minecraft:cyan_wall_banner": {
        "protocol_id": 441
    },
    "minecraft:purple_wall_banner": {
        "protocol_id": 442
    },
    "minecraft:blue_wall_banner": {
        "protocol_id": 443
    },
    "minecraft:brown_wall_banner": {
        "protocol_id": 444
    },
    "minecraft:green_wall_banner": {
        "protocol_id": 445
    },
    "minecraft:red_wall_banner": {
        "protocol_id": 446
    },
    "minecraft:black_wall_banner": {
        "protocol_id": 447
    },
    "minecraft:red_sandstone": {
        "protocol_id": 448
    },
    "minecraft:chiseled_red_sandstone": {
        "protocol_id": 449
    },
    "minecraft:cut_red_sandstone": {
        "protocol_id": 450
    },
    "minecraft:red_sandstone_stairs": {
        "protocol_id": 451
    },
    "minecraft:oak_slab": {
        "protocol_id": 452
    },
    "minecraft:spruce_slab": {
        "protocol_id": 453
    },
    "minecraft:birch_slab": {
        "protocol_id": 454
    },
    "minecraft:jungle_slab": {
        "protocol_id": 455
    },
    "minecraft:acacia_slab": {
        "protocol_id": 456
    },
    "minecraft:dark_oak_slab": {
        "protocol_id": 457
    },
    "minecraft:stone_slab": {
        "protocol_id": 458
    },
    "minecraft:smooth_stone_slab": {
        "protocol_id": 459
    },
    "minecraft:sandstone_slab": {
        "protocol_id": 460
    },
    "minecraft:cut_sandstone_slab": {
        "protocol_id": 461
    },
    "minecraft:petrified_oak_slab": {
        "protocol_id": 462
    },
    "minecraft:cobblestone_slab": {
        "protocol_id": 463
    },
    "minecraft:brick_slab": {
        "protocol_id": 464
    },
    "minecraft:stone_brick_slab": {
        "protocol_id": 465
    },
    "minecraft:nether_brick_slab": {
        "protocol_id": 466
    },
    "minecraft:quartz_slab": {
        "protocol_id": 467
    },
    "minecraft:red_sandstone_slab": {
        "protocol_id": 468
    },
    "minecraft:cut_red_sandstone_slab": {
        "protocol_id": 469
    },
    "minecraft:purpur_slab": {
        "protocol_id": 470
    },
    "minecraft:smooth_stone": {
        "protocol_id": 471
    },
    "minecraft:smooth_sandstone": {
        "protocol_id": 472
    },
    "minecraft:smooth_quartz": {
        "protocol_id": 473
    },
    "minecraft:smooth_red_sandstone": {
        "protocol_id": 474
    },
    "minecraft:spruce_fence_gate": {
        "protocol_id": 475
    },
    "minecraft:birch_fence_gate": {
        "protocol_id": 476
    },
    "minecraft:jungle_fence_gate": {
        "protocol_id": 477
    },
    "minecraft:acacia_fence_gate": {
        "protocol_id": 478
    },
    "minecraft:dark_oak_fence_gate": {
        "protocol_id": 479
    },
    "minecraft:spruce_fence": {
        "protocol_id": 480
    },
    "minecraft:birch_fence": {
        "protocol_id": 481
    },
    "minecraft:jungle_fence": {
        "protocol_id": 482
    },
    "minecraft:acacia_fence": {
        "protocol_id": 483
    },
    "minecraft:dark_oak_fence": {
        "protocol_id": 484
    },
    "minecraft:spruce_door": {
        "protocol_id": 485
    },
    "minecraft:birch_door": {
        "protocol_id": 486
    },
    "minecraft:jungle_door": {
        "protocol_id": 487
    },
    "minecraft:acacia_door": {
        "protocol_id": 488
    },
    "minecraft:dark_oak_door": {
        "protocol_id": 489
    },
    "minecraft:end_rod": {
        "protocol_id": 490
    },
    "minecraft:chorus_plant": {
        "protocol_id": 491
    },
    "minecraft:chorus_flower": {
        "protocol_id": 492
    },
    "minecraft:purpur_block": {
        "protocol_id": 493
    },
    "minecraft:purpur_pillar": {
        "protocol_id": 494
    },
    "minecraft:purpur_stairs": {
        "protocol_id": 495
    },
    "minecraft:end_stone_bricks": {
        "protocol_id": 496
    },
    "minecraft:beetroots": {
        "protocol_id": 497
    },
    "minecraft:grass_path": {
        "protocol_id": 498
    },
    "minecraft:end_gateway": {
        "protocol_id": 499
    },
    "minecraft:repeating_command_block": {
        "protocol_id": 500
    },
    "minecraft:chain_command_block": {
        "protocol_id": 501
    },
    "minecraft:frosted_ice": {
        "protocol_id": 502
    },
    "minecraft:magma_block": {
        "protocol_id": 503
    },
    "minecraft:nether_wart_block": {
        "protocol_id": 504
    },
    "minecraft:red_nether_bricks": {
        "protocol_id": 505
    },
    "minecraft:bone_block": {
        "protocol_id": 506
    },
    "minecraft:structure_void": {
        "protocol_id": 507
    },
    "minecraft:observer": {
        "protocol_id": 508
    },
    "minecraft:shulker_box": {
        "protocol_id": 509
    },
    "minecraft:white_shulker_box": {
        "protocol_id": 510
    },
    "minecraft:orange_shulker_box": {
        "protocol_id": 511
    },
    "minecraft:magenta_shulker_box": {
        "protocol_id": 512
    },
    "minecraft:light_blue_shulker_box": {
        "protocol_id": 513
    },
    "minecraft:yellow_shulker_box": {
        "protocol_id": 514
    },
    "minecraft:lime_shulker_box": {
        "protocol_id": 515
    },
    "minecraft:pink_shulker_box": {
        "protocol_id": 516
    },
    "minecraft:gray_shulker_box": {
        "protocol_id": 517
    },
    "minecraft:light_gray_shulker_box": {
        "protocol_id": 518
    },
    "minecraft:cyan_shulker_box": {
        "protocol_id": 519
    },
    "minecraft:purple_shulker_box": {
        "protocol_id": 520
    },
    "minecraft:blue_shulker_box": {
        "protocol_id": 521
    },
    "minecraft:brown_shulker_box": {
        "protocol_id": 522
    },
    "minecraft:green_shulker_box": {
        "protocol_id": 523
    },
    "minecraft:red_shulker_box": {
        "protocol_id": 524
    },
    "minecraft:black_shulker_box": {
        "protocol_id": 525
    },
    "minecraft:white_glazed_terracotta": {
        "protocol_id": 526
    },
    "minecraft:orange_glazed_terracotta": {
        "protocol_id": 527
    },
    "minecraft:magenta_glazed_terracotta": {
        "protocol_id": 528
    },
    "minecraft:light_blue_glazed_terracotta": {
        "protocol_id": 529
    },
    "minecraft:yellow_glazed_terracotta": {
        "protocol_id": 530
    },
    "minecraft:lime_glazed_terracotta": {
        "protocol_id": 531
    },
    "minecraft:pink_glazed_terracotta": {
        "protocol_id": 532
    },
    "minecraft:gray_glazed_terracotta": {
        "protocol_id": 533
    },
    "minecraft:light_gray_glazed_terracotta": {
        "protocol_id": 534
    },
    "minecraft:cyan_glazed_terracotta": {
        "protocol_id": 535
    },
    "minecraft:purple_glazed_terracotta": {
        "protocol_id": 536
    },
    "minecraft:blue_glazed_terracotta": {
        "protocol_id": 537
    },
    "minecraft:brown_glazed_terracotta": {
        "protocol_id": 538
    },
    "minecraft:green_glazed_terracotta": {
        "protocol_id": 539
    },
    "minecraft:red_glazed_terracotta": {
        "protocol_id": 540
    },
    "minecraft:black_glazed_terracotta": {
        "protocol_id": 541
    },
    "minecraft:white_concrete": {
        "protocol_id": 542
    },
    "minecraft:orange_concrete": {
        "protocol_id": 543
    },
    "minecraft:magenta_concrete": {
        "protocol_id": 544
    },
    "minecraft:light_blue_concrete": {
        "protocol_id": 545
    },
    "minecraft:yellow_concrete": {
        "protocol_id": 546
    },
    "minecraft:lime_concrete": {
        "protocol_id": 547
    },
    "minecraft:pink_concrete": {
        "protocol_id": 548
    },
    "minecraft:gray_concrete": {
        "protocol_id": 549
    },
    "minecraft:light_gray_concrete": {
        "protocol_id": 550
    },
    "minecraft:cyan_concrete": {
        "protocol_id": 551
    },
    "minecraft:purple_concrete": {
        "protocol_id": 552
    },
    "minecraft:blue_concrete": {
        "protocol_id": 553
    },
    "minecraft:brown_concrete": {
        "protocol_id": 554
    },
    "minecraft:green_concrete": {
        "protocol_id": 555
    },
    "minecraft:red_concrete": {
        "protocol_id": 556
    },
    "minecraft:black_concrete": {
        "protocol_id": 557
    },
    "minecraft:white_concrete_powder": {
        "protocol_id": 558
    },
    "minecraft:orange_concrete_powder": {
        "protocol_id": 559
    },
    "minecraft:magenta_concrete_powder": {
        "protocol_id": 560
    },
    "minecraft:light_blue_concrete_powder": {
        "protocol_id": 561
    },
    "minecraft:yellow_concrete_powder": {
        "protocol_id": 562
    },
    "minecraft:lime_concrete_powder": {
        "protocol_id": 563
    },
    "minecraft:pink_concrete_powder": {
        "protocol_id": 564
    },
    "minecraft:gray_concrete_powder": {
        "protocol_id": 565
    },
    "minecraft:light_gray_concrete_powder": {
        "protocol_id": 566
    },
    "minecraft:cyan_concrete_powder": {
        "protocol_id": 567
    },
    "minecraft:purple_concrete_powder": {
        "protocol_id": 568
    },
    "minecraft:blue_concrete_powder": {
        "protocol_id": 569
    },
    "minecraft:brown_concrete_powder": {
        "protocol_id": 570
    },
    "minecraft:green_concrete_powder": {
        "protocol_id": 571
    },
    "minecraft:red_concrete_powder": {
        "protocol_id": 572
    },
    "minecraft:black_concrete_powder": {
        "protocol_id": 573
    },
    "minecraft:kelp": {
        "protocol_id": 574
    },
    "minecraft:kelp_plant": {
        "protocol_id": 575
    },
    "minecraft:dried_kelp_block": {
        "protocol_id": 576
    },
    "minecraft:turtle_egg": {
        "protocol_id": 577
    },
    "minecraft:dead_tube_coral_block": {
        "protocol_id": 578
    },
    "minecraft:dead_brain_coral_block": {
        "protocol_id": 579
    },
    "minecraft:dead_bubble_coral_block": {
        "protocol_id": 580
    },
    "minecraft:dead_fire_coral_block": {
        "protocol_id": 581
    },
    "minecraft:dead_horn_coral_block": {
        "protocol_id": 582
    },
    "minecraft:tube_coral_block": {
        "protocol_id": 583
    },
    "minecraft:brain_coral_block": {
        "protocol_id": 584
    },
    "minecraft:bubble_coral_block": {
        "protocol_id": 585
    },
    "minecraft:fire_coral_block": {
        "protocol_id": 586
    },
    "minecraft:horn_coral_block": {
        "protocol_id": 587
    },
    "minecraft:dead_tube_coral": {
        "protocol_id": 588
    },
    "minecraft:dead_brain_coral": {
        "protocol_id": 589
    },
    "minecraft:dead_bubble_coral": {
        "protocol_id": 590
    },
    "minecraft:dead_fire_coral": {
        "protocol_id": 591
    },
    "minecraft:dead_horn_coral": {
        "protocol_id": 592
    },
    "minecraft:tube_coral": {
        "protocol_id": 593
    },
    "minecraft:brain_coral": {
        "protocol_id": 594
    },
    "minecraft:bubble_coral": {
        "protocol_id": 595
    },
    "minecraft:fire_coral": {
        "protocol_id": 596
    },
    "minecraft:horn_coral": {
        "protocol_id": 597
    },
    "minecraft:dead_tube_coral_fan": {
        "protocol_id": 598
    },
    "minecraft:dead_brain_coral_fan": {
        "protocol_id": 599
    },
    "minecraft:dead_bubble_coral_fan": {
        "protocol_id": 600
    },
    "minecraft:dead_fire_coral_fan": {
        "protocol_id": 601
    },
    "minecraft:dead_horn_coral_fan": {
        "protocol_id": 602
    },
    "minecraft:tube_coral_fan": {
        "protocol_id": 603
    },
    "minecraft:brain_coral_fan": {
        "protocol_id": 604
    },
    "minecraft:bubble_coral_fan": {
        "protocol_id": 605
    },
    "minecraft:fire_coral_fan": {
        "protocol_id": 606
    },
    "minecraft:horn_coral_fan": {
        "protocol_id": 607
    },
    "minecraft:dead_tube_coral_wall_fan": {
        "protocol_id": 608
    },
    "minecraft:dead_brain_coral_wall_fan": {
        "protocol_id": 609
    },
    "minecraft:dead_bubble_coral_wall_fan": {
        "protocol_id": 610
    },
    "minecraft:dead_fire_coral_wall_fan": {
        "protocol_id": 611
    },
    "minecraft:dead_horn_coral_wall_fan": {
        "protocol_id": 612
    },
    "minecraft:tube_coral_wall_fan": {
        "protocol_id": 613
    },
    "minecraft:brain_coral_wall_fan": {
        "protocol_id": 614
    },
    "minecraft:bubble_coral_wall_fan": {
        "protocol_id": 615
    },
    "minecraft:fire_coral_wall_fan": {
        "protocol_id": 616
    },
    "minecraft:horn_coral_wall_fan": {
        "protocol_id": 617
    },
    "minecraft:sea_pickle": {
        "protocol_id": 618
    },
    "minecraft:blue_ice": {
        "protocol_id": 619
    },
    "minecraft:conduit": {
        "protocol_id": 620
    },
    "minecraft:bamboo_sapling": {
        "protocol_id": 621
    },
    "minecraft:bamboo": {
        "protocol_id": 622
    },
    "minecraft:potted_bamboo": {
        "protocol_id": 623
    },
    "minecraft:void_air": {
        "protocol_id": 624
    },
    "minecraft:cave_air": {
        "protocol_id": 625
    },
    "minecraft:bubble_column": {
        "protocol_id": 626
    },
    "minecraft:polished_granite_stairs": {
        "protocol_id": 627
    },
    "minecraft:smooth_red_sandstone_stairs": {
        "protocol_id": 628
    },
    "minecraft:mossy_stone_brick_stairs": {
        "protocol_id": 629
    },
    "minecraft:polished_diorite_stairs": {
        "protocol_id": 630
    },
    "minecraft:mossy_cobblestone_stairs": {
        "protocol_id": 631
    },
    "minecraft:end_stone_brick_stairs": {
        "protocol_id": 632
    },
    "minecraft:stone_stairs": {
        "protocol_id": 633
    },
    "minecraft:smooth_sandstone_stairs": {
        "protocol_id": 634
    },
    "minecraft:smooth_quartz_stairs": {
        "protocol_id": 635
    },
    "minecraft:granite_stairs": {
        "protocol_id": 636
    },
    "minecraft:andesite_stairs": {
        "protocol_id": 637
    },
    "minecraft:red_nether_brick_stairs": {
        "protocol_id": 638
    },
    "minecraft:polished_andesite_stairs": {
        "protocol_id": 639
    },
    "minecraft:diorite_stairs": {
        "protocol_id": 640
    },
    "minecraft:polished_granite_slab": {
        "protocol_id": 641
    },
    "minecraft:smooth_red_sandstone_slab": {
        "protocol_id": 642
    },
    "minecraft:mossy_stone_brick_slab": {
        "protocol_id": 643
    },
    "minecraft:polished_diorite_slab": {
        "protocol_id": 644
    },
    "minecraft:mossy_cobblestone_slab": {
        "protocol_id": 645
    },
    "minecraft:end_stone_brick_slab": {
        "protocol_id": 646
    },
    "minecraft:smooth_sandstone_slab": {
        "protocol_id": 647
    },
    "minecraft:smooth_quartz_slab": {
        "protocol_id": 648
    },
    "minecraft:granite_slab": {
        "protocol_id": 649
    },
    "minecraft:andesite_slab": {
        "protocol_id": 650
    },
    "minecraft:red_nether_brick_slab": {
        "protocol_id": 651
    },
    "minecraft:polished_andesite_slab": {
        "protocol_id": 652
    },
    "minecraft:diorite_slab": {
        "protocol_id": 653
    },
    "minecraft:brick_wall": {
        "protocol_id": 654
    },
    "minecraft:prismarine_wall": {
        "protocol_id": 655
    },
    "minecraft:red_sandstone_wall": {
        "protocol_id": 656
    },
    "minecraft:mossy_stone_brick_wall": {
        "protocol_id": 657
    },
    "minecraft:granite_wall": {
        "protocol_id": 658
    },
    "minecraft:stone_brick_wall": {
        "protocol_id": 659
    },
    "minecraft:nether_brick_wall": {
        "protocol_id": 660
    },
    "minecraft:andesite_wall": {
        "protocol_id": 661
    },
    "minecraft:red_nether_brick_wall": {
        "protocol_id": 662
    },
    "minecraft:sandstone_wall": {
        "protocol_id": 663
    },
    "minecraft:end_stone_brick_wall": {
        "protocol_id": 664
    },
    "minecraft:diorite_wall": {
        "protocol_id": 665
    },
    "minecraft:scaffolding": {
        "protocol_id": 666
    },
    "minecraft:loom": {
        "protocol_id": 667
    },
    "minecraft:barrel": {
        "protocol_id": 668
    },
    "minecraft:smoker": {
        "protocol_id": 669
    },
    "minecraft:blast_furnace": {
        "protocol_id": 670
    },
    "minecraft:cartography_table": {
        "protocol_id": 671
    },
    "minecraft:fletching_table": {
        "protocol_id": 672
    },
    "minecraft:grindstone": {
        "protocol_id": 673
    },
    "minecraft:lectern": {
        "protocol_id": 674
    },
    "minecraft:smithing_table": {
        "protocol_id": 675
    },
    "minecraft:stonecutter": {
        "protocol_id": 676
    },
    "minecraft:bell": {
        "protocol_id": 677
    },
    "minecraft:lantern": {
        "protocol_id": 678
    },
    "minecraft:soul_lantern": {
        "protocol_id": 679
    },
    "minecraft:campfire": {
        "protocol_id": 680
    },
    "minecraft:soul_campfire": {
        "protocol_id": 681
    },
    "minecraft:sweet_berry_bush": {
        "protocol_id": 682
    },
    "minecraft:warped_stem": {
        "protocol_id": 683
    },
    "minecraft:stripped_warped_stem": {
        "protocol_id": 684
    },
    "minecraft:warped_hyphae": {
        "protocol_id": 685
    },
    "minecraft:stripped_warped_hyphae": {
        "protocol_id": 686
    },
    "minecraft:warped_nylium": {
        "protocol_id": 687
    },
    "minecraft:warped_fungus": {
        "protocol_id": 688
    },
    "minecraft:warped_wart_block": {
        "protocol_id": 689
    },
    "minecraft:warped_roots": {
        "protocol_id": 690
    },
    "minecraft:nether_sprouts": {
        "protocol_id": 691
    },
    "minecraft:crimson_stem": {
        "protocol_id": 692
    },
    "minecraft:stripped_crimson_stem": {
        "protocol_id": 693
    },
    "minecraft:crimson_hyphae": {
        "protocol_id": 694
    },
    "minecraft:stripped_crimson_hyphae": {
        "protocol_id": 695
    },
    "minecraft:crimson_nylium": {
        "protocol_id": 696
    },
    "minecraft:crimson_fungus": {
        "protocol_id": 697
    },
    "minecraft:shroomlight": {
        "protocol_id": 698
    },
    "minecraft:weeping_vines": {
        "protocol_id": 699
    },
    "minecraft:weeping_vines_plant": {
        "protocol_id": 700
    },
    "minecraft:twisting_vines": {
        "protocol_id": 701
    },
    "minecraft:twisting_vines_plant": {
        "protocol_id": 702
    },
    "minecraft:crimson_roots": {
        "protocol_id": 703
    },
    "minecraft:crimson_planks": {
        "protocol_id": 704
    },
    "minecraft:warped_planks": {
        "protocol_id": 705
    },
    "minecraft:crimson_slab": {
        "protocol_id": 706
    },
    "minecraft:warped_slab": {
        "protocol_id": 707
    },
    "minecraft:crimson_pressure_plate": {
        "protocol_id": 708
    },
    "minecraft:warped_pressure_plate": {
        "protocol_id": 709
    },
    "minecraft:crimson_fence": {
        "protocol_id": 710
    },
    "minecraft:warped_fence": {
        "protocol_id": 711
    },
    "minecraft:crimson_trapdoor": {
        "protocol_id": 712
    },
    "minecraft:warped_trapdoor": {
        "protocol_id": 713
    },
    "minecraft:crimson_fence_gate": {
        "protocol_id": 714
    },
    "minecraft:warped_fence_gate": {
        "protocol_id": 715
    },
    "minecraft:crimson_stairs": {
        "protocol_id": 716
    },
    "minecraft:warped_stairs": {
        "protocol_id": 717
    },
    "minecraft:crimson_button": {
        "protocol_id": 718
    },
    "minecraft:warped_button": {
        "protocol_id": 719
    },
    "minecraft:crimson_door": {
        "protocol_id": 720
    },
    "minecraft:warped_door": {
        "protocol_id": 721
    },
    "minecraft:crimson_sign": {
        "protocol_id": 722
    },
    "minecraft:warped_sign": {
        "protocol_id": 723
    },
    "minecraft:crimson_wall_sign": {
        "protocol_id": 724
    },
    "minecraft:warped_wall_sign": {
        "protocol_id": 725
    },
    "minecraft:structure_block": {
        "protocol_id": 726
    },
    "minecraft:jigsaw": {
        "protocol_id": 727
    },
    "minecraft:composter": {
        "protocol_id": 728
    },
    "minecraft:target": {
        "protocol_id": 729
    },
    "minecraft:bee_nest": {
        "protocol_id": 730
    },
    "minecraft:beehive": {
        "protocol_id": 731
    },
    "minecraft:honey_block": {
        "protocol_id": 732
    },
    "minecraft:honeycomb_block": {
        "protocol_id": 733
    },
    "minecraft:netherite_block": {
        "protocol_id": 734
    },
    "minecraft:ancient_debris": {
        "protocol_id": 735
    },
    "minecraft:crying_obsidian": {
        "protocol_id": 736
    },
    "minecraft:respawn_anchor": {
        "protocol_id": 737
    },
    "minecraft:potted_crimson_fungus": {
        "protocol_id": 738
    },
    "minecraft:potted_warped_fungus": {
        "protocol_id": 739
    },
    "minecraft:potted_crimson_roots": {
        "protocol_id": 740
    },
    "minecraft:potted_warped_roots": {
        "protocol_id": 741
    },
    "minecraft:lodestone": {
        "protocol_id": 742
    },
    "minecraft:blackstone": {
        "protocol_id": 743
    },
    "minecraft:blackstone_stairs": {
        "protocol_id": 744
    },
    "minecraft:blackstone_wall": {
        "protocol_id": 745
    },
    "minecraft:blackstone_slab": {
        "protocol_id": 746
    },
    "minecraft:polished_blackstone": {
        "protocol_id": 747
    },
    "minecraft:polished_blackstone_bricks": {
        "protocol_id": 748
    },
    "minecraft:cracked_polished_blackstone_bricks": {
        "protocol_id": 749
    },
    "minecraft:chiseled_polished_blackstone": {
        "protocol_id": 750
    },
    "minecraft:polished_blackstone_brick_slab": {
        "protocol_id": 751
    },
    "minecraft:polished_blackstone_brick_stairs": {
        "protocol_id": 752
    },
    "minecraft:polished_blackstone_brick_wall": {
        "protocol_id": 753
    },
    "minecraft:gilded_blackstone": {
        "protocol_id": 754
    },
    "minecraft:polished_blackstone_stairs": {
        "protocol_id": 755
    },
    "minecraft:polished_blackstone_slab": {
        "protocol_id": 756
    },
    "minecraft:polished_blackstone_pressure_plate": {
        "protocol_id": 757
    },
    "minecraft:polished_blackstone_button": {
        "protocol_id": 758
    },
    "minecraft:polished_blackstone_wall": {
        "protocol_id": 759
    },
    "minecraft:chiseled_nether_bricks": {
        "protocol_id": 760
    },
    "minecraft:cracked_nether_bricks": {
        "protocol_id": 761
    },
    "minecraft:quartz_bricks": {
        "protocol_id": 762
    }
}`
