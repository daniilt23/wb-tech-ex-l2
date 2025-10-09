package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "столик", "слиток", "стол"}
	anagramMap := make(map[string][]string)
	firstSeenMap := make(map[string]string)
	resultMap := make(map[string][]string)
	for _, val := range input {
		lowered := strings.ToLower(val)
		cVal := []rune(lowered)
		slices.Sort(cVal)
		s := string(cVal)
		if _, exists := firstSeenMap[s]; !exists {
			firstSeenMap[s] = lowered
		}
		anagramMap[s] = append(anagramMap[s], lowered)
	}

	for key, value := range anagramMap { // для того чтобы ключом было первое значение и удалять ключ с 1 значением
		if len(value) == 1 {
			delete(anagramMap, key)
			continue
		}
		slices.Sort(value) // здесь возможна nlog(n) в случае если группа все слова мапы
		resultMap[firstSeenMap[key]] = value
	}

	fmt.Println(resultMap)
}
