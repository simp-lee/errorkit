package errorkit

import (
	"fmt"
	"runtime/debug"
)

// Validate creates an error when a condition is not met.
func Validate(condition bool, format string, args ...any) error {
	if !condition {
		return fmt.Errorf(fmt.Sprintf(format, args...))
	}
	return nil
}

// SafeExec executes a function and returns the error value if it returns an error, nil otherwise.
// It also recovers from panics and converts them into errors.
func SafeExec(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v\nStack trace:\n%s", r, debug.Stack())
		}
	}()

	return fn()
}

// SafeExecWithNoResult has the same behavior as SafeExec, but fn returns no value.
func SafeExecWithNoResult(fn func()) error {
	return SafeExec(func() error {
		fn()
		return nil
	})
}

// SafeExecWithResult executes a function and returns a result and an error.
// It also recovers from panics and converts them into errors.
func SafeExecWithResult[T any](fn func() (T, error)) (result T, err error) {
	err = SafeExec(func() error {
		var fnErr error
		result, fnErr = fn()
		return fnErr
	})
	return
}

// SafeExecWithTwoResults executes a function and returns two results and an error.
// It also recovers from panics and converts them into errors.
func SafeExecWithTwoResults[T1, T2 any](fn func() (T1, T2, error)) (result1 T1, result2 T2, err error) {
	err = SafeExec(func() error {
		var fnErr error
		result1, result2, fnErr = fn()
		return fnErr
	})
	return
}

// SafeExecWithThreeResults executes a function and returns three results and an error.
// It also recovers from panics and converts them into errors.
func SafeExecWithThreeResults[T1, T2, T3 any](fn func() (T1, T2, T3, error)) (result1 T1, result2 T2, result3 T3, err error) {
	err = SafeExec(func() error {
		var fnErr error
		result1, result2, result3, fnErr = fn()
		return fnErr
	})
	return
}

// SafeExecWithHandler executes a function and calls the handler function if an error occurs.
func SafeExecWithHandler(fn func() error, handler func(error)) {
	err := SafeExec(fn)
	if err != nil {
		handler(err)
	}
}

// SafeExecWithHandler0 executes a function with no return value and calls the handler function if an error occurs.
func SafeExecWithHandler0(fn func(), handler func(error)) {
	if err := SafeExecWithNoResult(fn); err != nil {
		handler(err)
	}
}

// SafeExecWithHandlerWithResult executes a function returning a result and an error,
// and calls the handler function if an error occurs.
func SafeExecWithHandlerWithResult[T any](fn func() (T, error), handler func(error)) (result T) {
	result, err := SafeExecWithResult(fn)
	if err != nil {
		handler(err)
	}
	return
}

// SafeExecWithHandlerWithTwoResults executes a function returning two results and an error,
// and calls the handler function if an error occurs.
func SafeExecWithHandlerWithTwoResults[T1, T2 any](fn func() (T1, T2, error), handler func(error)) (result1 T1, result2 T2) {
	result1, result2, err := SafeExecWithTwoResults(fn)
	if err != nil {
		handler(err)
	}
	return
}

// SafeExecWithHandlerWithThreeResults executes a function returning three results and an error,
// and calls the handler function if an error occurs.
func SafeExecWithHandlerWithThreeResults[T1, T2, T3 any](fn func() (T1, T2, T3, error), handler func(error)) (result1 T1, result2 T2, result3 T3) {
	result1, result2, result3, err := SafeExecWithThreeResults(fn)
	if err != nil {
		handler(err)
	}
	return
}
