package cd

import (
	"log"
	"os"
)

func Run(path string) {
	err := cd(path)
	if err != nil {
		log.Println("Нет такого файла или директории...")
	}
}

func cd(path string) error {
	if err := os.Chdir(path); err != nil {
		return err
	}
	return nil
}
