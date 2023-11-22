package utilx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/spf13/cast"
)

// Base64Encode base64编码
func Base64Encode(v interface{}) string {
	var val []byte
	switch d := v.(type) {
	case []byte:
		val = d
	case string:
		val = []byte(d)
	default:
		val, _ = json.Marshal(d)
	}
	return base64.StdEncoding.EncodeToString(val)
}

// Base64Decode 解码
func Base64Decode(src string) string {
	b, _ := base64.StdEncoding.DecodeString(src)
	return string(b)
}

// Base64KeyEncode 主要对 OTS的NextNextStartPrimaryKey 进行base64编码
func Base64KeyEncode(v interface{}) string {
	var val []byte
	switch d := v.(type) {
	case []byte:
		val = d
	case string:
		val = []byte(d)
	default:
		val, _ = json.Marshal(d)
	}

	return base64.StdEncoding.EncodeToString(val)
}

// Base64KeyDecode 解码成*tablestore.PrimaryKey
func Base64KeyDecode(src string) *tablestore.PrimaryKey {
	b, _ := base64.StdEncoding.DecodeString(src)
	data := &tablestore.PrimaryKey{}
	d := json.NewDecoder(bytes.NewReader(b))
	// 精度
	d.UseNumber()
	d.Decode(data)
	for i, v := range data.PrimaryKeys {
		switch d := v.Value.(type) {
		case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64, float32, float64:
			data.PrimaryKeys[i].Value = cast.ToInt64(d)
		case json.Number:
			n, _ := d.Int64()
			data.PrimaryKeys[i].Value = n
		case *json.Number:
			n, _ := d.Int64()
			data.PrimaryKeys[i].Value = n
		}
	}

	return data
}
