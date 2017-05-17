package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/papaburgs/windowstime"
)

func main() {
	var (
		qdate time.Time
		err   error
	)

	if len(os.Args) == 1 {
		usage := `Usage:  windowstime <value> [--pwd]
where <value> is an LDAP or windows time value
ie 18 digit string representing number of 100's of nanoseconds since Jan 1, 1601

Example: 'windowstime 131395047950000000' Which result in May 17th 2017 around 14:30

 --pwd  Print the date 42 days after the input time
          - for a rough guess on when password needs to be reset
`
		fmt.Println(usage)
		return
	}
	qdate, err = windowstime.Convert(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 3 {
		if os.Args[2] == "--pwd" {
			qdate = qdate.Add(time.Duration(3628800000000000)) // 42 days in nanoseconds
			fmt.Println(qdate.Format("Mon, Jan 02 2006"))
			return
		}
	}
	fmt.Println(qdate)

}
