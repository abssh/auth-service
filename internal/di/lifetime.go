package di

import "sync"

func Singleton[T any](factory func() (T, error)) func() (T, error) {
    var (
        instance T
        err      error
        once     sync.Once
    )
    return func() (T, error) {
        once.Do(func() {
            instance, err = factory()
        })
        return instance, err
    }
}
