## Simple Pascal Interpreter in golang

This is going to be an implementation of pascal interpreter in golang.
I'm following [this tutorial](https://ruslanspivak.com/lsbasi-part1/).

### How to run
```sh
$ git clone https://github.com/riadafridishibly/spi.git
$ cd spi
$ go build
$ cat sampleprog.pas
PROGRAM Part10;
VAR
   number     : INTEGER;
   a, b, c, x : INTEGER;
   y          : REAL;

BEGIN {Part10}
   BEGIN
      number := 2;
      a := number;
      b := 10 * a + 10 * number DIV 4;
      c := a - - b
   END;
   x := 11;
   y := 20 / 7 + 3.14;
   { writeln('a = ', a); }
   { writeln('b = ', b); }
   { writeln('c = ', c); }
   { writeln('number = ', number); }
   { writeln('x = ', x); }
   { writeln('y = ', y); }
END.  {Part10}


$ ./spi sampleprog.pas
map[a:2 b:25 c:27 number:2 x:11 y:5.997142857142857]
```

### Next thing to do
- Implement AST :heavy_check_mark: 
- Support Unary Operator :heavy_check_mark:
- Update grammar to support Pascal Language :heavy_check_mark:
- Refactor code & write test 
- Read the next lesson
