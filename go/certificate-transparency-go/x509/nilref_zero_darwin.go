// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !arm && !arm64 && !ios && go1.10
// +build cgo,!arm,!arm64,!ios,go1.10

package x509

/*
#cgo CFLAGS: -mmacosx-version-min=10.6
#cgo LDFLAGS: -framework CoreFoundation -framework Security

#include <CoreFoundation/CoreFoundation.h>
*/
import "C"

// For Go versions >= 1.10, nil values for Apple's CoreFoundation
// CF*Ref types are represented by zero.  See:
//
//	https://github.com/golang/go/commit/b868616b63a8
func setNilCFRef(v *C.CFDataRef) {
	*v = 0
}

func isNilCFRef(v C.CFDataRef) bool {
	return v == 0
}
