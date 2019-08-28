jsonchan
========

[![GoDoc](https://godoc.org/github.com/trapped/jsonchan?status.svg)](https://godoc.org/github.com/trapped/jsonchan)

jsonchan streams JSON from a `chan interface{}` sink. It has no external dependencies.

Install with `go get -u github.com/trapped/jsonchan`.

## Usage

```go
c := make(chan interface{}, 0)
go func() {
	n := 0
	for n < 10 {
		c <- n
	}
	close(c)
}
jsonchan.Stream(os.Stdout, c)
```
