package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	slice, err := readFile("test.txt")
	if err != nil {
		log.Println(err)
		return
	}
	var uniqueSlice []string
	sort.Strings(slice)
	for _, item := range slice {
		uniqueSlice = findUnique(uniqueSlice, item)
	}
	printSlice(slice)
	fmt.Println()
	printSlice(uniqueSlice)
}

func readFile(filename string) ([]string, error) {
	var fileSlice []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileSlice = append(fileSlice, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error while reading file: %s", err))
	}
	return fileSlice, nil
}

func printSlice(slice []string) {
	for _, item := range slice {
		fmt.Println(item)
	}
}

func findUnique(slice []string, value string) []string {
	for _, item := range slice {
		if item == value {
			return slice
		}
	}
	return append(slice, value)
}
