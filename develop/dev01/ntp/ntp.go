package ntp

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

// Run runnable func
func Run() {
	currTime, err := CurrentTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Print(currTime)
}

// CurrentTime check curr time
func CurrentTime(host string) (string, error) {
	currTime, err := ntp.Time(host)
	if err != nil {
		return "", err
	}
	return currTime.UTC().Format(time.UnixDate), nil
}
