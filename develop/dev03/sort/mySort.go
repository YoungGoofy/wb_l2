package sorting

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

type flagEnv struct {
	isNumeric        *bool
	deleteDuplicates *bool
	isReverse        *bool
	filename         *string
}

type appEnv struct {
	data []string
}

func Run() error {
	var fl flagEnv
	var app appEnv
	if err := fl.parseArgs(&app); err != nil {
		return err
	}
	fl.sort(&app)
	outputData(app.data)
	return nil
}

func (fl flagEnv) sort(d *appEnv) {
	if *fl.deleteDuplicates {
		d.data = deleteDuplicates(d.data)
		fl.util(*d)
	} else {
		fl.util(*d)
	}
}

func (fl flagEnv) util(d appEnv) {
	if *fl.isReverse {
		d.reverseSort()
	} else {
		d.sortStrings()
	}
}
func (fl *flagEnv) parseArgs(d *appEnv) error {
	fl.filename = flag.String("f", "files/inNums.txt", "Filename to sort")
	fl.isNumeric = flag.Bool("n", false, "sort numerically")
	fl.isReverse = flag.Bool("r", false, "sort in reverse")
	fl.deleteDuplicates = flag.Bool("u", false, "delete duplicates")
	flag.Parse()

	if err := d.openFile(*fl.filename); err != nil {
		return err
	}
	return nil
}

func (app *appEnv) openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		app.data = append(app.data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (app *appEnv) sortStrings() {
	sort.Strings(app.data)
}

func (app *appEnv) reverseSort() {
	sort.Sort(sort.Reverse(sort.StringSlice(app.data)))
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

func outputData(d []string) {
	for _, item := range d {
		fmt.Println(item)
	}
}
