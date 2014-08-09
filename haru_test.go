// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package haru

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	if a, b := hpdfGetVersion(), GetVersion(); a != b {
		t.Fatalf("expect = %v, got = %v", a, b)
	}
}
