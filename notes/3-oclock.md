#OClock

Let create a simple time service

Create new folder for oclock project.

```bash
mkdir -p $GOPATH/src/oclock
```
Open editor

```bash
atom $GOPATH/src/oclock
```

##Test

Lets start with a test.

create main_test.go

```go
package main

import (
	"net/http/httptest"
	"testing"
)

func TestOClock(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 Go %d", w.Code)
	}

}
```

Run tests

```bash
go test -v oclock
```

```bash
# oclock
../go/src/oclock/main_test.go:11: undefined: OClock
FAIL	oclock [build failed]
```

##Implement

create oclock/main.go

```go
package main

import "net/http"

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {
}
```

```bash
go test -v oclock
```

```bash
=== RUN   TestOClock
--- PASS: TestOClock (0.00s)
PASS
ok  	oclock	0.011s
```

##Test

```go
package main

import (
	"net/http/httptest"
	"testing"
)

func TestOClock_Status(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 Go %d", w.Code)
	}

}

func TestOClock_Body(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)

  expected :=  "2017-02-14 21:40"
  result := w.Body.String()
	if result != expected {
		t.Errorf("Expected %q Result %q", expected,result)
	}

}
```



##Install

```bash
go install hello
```

##Run

```bash
hello
```


Ship It!
