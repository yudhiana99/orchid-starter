package constants

func GetStatusPkp(name string) (status int) {
	switch name {
	case "NON_PKP":
		return func() int {
			status := GetConstant(name)
			return int(status)
		}()
	case "PKP":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 1
			}
			return int(status)
		}()
	}
	return
}

func GetStatusProductPkp(name string) (status int) {
	switch name {
	case "PRODUCT_PKP_TAX_INPROGRESS":
		status := GetConstant(name)
		if status == 0 {
			return 1
		}
		return int(status)
	case "PRODUCT_PKP_TAX_FAILED":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 2
			}
			return int(status)
		}()
	case "PRODUCT_PKP_TAX_HALF_COMPLETED":
		status := GetConstant(name)
		if status == 0 {
			return 3
		}
		return int(status)
	case "PRODUCT_PKP_TAX_HALF_FAILED":
		status := GetConstant(name)
		if status == 0 {
			return 4
		}
		return int(status)
	case "PRODUCT_PKP_TAX_COMPLETED":
		status := GetConstant(name)
		if status == 0 {
			return 5
		}
		return int(status)
	}
	return
}
