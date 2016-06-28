// +build !appengine

package memcache

import (
	"unsafe"
)

var (
	MemcacheUseUnsafeStobs = true
)

func stobs(s string) []byte {
	if MemcacheUseUnsafeStobs {
		return *(*[]byte)(unsafe.Pointer(&s))
	}
	return []byte(s)
}
