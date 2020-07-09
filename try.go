package try

import (
	"errors"
	"log"
	"time"
)

// MaxRetries is the maximum number of retries before bailing.
var MaxRetries = 10

var errMaxRetriesReached = errors.New("exceeded retry limit")

// Func represents functions that can be retried.
type Func func(attempt int) (retry bool, err error)

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
// waiting for delay to elapse before retrying, and
// exponent to factor delay after each retry
func Do(delay, exponent time.Duration, fn Func) error {
	var err error
	var cont bool
	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}
		delay *= exponent
		log.Printf("retrying after %v", delay)
		time.Sleep(delay)
	}
	return err
}

// IsMaxRetries checks whether the error is due to hitting the
// maximum number of retries or not.
func IsMaxRetries(err error) bool {
	return err == errMaxRetriesReached
}
