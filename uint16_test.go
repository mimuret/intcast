package cast_test

import (
	"math"
	"testing"

	"github.com/mimuret/cast"
)

func TestUInt16(t *testing.T) {
	testcases := []testcase{
		// MaxInt*
		{"max8:int8", true, int8(math.MaxInt8)},
		{"max16:int16", true, int16(math.MaxInt16)},
		{"max32:int32", false, int32(math.MaxInt32)},
		{"max64:int64", false, int64(math.MaxInt64)},
		{"maxInt:int", false, int(math.MaxInt)},
		// 0
		{"0:int8", true, int8(0)},
		{"0:int16", true, int16(0)},
		{"0:int32", true, int32(0)},
		{"0:int64", true, int64(0)},
		{"0:int", true, int(0)},
		// -1
		{"min8:int8", false, int8(-1)},
		{"min16:int16", false, int16(-1)},
		{"min32:int32", false, int32(-1)},
		{"min64:int64", false, int64(-1)},
		{"minInt:int", false, int(-1)},
		// MinInt*
		{"min8:int8", false, int8(math.MinInt8)},
		{"min16:int16", false, int16(math.MinInt16)},
		{"min32:int32", false, int32(math.MinInt32)},
		{"min64:int64", false, int64(math.MinInt64)},
		{"minInt:int", false, int(math.MinInt)},
		// MaxUint8
		//{"math.MaxUint16:int8", true, int8(math.MaxUint16)},
		//{"math.MaxUint16:int16", true, int16(math.MaxUint16)},
		{"math.MaxUint16:int32", true, int32(math.MaxUint16)},
		{"math.MaxUint16:int64", true, int64(math.MaxUint16)},
		{"math.MaxUint16:int", true, int(math.MaxUint16)},
		// MaxUint8+1
		//{"math.MaxUint16+1:int8", false, int8(math.MaxUint16) + 1},
		//{"math.MaxUint16+1:int16", false, int16(math.MaxUint16 + 1)},
		{"math.MaxUint16+1:int32", false, int32(math.MaxUint16) + 1},
		{"math.MaxUint16+1:int64", false, int64(math.MaxUint16) + 1},
		{"math.MaxUint16+1:int", false, int(math.MaxUint16) + 1},
		// MaxUint*
		{"maxU8:uint8", true, uint8(math.MaxUint8)},
		{"maxU16:uint16", true, uint16(math.MaxUint16)},
		{"maxU32:uint32", false, uint32(math.MaxUint32)},
		{"maxU64:uint64", false, uint64(math.MaxUint64)},
		{"maxUInt:uint", false, uint(math.MaxUint)},
		// MinUint* = 0
		{"minU8:uint8", true, uint8(0)},
		{"minU16:uint16", true, uint16(0)},
		{"minU32:uint32", true, uint32(0)},
		{"minU64:uint64", true, uint64(0)},
		{"minUInt:uint", true, uint(0)},
		// MaxInt8
		//{"math.MaxUint16:uint8", true, uint8(math.MaxUint16)},
		{"math.MaxUint16:uint16", true, uint16(math.MaxUint16)},
		{"math.MaxUint16:uint32", true, uint32(math.MaxUint16)},
		{"math.MaxUint16:uint64", true, uint64(math.MaxUint16)},
		{"math.MaxUint16:uint", true, uint(math.MaxUint16)},
		// MaxInt8+1
		//{"math.MaxUint16+1:uint8", false, uint8(math.MaxUint16 + 1)},
		//{"math.MaxUint16+1:uint16", false, uint16(math.MaxUint16 + 1)},
		{"math.MaxUint16+1:uint32", false, uint32(math.MaxUint16 + 1)},
		{"math.MaxUint16+1:uint64", false, uint64(math.MaxUint16 + 1)},
		{"math.MaxUint16+1:uint", false, uint(math.MaxUint16 + 1)},
		// invalid
		{"invalid", false, "1"},
	}

	for _, tc := range testcases {
		_, err := cast.GetUInt16(tc.Value)
		if err != nil && tc.OK {
			t.Errorf("NG: %s error occurred", tc.Name)
		} else {
			t.Logf("OK: %s", tc.Name)
		}
		if err == nil && !tc.OK {
			t.Errorf("NG: %s error not occurred", tc.Name)
		} else {
			t.Logf("OK: %s", tc.Name)
		}
	}
}
