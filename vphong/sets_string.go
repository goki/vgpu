// Code generated by "stringer -type=Sets"; DO NOT EDIT.

package vphong

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MtxsSet-0]
	_ = x[ColorSet-1]
	_ = x[ViewMtxSet-2]
	_ = x[NLightSet-3]
	_ = x[LightSet-4]
	_ = x[TexSet-5]
	_ = x[SetsN-6]
}

const _Sets_name = "MtxsSetColorSetViewMtxSetNLightSetLightSetTexSetSetsN"

var _Sets_index = [...]uint8{0, 7, 15, 25, 34, 42, 48, 53}

func (i Sets) String() string {
	if i < 0 || i >= Sets(len(_Sets_index)-1) {
		return "Sets(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Sets_name[_Sets_index[i]:_Sets_index[i+1]]
}

func (i *Sets) FromString(s string) error {
	for j := 0; j < len(_Sets_index)-1; j++ {
		if s == _Sets_name[_Sets_index[j]:_Sets_index[j+1]] {
			*i = Sets(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Sets")
}