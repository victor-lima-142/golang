package basic

// HashMap demonstrates the use of Go's built-in map type.
func HashMap() {
	var hashmap map[string]string = make(map[string]string)
	hashmap["key1"] = "value1"
	hashmap["key2"] = "value2"
	hashmap["key3"] = "value3"
	hashmap["key4"] = "value4"

	for key, value := range hashmap {
		println(key, value)
	}
}
