# must

`must` is a tiny Go package that helps unwrap results from functions returning
`(..., error)` by panicking on non-nil errors. It reduces boilerplate when you
want concise code and are fine with panics on failure (e.g. in tests, quick
scripts, or initialization code).

---

## Install

```bash
go get github.com/linniksa/must
```

---

## Usage

### Example

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yourusername/must"
)

func main() {
	// Unwrap a single return + error
	port := must.P1(strconv.Atoi(os.Getenv("PORT")))
	fmt.Println("port:", port)

	// Unwrap two returns + error
	a, b := must.P2(someFunc())
	fmt.Println(a, b)

	// Unwrap three returns + error
	x, y, z := must.P3(otherFunc())
	fmt.Println(x, y, z)
}

func someFunc() (string, int, error) {
	return "ok", 42, nil
}

func otherFunc() (int, int, int, error) {
	return 1, 2, 3, nil
}
```

---

## API

- `P1(v1, err)` → `v1`
- `P2(v1, v2, err)` → `(v1, v2)`
- `P3(v1, v2, v3, err)` → `(v1, v2, v3)`
- `P4(v1, v2, v3, v4, err)` → `(v1, v2, v3, v4)`

Each function panics if `err != nil`.

---

## When to use

- ✅ In tests, prototypes, or CLI tools where panics are acceptable
- ✅ During initialization where an error is unrecoverable
- ❌ In production request paths where graceful error handling is required

---

## License

MIT
