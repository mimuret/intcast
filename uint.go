package intcast

import (
	"fmt"
	"math"
)

func GetUInt(arg interface{}) (uint, error) {
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
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint:
		return v, nil
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid uint type %v", arg)
	}
	switch arg.(type) {
	case int64, int, int32, int16, int8:
		if 0 <= i64 {
			return uint(i64), nil
		}
	case uint64:
		if u64 <= math.MaxUint {
			return uint(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid uint range value %v", arg)
}
