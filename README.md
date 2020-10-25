# go-calculator

## Usage

```
$ go run cmd/calculator/main.go
calculator> (2.5 - 1.35) * 2.0
2.3
calculator> -sin((-1+2.5)*pi)
1
calculator> 180*atan2(log(e), log10(10))/pi
45
calculator> exit
```

You can also use `calculator.Calculate()` in your application:
```go
package main

import (
	"fmt"
	"log"

	"github.com/mnogu/go-calculator"
)

func main() {
	val, err := calculator.Calculate("(2.5 - 1.35) * 2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 2.3

	val, err = calculator.Calculate("-sin((-1+2.5)*pi)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 1

	val, err = calculator.Calculate("180*atan2(log(e), log10(10))/pi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 45
}
```

## References

* [chibicc](https://github.com/rui314/chibicc): A small C compiler