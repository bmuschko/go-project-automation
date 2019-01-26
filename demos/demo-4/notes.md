# Demo 4

Demonstrates the purpose and usage of the tool [errcheck](https://github.com/kisielk/errcheck).

## Features

* Checks that program handles errors
* Typical example: file handling functions that return the file and/or and error

## Usage

```
$ errcheck
main.go:10:9:	f.Write([]byte("Hello, world."))
main.go:12:9:	f.Close()
```