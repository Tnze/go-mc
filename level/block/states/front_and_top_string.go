// Code generated by "stringer -type=FrontAndTop -output=front_and_top_string.go"; DO NOT EDIT.

package states

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DownEast-0]
	_ = x[DownNorth-1]
	_ = x[DownSouth-2]
	_ = x[DownWest-3]
	_ = x[UpEast-4]
	_ = x[UpNorth-5]
	_ = x[UpSouth-6]
	_ = x[UpWest-7]
	_ = x[WestUp-8]
	_ = x[EastUp-9]
	_ = x[NorthUp-10]
	_ = x[SouthUp-11]
}

const _FrontAndTop_name = "DownEastDownNorthDownSouthDownWestUpEastUpNorthUpSouthUpWestWestUpEastUpNorthUpSouthUp"

var _FrontAndTop_index = [...]uint8{0, 8, 17, 26, 34, 40, 47, 54, 60, 66, 72, 79, 86}

func (i FrontAndTop) String() string {
	if i >= FrontAndTop(len(_FrontAndTop_index)-1) {
		return "FrontAndTop(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FrontAndTop_name[_FrontAndTop_index[i]:_FrontAndTop_index[i+1]]
}

func (i FrontAndTop) Value() byte {
	return byte(i)
}
