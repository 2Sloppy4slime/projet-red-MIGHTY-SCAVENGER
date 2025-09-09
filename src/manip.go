package main

func IsLower(s string) bool {
	for _, value := range s {
		if value < 97 || value > 122 {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, value := range s {
		if value < 65 || value > 90 {
			return false
		}
	}
	return true
}

func ToLower(s string) string {
	str := ""
	for _, value := range s {
		if value < 65 || value > 90 {
			str += string(value)
		} else {
			str += string(value + 32)
		}
	}
	return str
}
func IsAlpha(s string) bool {
	for _, value := range s {
		if value < 65 || value > 90 {
			if value < 97 || value > 122 {
				if value < 48 || value > 57 {
					return false
				}
			}
		}
	}
	return true
}

func Capitalize(s string) string {
	capitalizenext := true
	res := ""
	for _, value := range s {
		if !IsAlpha(string(value)) {
			capitalizenext = true
			res += string(value)
		} else if capitalizenext && IsLower(string(value)) && !IsNumeric(string(value)) {
			res += string(value - 32)
			capitalizenext = false
		} else if !capitalizenext && IsUpper(string(value)) && !IsNumeric(string(value)) {
			res += string(value + 32)
			capitalizenext = false
		} else {
			res += string(value)
			capitalizenext = false
		}
	}
	return res
}
func IsNumeric(s string) bool {
	for _, value := range s {
		if value < 48 || value > 57 {
			return false
		}
	}
	return true
}
