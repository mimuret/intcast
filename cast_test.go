package intcast_test

import (
	"math"
	"strings"
	"testing"

	"github.com/mimuret/intcast"
)

type castTestCase struct {
	Name  string
	OK    bool
	Value interface{}
	Dst   interface{}
}

func TestCast(t *testing.T) {
	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
		u   uint
		u8  uint8
		u16 uint16
		u32 uint32
		u64 uint64
	)
	testcase := []castTestCase{
		{"int2i8", false, int(math.MaxInt), &i8},
		{"int2i16", false, int(math.MaxInt), &i16},
		{"int2i32", false, int(math.MaxInt), &i32},
		{"int2i64", true, int(math.MaxInt), &i64},
		{"int2int", true, int(math.MaxInt), &i},
		{"int2u8", false, int(math.MaxInt), &u8},
		{"int2u16", false, int(math.MaxInt), &u16},
		{"int2u32", false, int(math.MaxInt), &u32},
		{"int2u64", true, int(math.MaxInt), &u64},
		{"int2uint", true, int(math.MaxInt), &u},
	}
	// invalid type
	err := intcast.Cast(1, 1)

	if !strings.Contains(err.Error(), "invalid int type") {
		t.Errorf("NG: invalid type")
	}
	for _, tc := range testcase {
		err := intcast.Cast(tc.Value, tc.Dst)
		if err != nil {
			if strings.Contains(err.Error(), "invalid int type") {
				t.Errorf("NG: %s error occurred: %s", tc.Name, err.Error())
			} else if tc.OK {
				t.Errorf("NG: %s error occurred: %s", tc.Name, err.Error())
			} else {
				t.Logf("OK: %s", tc.Name)
			}
		} else if !tc.OK {
			t.Errorf("NG: %s error not occurred: %s", tc.Name, err.Error())
		} else {
			t.Logf("OK: %s", tc.Name)
		}
	}
}
