package main

func CountingSort(contacts []string) map[string]int {
	frequency := make(map[string]int)

	for _, contact := range contacts {
		frequency[contact]++
	}
	return frequency
}
