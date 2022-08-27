package intcast

import (
	"fmt"
)

func Cast(src, dst interface{}) error {
	var err error
	switch v := dst.(type) {
	case *int:
		*v, err = GetInt(src)
	case *int8:
		*v, err = GetInt8(src)
	case *int16:
		*v, err = GetInt16(src)
	case *int32:
		*v, err = GetInt32(src)
	case *int64:
		*v, err = GetInt64(src)
	case *uint:
		*v, err = GetUInt(src)
	case *uint8:
		*v, err = GetUInt8(src)
	case *uint16:
		*v, err = GetUInt16(src)
	case *uint32:
		*v, err = GetUInt32(src)
	case *uint64:
		*v, err = GetUInt64(src)
	default:
		err = fmt.Errorf("invalid int type %v", dst)
	}
	return err
}
