Termhooks for Go
================

Go library for running a function when a app signal is caught.
Mostly used for running cleanup methods on terminate, Ctrl+C.

[![Build Status](https://secure.travis-ci.org/jeremyje/gotermhook.png)](http://travis-ci.org/jeremyje/gotermhook)

Install
-------
```
go get github.com/jeremyje/gotermhook
```

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
