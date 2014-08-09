// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go bindings for libharu.
//
// See http://libharu.org/
package haru

/*
#include "hpdf.h"
*/
import "C"
import (
//"unsafe"
)

const (
	hpdfVersion = "2.4.0dev"
)

// for test
func getVersion() string {
	p := C.HPDF_GetVersion()
	s := C.GoString(p)
	return s
}

func GetVersion() string {
	return hpdfVersion
}
