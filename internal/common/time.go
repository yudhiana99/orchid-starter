package common

import (
	"fmt"
	"time"
)

func GetIndonesianTimeZone(timeZoneCode string) (*time.Location, error) {
	switch timeZoneCode {
	case "WIB":
		return time.LoadLocation("Asia/Jakarta")
	case "WITA":
		return time.LoadLocation("Asia/Makassar")
	case "WIT":
		return time.LoadLocation("Asia/Jayapura")
	default:
		return nil, fmt.Errorf("unknown Indonesian time zone: %s", timeZoneCode)
	}
}
