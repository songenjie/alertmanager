// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package width

import "strconv"

const _Kind_name = "NeutralEastAsianAmbiguousEastAsianWideEastAsianNarrowEastAsianFullwidthEastAsianHalfwidth"

var _Kind_index = [...]uint8{0, 7, 25, 38, 53, 71, 89}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
