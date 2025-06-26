// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slboolcore

import (
	"cogentcore.org/core/core"
	"goki.dev/vgpu/v2/gosl/slbool"
)

func init() {
	core.AddValueType[slbool.Bool, core.Switch]()
}
