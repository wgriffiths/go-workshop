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
		t.Errorf("Expected 200 Result was %d", w.Code)
	}

}

func TestOClock_Body(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()
	OClock(w, req)
	expected := time.Now().Format("2006-01-02T15:04")
	result := w.Body.String()
	if result != expected {
		t.Errorf("Expected %q Result %q", expected, result)
	}

}

```

```bash
=== RUN   TestOClock
--- PASS: TestOClock (0.00s)
=== RUN   TestOClock_Body
--- FAIL: TestOClock_Body (0.00s)
	main_test.go:27: Expected "2017-02-15T20:27" Result ""
FAIL
exit status 1
FAIL	oclock	0.011s
```

##Implement

modify oclock/main.go

```go
package main

import "net/http"

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format("2006-01-02T15:04")))
}
```

##Test

```bash
go test -v oclock
```

```bash
=== RUN   TestOClock
--- PASS: TestOClock (0.00s)
=== RUN   TestOClock_Body
--- PASS: TestOClock_Body (0.00s)
PASS
ok  	oclock	0.010s
```

##Implements

No that we have a handler lets add it to a server

```go
package main

import (
	"log"
	"net/http"
	"time"
)

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format("2006-01-02T15:04")))
}

func main() {
	http.HandleFunc("/oclock", OClock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```


##Install

```bash
go install oclock
```

##Run

```bash
oclock
```

http://localhost:8080/oclock

Ship It!
