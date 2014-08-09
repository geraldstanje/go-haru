// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package haru

import (
	"testing"
)

func TextGetVersion(t *testing.T) {
	if s := GetVersion(); s != "" {
		t.Fatalf("got %v", s)
	}
}
