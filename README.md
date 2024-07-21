# ErrorKit

`errorkit` is a Go package designed to simplify error handling and panic recovery in Go applications. It provides a set of utility functions to execute functions safely, handle errors gracefully, and recover from panics, enhancing the robustness and maintainability of your code.

## Features

- **Safe Execution**: Execute functions safely, catching panics and converting them into errors.
- **Error Handling**: Handle errors with optional catch functions.
- **Panic Recovery**: Recover from panics and handle them gracefully.
- **Stack Traces**: Obtain stack traces for errors and panics.
- **Generic Support**: Utilize generic types for flexible and type-safe results.
- **Simple API**: A straightforward and intuitive API for ease of use.

## Installation

To install `errorkit`, use the following command:

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

### Try

Execute a function safely, catching any panics:

```Go
err := errorkit.Try(func() error {
	panic("unexpected panic")
})
if err != nil {
	fmt.Println(err) // Output: panic occurred: unexpected panic
}
```

### Try0

Execute a function with no return value safely:

```Go
err := errorkit.Try0(func() {
	panic("unexpected panic")
})
if err != nil {
	fmt.Println(err) // Output: panic occurred: unexpected panic
}
```

### Try1

Execute a function and return a result safely:

```Go
result, err := errorkit.Try1(func() (int, error) {
	return 42, nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result) // Output: 42
}
```

### Try2

Execute a function and return two results safely:

```go
result1, result2, err := errorkit.Try2(func() (int, string, error) {
	return 42, "hello", nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result1, result2) // Output: 42 hello
}
```

### Try3

Execute a function and return three results safely:

```go
result1, result2, result3, err := errorkit.Try3(func() (int, string, bool, error) {
	return 42, "hello", true, nil
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result1, result2, result3) // Output: 42 hello true
}
```

### TryCatch

Execute a function and call a catch function if an error occurs:

```go
errorkit.TryCatch(func() error {
	return errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
```

### Try1Catch

Execute a function returning a result and call a catch function if an error occurs:

```go
result := errorkit.Try1Catch(func() (int, error) {
	return 42, errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
fmt.Println(result) // Output: 0
```

### Try2Catch

Execute a function returning two results and call a catch function if an error occurs:

```go
result1, result2 := errorkit.Try2Catch(func() (int, string, error) {
	return 42, "hello", errors.New("test error")
}, func(err error) {
	fmt.Println(err) // Output: test error
})
fmt.Println(result1, result2) // Output: 0
```

### Try3Catch

Execute a function returning three results and call a catch function if an error occurs:

```go
result1, result2, result3 := errorkit.Try3Catch(func() (int, string, bool, error) {
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