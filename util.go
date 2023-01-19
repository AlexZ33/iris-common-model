package iris_common_model

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	"log"
	"strconv"
	"strings"
	"time"
)

func ParseString(value interface{}, values ...string) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float64:
		return strconv.FormatFloat(value.(float64), "f", -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	case []string:
		return strings.Join(value.([]string), ",")
	case []byte:
		return string(value.([]byte))
	case time.Time:
		return StringifyTime(value.(time.Time))
	case []int64:
		numbers := make([]string, 0)
		for _, number := range value.([]int64) {
			numbers = append(numbers, strconv.FormatInt(number, 10))
		}
		return strings.Join(numbers, ",")
	case []uint64:
		numbers := make([]string, 0)
		for _, number := range value.([]uint64) {
			numbers = append(numbers, strconv.FormatUint(number, 10))
		}
		return strings.Join(numbers, ",")
	case []float64:
		numbers := make([]string, 0)
		for _, number := range value.([]float64) {
			numbers = append(numbers, strconv.FormatFloat(number, "f", -1, 64))
		}
		return strings.Join(numbers, ",")
	case []interface{}:
		values := make([]string, 0)
		for _, str := range value.([]interface{}) {
			values = append(values, ParseString(str))
		}
		return "[" + strings.Join(values, ",") + "]"
	default:
		if value != nil {
			return string(GetJSON(value))
		}
		if len(values) > 0 {
			return values[0]
		}
		return ""
	}
}

func GetJSON(value interface{}) []byte {
	if value == nil {
		return make([]byte, 0)
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
	}
	return bytes
}

func StringifyTime(t time.Time) string {
	layout := "2006-01-02T15:04:05.999999-07:00"
	return t.Format(layout)
}

func Id() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
	}
	return id.String()
}

func CheckId(s string) bool {
	if _, err := uuid.Parse(s); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Key(phrase string) string {
	return HS256(phrase, Base64Encode(phrase))
}

func AccessKey() string {
	iv := make([]byte, sha1.Size)
	if _, err := rand.Read(iv); err != nil {
		log.Println(err)
	}
	return Base64Encode(string(iv))
}

func GetTree(tree *toml.Tree, key string) *toml.Tree {
	if value, ok := tree.Get(key).(*toml.Tree); ok {
		return value
	}
	return new(toml.Tree)
}

func GetString(tree *toml.Tree, key string, values ...string) string {
	value := tree.Get(key)
	if value != nil {
		return ParseString(value)
	} else if len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetFloat64(tree *toml.Tree, key string, values ...float64) float64 {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case float64:
			return value.(float64)
		case int64:
			return float64(value.(int64))
		case uint64:
			return float64(value.(uint64))
		case string:
			value, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				log.Println(err)
			} else {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0.0
}

func GetInt(tree *toml.Tree, key string, values ...int) int {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case int64:
			return int(value.(int64))
		case uint64:
			return int(value.(uint64))
		case float64:
			return int(value.(float64))
		case string:
			value, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				return int(value)
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetInt64(tree *toml.Tree, key string, values ...int64) int64 {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case int64:
			return value.(int64)
		case uint64:
			return int64(value.(uint64))
		case float64:
			return int64(value.(float64))
		case string:
			value, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetUint64(tree *toml.Tree, key string, values ...uint64) uint64 {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case uint64:
			return value.(uint64)
		case int64:
			return uint64(value.(int64))
		case float64:
			return uint64(value.(float64))
		case string:
			value, err := strconv.ParseUint(value.(string), 10, 64)
			if err != nil {
				log.Println(err)
			} else {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetBool(tree *toml.Tree, key string, values ...bool) bool {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case bool:
			return value.(bool)
		case string:
			value, err := strconv.ParseBool(value.(string))
			if err != nil {
				log.Println(err)
			} else {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return false
}

func GetDuration(tree *toml.Tree, key string, values ...time.Duration) time.Duration {
	value := tree.Get(key)
	if value != nil {
		switch value.(type) {
		case string:
			duration, err := time.ParseDuration(value.(string))
			if err != nil {
				log.Println(err)
			} else {
				return duration
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0 * time.Second
}

func GetStringArray(tree *toml.Tree, key string, values ...[]string) []string {
	strings := make([]string, 0)
	if array, ok := tree.Get(key).([]interface{}); ok {
		for _, value := range array {
			strings = append(strings, ParseString(value))
		}
	}
	if len(strings) == 0 && len(values) > 0 {
		return values[0]
	}
	return strings
}
