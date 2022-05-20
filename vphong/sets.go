// Copyright (c) 2022, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vphong

import "github.com/goki/ki/kit"

// Sets are variable set numbers - must coordinate with System sets!
type Sets int

const (
	MtxsSet Sets = iota
	ColorSet
	ViewMtxSet
	NLightSet
	LightSet
	TexSet
	SetsN
)

//go:generate stringer -type=Sets

var KiT_Sets = kit.Enums.AddEnum(SetsN, kit.NotBitFlag, nil)