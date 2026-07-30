package assert

import "testing"

type ErrorAssertion struct {
    t   *testing.T
    err error
}

func AssertError(t *testing.T, err error) ErrorAssertion {
    return ErrorAssertion{
        t:   t,
        err: err,
    }
}

func (a ErrorAssertion) NoError() {
    a.t.Helper()

    if a.err != nil {
        a.t.Fatalf("unexpected error: %v", a.err)
    }
}

func (a ErrorAssertion) Error() {
    a.t.Helper()

    if a.err == nil {
        a.t.Fatal("expected an error")
    }
}