package utilx

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5加密
// salt 盐值
func Md5Encode(val string, salt ...string) string {
	if val == "" {
		return ""
	}
	buf := bytes.NewBufferString(val)
	buf.WriteString(strings.Join(salt, ""))

	h := md5.New()
	h.Write(buf.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}
