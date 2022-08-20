package intcast

import (
	"fmt"
	"math"
)

func GetInt64(arg interface{}) (int64, error) {
	var u64 uint64
	switch v := arg.(type) {
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint:
		u64 = uint64(v)
	case uint64:
		u64 = v
	default:
		return 0, fmt.Errorf("invalid int64 type %v", arg)
	}
	switch arg.(type) {
	case uint64, uint:
		if u64 <= math.MaxInt64 {
			return int64(u64), nil
		}
	}
	return 0, fmt.Errorf("invalid int64 range value %v", arg)
}
