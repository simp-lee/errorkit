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

// Try executes a function and returns the error value if it returns an error, nil otherwise.
// It also recovers from panics and converts them into errors.
func Try(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v\nStack trace:\n%s", r, debug.Stack())
		}
	}()

	return fn()
}

// Try0 has the same behavior as Try, but fn returns no value.
func Try0(fn func()) error {
	return Try(func() error {
		fn()
		return nil
	})
}

// Try1 executes a function and returns a result and an error.
func Try1[T any](fn func() (T, error)) (result T, err error) {
	err = Try(func() error {
		var fnErr error
		result, fnErr = fn()
		return fnErr
	})
	return
}

// Try2 executes a function and returns two results and an error.
func Try2[T1, T2 any](fn func() (T1, T2, error)) (result1 T1, result2 T2, err error) {
	err = Try(func() error {
		var fnErr error
		result1, result2, fnErr = fn()
		return fnErr
	})
	return
}

// Try3 executes a function and returns three results and an error.
func Try3[T1, T2, T3 any](fn func() (T1, T2, T3, error)) (result1 T1, result2 T2, result3 T3, err error) {
	err = Try(func() error {
		var fnErr error
		result1, result2, result3, fnErr = fn()
		return fnErr
	})
	return
}

// TryCatch executes a function and calls the catch function if an error occurs.
func TryCatch(fn func() error, catch func(error)) {
	err := Try(fn)
	if err != nil {
		catch(err)
	}
}

// Try0Catch executes a function with no return value and calls the catch function if an error occurs.
func Try0Catch(fn func(), catch func(error)) {
	if err := Try0(fn); err != nil {
		catch(err)
	}
}

// Try1Catch executes a function returning a result and an error,
// and calls the catch function if an error occurs.
func Try1Catch[T any](fn func() (T, error), catch func(error)) (result T) {
	result, err := Try1(fn)
	if err != nil {
		catch(err)
	}
	return
}

// Try2Catch executes a function returning two results and an error,
// and calls the catch function if an error occurs.
func Try2Catch[T1, T2 any](fn func() (T1, T2, error), catch func(error)) (result1 T1, result2 T2) {
	result1, result2, err := Try2(fn)
	if err != nil {
		catch(err)
	}
	return
}

// Try3Catch executes a function returning three results and an error,
// and calls the catch function if an error occurs.
func Try3Catch[T1, T2, T3 any](fn func() (T1, T2, T3, error), catch func(error)) (result1 T1, result2 T2, result3 T3) {
	result1, result2, result3, err := Try3(fn)
	if err != nil {
		catch(err)
	}
	return
}
