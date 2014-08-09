// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package haru_test

import (
	"fmt"

	"github.com/chai2010/go-haru"
)

func ExampleGetVersion() {
	fmt.Printf("haru version: %v\n", haru.GetVersion())
	// Output:
	// haru version: 2.4.0dev
}
