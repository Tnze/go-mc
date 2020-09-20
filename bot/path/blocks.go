package path

import (
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/block"
)

var (
	safeStepBlocks = make(map[world.BlockStatus]struct{}, 1024)
	blocks         = []block.Block{
		block.Stone,
		block.Granite,
		block.PolishedGranite,
		block.Diorite,
		block.PolishedDiorite,
		block.Andesite,
		block.PolishedAndesite,
		block.GrassBlock,
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
	}
)

func init() {
	for _, b := range blocks {
		if b.MinStateID == b.MaxStateID {
			safeStepBlocks[world.BlockStatus(b.MinStateID)] = struct{}{}
		} else {
			for i := b.MinStateID; i <= b.MaxStateID; i++ {
				safeStepBlocks[world.BlockStatus(i)] = struct{}{}
			}
		}
	}
}
