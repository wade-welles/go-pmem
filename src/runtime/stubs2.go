// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9
// +build !solaris
// +build !windows
// +build !nacl
// +build !js
// +build !darwin
// +build !aix

package runtime

import "unsafe"

func read(fd int32, p unsafe.Pointer, n int32) int32
func closefd(fd int32) int32

func exit(code int32)
func usleep(usec uint32)

//go:noescape
func write(fd uintptr, p unsafe.Pointer, n int32) int32

func ftruncate(fd, len uintptr) int32
func fallocate(fd, mode, offset, len uintptr) int32
func fstat(fd, stat uintptr) int32
func unlinkat(fd, path, flags uintptr) int32
func msync(addr, len, flags uintptr) int32
func readlink(path, buf, len uintptr) int32

//go:noescape
func open(name *byte, mode, perm int32) int32

// return value is only set on linux to be used in osinit()
func madvise(addr unsafe.Pointer, n uintptr, flags int32) int32

// exitThread terminates the current thread, writing *wait = 0 when
// the stack is safe to reclaim.
//
//go:noescape
func exitThread(wait *uint32)
