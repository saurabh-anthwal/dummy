package utils

import (
	"github.com/cenkalti/backoff"
	"go.uber.org/zap"
	"time"
)

// Permanent wraps the given err in a *PermanentError.
func Permanent(err error) *backoff.PermanentError {
	return &backoff.PermanentError{
		Err: err,
	}
}

// ExponentialRetry executes the given function and retries exponentially on errors.
// Usage
// 	utils.ExponentialRetry(r.log, func() error {
//		// do what ever you want to retry
//	})
func ExponentialRetry(log *zap.SugaredLogger, F func() error) {
	retried := false
	t := time.Now()

	operation := func() error {
		if retried {
			log.Debug("ExponentialRetry retrying! time elapsed: ", time.Now().Sub(t))
		}

		if err := F(); err != nil {
			retried = true
			return err
		}
		return nil
	}

	if err := backoff.Retry(operation, backoff.NewExponentialBackOff()); err != nil {
		log.Error("ExponentialRetry got error: ", err)
	}
}
