package block

var (
	Attached               = NewPropertyBoolean("attached")
	BottomProperty         = NewPropertyBoolean("bottom")
	Conditional            = NewPropertyBoolean("conditional")
	Disarmed               = NewPropertyBoolean("disarmed")
	Drag                   = NewPropertyBoolean("drag")
	Enabled                = NewPropertyBoolean("enabled")
	Extended               = NewPropertyBoolean("extended")
	Eye                    = NewPropertyBoolean("eye")
	Falling                = NewPropertyBoolean("falling")
	Hanging                = NewPropertyBoolean("hanging")
	HasBottle0             = NewPropertyBoolean("has_bottle_0")
	HasBottle1             = NewPropertyBoolean("has_bottle_1")
	HasBottle2             = NewPropertyBoolean("has_bottle_2")
	HasRecord              = NewPropertyBoolean("has_record")
	HasBook                = NewPropertyBoolean("has_book")
	Inverted               = NewPropertyBoolean("inverted")
	InWall                 = NewPropertyBoolean("in_wall")
	Lit                    = NewPropertyBoolean("lit")
	Locked                 = NewPropertyBoolean("locked")
	Occupied               = NewPropertyBoolean("occupied")
	Open                   = NewPropertyBoolean("open")
	Persistent             = NewPropertyBoolean("persistent")
	Powered                = NewPropertyBoolean("powered")
	Short                  = NewPropertyBoolean("short")
	SignalFire             = NewPropertyBoolean("signal_fire")
	Snowy                  = NewPropertyBoolean("snowy")
	Triggered              = NewPropertyBoolean("triggered")
	Unstable               = NewPropertyBoolean("unstable")
	Waterlogged            = NewPropertyBoolean("waterlogged")
	VineEnd                = NewPropertyBoolean("vine_end")
	Berries                = NewPropertyBoolean("berries")
	Bloom                  = NewPropertyBoolean("bloom")
	Shrieking              = NewPropertyBoolean("shrieking")
	CanSummon              = NewPropertyBoolean("can_summon")
	HorizontalAxisProperty = NewPropertyEnum[Axis]("axis", map[string]Axis{
		"x": X,
		"z": Z,
	})
	AxisProperty = NewPropertyEnum[Axis]("axis", map[string]Axis{
		"x": X,
		"y": Y,
		"z": Z,
	})
	UpProperty     = NewPropertyBoolean("up")
	DownProperty   = NewPropertyBoolean("down")
	NorthProperty  = NewPropertyBoolean("north")
	SouthProperty  = NewPropertyBoolean("south")
	WestProperty   = NewPropertyBoolean("west")
	EastProperty   = NewPropertyBoolean("east")
	FacingProperty = NewPropertyEnum[Direction]("facing", map[string]Direction{
		"north": North,
		"east":  East,
		"south": South,
		"west":  West,
	})
	HorizontalFacingProperty = NewPropertyEnum[Direction]("facing", map[string]Direction{
		"north": North,
		"east":  East,
		"south": South,
		"west":  West,
	})
	Orientation = NewPropertyEnum[FrontAndTop]("orientation", map[string]FrontAndTop{
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
		"top":    Top,
		"bottom": Bottom,
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
	Age1Property                 = NewPropertyInteger("age", 0, 1)
	Age2Property                 = NewPropertyInteger("age", 0, 2)
	Age3Property                 = NewPropertyInteger("age", 0, 3)
	Age5Property                 = NewPropertyInteger("age", 0, 5)
	Age7Property                 = NewPropertyInteger("age", 0, 7)
	Age15Property                = NewPropertyInteger("age", 0, 15)
	Age25Property                = NewPropertyInteger("age", 0, 25)
	BitesProperty                = NewPropertyInteger("bites", 0, 6)
	CandlesProperty              = NewPropertyInteger("candles", 1, 4)
	DelayProperty                = NewPropertyInteger("delay", 1, 4)
	DistanceProperty             = NewPropertyInteger("distance", 1, 7)
	EggsProperty                 = NewPropertyInteger("eggs", 1, 4)
	HatchProperty                = NewPropertyInteger("hatch", 0, 2)
	LayersProperty               = NewPropertyInteger("layers", 1, 8)
	LevelCauldronProperty        = NewPropertyInteger("level", 1, 3)
	LevelComposterProperty       = NewPropertyInteger("level", 0, 8)
	LevelFlowingProperty         = NewPropertyInteger("level", 1, 8)
	LevelHoneyProperty           = NewPropertyInteger("level", 0, 5)
	LevelProperty                = NewPropertyInteger("level", 0, 15)
	MoistureProperty             = NewPropertyInteger("moisture", 0, 7)
	NoteProperty                 = NewPropertyInteger("note", 0, 24)
	PicklesProperty              = NewPropertyInteger("pickles", 1, 4)
	RedstoneSignalProperty       = NewPropertyInteger("signal", 0, 15)
	StagesProperty               = NewPropertyInteger("stage", 0, 1)
	StabilityProperty            = NewPropertyInteger("distance", 0, 2)
	RespawnAnchorChargesProperty = NewPropertyInteger("charges", 0, 4)
	Rotation16Property           = NewPropertyInteger("rotation", 0, 15)
	BedPartProperty              = NewPropertyEnum[BedPart]("part", map[string]BedPart{
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
	NoteInstrumentProperty = NewPropertyEnum[NoteBlockInstrument]("instrument", map[string]NoteBlockInstrument{
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
		"normal": PistonTypeNormal,
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
	StructureBlockModeProperty = NewPropertyEnum[StructureMode]("mode", map[string]StructureMode{
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
		"up":   Up,
		"down": Down,
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
