package math

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	dataArr := strings.Split("5A A5 11 01 01 00 10 10 01 82 10 01 11 01 00 11 01", " ")
	//tempArr := []
	for i := 1; i < 9; i++ {
		j := i * 3
		data := fmt.Sprintf("%v%v%v", dataArr[j], dataArr[j+1], dataArr[j+2])
		//fmt.Println(data)
		vadStr, _ := ConvertInt(data, 16, 10)
		vad, _ := strconv.Atoi(vadStr)
		a := float64(vad * 1000.0)
		b := math.Pow(2, 24)
		c := b - float64(vad)
		rpt := Round(a/c, 3)
		if 185.201 <= rpt && rpt <= 4358.275 {

		}

	}
}
