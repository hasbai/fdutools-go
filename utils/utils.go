package utils

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	OptIgnore    = "-"
	OptOmitempty = "omitempty"
	OptDive      = "dive"
)

const (
	flagIgnore = 1 << iota
	flagOmiEmpty
	flagDive
)

func readTag(f reflect.StructField, tag string) (string, int) {
	val, ok := f.Tag.Lookup(tag)
	fieldTag := ""
	flag := 0

	// no tag, use field name
	if !ok {
		return f.Name, flag
	}
	opts := strings.Split(val, ",")

	fieldTag = opts[0]
	for i := 1; i < len(opts); i++ {
		switch opts[i] {
		case OptIgnore:
			flag |= flagIgnore
		case OptOmitempty:
			flag |= flagOmiEmpty
		case OptDive:
			flag |= flagDive
		}
	}
	return fieldTag, flag
}

func StructToMap(obj any) map[string]string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		log.Panicf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	t := reflect.TypeOf(obj)
	m := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("json")
		if tag == "" {
			continue
		}
		fieldValue := v.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			m[tag] = fieldValue.String()
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			m[tag] = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			m[tag] = strconv.FormatUint(fieldValue.Uint(), 10)
		default:
			m[tag] = fmt.Sprintf("%v", fieldValue)
		}
	}
	return m
}

func ReadBody(body io.ReadCloser) string {
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}(body)
	data, err := io.ReadAll(body)
	if err != nil {
		log.Panicln("read body error", err)
		return ""
	}
	return string(data)
}

func UnmarshalResponse(body io.ReadCloser, v any) error {
	//goland:noinspection GoUnhandledErrorResult
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

var pattern = regexp.MustCompile(`([a-zA-Z]+?):`)

func NormalizeJSON(json string) string {
	json = pattern.ReplaceAllString(json, "\"$1\":")
	json = strings.ReplaceAll(json, "'", "\"")
	return json
}
