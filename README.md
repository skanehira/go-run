# go-run
`go-run` is tool that running go source code from stdin or arg.

![](https://i.imgur.com/WHqgenv.gif)

## Installation
```sh
$ git clone https://github.com/skanehira/go-run
$ go install
```

## Usage
```sh
$ go-run -h
Usage of go-run:
  -c string
        source code
  -debug
        print debug log

# run code use -c
$ go-run -c '
package main
func main() {
println("hello gorilla")
}'
hello gorilla

# run code use pipeline
echo 'package main
func main() {
println("hello gorilla")
}
' | go-run
hello gorilla
```

Also, you can run code from vim buffer.

```vim
:w !go-run
```
