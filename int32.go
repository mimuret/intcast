package cast

import (
	"fmt"
	"math"
)

func GetInt32(arg interface{}) (int32, error) {
	var i64 int64
	var u64 uint64
	switch v := arg.(type) {
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int:
		i64 = int64(v)
	case int64:
		i64 = v
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		u64 = uint64(v)
	case uint:
		u64 = uint64(v)
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid int32 type %v", arg)
	}
	switch arg.(type) {
	case int64, int:
		if math.MinInt32 <= i64 && i64 <= math.MaxInt32 {
			return int32(i64), nil
		}
	case uint64, uint, uint32:
		if u64 <= math.MaxInt32 {
			return int32(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid int32 range value %v", arg)
}
