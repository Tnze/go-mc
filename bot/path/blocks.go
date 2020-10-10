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
