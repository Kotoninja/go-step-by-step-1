package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	result := make(map[string]string)

	for key, value := range m {
		result[value] = key
	}
	return result
}
