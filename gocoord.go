package gocoord

import (
	"math"
)

/*

BD09与GCJ02与WGS84点的坐标的相互转换;


*/

// PI ...
const PI = math.Pi

// WGS84
const WGS84 = "WGS84"
const WGS1984 = WGS84
const EPSG4326 = WGS84

// GCJ02
const GCJ02 = "GCJ02"
const AMap = GCJ02

// BD09
const BD09 = "BD09"
const BD09LL = BD09
const Baidu = BD09
const BMap = BD09
const baiduFactor = math.Pi * 3000.0 / 180.0

// BD09MC
const BD09MC = "BD09MC"
const BD09Meter = BD09MC

// EPSG3857
const EPSG3857 = "EPSG3857"
const EPSG900913 = EPSG3857
const EPSG102100 = EPSG3857
const WebMercator = EPSG3857
const WM = EPSG3857

const a = 6378245
const ee = 0.006693421622965823

// MCBAND ...
var MCBAND = []float64{
	12890594.86,
	8362377.87,
	5591021,
	3481989.83,
	1678043.12,
	0.0,
}

// MC2LL ...
var MC2LL = [][]float64{
	[]float64{1.410526172116255e-8,
		0.00000898305509648872,
		-1.9939833816331,
		200.9824383106796,
		-187.2403703815547,
		91.6087516669843,
		-23.38765649603339,
		2.57121317296198,
		-0.03801003308653,
		17337981.2,
	},
	[]float64{
		-7.435856389565537e-9,
		0.000008983055097726239,
		-0.78625201886289,
		96.32687599759846,
		-1.85204757529826,
		-59.36935905485877,
		47.40033549296737,
		-16.50741931063887,
		2.28786674699375,
		10260144.86,
	},
	[]float64{
		-3.030883460898826e-8,
		0.00000898305509983578,
		0.30071316287616,
		59.74293618442277,
		7.357984074871,
		-25.38371002664745,
		13.45380521110908,
		-3.29883767235584,
		0.32710905363475,
		6856817.37,
	},
	[]float64{
		-1.981981304930552e-8,
		0.000008983055099779535,
		0.03278182852591,
		40.31678527705744,
		0.65659298677277,
		-4.44255534477492,
		0.85341911805263,
		0.12923347998204,
		-0.04625736007561,
		4482777.06,
	},
	[]float64{
		3.09191371068437e-9,
		0.000008983055096812155,
		0.00006995724062,
		23.10934304144901,
		-0.00023663490511,
		-0.6321817810242,
		-0.00663494467273,
		0.03430082397953,
		-0.00466043876332,
		2555164.4,
	},
	[]float64{
		2.890871144776878e-9,
		0.000008983055095805407,
		-3.068298e-8,
		7.47137025468032,
		-0.00000353937994,
		-0.02145144861037,
		-0.00001234426596,
		0.00010322952773,
		-0.00000323890364,
		826088.5,
	},
}

// LL2MC ...
var LL2MC = [][]float64{
	[]float64{
		-0.0015702102444,
		111320.7020616939,
		1704480524535203,
		-10338987376042340,
		26112667856603880,
		-35149669176653700,
		26595700718403920,
		-10725012454188240,
		1800819912950474,
		82.5,
	},
	[]float64{
		0.0008277824516172526,
		111320.7020463578,
		647795574.6671607,
		-4082003173.641316,
		10774905663.51142,
		-15171875531.51559,
		12053065338.62167,
		-5124939663.577472,
		913311935.9512032,
		67.5,
	},
	[]float64{
		0.00337398766765,
		111320.7020202162,
		4481351.045890365,
		-23393751.19931662,
		79682215.47186455,
		-115964993.2797253,
		97236711.15602145,
		-43661946.33752821,
		8477230.501135234,
		52.5,
	},
	[]float64{
		0.00220636496208,
		111320.7020209128,
		51751.86112841131,
		3796837.749470245,
		992013.7397791013,
		-1221952.21711287,
		1340652.697009075,
		-620943.6990984312,
		144416.9293806241,
		37.5,
	},
	[]float64{-0.0003441963504368392,
		111320.7020576856,
		278.2353980772752,
		2485758.690035394,
		6070.750963243378,
		54821.18345352118,
		9540.606633304236,
		-2710.55326746645,
		1405.483844121726,
		22.5,
	},
	[]float64{-0.0003218135878613132,
		111320.7020701615,
		0.00369383431289,
		823725.6402795718,
		0.46104986909093,
		2351.343141331292,
		1.58060784298199,
		8.77738589078284,
		0.37238884252424,
		7.45,
	},
}

// LLBAND ...
var LLBAND = []float64{75.0, 60.0, 45.0, 30.0, 15.0, 0.0}

// Position struct
type Position struct {
	Lon float64
	Lat float64
}

// IsInChina GCJ02 标准下,是否在中国
func IsInChina(coord Position) bool {
	return coord.Lon >= 72.004 && coord.Lon <= 137.8347 && coord.Lat >= 0.8293 && coord.Lat <= 55.8271
}

// delta ...
func delta(lon, lat float64) []float64 {
	dLon := transformLon(lon-105, lat-35)
	dLat := transformLat(lon-105, lat-35)
	radLat := lat / 180 * PI
	magic := math.Sin(radLat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)

	dLon = (dLon * 180) / (a / sqrtMagic * math.Cos(radLat) * PI)
	dLat = (dLat * 180) / ((a * (1 - ee)) / (magic * sqrtMagic) * PI)

	return []float64{dLon, dLat}
}

func transformLat(x, y float64) float64 {
	ret := -100 + 2*x + 3*y + 0.2*y*y + 0.1*x*y + 0.2*math.Sqrt(math.Abs(x))
	ret += (20*math.Sin(6*x*PI) + 20*math.Sin(2*x*PI)) * 2 / 3
	ret += (20*math.Sin(y*PI) + 40*math.Sin(y/3*PI)) * 2 / 3
	ret += (160*math.Sin(y/12*PI) + 320*math.Sin(y*PI/30)) * 2 / 3
	return ret
}

func transformLon(x, y float64) float64 {
	ret := 300 + x + 2*y + 0.1*x*x + 0.1*x*y + 0.1*math.Sqrt(math.Abs(x))
	ret += (20*math.Sin(6*x*PI) + 20*math.Sin(2*x*PI)) * 2 / 3
	ret += (20*math.Sin(x*PI) + 40*math.Sin(x/3*PI)) * 2 / 3
	ret += (150*math.Sin(x/12*PI) + 300*math.Sin(x/30*PI)) * 2 / 3
	return ret
}

func transform(x, y float64, factors []float64) Position {

	cc := math.Abs(y) / factors[9]
	xt := factors[0] + factors[1]*math.Abs(x)
	yt :=
		factors[2] +
			factors[3]*cc +
			factors[4]*math.Pow(cc, 2) +
			factors[5]*math.Pow(cc, 3) +
			factors[6]*math.Pow(cc, 4) +
			factors[7]*math.Pow(cc, 5) +
			factors[8]*math.Pow(cc, 6)

	if x < 0 {

		xt *= -1
	} else {
		xt *= 1
	}

	if y < 0 {

		yt *= -1
	} else {
		yt *= 1
	}

	return Position{Lon: xt, Lat: yt}
}

// GCJ02ToWGS84 ...
func GCJ02ToWGS84(coord Position) Position {

	if !IsInChina(coord) {
		return coord

	}
	lon, lat := coord.Lon, coord.Lat
	wgsLon, wgsLat := coord.Lon, coord.Lat
	tempPoint := WGS84ToGCJ02(coord)

	dx := tempPoint.Lon - lon
	dy := tempPoint.Lat - lat

	for ok := true; ok == true; ok = math.Abs(dx) > 1e-6 || math.Abs(dy) > 1e-6 {
		wgsLon -= dx
		wgsLat -= dy

		tempPoint = WGS84ToGCJ02(Position{Lon: wgsLon, Lat: wgsLat})
		dx = tempPoint.Lon - lon
		dy = tempPoint.Lat - lat
	}

	return Position{Lon: wgsLon, Lat: wgsLat}

}

// WGS84ToGCJ02 ...
func WGS84ToGCJ02(coord Position) Position {

	if !IsInChina(coord) {
		return coord

	}
	lon, lat := coord.Lon, coord.Lat

	d := delta(lon, lat)

	return Position{Lon: lon + d[0], Lat: lat + d[1]}
}

// BD09ToGCJ02 ...
func BD09ToGCJ02(coord Position) Position {

	lon, lat := coord.Lon, coord.Lat
	x := lon - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*baiduFactor)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*baiduFactor)
	newLon := z * math.Cos(theta)
	newLat := z * math.Sin(theta)
	return Position{Lon: newLon, Lat: newLat}
}

// GCJ02ToBD09 ...
func GCJ02ToBD09(coord Position) Position {

	lon, lat := coord.Lon, coord.Lat
	x := lon
	y := lat
	z := math.Sqrt(x*x+y*y) + 0.00002*math.Sin(y*baiduFactor)
	theta := math.Atan2(y, x) + 0.000003*math.Cos(x*baiduFactor)
	newLon := z*math.Cos(theta) + 0.0065
	newLat := z*math.Sin(theta) + 0.006

	return Position{Lon: newLon, Lat: newLat}
}

// BD09toBD09MC ...
func BD09toBD09MC(coord Position) Position {

	lon, lat := coord.Lon, coord.Lat
	factors := []float64{}

	for i := 0; i < len(LLBAND); i++ {
		if math.Abs(lat) > LLBAND[i] {
			factors = LL2MC[i]
			break
		}
	}
	return transform(lon, lat, factors)
}

// BD09MCtoBD09 ...
func BD09MCtoBD09(coord Position) Position {
	x, y := coord.Lon, coord.Lat
	factors := []float64{}
	for i := 0; i < len(MCBAND); i++ {
		if y >= MCBAND[i] {
			factors = MC2LL[i]
			break
		}
	}
	return transform(x, y, factors)
}

// WGS84ToBD09 ...
func WGS84ToBD09(coord Position) Position {

	return GCJ02ToBD09(WGS84ToGCJ02(coord))

}

// BD09ToWGS84 ...
func BD09ToWGS84(coord Position) Position {

	return GCJ02ToWGS84(BD09ToGCJ02(coord))
}
