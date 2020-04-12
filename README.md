## Simple Pascal Interpreter in golang

This is going to be an implementation of pascal interpreter in golang.
I'm following [this tutorial](https://ruslanspivak.com/lsbasi-part1/).

### How to run
```sh
$ git clone https://github.com/riadafridishibly/spi.git
$ cd spi
$ go build
$ ./spi

spi> 1 + 2
3
spi> 1 - - 2
3
spi> (1 + 3) / 3  # currently supports only integer division
1
spi> 3 - (3 - ( 2 * 4 ) / 4)
2
spi>

```

### Next thing to do
- Implement AST :heavy_check_mark: 
- Support Unary Operator :heavy_check_mark:
- Update grammar to support Pascal Language
