package intcast

import (
	"fmt"
)

func GetUInt64(arg interface{}) (uint64, error) {
	var i64 int64
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
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid uint64 type %v", arg)
	}
	switch arg.(type) {
	case int64, int, int32, int16, int8:
		if 0 <= i64 {
			return uint64(i64), nil
		}
	}
	return 0, fmt.Errorf("invalid uint64 range value %v", arg)
}
