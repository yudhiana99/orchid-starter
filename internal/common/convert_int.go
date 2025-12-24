package common

import (
	"fmt"
	"math"
	"strconv"
)

// ================= int8

// ConvertInt64ToInt8 is a helper function that converts an int64 to an int8.
// It will return an error if the int64 value is not within the range of int8.
func ConvertInt64ToInt8(int64Value int64) (int8, error) {
	if int64Value >= math.MinInt8 && int64Value <= math.MaxInt8 {
		// int64Value is within the range of int8
		return int8(int64Value), nil
	}
	// int64Value is out of range of int8
	return 0, fmt.Errorf("%v is out of range int8", int64Value)
}

// ConvertIntToInt8 is a helper function that converts an int to an int8.
// It will return an error if the int value is not within the range of int8.
func ConvertIntToInt8(intValue int) (int8, error) {
	if intValue >= math.MinInt8 && intValue <= math.MaxInt8 {
		return int8(intValue), nil
	}
	return 0, fmt.Errorf("%v is out of range int8", intValue)
}

// ================= uint8
func ConvertInt8ToUint8(int8Value int8) (uint8, error) {
	if int8Value >= 0 {
		return uint8(int8Value), nil
	}
	return 0, fmt.Errorf("%v is out of range uint8", int8Value)
}

func ConvertIntToUint8(intValue int) (uint8, error) {
	if intValue >= 0 && intValue <= math.MaxUint8 {
		return uint8(intValue), nil
	}
	return 0, fmt.Errorf("%v is out of range uint8", intValue)
}

func ConvertInt64ToUint8(intValue int64) (uint8, error) {
	if intValue >= 0 && intValue <= math.MaxUint8 {
		return uint8(intValue), nil
	}
	return 0, fmt.Errorf("%v is out of range uint8", intValue)
}

// ================= int

func ConvertInt8ToInt(int8Value int8) (int, error) {
	return int(int8Value), nil
}

// ConvertInt16ToInt is a helper function that converts an int16 to an int.
// It will return an error if the int16 value is not within the range of int.
func ConvertInt16ToInt(int16Value int16) (int, error) {
	return int(int16Value), nil
}

// ConvertInt32ToInt is a helper function that converts an int32 to an int.
// It will return an error if the int32 value is not within the range of int.
func ConvertInt32ToInt(int32Value int32) (int, error) {
	return int(int32Value), nil
}

func ConvertInt64ToInt(int64Value int64) (int, error) {
	if int64Value >= math.MinInt && int64Value <= math.MaxInt {
		// int64Value is within the range of int
		return int(int64Value), nil
	}
	// int64Value is out of range of int
	return 0, fmt.Errorf("%v is out of range int", int64Value)
}

func ConvertUintToInt(uintValue uint) (int, error) {
	if uintValue > uint(math.MaxInt) {
		return 0, fmt.Errorf("%v is out of range int", uintValue)
	}
	return int(uintValue), nil
}

// ConvertUint8ToInt is a helper function that converts a uint8 to an int.
func ConvertUint8ToInt(uint8Value uint8) (int, error) {
	return int(uint8Value), nil
}

// ConvertUint16ToInt is a helper function that converts a uint16 to an int.
func ConvertUint16ToInt(uint16Value uint16) (int, error) {
	return int(uint16Value), nil
}

// ConvertUint32ToInt is a helper function that converts a uint32 to an int.
func ConvertUint32ToInt(uint32Value uint32) (int, error) {
	if uint32Value <= math.MaxInt32 {
		return int(uint32Value), nil
	}
	return 0, fmt.Errorf("%v is out of range int", uint32Value)
}

// ConvertUint64ToInt is a helper function that converts a uint64 to an int.
func ConvertUint64ToInt(uint64Value uint64) (int, error) {
	if uint64Value <= math.MaxInt {
		return int(uint64Value), nil
	}
	return 0, fmt.Errorf("%v is out of range int", uint64Value)
}

// ConvertFloat32ToInt is a helper function that converts a float32 to an int.
func ConvertFloat32ToInt(float32Value float32) (int, error) {
	if (float32Value >= math.MinInt) && (float32Value <= math.MaxInt) {
		return int(float32Value), nil
	}
	return 0, fmt.Errorf("%v is out of range int", float32Value)
}

// ConvertFloat64ToInt is a helper function that converts a float64 to an int.
func ConvertFloat64ToInt(float64Value float64) (int, error) {
	if float64Value >= math.MinInt && float64Value <= math.MaxInt {
		return int(float64Value), nil
	}
	return 0, fmt.Errorf("%v is out of range int", float64Value)
}

// ================= uint64
func ConvertInt64ToUint64(int64Value int64) (uint64, error) {
	if int64Value >= 0 {
		// uint64Value is very large number which is larger than maximum int64 value
		return uint64(int64Value), nil
	}
	// int64Value is out of range of uint64
	return 0, fmt.Errorf("%v is out of range uint64", int64Value)
}
func ConvertIntToUint64(value int) (uint64, error) {
	if value >= 0 {
		// uint64Value is very large number which is larger than maximum int64 value
		return uint64(value), nil
	}
	// int64Value is out of range of uint64
	return 0, fmt.Errorf("%v is out of range uint64", value)
}

func ConvertStringToUint64(strValue string) (uint64, error) {
	uint64Value, err := strconv.ParseUint(strValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64Value, nil
}

// ================= int64
func ConvertUInt64ToInt64(uint64Value uint64) (int64, error) {
	if uint64Value > math.MaxInt64 {
		return 0, fmt.Errorf("%v is out of range int64", uint64Value)
	}
	return int64(uint64Value), nil
}

func ConvertStringToInt64(strValue string) (int64, error) {
	int64Value, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64Value, nil
}
