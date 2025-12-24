package common

import (
	"os"
	"strconv"

	"crypto/rand"
	"encoding/base64"
)

func RandomStringMixedNumber(length int) (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b)[:length], nil
}

func IsSuccess(status int) bool {
	return status >= 200 && status < 300
}

func GetIntEnv(env string, fallback ...int) int {
	valueEnv := os.Getenv(env)
	fallbackValue := int(0)
	if len(fallback) > 0 {
		fallbackValue = fallback[0]
	}

	if valueEnv == "" {
		return fallbackValue
	}

	result, err := strconv.Atoi(valueEnv)
	if err != nil {
		return fallbackValue
	}
	return result
}

func GetInt64Env(env string, fallback ...int64) int64 {
	valueEnv := os.Getenv(env)
	fallbackValue := int64(0)
	if len(fallback) > 0 {
		fallbackValue = fallback[0]
	}

	if valueEnv == "" {
		return fallbackValue
	}

	result, err := ConvertStringToInt64(valueEnv)
	if err != nil {
		return fallbackValue
	}
	return result
}

func GetUint64Env(env string, fallback ...uint64) uint64 {
	valueEnv := os.Getenv(env)
	fallbackValue := uint64(0)
	if len(fallback) > 0 {
		fallbackValue = fallback[0]
	}

	if valueEnv == "" {
		return fallbackValue
	}

	result, err := ConvertStringToUint64(valueEnv)
	if err != nil {
		return fallbackValue
	}
	return result
}

func GetBoolEnv(env string, fallback ...bool) bool {
	valueEnv := os.Getenv(env)
	fallbackValue := false
	if len(fallback) > 0 {
		fallbackValue = fallback[0]
	}

	if valueEnv == "" {
		return fallbackValue
	}

	result, err := strconv.ParseBool(valueEnv)
	if err != nil {
		return fallbackValue
	}
	return result
}

func GetEnvWithDefault(env string, fallback ...string) string {
	valueEnv := os.Getenv(env)
	if valueEnv == "" {
		if len(fallback) > 0 {
			valueEnv = fallback[0]
		}
	}
	return valueEnv
}
