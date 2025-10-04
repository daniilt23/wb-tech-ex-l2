package main

// golangci-lint run 0 issues

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func getNTPTime() (time.Time, error) {
	ntpTime, err := ntp.Time("ntp0.ntp-servers.net") // реализуем взятие московского ntp времени
	if err != nil {
		return time.Time{}, err
	}

	return ntpTime, nil
}

func main() {
	time, err := getNTPTime()
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // завершаем работу с кодом 1
	}

	fmt.Println(time)
}
