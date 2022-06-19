// Package withgomega wraps gomega global functions into structs.
package withgomega

//go:generate sh -c "go run internal/gen/main.go > matcher.go"
//go:generate gofmt -w matcher.go
