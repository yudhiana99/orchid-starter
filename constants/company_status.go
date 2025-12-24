package constants

func GetCompanyStatus(name string) (status int) {
	switch name {
	case "TYPE_CUSTOMER":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 1
			}
			return int(status)
		}()
	case "TYPE_VENDOR":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 2
			}
			return int(status)
		}()
	case "TYPE_BOTH":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 3
			}
			return int(status)
		}()
	}
	return
}

func GetActiveStatus(name string) (status int) {
	switch name {
	case "STATUS_ACTIVE":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 99
			}
			return int(status)
		}()
	case "STATUS_INACTIVE":
		return func() int {
			status := GetConstant(name)
			if status == 0 {
				return 98
			}
			return int(status)
		}()
	}
	return
}
