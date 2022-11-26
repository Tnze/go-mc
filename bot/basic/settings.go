package basic

// Settings of client
type Settings struct {
	Locale             string // 地区
	ViewDistance       int    // 视距
	ChatMode           int    // 聊天模式
	ChatColors         bool   // 聊天颜色
	DisplayedSkinParts uint8  // 皮肤显示
	MainHand           int    // 主手

	// Enables filtering of text on signs and written book titles.
	// Currently, always false (i.e. the filtering is disabled)
	EnableTextFiltering bool
	AllowListing        bool

	// The brand string presented to the server.
	Brand string
}

// Used by Settings.DisplayedSkinParts.
// For each bit set if shows match part.
const (
	_ = 1 << iota
	Jacket
	LeftSleeve
	RightSleeve
	LeftPantsLeg
	RightPantsLeg
	Hat
)

// DefaultSettings are the default settings of client
var DefaultSettings = Settings{
	Locale:             "zh_CN", // ^_^
	ViewDistance:       15,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,

	EnableTextFiltering: false,
	AllowListing:        true,

	Brand: "vanilla",
}
