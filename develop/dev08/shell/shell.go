package shell

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/YoungGoofy/wb_l2/develop/dev08/shell/cd"
	"github.com/YoungGoofy/wb_l2/develop/dev08/shell/echo"
	"github.com/YoungGoofy/wb_l2/develop/dev08/shell/kill"
	"github.com/YoungGoofy/wb_l2/develop/dev08/shell/ps"
	"github.com/YoungGoofy/wb_l2/develop/dev08/shell/pwd"
)

func Run() {
	for {
		root := customPwd()
		fmt.Print(root, " >>> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		parsing := strings.Split(text, " ")
		switch parsing[0] {
		case "/exit":
			os.Exit(0)
		case "/help":
			help()
		case "cd":
			cdUtil(parsing)
		case "pwd":
			pwd.Run()
		case "echo":
			echo.Run(parsing)
		case "ps":
			ps.Run()
		case "kill":
			item, err := strconv.Atoi(parsing[1])
			if err != nil {
				log.Println(err)
			}
			kill.Run(item)
		default:
			fmt.Println("Для помощи введи команду /help")
		}
	}
}

func customPwd() string {
	var shortRoot string
	root, _ := os.Getwd()
	userHomeDir, _ := os.UserHomeDir()
	root = strings.Replace(root, userHomeDir, "~", 1)
	rootSlice := strings.Split(root, "/")
	if len(rootSlice) > 2 {
		for i := 0; i < len(rootSlice)-2; i++ {
			shortRoot += string(rootSlice[i][0]) + "/"
		}
	}
	return shortRoot + rootSlice[len(rootSlice)-2] + "/" + rootSlice[len(rootSlice)-1]
}

func cdUtil(args []string) {
	count := len(args)
	switch count {
	case 1:
		userHomeDir, _ := os.UserHomeDir()
		cd.Run(userHomeDir)
	default:
		cd.Run(args[1])
	}
}

func help() {
	fmt.Println()
	fmt.Println(`Руководство:
1. /exit - выход из программы
2. /help - помощь
3. cd <dir> - перемещение по директориям
4. pwd - показывает директорию, в которой ты находишься
5. echo <text> - выводит в консоль текст
6. ps - выводит запущенные процессы
7. kill <pid> - убивает запущенный процесс`)
	fmt.Println()
}
