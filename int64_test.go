package intcast_test

import (
	"math"
	"testing"

	cast "github.com/mimuret/intcast"
)

func TestInt64(t *testing.T) {
	testcases := []testcase{
		// MaxInt*
		{"max8:int8", true, int8(math.MaxInt8)},
		{"max16:int16", true, int16(math.MaxInt16)},
		{"max32:int32", true, int32(math.MaxInt32)},
		{"max64:int64", true, int64(math.MaxInt64)},
		{"maxInt:int", true, int(math.MaxInt)},
		// MinInt*
		{"min8:int8", true, int8(math.MinInt8)},
		{"min16:int16", true, int16(math.MinInt16)},
		{"min32:int32", true, int32(math.MinInt32)},
		{"min64:int64", true, int64(math.MinInt64)},
		{"minInt:int", true, int(math.MinInt)},
		// MaxInt64
		//{"math.MaxInt64:int8", true, int8(math.MaxInt64)},
		//{"math.MaxInt64:int16", true, int16(math.MaxInt64)},
		//{"math.MaxInt64:int32", true, int32(math.MaxInt64)},
		{"math.MaxInt64:int64", true, int64(math.MaxInt64)},
		// MinInt64
		//{"math.MinInt64:int8", true, int8(math.MinInt64)},
		//{"math.MinInt64:int16", true, int16(math.MinInt64)},
		//{"math.MinInt64:int32", true, int32(math.MinInt64)},
		{"math.MinInt64:int64", true, int64(math.MinInt64)},
		// MaxInt32+1
		//	{"math.MaxInt64+1:int8", false, int8(math.MaxInt64+1) },
		//	{"math.MaxInt64+1:int16", false, int16(math.MaxInt64+1)},
		//	{"math.MaxInt64+1:int32", false, int32(math.MaxInt64 + 1)},
		//	{"math.MaxInt64+1:int64", false, int64(math.MaxInt64 + 1)},
		//	{"math.MaxInt64+1:int", false, int(math.MaxInt64 + 1)},
		// MinInt32-1
		//	{"math.MinInt64-1:int8", false, int8(math.MinInt64) - 1},
		//	{"math.MinInt64-1:int16", false, int16(math.MinInt64) - 1},
		//	{"math.MinInt64-1:int32", false, int32(math.MinInt64 - 1)},
		//	{"math.MinInt64-1:int64", false, int64(math.MinInt64 - 1)},
		//	{"math.MinInt64-1:int", false, int(math.MinInt64 - 1)},
		// MaxUint*
		{"maxU8:uint8", true, uint8(math.MaxUint8)},
		{"maxU16:uint16", true, uint16(math.MaxUint16)},
		{"maxU32:uint32", true, uint32(math.MaxUint32)},
		{"maxU64:uint64", false, uint64(math.MaxUint64)},
		// MinUint* = 0
		{"minU8:uint8", true, uint8(0)},
		{"minU16:uint16", true, uint16(0)},
		{"minU32:uint32", true, uint32(0)},
		{"minU64:uint64", true, uint64(0)},
		{"minUInt:uint", true, uint(0)},
		// MaxInt32
		//{"math.MaxInt64:uint8", true, uint8(math.MaxInt64)},
		//{"math.MaxInt64:uint16", true, uint16(math.MaxInt64)},
		//{"math.MaxInt64:uint32", true, uint32(math.MaxInt64)},
		{"math.MaxInt64:uint64", true, uint64(math.MaxInt64)},
		{"math.MaxInt64:uint", true, uint(math.MaxInt64)},
		// MaxInt32+1
		//{"math.MaxInt64:uint8", false, uint8(math.MaxInt64 + 1)},
		//{"math.MaxInt64:uint16", false, uint16(math.MaxInt64 + 1)},
		//{"math.MaxInt64:uint32", false, uint32(math.MaxInt64 + 1)},
		{"math.MaxInt64+1:uint64", false, uint64(math.MaxInt64) + 1},
		// invalid
		{"invalid", false, "1"},
	}
	if math.MaxInt == math.MaxInt64 {
		testcases = append(testcases,
			testcase{"math.MaxInt64:int", true, int(math.MaxInt64)},
			testcase{"math.MinInt64:int", true, int(math.MinInt64)},
		)
	}
	if math.MaxUint == math.MaxUint64 {
		testcases = append(testcases,
			testcase{"maxUInt:uint", false, uint(math.MaxUint)},
			testcase{"math.MaxInt64+1:uint", false, uint(math.MaxInt64) + 1},
		)
	} else {
		testcases = append(testcases,
			testcase{"maxUInt:uint", true, uint(math.MaxUint)},
		)
	}
	for _, tc := range testcases {
		_, err := cast.GetInt64(tc.Value)
		if err != nil && tc.OK {
			t.Errorf("NG: %s error occurred: %s", tc.Name, err.Error())
		} else {
			t.Logf("OK: %s", tc.Name)
		}
		if err == nil && !tc.OK {
			t.Errorf("NG: %s error not occurred: %s", tc.Name, err.Error())
		} else {
			t.Logf("OK: %s", tc.Name)
		}
	}
}
