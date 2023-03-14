package pwd

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	res, err := pwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}

func pwd() (string, error) {
	return os.Getwd()
}
