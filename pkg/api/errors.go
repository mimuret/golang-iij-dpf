package api

func IsNotFound(err error) (bool, string) {
	return IsErrorCode(err, "not_found")
}

func IsErrorCode(err error, code string) (bool, string) {
	bad, ok := err.(*BadResponse)
	if !ok {
		return false, ""
	}
	for _, errDetail := range bad.ErrorDetails {
		if errDetail.Code == code {
			return true, errDetail.Attribute
		}
	}
	return false, ""
}

func IsAuthError(err error) bool {
	return IsErrorCodeAttribute(err, "invalid", "access_token")
}
func IsInvalidSchema(err error) bool {
	return IsErrorCodeAttribute(err, "invalid", "schema")
}

func IsErrorCodeAttribute(err error, code string, attribute string) bool {
	bad, ok := err.(*BadResponse)
	if !ok {
		return false
	}
	for _, errDetail := range bad.ErrorDetails {
		if errDetail.Code == code {
			return true
		}
	}
	return false
}
