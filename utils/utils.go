package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ToString(raw any) (string, error) {
	if raw == nil {
		return "", errors.New("nil value")
	}

	// 检查是否实现了fmt.Stringer接口
	if s, ok := raw.(fmt.Stringer); ok {
		val := reflect.ValueOf(raw)
		if val.Kind() == reflect.Ptr && val.IsNil() {
			return "", errors.New("nil Stringer pointer")
		}
		return s.String(), nil
	}

	val := reflect.ValueOf(raw)

	// 解引用指针直到非指针类型
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "", errors.New("nil pointer")
		}
		val = val.Elem()
	}

	// 再次检查解引用后的值是否实现了Stringer接口
	if s, ok := val.Interface().(fmt.Stringer); ok {
		return s.String(), nil
	}

	// 处理基础类型
	switch val.Kind() {
	case reflect.String:
		return val.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(val.Bool()), nil
	default:
		return "", fmt.Errorf("unsupported type: %s", val.Type())
	}
}
