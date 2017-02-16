#Channels

Let create a simple example of using channels

Create new folder for five project.

```bash
mkdir -p $GOPATH/src/five
```
Open editor

```bash
atom $GOPATH/src/five
```


##Implement

create five/main.go

Will not work

```go
package main

import "fmt"

func main() {
	c := make(chan string)
  // write to a channel

	c <- "The Gadget Show"

	// read from a channel
	val := <-c

	//Print Out Result
	fmt.Printf("Watching %s\n", val)

}

```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/Users/user/go/src/five/main.go:9 +0x89
```

#Fix

```go
package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		// write to a channel
		c <- "The Gadget Show"
	}()

	// read from a channel
	val := <-c

	//Print Out Result
	fmt.Printf("Watching %q\n", val)

}
```
