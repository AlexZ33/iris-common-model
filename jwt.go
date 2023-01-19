package iris_common_model

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/request"
	"github.com/pelletier/go-toml"
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

func Token(header string, signature string) string {
	return Base64Encode(header) + "." + signature
}

func MD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

func Sha256(str string) string {
	sum := sha256.Sum256([]byte(str))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func Sha384(str string) string {
	sum := sha512.Sum384([]byte(str))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func Sha512(str string) string {
	sum := sha512.Sum512([]byte(str))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func HS256(str string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func HS384(str string, key string) string {
	mac := hmac.New(sha512.New384, []byte(key))
	mac.Write([]byte(str))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func HS512(str string, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(str))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func Hash(str string, config *toml.Tree, flag bool) string {
	hash := strings.ToUpper(GetString(config, "hash", "SHA256"))
	if flag && GetBool(config, "use-hmac-hash") {
		key := GetString(config, "key")
		if key != "" {
			switch hash {
			case "SHA256":
				return HS256(str, key)
			case "SHA384":
				return HS384(str, key)
			case "SHA512":
				return HS512(str, key)
			default:
				return HS256(str, key)
			}
		}
	}
	switch hash {
	case "SHA256":
		return Sha256(str)
	case "SHA384":
		return Sha384(str)
	case "SHA512":
		return Sha512(str)
	default:
		return Sha256(str)
	}
}

func Sign256(claims jwt.Claims, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}
	return str
}

func Sign384(claims jwt.Claims, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}
	return str
}

func Sign512(claims jwt.MapClaims, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
	}
	return str
}

func SignClaims(claims jwt.MapClaims, config *toml.Tree, key ...string) string {
	secret := ""
	if GetBool(config, "use-global-key") {
		secret = GetString(config, "key")
	} else if len(key) > 0 {
		secret = key[0]
	}
	if claims["exp"] == nil && config.Has("max-age") {
		maxAge := GetDuration(config, "max-age")
		claims["exp"] = time.Now().Add(maxAge).Unix()
	}
	if claims["iat"] == nil {
		claims["iat"] = time.Now().Unix()
	}
	if claims["nbf"] == nil && config.Has("lockup-period") {
		period := GetDuration(config, "lockup-period")
		claims["nbf"] = time.Now().Add(period).Unix()
	}
	hash := strings.ToUpper(GetString(config, "hash", "HS256"))
	switch hash {
	case "HS256":
		return Sign256(claims, secret)
	case "HS384":
		return Sign384(claims, secret)
	case "HS512":
		return Sign512(claims, secret)
	default:
		return Sign256(claims, secret)
	}
}

func ExtractToken(req *http.Request, config *toml.Tree) string {
	extractor := request.MultiExtractor{}
	if config.Has("request-arguments") {
		argumentExtractor := request.ArgumentExtractor{}
		arguments := GetStringArray(config, "request-arguments")
		argumentExtractor = append(argumentExtractor, arguments...)
		extractor = append(extractor, argumentExtractor)
	}
	if config.Has("request-headers") {
		headerExtractor := request.HeaderExtractor{}
		headers := GetStringArray(config, "request-headers")
		headerExtractor = append(headerExtractor, headers...)
		extractor = append(extractor, headerExtractor)
	}
	extractor = append(extractor, request.AuthorizationHeaderExtractor)
	token, err := extractor.ExtractToken(req)
	if err != nil {
		log.Println(err)
	}
	if token == "" {
		cookieName := GetString(config, "cookie-name")
		if cookieName != "" {
			if cookie, err := req.Cookie(cookieName); err != nil {
				log.Println(err)
			} else {
				token = cookie.Value
			}
		}
	}
	return token
}

func ParseToken(str string, key string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Println(err)
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return jwt.MapClaims{}, false
}

func GetClaims(str string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	parts := strings.Split(str, ".")
	if len(parts) == 3 {
		if data, err := jwt.DecodeSegment(parts[1]); err != nil {
			log.Println(err)
		} else {
			dec := json.NewDecoder(bytes.NewBuffer(data))
			if err := dec.Decode(&claims); err != nil {
				log.Println(err)
			}
		}
	}
	return claims
}
