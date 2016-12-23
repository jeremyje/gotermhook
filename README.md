Termhooks for Go
================

Go library for running a function when a app signal is caught.
Mostly used for running cleanup methods on terminate, Ctrl+C.

Example
-------

```
package main

import (
    "fmt"
    "net/http"
    "github.com/jeremyje/gotermhook"
)

func main() {
    gotermhook.Add(func() {
        fmt.Println("Cleanup Done")
    })
    fmt.Println("Press Ctrl+C to close.")
    http.ListenAndServe(":8080", nil)
}
```