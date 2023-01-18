package iris_common_model

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"log"
	"strings"
)

func Base64Encode(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}
func Base64Decode(str string) string {
	base64Encoding := base64.RawURLEncoding
	if strings.ContainsAny(str, "+/") {
		base64Encoding = base64.RawStdEncoding
	}
	if strings.HasSuffix(str, "=") {
		str = strings.TrimRight(str, "=")
	}
	data, err := base64Encoding.DecodeString(str)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func Token(header string, signature string) string {
	return Base64Encode(header) + "." + signature
}

func HexEncode(str string) string {
	return hex.EncodeToString([]byte(str))
}

func HexDecode(str string) string {
	data, err := hex.DecodeString(str)
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func HS512(str string, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(str))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func HS256(str string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
