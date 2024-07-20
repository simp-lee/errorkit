# ErrorKit

`errorkit` is a Go package designed to simplify error handling and panic recovery in Go applications. It provides a set of utility functions to execute functions safely, handle errors gracefully, and recover from panics, enhancing the robustness and maintainability of your code.

## Features

- **Safe Execution**: Execute functions safely, catching panics and converting them into errors.
- **Error Handling**: Handle errors with optional handler functions.
- **Panic Recovery**: Recover from panics and handle them gracefully.
- **Stack Traces**: Obtain stack traces for errors and panics.
- **Generic Support**: Utilize generic types for flexible and type-safe results.
- **Simple API**: A straightforward and intuitive API for ease of use.

## Installation

To install ErrorKit, use the following command:

```shell
go get github.com/simp-lee/errorkit
```

## Usage

Below are examples demonstrating how to use various functions from the `errorkit` package:

### Validate

Create an error when a condition is not met:

```Go
err := errorkit.Validate(false, "test error: %s", "condition failed")
if err != nil {
	fmt.Println(err) // Output: test error: condition failed
}
```

### SafeExec

Execute a function safely, catching any panics:

```Go
err := errorkit.SafeExec(func() error {
	panic("unexpected panic")
})
if err != nil {
	fmt.Println(err) // Output: panic occurred: unexpected panic
}
```

### SafeExecWithNoResult

Execute a function with no return value safely:

```Go
err := errorkit.SafeExecWithNoResult(func() {
	panic("unexpected panic")
})
if err != nil {
	fmt.Println(err) // Output: panic occurred: unexpected panic
}
```

### SafeExecWithResult

Execute a function and return a result safely:

```Go
result, err := errorkit.SafeExecWithResult(func() (int, error) {
	return 42, nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result) // Output: 42
}
```

### SafeExecWithTwoResults

Execute a function and return two results safely:

```go
result1, result2, err := errorkit.SafeExecWithTwoResults(func() (int, string, error) {
	return 42, "hello", nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result1, result2) // Output: 42 hello
}
```

### SafeExecWithThreeResults

Execute a function and return three results safely:

```go
result1, result2, result3, err := errorkit.SafeExecWithThreeResults(func() (int, string, bool, error) {
	return 42, "hello", true, nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result1, result2, result3) // Output: 42 hello true
}
```

### SafeExecWithHandler

Execute a function and call a handler if an error occurs:

```go
errorkit.SafeExecWithHandler(func() error {
	return errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
```

### SafeExecWithHandlerWithResult

Execute a function returning a result and call a handler if an error occurs:

```go
result := errorkit.SafeExecWithHandlerWithResult(func() (int, error) {
	return 42, errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
fmt.Println(result) // Output: 0
```

### SafeExecWithHandlerWithTwoResults

Execute a function returning two results and call a handler if an error occurs:

```go
result1, result2 := errorkit.SafeExecWithHandlerWithTwoResults(func() (int, string, error) {
	return 42, "hello", errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
fmt.Println(result1, result2) // Output: 0
```

### SafeExecWithHandlerWithThreeResults

Execute a function returning three results and call a handler if an error occurs:

```go
result1, result2, result3 := errorkit.SafeExecWithHandlerWithThreeResults(func() (int, string, bool, error) {
	return 42, "hello", true, errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
fmt.Println(result1, result2, result3) // Output: 0
```

## Contributing

Contributions are welcome! To contribute, please open an issue or submit a pull request on GitHub.

## License

`errorkit` is licensed under the MIT License.