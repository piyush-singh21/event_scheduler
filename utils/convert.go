package auth

func Convert(val interface{}) int {
	if v, ok := val.(int); ok {
		return v
	}
	return 0
}
