package errorkit

import (
	"errors"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		format    string
		args      []interface{}
		wantErr   string
	}{
		{"false condition", false, "test error: %s", []interface{}{"condition failed"}, "test error: condition failed"},
		{"true condition", true, "this should not be returned", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.condition, tt.format, tt.args...)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err == nil || err.Error() != tt.wantErr {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTry(t *testing.T) {
	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{"return error", func() error { return errors.New("test error") }, "test error"},
		{"panic", func() error { panic("panic test") }, "panic occurred"},
		{"no error", func() error { return nil }, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Try(tt.fn)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("SafeExec() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("SafeExec() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTry0(t *testing.T) {
	tests := []struct {
		name    string
		fn      func()
		wantErr string
	}{
		{"panic", func() { panic("panic test") }, "panic occurred"},
		{"no error", func() {}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Try0(tt.fn)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("SafeExecWithNoResult() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("SafeExecWithNoResult() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTry1(t *testing.T) {
	tests := []struct {
		name    string
		fn      func() (int, error)
		want    int
		wantErr string
	}{
		{"return result and no error", func() (int, error) { return 42, nil }, 42, ""},
		{"return error", func() (int, error) { return 0, errors.New("test error") }, 0, "test error"},
		{"panic", func() (int, error) { panic("panic test") }, 0, "panic occurred"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Try1(tt.fn)
			if tt.wantErr == "" {
				if err != nil || result != tt.want {
					t.Errorf("SafeExecWithResult() = %v, %v, want %v, nil", result, err, tt.want)
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("SafeExecWithResult() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTry2(t *testing.T) {
	tests := []struct {
		name    string
		fn      func() (int, string, error)
		want1   int
		want2   string
		wantErr string
	}{
		{"return results and no error", func() (int, string, error) { return 42, "hello", nil }, 42, "hello", ""},
		{"return error", func() (int, string, error) { return 0, "", errors.New("test error") }, 0, "", "test error"},
		{"panic", func() (int, string, error) { panic("panic test") }, 0, "", "panic occurred"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2, err := Try2(tt.fn)
			if tt.wantErr == "" {
				if err != nil || result1 != tt.want1 || result2 != tt.want2 {
					t.Errorf("SafeExecWithTwoResults() = %v, %v, %v, want %v, %v, nil", result1, result2, err, tt.want1, tt.want2)
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("SafeExecWithTwoResults() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTry3(t *testing.T) {
	tests := []struct {
		name    string
		fn      func() (int, string, bool, error)
		want1   int
		want2   string
		want3   bool
		wantErr string
	}{
		{"return results and no error", func() (int, string, bool, error) { return 42, "hello", true, nil }, 42, "hello", true, ""},
		{"return error", func() (int, string, bool, error) { return 0, "", false, errors.New("test error") }, 0, "", false, "test error"},
		{"panic", func() (int, string, bool, error) { panic("panic test") }, 0, "", false, "panic occurred"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2, result3, err := Try3(tt.fn)
			if tt.wantErr == "" {
				if err != nil || result1 != tt.want1 || result2 != tt.want2 || result3 != tt.want3 {
					t.Errorf("SafeExecWithThreeResults() = %v, %v, %v, %v, want %v, %v, %v, nil", result1, result2, result3, err, tt.want1, tt.want2, tt.want3)
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("SafeExecWithThreeResults() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestTryCatch(t *testing.T) {
	tests := []struct {
		name          string
		fn            func() error
		handlerCalled *bool
	}{
		{"return error", func() error { return errors.New("test error") }, new(bool)},
		{"no error", func() error { return nil }, new(bool)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(err error) {
				*tt.handlerCalled = true
			}

			TryCatch(tt.fn, handler)
			if tt.fn() != nil && !*tt.handlerCalled {
				t.Errorf("SafeExecWithHandler() handler not called")
			}
			if tt.fn() == nil && *tt.handlerCalled {
				t.Errorf("SafeExecWithHandler() handler called unexpectedly")
			}
		})
	}
}

func TestTry0Catch(t *testing.T) {
	tests := []struct {
		name          string
		fn            func()
		handlerCalled *bool
	}{
		{"panic", func() { panic("panic test") }, new(bool)},
		{"no error", func() {}, new(bool)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(err error) {
				*tt.handlerCalled = true
			}

			Try0Catch(tt.fn, handler)
			if tt.name == "panic" && !*tt.handlerCalled {
				t.Errorf("SafeExecWithHandler0() handler not called")
			}
			if tt.name == "no error" && *tt.handlerCalled {
				t.Errorf("SafeExecWithHandler0() handler called unexpectedly")
			}
		})
	}
}

func TestTry1Catch(t *testing.T) {
	tests := []struct {
		name          string
		fn            func() (int, error)
		handlerCalled *bool
		want          int
		wantErr       bool
	}{
		{"return result and no error", func() (int, error) { return 42, nil }, new(bool), 42, false},
		{"return error", func() (int, error) { return 0, errors.New("test error") }, new(bool), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(err error) {
				*tt.handlerCalled = true
			}

			result := Try1Catch(tt.fn, handler)
			if tt.wantErr && !*tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithResult() handler not called")
			}
			if !tt.wantErr && *tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithResult() handler called unexpectedly")
			}
			if result != tt.want {
				t.Errorf("SafeExecWithHandlerWithResult() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestTry2Catch(t *testing.T) {
	tests := []struct {
		name          string
		fn            func() (int, string, error)
		handlerCalled *bool
		want1         int
		want2         string
		wantErr       bool
	}{
		{"return results and no error", func() (int, string, error) { return 42, "hello", nil }, new(bool), 42, "hello", false},
		{"return error", func() (int, string, error) { return 0, "", errors.New("test error") }, new(bool), 0, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(err error) {
				*tt.handlerCalled = true
			}

			result1, result2 := Try2Catch(tt.fn, handler)
			if tt.wantErr && !*tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithTwoResults() handler not called")
			}
			if !tt.wantErr && *tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithTwoResults() handler called unexpectedly")
			}
			if result1 != tt.want1 || result2 != tt.want2 {
				t.Errorf("SafeExecWithHandlerWithTwoResults() = %v, %v, want %v, %v", result1, result2, tt.want1, tt.want2)
			}
		})
	}
}

func TestTry3Catch(t *testing.T) {
	tests := []struct {
		name          string
		fn            func() (int, string, bool, error)
		handlerCalled *bool
		want1         int
		want2         string
		want3         bool
		wantErr       bool
	}{
		{"return results and no error", func() (int, string, bool, error) { return 42, "hello", true, nil }, new(bool), 42, "hello", true, false},
		{"return error", func() (int, string, bool, error) { return 0, "", false, errors.New("test error") }, new(bool), 0, "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := func(err error) {
				*tt.handlerCalled = true
			}

			result1, result2, result3 := Try3Catch(tt.fn, handler)
			if tt.wantErr && !*tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithThreeResults() handler not called")
			}
			if !tt.wantErr && *tt.handlerCalled {
				t.Errorf("SafeExecWithHandlerWithThreeResults() handler called unexpectedly")
			}
			if result1 != tt.want1 || result2 != tt.want2 || result3 != tt.want3 {
				t.Errorf("SafeExecWithHandlerWithThreeResults() = %v, %v, %v, want %v, %v, %v", result1, result2, result3, tt.want1, tt.want2, tt.want3)
			}
		})
	}
}
