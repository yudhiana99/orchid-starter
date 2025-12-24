package common

import (
	"fmt"
	"strconv"

	"encoding/json"
)

func ConvertToIntFromAny(v any) (data int, err error) {
	switch value := v.(type) {
	case int:
		data = value
	case int8:
		data, err = ConvertInt8ToInt(value)
	case int16:
		data, err = ConvertInt16ToInt(value)
	case int32:
		data, err = ConvertInt32ToInt(value)
	case int64:
		data, err = ConvertInt64ToInt(value)
	case uint:
		data, err = ConvertUintToInt(value)
	case uint8:
		data, err = ConvertUint8ToInt(value)
	case uint16:
		data, err = ConvertUint16ToInt(value)
	case uint32:
		data, err = ConvertUint32ToInt(value)
	case uint64:
		data, err = ConvertUint64ToInt(value)
	case json.Number:
		data, err = strconv.Atoi(value.String())
	case string:
		data, err = strconv.Atoi(value)
	case float32:
		data, err = ConvertFloat32ToInt(value)
	case float64:
		data, err = ConvertFloat64ToInt(value)
	default:
		err = fmt.Errorf("can't convert to int caused unknown type")
	}
	return
}
