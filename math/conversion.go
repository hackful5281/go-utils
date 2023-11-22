package math

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

/**
二进制 	Binary 		BIN	[0-1]			%b
八进制 	Octal 		OCT [0-7]			%o		以 0 开头
十进制 	Decimal 	DEC [0-9]			%d
十六进制 Hexadecimal HEX [0-9] + [A-F]	%x|%X	以 0x｜0X 开头
*/

// ConvertInt 进制转换
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func Str2Hex(str string) []byte {

	//_str := strings.TrimSpace(message)
	//str := strings.ReplaceAll(_str, "\\u000", "")

	// 拆分字符串
	strArr := strings.Split(str, " ")

	// 转换为字节数组
	var bytes []byte
	for _, s := range strArr {
		// 将十六进制字符串转换为十进制数值
		num, err := strconv.ParseUint(s, 16, 8)
		if err != nil {
			log.Println("Str2Hex")
			return nil
		}

		// 将十进制数值转换为字节类型
		bytes = append(bytes, byte(num))
	}
	return bytes
}

func Hex2Str(data []byte) string {
	var result string
	for _, b := range data {
		result += fmt.Sprintf("%02x ", b)
	}
	return strings.TrimSpace(result)
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
