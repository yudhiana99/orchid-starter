package constants

func GetStatusSelling(name string) (status int) {
	switch name {
	case "SELLING_PRODUCT_LIVE":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 99
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_ARCHIVED":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 98
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_OUT_OF_STOCK":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 97
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_BLOCKED":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 96
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_DELETED":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 94
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_WAITING":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 93
			}
			return int(status)
		}()

	case "SELLING_PRODUCT_ARCHIVED_BY_ADMIN":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 92
			}
			return int(status)
		}()
	}
	return
}
