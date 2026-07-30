package assert

import (
	"reflect"
	"testing"
)

type Assertion[T any] struct {
	t      *testing.T
	actual T
}

func Assert[T any](t *testing.T, value T) *Assertion[T] {
	return &Assertion[T]{
		t:      t,
		actual: value,
	}
}

func isValueNil(i interface{}) bool {

	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

func (a Assertion[T]) IsNil() {
	a.t.Helper()

	if !isValueNil(a.actual) {
		a.t.Fatalf("assertion failed: expected value to be non-nil, but got %v", a.actual)
	}
}

func (a Assertion[T]) IsNotNil() {
	a.t.Helper()

	if isValueNil(a.actual) {
		a.t.Fatalf("assertion failed: expected value to be non-nil, but got nil")
	}
}

func (a Assertion[T]) Error() {
	a.t.Helper()

}

func (a Assertion[T]) IsEqualTo(expected T) {
	a.t.Helper()
	if !reflect.DeepEqual(a.actual, expected) {
		a.t.Fatalf(
			"assertion failed:\nexpected: %#v\nactual:   %#v",
			expected,
			a.actual,
		)
	}
}

func (a Assertion[T]) IsNotEqualTo(unexpected T) {
	a.t.Helper()
	if reflect.DeepEqual(a.actual, unexpected) {
		a.t.Fatalf(
			"assertion failed:\nexpected: %#v\nactual:   %#v",
			unexpected,
			a.actual,
		)
	}
}
