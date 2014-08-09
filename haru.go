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
	"unsafe"
)

func GetVersion() string {
	p := C.HPDF_GetVersion()
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}
