package tools

func Empty(v interface{}) bool {
	switch v.(type) {
	case string:
		if len(v.(string)) <= 0 {
			return true
		} else {
			return false
		}
	case int:
		if v.(int) == 0 {
			return true
		} else {
			return false
		}
	case int32:
		if v.(int32) == 0 {
			return true
		} else {
			return false
		}
	case int64:
		if v.(int64) == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
