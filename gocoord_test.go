package gocoord_test

import (
	"math"
	"testing"

	"github.com/suifengtec/gocoord"
)

var testsCases = []struct {
	from          gocoord.Position
	expectedValue gocoord.Position
}{
	{ // BD09ToGCJ02
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.73337197243862, Lat: 34.350630274732744},
	},
	{ //BD09ToWGS84
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.7272281721665, Lat: 34.351951705458674},
	},
	{ //WGS84ToBD09
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.75246057307969, Lat: 34.3616443275599},
	},
	{ //WGS84ToGCJ02
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.74601884470827, Lat: 34.35538495876689},
	},
	{ //GCJ02ToBD09
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.7463400404353, Lat: 34.36287007240432},
	},
	{ //GCJ02ToWGS84
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 113.73372688226371, Lat: 34.358009517422026},
	},

	{ //BD09toBD09MC
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 1.266160251233583e+07, Lat: 4.052461800944403e+06},
	},
	{ //BD09MCtoBD09
		gocoord.Position{Lon: 113.739873, Lat: 34.356696},
		gocoord.Position{Lon: 0.0010217344366200546, Lat: 0.000310700669093249},
	},
}

func TestBD09ToGCJ02(t *testing.T) {

	input := testsCases[0]
	ret := gocoord.BD09ToGCJ02(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-6 {
		t.Errorf("FAIL: BD09ToGCJ02 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
func TestBD09ToWGS84(t *testing.T) {

	input := testsCases[1]
	ret := gocoord.BD09ToWGS84(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-6 {
		t.Errorf("FAIL: BD09ToWGS84 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
func TestWGS84ToBD09(t *testing.T) {

	input := testsCases[2]
	ret := gocoord.WGS84ToBD09(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: WGS84ToBD09 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}

func TestWGS84ToGCJ02(t *testing.T) {

	input := testsCases[3]
	ret := gocoord.WGS84ToGCJ02(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: WGS84ToGCJ02 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
func TestGCJ02ToBD09(t *testing.T) {

	input := testsCases[4]
	ret := gocoord.GCJ02ToBD09(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: GCJ02ToBD09 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
func TestGCJ02ToWGS84(t *testing.T) {

	input := testsCases[5]
	ret := gocoord.GCJ02ToWGS84(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: GCJ02ToWGS84 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}

func TestBD09toBD09MC(t *testing.T) {

	input := testsCases[6]
	ret := gocoord.BD09toBD09MC(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: BD09toBD09MC 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
func TestBD09MCtoBD09(t *testing.T) {

	input := testsCases[7]
	ret := gocoord.BD09MCtoBD09(input.from)

	if math.Abs(ret.Lat-input.expectedValue.Lat) > 1e-3 {

		t.Errorf("FAIL: BD09MCtoBD09 测试的输入是 %#v ,期望得到 : %#v ,但计算结果为 %#v",
			input.from,
			input.expectedValue,
			ret,
		)
	}

}
