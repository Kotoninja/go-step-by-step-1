package main

import (
	"fmt"
	"sort"
	"strings"
)

func getArray(text string) []string {
	words := []string{}

	var interWord string
	for _, letter := range text {
		if strings.Contains(".,!? ", string(letter)) {
			if interWord != "" {
				words = append(words, interWord)
			}
			interWord = ""
		} else {
			interWord += strings.ToLower(string(letter))
		}
	}
	if interWord != "" {
		words = append(words, interWord)
	}
	return words
}

func getMap(words []string) map[string]int {
	result := map[string]int{}

	for _, v := range words {
		result[v]++
	}

	return result
}

func uniqueWords(words map[string]int) int {
	unique := []int{}

	for _, v := range words {
		unique = append(unique, v)
	}
	return len(unique)
}

func mostFrequencyWord(words map[string]int) (string, int) {
	word := ""
	frequency := 0

	for k, v := range words {
		if v > frequency {
			word = k
			frequency = v
		}
	}

	return word, frequency
}

func getTopWords(wordMap map[string]int, n int) []string {
	frequency := []int{}

	for _, v := range wordMap {
		frequency = append(frequency, v)
	}

	sort.Ints(frequency)

	result := []string{}

	i := len(frequency) - 1
	for n > 0 {
		for k, v := range wordMap {
			if v == frequency[i] {
				result = append(result, k)
			}
		}
		n--
		i--
	}
	return result
}

func AnalyzeText(text string) {
	array := getArray(text)
	wordsMap := getMap(array)
	countWords := len(array)
	uniqueWordsCount := uniqueWords(wordsMap)
	frequencyWord, frequencyWordCount := mostFrequencyWord(wordsMap)
	topFive := getTopWords(wordsMap, 5)

	fmt.Printf("Количество слов: %d\n", countWords)
	fmt.Printf("Количество уникальных слов: %d\n", uniqueWordsCount)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", frequencyWord, frequencyWordCount)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, v := range topFive {
		fmt.Printf("\"%s\": %d раз\n", v, wordsMap[v])
	}
}

func main() {
	AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю. Море! Я море!!! ЛЮБЛЮ МОРЕ")

}
