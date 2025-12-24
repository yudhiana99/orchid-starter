package constants

func GetStatusSkuImageType(name string) (status int64) {
	switch name {
	case "PRODUCT_IMAGE_TYPE":
		return func() int64 {
			status := GetConstant(name)
			if status == 0 {
				return 1
			}
			return status
		}()
	case "SKU_IMAGE_TYPE":
		return func() int64 {
			status := GetConstant(name)
			if status == 0 {
				return 2
			}
			return status
		}()
	case "SIZE_GUIDE_IMAGE_TYPE":
		return func() int64 {
			status := GetConstant(name)
			if status == 0 {
				return 3
			}
			return status
		}()
	}
	return
}
