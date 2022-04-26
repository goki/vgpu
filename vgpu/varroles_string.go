// Code generated by "stringer -type=VarRoles"; DO NOT EDIT.

package vgpu

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UndefVarRole-0]
	_ = x[Vertex-1]
	_ = x[Index-2]
	_ = x[Uniform-3]
	_ = x[Storage-4]
	_ = x[UniformTexel-5]
	_ = x[StorageTexel-6]
	_ = x[StorageImage-7]
	_ = x[SamplerVar-8]
	_ = x[SampledImage-9]
	_ = x[CombinedImage-10]
	_ = x[VarRolesN-11]
}

const _VarRoles_name = "UndefVarRoleVertexIndexUniformStorageUniformTexelStorageTexelStorageImageSamplerVarSampledImageCombinedImageVarRolesN"

var _VarRoles_index = [...]uint8{0, 12, 18, 23, 30, 37, 49, 61, 73, 83, 95, 108, 117}

func (i VarRoles) String() string {
	if i < 0 || i >= VarRoles(len(_VarRoles_index)-1) {
		return "VarRoles(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VarRoles_name[_VarRoles_index[i]:_VarRoles_index[i+1]]
}

func (i *VarRoles) FromString(s string) error {
	for j := 0; j < len(_VarRoles_index)-1; j++ {
		if s == _VarRoles_name[_VarRoles_index[j]:_VarRoles_index[j+1]] {
			*i = VarRoles(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: VarRoles")
}