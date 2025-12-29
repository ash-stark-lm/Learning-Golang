# handleError — Simple Error-Aware Greeting Package (Go)

`handleError` is a **small Go library package** that demonstrates **idiomatic Go error handling** by returning a value along with an `error` instead of using exceptions or panics.

This package does **not** define a `main` function and is intended to be **imported and used by other Go applications or libraries**.

---

## 📦 Package Overview

- Package name: `handleError`
- Type: **Library (not executable)**
- Language: Go
- Error handling style: **Return `(value, error)`**

---

## 📁 File Structure

```
handleError/
├── handleError.go
|── go.mod
└── README.md
```

---

## 📄 handleError.go

```go
package handleError

import (
	"errors"
	"fmt"
)
```

### What this package provides

The package exports a single function:

```go
func Hello(name string) (string, error)
```

---

## 🔧 Function Behavior

### Signature

```go
func Hello(name string) (string, error)
```

### Rules

| Condition    | Return Value                   |
| ------------ | ------------------------------ |
| `name == ""` | `("", error)`                  |
| Valid name   | `("Hi <name>. Welcome!", nil)` |

---

### Implementation

```go
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Please provide a name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
```

---

## 🧠 Go Error-Handling Philosophy

This package follows Go’s core principle:

> **Errors are values and should be returned, not thrown.**

- No exceptions
- No `try/catch`
- No `panic`
- Caller decides how to handle errors

---

## ▶️ How to Use

### Import the package

```go
import "example.com/handleError"
```

### Call the function

```go
message, err := handleError.Hello("Ashish")
if err != nil {
	log.Fatal(err)
}

fmt.Println(message)
```

---

## ❌ Empty Input Example

```go
message, err := handleError.Hello("")
```

Returned values:

```text
message = ""
err     = "Please provide a name"
```

---

## ✅ Valid Input Example

```go
message, err := handleError.Hello("Ashish")
```

Output:

```
Hi, Ashish. Welcome!
```
