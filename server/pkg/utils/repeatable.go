package repeatable

import "time"

// DoWithTries executes the given function fn with the given number of attempts and a delay between attempts.
// If the function returns an error, the function will sleep for the given delay and retry the function up to the given number of attempts.
// If the function does not return an error after the given number of attempts, the function will return the last error encountered.
func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--
			continue
		}
		return nil
	}
	return
}