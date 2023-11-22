package utilx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// AesEny aes对称加密
func AesEny(key, iv string, content interface{}) (string, error) {
	key = paddingAes(key)
	iv = paddingAes(iv)
	plaintext, err := json.Marshal(content)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	stream := cipher.NewCTR(block, []byte(iv))
	stream.XORKeyStream(plaintext, plaintext)
	return base64.StdEncoding.EncodeToString(plaintext), nil
}

// AesDec 解密
func AesDec(key, iv, content string, v interface{}) error {
	key = paddingAes(key)
	iv = paddingAes(iv)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	ciptext, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}
	stream := cipher.NewCTR(block, []byte(iv))
	stream.XORKeyStream(ciptext, ciptext)
	return json.Unmarshal(ciptext, v)
}

// paddingAes key iv 必须 16位
// 不足在前面补充0
// 超过裁剪
func paddingAes(s string) string {
	length := len(s)
	if length == 16 {
		return s
	}
	if length < 16 {
		return fmt.Sprintf("%016s", s)
	}
	return s[:16]
}
