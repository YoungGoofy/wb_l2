package ps

import (
	"fmt"
	"log"
	"os/exec"
)

func Run() {
	out, err := pc()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(out)
}

func pc() (string, error) {
	cmd := exec.Command("ps")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
