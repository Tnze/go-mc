package states

var (
	HorizontalAxisPropertyProperty = NewPropertyEnum[Axis]("axis", map[string]Axis{
		"x": X,
		"z": Z,
	})
	AxisProperty = NewPropertyEnum[Axis]("axis", map[string]Axis{
		"x": X,
		"y": Y,
		"z": Z,
	})
	FacingProperty = NewPropertyEnum[Direction]("facing", map[string]Direction{
		"north": DirectionNorth,
		"east":  DirectionEast,
		"south": DirectionSouth,
		"west":  DirectionWest,
	})
	HorizontalFacingProperty = NewPropertyEnum[Direction]("facing", map[string]Direction{
		"north": DirectionNorth,
		"east":  DirectionEast,
		"south": DirectionSouth,
		"west":  DirectionWest,
	})
	OrientationProperty = NewPropertyEnum[FrontAndTop]("orientation", map[string]FrontAndTop{
		"down_east":  DownEast,
		"down_north": DownNorth,
		"down_south": DownSouth,
		"down_west":  DownWest,
		"up_east":    UpEast,
		"up_north":   UpNorth,
		"up_south":   UpSouth,
		"up_west":    UpWest,
		"west_up":    WestUp,
		"east_up":    EastUp,
		"north_up":   NorthUp,
		"south_up":   SouthUp,
	})
	AttachFaceProperty = NewPropertyEnum[AttachFace]("face", map[string]AttachFace{
		"floor":   AttachFaceFloor,
		"wall":    AttachFaceWall,
		"ceiling": AttachFaceCeiling,
	})
	BellAttachmentProperty = NewPropertyEnum[BellAttachType]("attachment", map[string]BellAttachType{
		"floor":       BellAttachTypeFloor,
		"ceiling":     BellAttachTypeCeiling,
		"single_wall": BellAttachTypeSingleWall,
		"double_wall": BellAttachTypeDoubleWall,
	})
	EastWallProperty = NewPropertyEnum[WallSide]("east", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	NorthWallProperty = NewPropertyEnum[WallSide]("north", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	SouthWallProperty = NewPropertyEnum[WallSide]("south", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	WestWallProperty = NewPropertyEnum[WallSide]("west", map[string]WallSide{
		"none": WallSideNone,
		"low":  WallSideLow,
		"tall": WallSideTall,
	})
	EastRedstoneProperty = NewPropertyEnum[RedstoneSide]("east", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	NorthRedstoneProperty = NewPropertyEnum[RedstoneSide]("north", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	SouthRedstoneProperty = NewPropertyEnum[RedstoneSide]("south", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	WestRedstoneProperty = NewPropertyEnum[RedstoneSide]("west", map[string]RedstoneSide{
		"none": RedstoneSideNone,
		"side": RedstoneSideSide,
		"up":   RedstoneSideUp,
	})
	DoubleBlockHalfProperty = NewPropertyEnum[DoubleBlockHalf]("half", map[string]DoubleBlockHalf{
		"lower": DoubleBlockHalfLower,
		"upper": DoubleBlockHalfUpper,
	})
	HalfProperty = NewPropertyEnum[Half]("half", map[string]Half{
		"top":    HalfTop,
		"bottom": HalfBottom,
	})
	RailShapeProperty = NewPropertyEnum[RailShape]("shape", map[string]RailShape{
		"north_south":     RailShapeNorthSouth,
		"east_west":       RailShapeEastWest,
		"ascending_east":  RailShapeAscendingEast,
		"ascending_west":  RailShapeAscendingWest,
		"ascending_north": RailShapeAscendingNorth,
		"ascending_south": RailShapeAscendingSouth,
		"south_east":      RailShapeSouthEast,
		"south_west":      RailShapeSouthWest,
		"north_west":      RailShapeNorthWest,
		"north_east":      RailShapeNorthEast,
	})
	RailShapeStraightProperty = NewPropertyEnum[RailShape]("shape", map[string]RailShape{
		"north_south":     RailShapeNorthSouth,
		"east_west":       RailShapeEastWest,
		"ascending_east":  RailShapeAscendingEast,
		"ascending_west":  RailShapeAscendingWest,
		"ascending_north": RailShapeAscendingNorth,
		"ascending_south": RailShapeAscendingSouth,
	})
	/*LevelCauldronProperty        = block.NewPropertyInteger("level", 1, 3)
	LevelComposterProperty       = block.NewPropertyInteger("level", 0, 8)
	LevelFlowingProperty         = block.NewPropertyInteger("level", 1, 8)
	LevelHoneyProperty           = block.NewPropertyInteger("level", 0, 5)
	RedstoneSignalProperty       = block.NewPropertyInteger("signal", 0, 15)
	StagesProperty               = block.NewPropertyInteger("stage", 0, 1)
	StabilityProperty            = block.NewPropertyInteger("distance", 0, 2)
	RespawnAnchorChargesProperty = block.NewPropertyInteger("charges", 0, 4)
	Rotation16Property           = block.NewPropertyInteger("rotation", 0, 15)*/
	BedPartProperty = NewPropertyEnum[BedPart]("part", map[string]BedPart{
		"head": BedPartHead,
		"foot": BedPartFoot,
	})
	ChestTypeProperty = NewPropertyEnum[ChestType]("type", map[string]ChestType{
		"single": ChestTypeSingle,
		"left":   ChestTypeLeft,
		"right":  ChestTypeRight,
	})
	ComparatorModeProperty = NewPropertyEnum[ComparatorMode]("mode", map[string]ComparatorMode{
		"compare":  ComparatorModeCompare,
		"subtract": ComparatorModeSubtract,
	})
	DoorHingeProperty = NewPropertyEnum[DoorHingeSide]("hinge", map[string]DoorHingeSide{
		"left":  DoorHingeSideLeft,
		"right": DoorHingeSideRight,
	})
	InstrumentProperty = NewPropertyEnum[NoteBlockInstrument]("instrument", map[string]NoteBlockInstrument{
		"harp":           NoteBlockInstrumentHarp,
		"basedrum":       NoteBlockInstrumentBasedrum,
		"snare":          NoteBlockInstrumentSnare,
		"hat":            NoteBlockInstrumentHat,
		"bass":           NoteBlockInstrumentBass,
		"flute":          NoteBlockInstrumentFlute,
		"bell":           NoteBlockInstrumentBell,
		"guitar":         NoteBlockInstrumentGuitar,
		"chime":          NoteBlockInstrumentChime,
		"xylophone":      NoteBlockInstrumentXylophone,
		"iron_xylophone": NoteBlockInstrumentIronXylophone,
		"cow_bell":       NoteBlockInstrumentCowBell,
		"didgeridoo":     NoteBlockInstrumentDidgeridoo,
		"bit":            NoteBlockInstrumentBit,
		"banjo":          NoteBlockInstrumentBanjo,
		"pling":          NoteBlockInstrumentPling,
	})
	PistonTypeProperty = NewPropertyEnum[PistonType]("type", map[string]PistonType{
		"normal": PistonTypeDefault,
		"sticky": PistonTypeSticky,
	})
	SlabTypeProperty = NewPropertyEnum[SlabType]("type", map[string]SlabType{
		"bottom": SlabTypeBottom,
		"top":    SlabTypeTop,
		"double": SlabTypeDouble,
	})
	StairsShapeProperty = NewPropertyEnum[StairsShape]("shape", map[string]StairsShape{
		"straight":    StairsShapeStraight,
		"inner_left":  StairsShapeInnerLeft,
		"inner_right": StairsShapeInnerRight,
		"outer_left":  StairsShapeOuterLeft,
		"outer_right": StairsShapeOuterRight,
	})
	StructureModeProperty = NewPropertyEnum[StructureMode]("mode", map[string]StructureMode{
		"save":   StructureModeSave,
		"load":   StructureModeLoad,
		"corner": StructureModeCorner,
		"data":   StructureModeData,
	})
	BambooLeavesProperty = NewPropertyEnum[BambooLeaves]("leaves", map[string]BambooLeaves{
		"none":  BambooLeavesNone,
		"small": BambooLeavesSmall,
		"large": BambooLeavesLarge,
	})
	TiltProperty = NewPropertyEnum[Tilt]("tilt", map[string]Tilt{
		"none":     TiltNone,
		"unstable": TiltUnstable,
		"partial":  TiltPartial,
		"full":     TiltFull,
	})
	VerticalDirectionProperty = NewPropertyEnum[Direction]("direction", map[string]Direction{
		"up":   DirectionUp,
		"down": DirectionDown,
	})
	DripstoneThicknessProperty = NewPropertyEnum[DripstoneThickness]("thickness", map[string]DripstoneThickness{
		"tip_merge": DripstoneThicknessTipMerge,
		"tip":       DripstoneThicknessTip,
		"frustum":   DripstoneThicknessFrustum,
		"middle":    DripstoneThicknessMiddle,
		"base":      DripstoneThicknessBase,
	})
	SculkSensorPhaseProperty = NewPropertyEnum[SculkSensorPhase]("phase", map[string]SculkSensorPhase{
		"inactive": SculkSensorPhaseInactive,
		"active":   SculkSensorPhaseActive,
		"cooldown": SculkSensorPhaseCooldown,
	})
)
