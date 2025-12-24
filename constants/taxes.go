package constants

func GetStatusTax(name string) (status int64) {
	switch name {
	case "TAX_POSITIVE":
		return func() int64 {
			status := GetConstant(name)
			if status == 0 {
				return 1
			}
			return status
		}()
	case "TAX_NEGATIVE":
		return func() int64 {
			status := GetConstant(name)
			if status == 0 {
				return 2
			}
			return status
		}()
	}
	return
}
