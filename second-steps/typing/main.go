package typing

/* verify a variable type */
func CheckType(variable any) string {
	switch variable.(type) {
	case int:
		return "Integer"
	case float64:
		return "Float"
	case string:
		return "String"
	default:
		return "Unknown"
	}
}
