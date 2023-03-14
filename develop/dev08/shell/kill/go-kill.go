package kill

import (
	"log"
	"os"
)

func Run(pid int) {
	err := kill(pid)
	if err != nil {
		log.Println(err)
	}
}

func kill(pid int) error {
	fProcess, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	fProcess.Kill()
	return nil
}
