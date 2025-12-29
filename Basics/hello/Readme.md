# Hello with Logging, Error Handling, Modules, Replace Directive & Unicode

This folder contains a **standalone Go executable program** that demonstrates **real-world Go application practices**, including:

- Structured **logging with `log`**
- Idiomatic **error handling**
- Use of `package main`
- Importing **external modules**
- Importing **local modules**
- Using the `replace` directive for local development
- Dependency management with `go mod tidy`
- Native **Unicode (UTF-8)** support
- Go’s **strict import rules**

---

## 🧾 Logging & Error Handling (NEW)

This program uses Go’s built-in `log` package for **clean, production-style error handling**.

### Logger Configuration

```go
log.SetPrefix("app: ")
log.SetFlags(0)
```

- Adds a consistent prefix to log messages
- Disables timestamps and file info for clean output

---

### Error Handling Pattern

```go
result, err := handleError.Hello("")
if err != nil {
	log.Fatal(err)
}
```

**Important Go principle:**

- Go does **not** use `try/catch`
- Functions return `(value, error)`
- Errors must be checked explicitly
- `log.Fatal`:

  - Prints the error
  - Applies the logger prefix
  - Exits the program with a non-zero status

---

### Why `log.Fatal` is used here

| Scenario                    | Recommended            |
| --------------------------- | ---------------------- |
| Application startup failure | `log.Fatal`            |
| Recoverable error           | return `error`         |
| Libraries                   | NEVER call `log.Fatal` |

---

## 📦 What this program is

- A **Go application** (not a library)
- Execution begins at `func main()`
- Combines:

  - Standard library → `fmt`, `log`
  - External module → `rsc.io/quote`
  - Local modules → `example.com/greetings`, `example.com/handleError`

---

## 📁 Folder Structure

```
hello/
├── go.mod
├── hello.go
└── README.md
```

---

## 📄 main.go (Code Explained)

```go
package main
```

If a Go program uses `package main`, Go treats it as a **standalone executable program**.

---

```go
import (
	"fmt"
	"log"

	"example.com/greetings"
	"example.com/handleError"
	"rsc.io/quote"
)
```

### Strict Import Rules

Go **will not compile** if:

- An imported package is **missing**
- An imported package is **unused**

This strict requirement prevents unused dependencies from accumulating as programs evolve.

---

```go
func main() {
	log.SetPrefix("app: ")
	log.SetFlags(0)

	fmt.Println(quote.Hello())

	message := greetings.Hello("Ashish")
	fmt.Println(message)

	result, err := handleError.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println("Hello, 世界")
}
```

---

## 🌍 Unicode Support

```go
fmt.Println("Hello, 世界")
```

Go natively supports **UTF-8 / Unicode**, allowing multilingual text without additional configuration.

---

## 🔁 Using a Local Module with `replace`

This program imports:

```
example.com/greetings
example.com/handleError
```

Since these modules are **not published**, Go must be told where to find them locally.

---

### Step 1: Add a replace directive

From the **hello directory**, run:

```bash
go mod edit -replace example.com/greetings=../greetings
go mod edit -replace example.com/handleError=../handleError
```

This tells Go to resolve these imports locally.

---

### Step 2: Synchronize dependencies

Run:

```bash
go mod tidy
```

This command:

- Scans all imports
- Adds missing dependencies
- Removes unused ones

---

### Final `go.mod` (example)

```go
module example.com/hello

go 1.25.4

replace example.com/greetings => ../greetings
replace example.com/handleError => ../handleError

require (
	example.com/greetings v0.0.0
	example.com/handleError v0.0.0
	rsc.io/quote v1.5.2
)
```

---

## 🧠 About Pseudo-Versions

```
v0.0.0
```

This is a **pseudo-version**, used when:

- A module has no Git tags
- The dependency is resolved locally

This is normal during development.

---

## 🚀 Production Note

In production:

- `replace` directives are removed
- Real semantic versions are used

```go
require example.com/greetings v1.1.0
```

---

## ▶️ How to Run

From inside the `hello` directory:

```bash
go run .
```

---

## 🖥 Example Output

```
Ahoy, world!
Hi, Ashish. Welcome!
app: Please provide a name
```

(The program exits after the logged error.)

---

## 🎯 Learning Goals

This example teaches:

- Idiomatic Go error handling
- Logging best practices
- Go application structure
- Module resolution
- Local development with `replace`
- Unicode support
- Why Go enforces strict imports

---

## ✅ Key Takeaways

- `package main` + `func main()` → executable program
- Errors are returned, not thrown
- Applications log and exit; libraries return errors
- `replace` is for local development only
- `go mod tidy` keeps dependencies clean
- Go supports Unicode out of the box

---
