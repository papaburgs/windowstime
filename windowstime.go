// Package windowstime is a library to convert a ldap time string into a go time.Time
package windowstime

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ErrInvalidInput is thrown on invalid input, ie, can't convert to integer
var ErrInvalidInput = errors.New("Invalid input")

// Convert takes an 18 digit number representing the 100s of nanoseconds since Jan 1, 1601
// and returns a time in UTC
// returns Jan 1, 1601 as time on error
func Convert(input string) (time.Time, error) {
	var (
		nanosecs             int
		tenthOfInputDuration time.Duration
		qdate                time.Time
		err                  error
	)
	qdate = time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)

	// Durations can only be an int64 variable of nanoseconds,
	//  this limits us to about 295 years
	//  our refernece year is over 400 years
	//  input is in 100's of nanoseconds
	//  so adding on more zero and then pretending it is nanoseconds
	//  this is effectivly dividing the input by 10
	//  so we need to add it to the 'epoch' time 10 times
	nanosecs, err = strconv.Atoi(fmt.Sprintf("%s0", input))
	if err != nil {
		return qdate, ErrInvalidInput
	}
	tenthOfInputDuration = time.Duration(nanosecs)

	for x := 0; x < 10; x++ {
		qdate = qdate.Add(tenthOfInputDuration)
	}
	return qdate, nil
}
