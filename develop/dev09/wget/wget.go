package wget

import (
	"flag"
	"log"
	"os"
	"strings"
)

type Args struct {
	path *string
	url  *string
}

func (args *Args) parsingArgs() {
	args.url = flag.String("u", "", "path to source")
	args.path = flag.String("p", "", "path to source")

	flag.Parse()
}

func (args *Args) mkDir() error {
	if *args.path == "" {
		args.path = &strings.Split(*args.url, "//")[1]
	}
	if err := os.MkdirAll(*args.path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (args *Args) run() error {
	args.parsingArgs()
	if err := args.mkDir(); err != nil {
		return err
	}

	html, err := getHtml(*args.url)
	if err != nil {
		return err
	}
	if err := saveInFile(html, *args.path+"/index.html"); err != nil {
		return err
	}
	return nil
}

func deleteFromSlice(s []string, elem string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == elem {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

func Run() {
	var args Args
	if err := args.run(); err != nil {
		log.Println(err)
		return
	}
}
