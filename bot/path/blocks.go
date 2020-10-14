package path

import (
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/block"
)

var (
	safeStepBlocks = make(map[world.BlockStatus]struct{}, 1024)
	stepBlocks     = []block.Block{
		block.Stone,
		block.Granite,
		block.PolishedGranite,
		block.Diorite,
		block.PolishedDiorite,
		block.Andesite,
		block.PolishedAndesite,
		block.GrassBlock,
		block.GrassPath,
		block.Dirt,
		block.CoarseDirt,
		block.Cobblestone,
		block.OakPlanks,
		block.SprucePlanks,
		block.BirchPlanks,
		block.JunglePlanks,
		block.AcaciaPlanks,
		block.DarkOakPlanks,
		block.Bedrock,
		block.GoldOre,
		block.IronOre,
		block.CoalOre,
		block.Glass,
		block.LapisOre,
		block.Sandstone,
		block.RedstoneOre,
		block.OakStairs,
		block.AcaciaStairs,
		block.DarkOakStairs,
		block.RedSandstoneStairs,
		block.PolishedGraniteStairs,
		block.SmoothRedSandstoneStairs,
		block.MossyStoneBrickStairs,
		block.PolishedDioriteStairs,
		block.MossyCobblestoneStairs,
		block.EndStoneBrickStairs,
		block.StoneStairs,
		block.SmoothSandstoneStairs,
		block.SmoothQuartzStairs,
		block.GraniteStairs,
		block.AndesiteStairs,
		block.RedNetherBrickStairs,
		block.PolishedAndesiteStairs,
		block.OakSlab,
		block.AcaciaSlab,
		block.DarkOakSlab,
		block.RedSandstoneSlab,
		block.PolishedGraniteSlab,
		block.SmoothRedSandstoneSlab,
		block.MossyStoneBrickSlab,
		block.PolishedDioriteSlab,
		block.MossyCobblestoneSlab,
		block.EndStoneBrickSlab,
		block.StoneSlab,
		block.SmoothSandstoneSlab,
		block.SmoothQuartzSlab,
		block.GraniteSlab,
		block.AndesiteSlab,
		block.RedNetherBrickSlab,
		block.PolishedAndesiteSlab,
	}

	slabs = map[block.ID]struct{}{
		block.OakSlab.ID:                struct{}{},
		block.AcaciaSlab.ID:             struct{}{},
		block.DarkOakSlab.ID:            struct{}{},
		block.RedSandstoneSlab.ID:       struct{}{},
		block.PolishedGraniteSlab.ID:    struct{}{},
		block.SmoothRedSandstoneSlab.ID: struct{}{},
		block.MossyStoneBrickSlab.ID:    struct{}{},
		block.PolishedDioriteSlab.ID:    struct{}{},
		block.MossyCobblestoneSlab.ID:   struct{}{},
		block.EndStoneBrickSlab.ID:      struct{}{},
		block.StoneSlab.ID:              struct{}{},
		block.SmoothSandstoneSlab.ID:    struct{}{},
		block.SmoothQuartzSlab.ID:       struct{}{},
		block.GraniteSlab.ID:            struct{}{},
		block.AndesiteSlab.ID:           struct{}{},
		block.RedNetherBrickSlab.ID:     struct{}{},
		block.PolishedAndesiteSlab.ID:   struct{}{},
	}
	stairs = map[block.ID]struct{}{
		block.OakStairs.ID:                struct{}{},
		block.AcaciaStairs.ID:             struct{}{},
		block.DarkOakStairs.ID:            struct{}{},
		block.RedSandstoneStairs.ID:       struct{}{},
		block.PolishedGraniteStairs.ID:    struct{}{},
		block.SmoothRedSandstoneStairs.ID: struct{}{},
		block.MossyStoneBrickStairs.ID:    struct{}{},
		block.PolishedDioriteStairs.ID:    struct{}{},
		block.MossyCobblestoneStairs.ID:   struct{}{},
		block.EndStoneBrickStairs.ID:      struct{}{},
		block.StoneStairs.ID:              struct{}{},
		block.SmoothSandstoneStairs.ID:    struct{}{},
		block.SmoothQuartzStairs.ID:       struct{}{},
		block.GraniteStairs.ID:            struct{}{},
		block.AndesiteStairs.ID:           struct{}{},
		block.RedNetherBrickStairs.ID:     struct{}{},
		block.PolishedAndesiteStairs.ID:   struct{}{},
	}

	safeWalkBlocks = make(map[world.BlockStatus]struct{}, 128)
	walkBlocks     = []block.Block{
		block.Air,
		block.CaveAir,
		block.Grass,
		block.Torch,
		block.OakSign,
		block.SpruceSign,
		block.BirchSign,
		block.AcaciaSign,
		block.JungleSign,
		block.DarkOakSign,
		block.OakWallSign,
		block.SpruceWallSign,
		block.BirchWallSign,
		block.AcaciaWallSign,
		block.JungleWallSign,
		block.DarkOakWallSign,
		block.Cornflower,
		block.TallGrass,
	}

	additionalCostBlocks = map[*block.Block]int{
		&block.Rail:        120,
		&block.PoweredRail: 200,
	}
)

func init() {
	for _, b := range stepBlocks {
		if b.MinStateID == b.MaxStateID {
			safeStepBlocks[world.BlockStatus(b.MinStateID)] = struct{}{}
		} else {
			for i := b.MinStateID; i <= b.MaxStateID; i++ {
				safeStepBlocks[world.BlockStatus(i)] = struct{}{}
			}
		}
	}

	for _, b := range walkBlocks {
		if b.MinStateID == b.MaxStateID {
			safeWalkBlocks[world.BlockStatus(b.MinStateID)] = struct{}{}
		} else {
			for i := b.MinStateID; i <= b.MaxStateID; i++ {
				safeWalkBlocks[world.BlockStatus(i)] = struct{}{}
			}
		}
	}
}

func SteppableBlock(bID world.BlockStatus) bool {
	_, ok := safeStepBlocks[bID]
	return ok
}

func AirLikeBlock(bID world.BlockStatus) bool {
	_, ok := safeWalkBlocks[bID]
	return ok
}

func IsLadder(bID world.BlockStatus) bool {
	return uint32(bID) >= block.Ladder.MinStateID && uint32(bID) <= block.Ladder.MaxStateID
}

func IsSlab(bID world.BlockStatus) bool {
	_, isSlab := slabs[block.StateID[uint32(bID)]]
	return isSlab
}
