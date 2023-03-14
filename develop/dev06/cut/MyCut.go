package cut

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	delimiter *string
	separated *bool
	filename  *string
	fields    *string
}

type appEnv struct {
	text []string
}

func Run() error {
	var args Args
	var app appEnv

	if err := args.parseArgs(&app); err != nil {
		return err
	}
	if err := args.run(app); err != nil {
		return err
	}
	return nil
}

func (args *Args) parseArgs(app *appEnv) error {
	args.fields = flag.String("f", "", "write the column numbers separated by commas")
	args.filename = flag.String("file", "input.txt", "filename")
	args.delimiter = flag.String("d", "\t", "delimiter")
	args.separated = flag.Bool("s", false, "separated")
	flag.Parse()
	if err := app.openFile(*args.filename); err != nil {
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
		app.text = append(app.text, scanner.Text())
	}
	return nil
}

func (args *Args) run(app appEnv) error {
	for _, offer := range app.text {
		res := strings.Split(offer, *args.delimiter)
		l := len(res)
		if l == 1 && *args.separated {
			continue
		}
		if *args.fields != "" && *args.delimiter != "\t" {
			data, err := args.Fields(res, " ")
			if err != nil {
				return err
			}
			fmt.Println(data)
		} else {
			data, err := args.Fields(app.text, "\n")
			if err != nil {
				return err
			}
			fmt.Println(data)
			break
		}

	}
	return nil
}

func (args *Args) Fields(offer []string, delimiter string) (string, error) {
	var res string
	f := strings.Split(*args.fields, ",")

	for _, item := range f {
		i, err := strconv.Atoi(item)
		if err != nil {
			return "", err
		}
		if i > len(offer) {
			break
		} else {
			res += offer[i-1] + delimiter
		}
	}

	return res, nil
}
