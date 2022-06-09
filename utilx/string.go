package utilx

import "strings"

// Concat 合并字符串
func Concat(ss ...string) string {
	return strings.Join(ss, "")
}
