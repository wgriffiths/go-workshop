#Hello

Create new folder for hello project.

```bash
mkdir -p $GOPATH/src/hello
```
Open editor

```bash
atom $GOPATH/src/hello
```

##Test

Lets start with a test.

create main_test.go

```go
package main

func ExampleHello() {
	main()
	// Output: Hello World
}
```

Run tests

```bash
go test -v hello
```

```bash
# hello
hello/main_test.go:4: undefined: main
FAIL	hello [build failed]
```

##Implement

create hello/main.go

```go
import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

```bash
go test -v hello
```

```bash
=== RUN   ExampleHello
--- PASS: ExampleHello (0.00s)
PASS
ok  	hello	0.006s
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
