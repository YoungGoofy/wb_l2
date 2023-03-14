package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "краска", "сакрал"}
	anagramSets := *Anagram(&words)
	fmt.Println(anagramSets)
}

func Anagram(slice *[]string) *map[string][]string {
	result := make(map[string][]string)
	*slice = deleteDuplicates(*slice)
	for _, item := range *slice {
		sortedItem := sortString(item)
		_, ok := result[sortedItem]
		if ok {
			result[sortedItem] = append(result[sortedItem], item)
		} else {
			result[sortedItem] = []string{item}
		}
	}

	for key, value := range result {
		if len(value) == 1 {
			delete(result, key)
		} else {
			sort.Strings(value)
		}
	}
	return &result
}

func sortString(item string) string {
	s := strings.Split(item, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func deleteDuplicates(strSlice []string) []string {
	sort.Strings(strSlice)
	var newSlice []string
	for _, str := range strSlice {
		newSlice = withoutDubles(newSlice, str)
	}
	return newSlice
}

func withoutDubles(str []string, item string) []string {
	for _, s := range str {
		if s == item {
			return str
		}
	}
	return append(str, item)
}
