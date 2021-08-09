package base

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var host = "0.beevik-ntp.pool.ntp.org"

func Get() time.Time {
	ntpTime, err := ntp.Time(host)
	if err != nil {
		println(err) // пишем в Stderr
		os.Exit(2)
	}

	fmt.Println("NTP time ", ntpTime)
	fmt.Println("Local time ", time.Now())
	return ntpTime
}
