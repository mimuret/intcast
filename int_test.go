package cast_test

import (
	"math"
	"testing"

	"github.com/mimuret/cast"
)

func TestInt(t *testing.T) {
	targetMax := math.MaxInt
	targetMin := math.MinInt
	is64 := false
	if math.MaxInt == math.MaxInt64 {
		is64 = true
	}
	testcases := []testcase{
		// MaxInt*
		{"max8:int8", true, int8(math.MaxInt8)},
		{"max16:int16", true, int16(math.MaxInt16)},
		{"max32:int32", true, int32(math.MaxInt32)},
		{"max64:int64", is64, int64(math.MaxInt64)},
		{"maxInt:int", true, int(math.MaxInt)},
		// MinInt*
		{"min8:int8", true, int8(math.MinInt8)},
		{"min16:int16", true, int16(math.MinInt16)},
		{"min32:int32", true, int32(math.MinInt32)},
		{"min64:int64", is64, int64(math.MinInt64)},
		{"minInt:int", true, int(math.MinInt)},
		// MaxInt64
		//	{"targetMax:int8", true, int8(targetMax)},
		//	{"targetMax:int16", true, int16(targetMax)},
		{"targetMax:int64", true, int64(targetMax)},
		{"targetMax:int", true, int(targetMax)},
		// MinInt64
		//		{"targetMin:int8", true, int8(targetMin)},
		//	{"targetMin:int16", true, int16(targetMin)},
		{"targetMin:int64", true, int64(targetMin)},
		{"targetMin:int", true, int(targetMin)},
		// MaxInt32+1
		//	{"targetMax+1:int8", false, int8(targetMax+1) },
		//	{"targetMax+1:int16", false, int16(targetMax+1)},
		//	{"targetMax+1:int32", false, int32(targetMax + 1)},
		//	{"targetMax+1:int", false, int(targetMax + 1)},
		// MinInt32-1
		//	{"targetMin-1:int8", false, int8(targetMin) - 1},
		//	{"targetMin-1:int16", false, int16(targetMin) - 1},
		//	{"targetMin-1:int32", false, int32(targetMin - 1)},
		//	{"targetMin-1:int", false, int(targetMin - 1)},
		// MaxUint*
		{"maxU8:uint8", true, uint8(math.MaxUint8)},
		{"maxU16:uint16", true, uint16(math.MaxUint16)},
		{"maxU32:uint32", is64, uint32(math.MaxUint32)},
		{"maxU64:uint64", false, uint64(math.MaxUint64)},
		{"maxUInt:uint", false, uint(math.MaxUint)},
		// MinUint* = 0
		{"minU8:uint8", true, uint8(0)},
		{"minU16:uint16", true, uint16(0)},
		{"minU32:uint32", true, uint32(0)},
		{"minU64:uint64", true, uint64(0)},
		{"minUInt:uint", true, uint(0)},
		// MaxInt32
		//{"targetMax:uint8", true, uint8(targetMax)},
		//{"targetMax:uint16", true, uint16(targetMax)},
		{"targetMax:uint32", true, uint32(targetMax)},
		{"targetMax:uint64", true, uint64(targetMax)},
		{"targetMax:uint", true, uint(targetMax)},
		// MaxInt32+1
		//{"targetMax:uint8", false, uint8(targetMax + 1)},
		//{"targetMax:uint16", false, uint16(targetMax + 1)},
		{"targetMax+1:uint64", false, uint64(targetMax + 1)},
		{"targetMax+1:uint", false, uint64(targetMax + 1)},
		{"invalid", false, "1"},
	}
	if math.MaxInt == math.MaxInt32 {
		testcases = append(testcases,
			testcase{"targetMax:int32", true, int32(targetMax)},
			testcase{"targetMin:int32", true, int32(targetMin)},
			testcase{"targetMax+1:int64", false, int64(targetMax + 1)},
			testcase{"targetMin-1:int64", false, int64(targetMin - 1)},
			testcase{"targetMax:uint32", false, uint32(targetMax + 1)},
		)
	}
	for _, tc := range testcases {
		_, err := cast.GetInt(tc.Value)
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
