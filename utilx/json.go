package utilx

import "encoding/json"

// ToJson 转json
func ToJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func ToJsonBytes(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

// ToObj json转对象
func ToObj(b string, v interface{}) error {
	return json.Unmarshal([]byte(b), v)
}

func BytesToObj(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
