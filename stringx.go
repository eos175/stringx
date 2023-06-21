package stringx

import (
	"reflect"
	"unsafe"
)

var mask = [8]uint64{
	0x0000000000000000,
	0x00000000000000ff,
	0x000000000000ffff,
	0x0000000000ffffff,
	0x00000000ffffffff,
	0x000000ffffffffff,
	0x0000ffffffffffff,
	0x00ffffffffffffff,
}

type ptr = unsafe.Pointer

func HasPrefixFast(s, prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if len(prefix) > len(s) {
		return false
	}

	sH := *(*uint64)(ptr(((*reflect.StringHeader)(ptr(&s)).Data)))
	preH := *(*uint64)(ptr(((*reflect.StringHeader)(ptr(&prefix)).Data)))

	if len(prefix) > 8 {
		return sH == preH && s[0:len(prefix)] == prefix
	}

	if len(prefix) == 8 {
		return sH == preH
	}

	//mask := ^(uint64(0xffffffffffffffff) << uint64(8*len(prefix)))
	//mask := uint64(1<<uint64(8*len(prefix))) - 1

	//return (sH^preH)&mask == 0
	return (sH^preH)&mask[len(prefix)] == 0
}
