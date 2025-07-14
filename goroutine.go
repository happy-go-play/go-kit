package gokit

import (
	"log/slog"
	"runtime/debug"
)

func Go(f func()) {
	go func() {
		defer func() {
			// recover from panic
			if r := recover(); r != nil {
				slog.Error("Go panic recovered", slog.Any("panic", r), slog.String("stack", string(debug.Stack())))

				panic(r) // re-panic
			}
		}()

		f()
	}()
}

func SafeGo(f func()) {
	go func() {
		defer func() {
			// recover from panic
			if r := recover(); r != nil {
				slog.Error("Go panic recovered", slog.Any("panic", r), slog.String("stack", string(debug.Stack())))
			}
		}()

		f()
	}()
}
