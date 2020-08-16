#  ![](docs/logo.svg)

**php2go** is a transpiler from a small subset of PHP to Go.

<img src="docs/php2go.jpg" width="500px"/>

*(photo by [Iskander Sharipov](https://github.com/quasilyte))*



This tool is written in [Go](https://golang.org/) and uses [z7zmey/php-parser](https://github.com/z7zmey/php-parser).

### Install

```
go get github.com/i582/php2go
```

### CLI

```
php2go [flags]
```

| flag | type   | description |
| ---- | ------ | ----------- |
| -i   | string | input file  |
| -o   | string | output file |

## What is currently supported.

**Types**:

1. Integer
2. Float
3. String
4. Bool

Union types composed of the types above are also supported.

**For union types supported**:

1. Comparison with any type above;
2. Use in conditions and boolean expressions;
3. Printing.

**Operators**

Supported arithmetic operators (`+`,`-`,`*`,`/` ,`.`,`++`,`--`) and 

boolean (`<`,`>`,`<=`,`>=`, `&&`, `||`) 

**Arrays**

Arrays are supported, both regular and associative, but they **must** consist of elements of the same type and with the same type of keys.

It supports index/key access, assignment to an element and a construction like `$arr[] = Elem;`

**Constructs**

The following language constructs are available:

1. `if-else`
2. `for`
3. `while`

**Currently only code in function is supported.**

**Output**

The `echo` operator is supported for output.



## TODO

1. Add support for all operators;
2. Add foreach support;
3. Add support for other types;
4. Add multi-file support;
5. And much more...



## Example

```php
<?php

function Foo() {
  // int
  $a = 100;
  // float
  $b = 1.5;
  // string
  $c = "Hello";
  // bool
  $d = true;
  // output
  echo $a;
  echo $b;
  echo $c;
  echo $d;
  // if-else
  if ($a == 100) {
    // variable inside
    $e = 10;
  } else {
    // and here, so variable should be inside
    $e = 10;
  }
  // output this variable
  echo $e;
  // another if-else
  if ($a == 100) {
    // variable with int type
    $f = 10;
  } else {
    // here is string
    $f = "10";
  }
  // so $f has type int|string
  echo $f;
  // support only single-type array
  $f = [1,2,3];
  echo $f;
  // fetch by index
  echo $f[1];
  // assign by index
  $f[1] = 10;
  echo $f;
  // simple array
  $g = [1,2,3];
  echo $g;
  // adding element
  $g[] = 100;
  echo $g;
  // and associative array with single-type keys
  $f = ["Key1" => 1, "Key2" => 2, "Key3" => 3];
  echo $f;
  // fetch by key
  echo $f["Key1"];
  // assign by key
  $f["Key1"] = 5;
  echo $f;
  // while
  $i = 0;
  while ($i < 20) {
    echo $i;
    $i++;
  }
  // for
  for ($i = 0; $i < 20; $i++) {
    echo $i + 5;
  }
  $qw = 1.5;
  // different operators
  echo $qw + 5 - 56.56 * 6 / 56;
}
```

Output:

```go
// Code generated by php2go. PLEASE DO NOT EDIT.
package index

import (
   "fmt"
)

// ...
// defining structures for types. See program/full.go
// ...

func Foo() {
	a := 100
	b := 1.5
	c := "Hello"
	d := true
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
	var e int64
	var f Var
	if a == 100 {
		e = 10
	} else {
		e = 10
	}
	fmt.Print(e)
	if a == 100 {
		f.Setint64(10)
	} else {
		f.Setstring("10")
	}
	fmt.Print(f.String())
	f.SetElementTypeint64([]int64{1, 2, 3})
	fmt.Print(f.GetElementTypeint64())
	fmt.Print(f.GetElementTypeint64()[1])
	f.GetElementTypeint64()[1] = 10
	fmt.Print(f.GetElementTypeint64())
	g := []int64{1, 2, 3}
	fmt.Print(g)
	g = append(g, 100)
	fmt.Print(g)
	f.SetmapWithKeystringWithValueint64(map[string]int64{"Key1": 1, "Key2": 2, "Key3": 3})
	fmt.Print(f.GetmapWithKeystringWithValueint64())
	fmt.Print(f.GetmapWithKeystringWithValueint64()["Key1"])
	f.GetmapWithKeystringWithValueint64()["Key1"] = 5
	fmt.Print(f.GetmapWithKeystringWithValueint64())
	i := 0
	for i < 20 {
		fmt.Print(i)
		i++
	}
	for i = 0; i < 20; i++ {
		fmt.Print(i + 5)
	}
	qw := 1.5
	fmt.Print(qw + float64(5) - 56.56 * float64(6) / float64(56))
}
```

The `Var` structure is a container for Union types.



## Contact
Name: Petr Makhnev

E-Mail: mr.makhneff@gmail.com

Telegram: @petr_makhnev

## License

This library is released under the [MIT](https://github.com/i582/component-sdl2/blob/master/LICENSE) license. For more information refer to the [LICENSE](https://github.com/i582/component-sdl2/blob/master/LICENSE) file provided with this project.