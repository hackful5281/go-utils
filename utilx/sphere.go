package utilx

import (
	"math"
	"strconv"
	"strings"
)

// R 地球半径，单位米
const R = 6378137.0

// Sphere 计算两地间距
// latA, lngA分别为A点的纬度和经度
// 返回的距离单位为米
func Sphere(latA, lngA, latB, lngB float64) float64 {
	rad := math.Pi / 180.0
	latA = latA * rad
	lngA = lngA * rad
	latB = latB * rad
	lngB = lngB * rad
	theta := lngB - lngA
	dist := math.Acos(math.Sin(latA)*math.Sin(latB) + math.Cos(latA)*math.Cos(latB)*math.Cos(theta))
	return dist * R
}

func FmtLonLat(sA, sB string) (float64, float64, float64, float64) {
	scliceA := strings.Split(sA, ",")
	scliceB := strings.Split(sB, ",")
	lonA, _ := strconv.ParseFloat(scliceA[0], 64)
	latA, _ := strconv.ParseFloat(scliceA[1], 64)
	lonB, _ := strconv.ParseFloat(scliceB[0], 64)
	latB, _ := strconv.ParseFloat(scliceB[1], 64)

	return latA, lonA, latB, lonB
}

// GetApart 获取两点间距
// return: -1    前一位人没有定位信息
// return: -2    后一位人没有定位信息
// return: 正常距离，单位 米
func GetApart(sA, sB string) float64 {
	if sA == "" {
		return -1
	} else if sB == "" {
		return -2
	} else if sA == sB {
		// 校验是否为同一地点，防止无穷大 NaN
		return 0
	} else {
		l := Sphere(FmtLonLat(sA, sB))
		if math.IsNaN(l) {
			l = 0
		}
		return l
	}
}
