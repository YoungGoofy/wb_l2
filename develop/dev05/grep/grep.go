package grep

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Args struct {
	after      *int
	before     *int
	context    *int
	count      *int
	ignoreCase *bool
	invert     *bool
	fixed      *bool
	lineNum    *bool
}

type appEnv struct {
	data    []string
	pattern string
}

type Container struct {
	args Args
	app  appEnv
}

func Run() {
	var c Container
	if err := c.run(); err != nil {
		log.Println(err)
	}
}

func (c *Container) run() error {
	err := c.parseArgs()
	if err != nil {
		return err
	}
	c.findString(*regexp.MustCompile(c.app.pattern))

	return nil
}

func (c *Container) parseArgs() error {
	c.args.after = flag.Int("A", 0, "Print N lines after match")
	c.args.before = flag.Int("B", 0, "Print N lines before match")
	c.args.context = flag.Int("C", 0, "Print N lines around match")
	c.args.count = flag.Int("c", -1, "Print count of matching lines")
	c.args.ignoreCase = flag.Bool("i", false, "Ignore case")
	c.args.invert = flag.Bool("v", false, "Invert match")
	c.args.fixed = flag.Bool("F", false, "Fixed string match")
	c.args.lineNum = flag.Bool("n", false, "Print line number")

	flag.Parse()

	a := flag.Args()
	if len(a) < 2 {
		return fmt.Errorf("usage: go-grep [OPTIONS] PATTERN FILE")
	}

	c.app.pattern = a[0]
	filename := a[1]

	if err := c.openFile(filename); err != nil {
		return err
	}

	if *c.args.after == 0 {
		c.args.after = c.args.context
	}

	if *c.args.before == 0 {
		c.args.before = c.args.context
	}

	if *c.args.fixed {
		c.app.pattern = fmt.Sprintf("^%s$", c.app.pattern)
	}

	if *c.args.ignoreCase {
		c.app.pattern = "(?i)" + c.app.pattern
	}

	return nil
}

func (c *Container) openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error [%v] in func openFile", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.app.data = append(c.app.data, scanner.Text())
	}
	return nil
}

func (c *Container) findString(re regexp.Regexp) {
	matched := make(map[int]struct{})
	printed := make(map[int]struct{})
	for index, item := range c.app.data {
		if *c.args.count == 0 {
			break
		}
		if re.MatchString(item) && !*c.args.invert || !re.MatchString(item) && *c.args.invert {
			matched[index] = struct{}{}
			*c.args.count--
		}
	}
	var lastPrinted int
	for key := range matched {
		if *c.args.after > 0 || *c.args.before > 0 {
			if key-lastPrinted > 2 {
				fmt.Println("--")
			}

			if _, ok := printed[key]; ok {
				continue
			}

			start := key - *c.args.before
			if start < 0 {
				start = 0
			}

			finish := key + *c.args.after
			if finish > len(c.app.data)-1 {
				finish = len(c.app.data) - 1
			}

			for ; start <= finish; start++ {
				if *c.args.lineNum {
					fmt.Printf("line:%v result:%v\n", start+1, c.app.data[start])
					printed[start] = struct{}{}
					lastPrinted = start
					continue
				}
				fmt.Println(c.app.data[start])
				printed[start] = struct{}{}
				lastPrinted = start
			}
			continue
		}
		if *c.args.lineNum {
			fmt.Printf("line:%v result:%v\n", key+1, c.app.data[key])
			continue
		}
		fmt.Println(key)
	}
}
