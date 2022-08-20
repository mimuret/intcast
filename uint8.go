package intcast

import (
	"fmt"
	"math"
)

func GetUInt8(arg interface{}) (uint8, error) {
	var i64 int64
	var u64 uint64
	switch v := arg.(type) {
	case int8:
		i64 = int64(v)
	case int16:
		i64 = int64(v)
	case int:
		i64 = int64(v)
	case int32:
		i64 = int64(v)
	case int64:
		i64 = int64(v)
	case uint8:
		return v, nil
	case uint16:
		u64 = uint64(v)
	case uint32:
		u64 = uint64(v)
	case uint:
		u64 = uint64(v)
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid uint8 type %v", arg)
	}
	switch arg.(type) {
	case int64, int, int32, int16, int8:
		if 0 <= i64 && i64 <= math.MaxUint8 {
			return uint8(i64), nil
		}
	case uint64, uint, uint32, uint16:
		if u64 <= math.MaxUint8 {
			return uint8(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid uint8 range value %v", arg)
}
