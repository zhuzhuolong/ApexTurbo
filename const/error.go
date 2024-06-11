package _const

import (
	"errors"
	"regexp"
)

var (
	// ErrMethodMismatch is returned when the method in the request does not match
	// the method defined against the route.
	ErrMethodMismatch = errors.New("method is not allowed")
	// ErrNotFound is returned when no route match is found.
	ErrNotFound = errors.New("no matching route was found")
	// RegexpCompileFunc aliases regexp.Compile and enables overriding it.
	// Do not run this function from `init()` in importable packages.
	// Changing this value is not safe for concurrent use.
	RegexpCompileFunc = regexp.Compile
)
