package intcast

import (
	"fmt"
	"math"
)

func GetInt(arg interface{}) (int, error) {
	var i64 int64
	var u64 uint64
	switch v := arg.(type) {
	case int:
		return int(v), nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		i64 = v
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		u64 = uint64(v)
	case uint:
		u64 = uint64(v)
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid int type %v", arg)
	}
	switch arg.(type) {
	case int64:
		if math.MinInt <= i64 && i64 <= math.MaxInt {
			return int(i64), nil
		}
	case uint64, uint, uint32:
		if u64 <= math.MaxInt {
			return int(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid int range value %v", arg)
}
