package intcast_test

import (
	"math"
	"testing"

	cast "github.com/mimuret/intcast"
)

func TestInt8(t *testing.T) {
	testcases := []testcase{
		// MaxInt*
		{"max8:int8", true, int8(math.MaxInt8)},
		{"max16:int16", false, int16(math.MaxInt16)},
		{"max32:int32", false, int32(math.MaxInt32)},
		{"max64:int64", false, int64(math.MaxInt64)},
		{"maxInt:int", false, int(math.MaxInt)},
		// MinInt*
		{"min8:int8", true, int8(math.MinInt8)},
		{"min16:int16", false, int16(math.MinInt16)},
		{"min32:int32", false, int32(math.MinInt32)},
		{"min64:int64", false, int64(math.MinInt64)},
		{"minInt:int", false, int(math.MinInt)},
		// MaxInt8
		{"math.MaxInt8:int8", true, int8(math.MaxInt8)},
		{"math.MaxInt8:int16", true, int16(math.MaxInt8)},
		{"math.MaxInt8:int32", true, int32(math.MaxInt8)},
		{"math.MaxInt8:int64", true, int64(math.MaxInt8)},
		{"math.MaxInt8:int", true, int(math.MaxInt8)},
		// MinInt8
		{"math.MinInt8:int8", true, int8(math.MinInt8)},
		{"math.MinInt8:int16", true, int16(math.MinInt8)},
		{"math.MinInt8:int32", true, int32(math.MinInt8)},
		{"math.MinInt8:int64", true, int64(math.MinInt8)},
		{"math.MinInt8:int", true, int(math.MinInt8)},
		// MaxInt8+1
		// {"math.MaxInt8+1:int8", false, int8(math.MaxInt8) + 1},
		{"math.MaxInt8+1:int16", false, int16(math.MaxInt8 + 1)},
		{"math.MaxInt8+1:int32", false, int32(math.MaxInt8) + 1},
		{"math.MaxInt8+1:int64", false, int64(math.MaxInt8) + 1},
		{"math.MaxInt8+1:int", false, int(math.MaxInt8) + 1},
		// MinInt8 -1
		// {"math.MinInt8-1:int8", false, int8(math.MinInt8-1)},
		{"math.MinInt8-1:int16", false, int16(math.MinInt8 - 1)},
		{"math.MinInt8-1:int32", false, int32(math.MinInt8 - 1)},
		{"math.MinInt8-1:int64", false, int64(math.MinInt8 - 1)},
		{"math.MinInt8-1:int", false, int(math.MinInt8) - 1},
		// MaxUint*
		{"maxU8:uint8", false, uint8(math.MaxUint8)},
		{"maxU16:uint16", false, uint16(math.MaxUint16)},
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
		{"math.MaxInt8:uint8", true, uint8(math.MaxInt8)},
		{"math.MaxInt8:uint16", true, uint16(math.MaxInt8)},
		{"math.MaxInt8:uint32", true, uint32(math.MaxInt8)},
		{"math.MaxInt8:uint64", true, uint64(math.MaxInt8)},
		{"math.MaxInt8:uint", true, uint(math.MaxInt8)},
		// MaxInt8+1
		{"math.MaxInt8+1:uint8", false, uint8(math.MaxInt8 + 1)},
		{"math.MaxInt8+1:uint16", false, uint16(math.MaxInt8 + 1)},
		{"math.MaxInt8+1:uint32", false, uint32(math.MaxInt8 + 1)},
		{"math.MaxInt8+1:uint64", false, uint64(math.MaxInt8 + 1)},
		{"math.MaxInt8+1:uint", false, uint(math.MaxInt8 + 1)},
		// invalid
		{"invalid", false, "1"},
	}

	for _, tc := range testcases {
		_, err := cast.GetInt8(tc.Value)
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
