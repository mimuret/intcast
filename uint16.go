package cast

import (
	"fmt"
	"math"
)

func GetUInt16(arg interface{}) (uint16, error) {
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
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		u64 = uint64(v)
	case uint:
		u64 = uint64(v)
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid uint16 type %v", arg)
	}
	switch arg.(type) {
	case int64, int, int32, int16, int8:
		if 0 <= i64 && i64 <= math.MaxUint16 {
			return uint16(i64), nil
		}
	case uint64, uint, uint32:
		if u64 <= math.MaxUint16 {
			return uint16(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid uint16 range value %v", arg)
}
