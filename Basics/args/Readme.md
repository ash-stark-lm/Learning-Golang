# Command-Line Arguments in Go (Echo Example)

Most programs **process input to produce output**—that is the essence of computing.
In Go, input can come from many external sources such as **files, networks, user input, or command-line arguments**.

This example focuses on **command-line arguments** and demonstrates **three different, idiomatic ways** to process them in Go.

---

## 📌 What This Program Does

- Reads command-line arguments using `os.Args`
- Skips the program name (`os.Args[0]`)
- Joins all arguments into a **single space-separated string**
- Prints the result **three times**, using different approaches

This program is inspired by the Unix `echo` command.

---

## 📁 Project Structure

```
args/
├── main.go
├── go.mod
└── README.md
```

---

## 📄 main.go

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echoClassic()
	echoRange()
	echoJoin()
}

func echoClassic() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoRange() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
```

---

## 🧠 Key Concepts Explained

### 1️⃣ `os.Args`

The `os` package provides platform-independent access to operating system features.

`os.Args` is a **slice of strings** containing command-line arguments.

| Index         | Meaning            |
| ------------- | ------------------ |
| `os.Args[0]`  | Program name       |
| `os.Args[1:]` | User-provided args |

Go uses **half-open slicing**, so:

```go
os.Args[1:]
```

means “all arguments except the program name”.

---

### 2️⃣ Slices in Go

A **slice** is a dynamically sized sequence of elements.

- Access element → `s[i]`
- Sub-slice → `s[m:n]`
- Length → `len(s)`

Half-open intervals simplify logic and are standard in Go.

---

## 🔁 Three Loop Approaches Used

### ✅ Version 1: Classic `for` loop

```go
for i := 1; i < len(os.Args); i++ {
```

- Uses explicit indexing
- Closest to C/C++ style
- Good for learning fundamentals

---

### ✅ Version 2: `range` loop (idiomatic Go)

```go
for _, arg := range os.Args[1:] {
```

- Uses Go’s `range`
- Ignores index using `_`
- Cleaner and safer than manual indexing

---

### ✅ Version 3: `strings.Join` (most efficient)

```go
strings.Join(os.Args[1:], " ")
```

- Fastest and simplest
- Avoids repeated string allocations
- Recommended for real-world usage

---

## 🧩 Choosing Between Variable Declaration Forms

Go allows multiple equivalent ways to declare strings:

```go
s := ""
var s string
var s = ""
var s string = ""
```

### When to use which

| Form                | Use case                              |
| ------------------- | ------------------------------------- |
| `s := ""`           | Most compact; use inside functions    |
| `var s string`      | When the initial value doesn’t matter |
| `var s = ""`        | Rare; mostly for grouped declarations |
| `var s string = ""` | When the type must be explicit        |

**In practice**:

- Use **explicit initialization** when the initial value matters
- Use **implicit initialization** when it does not

---

## 🔁 Why `+=` Works but Can Be Costly

In the first two versions:

```go
s += sep + arg
```

This is equivalent to:

```go
s = s + sep + arg
```

Each iteration:

- Allocates a new string
- Copies old contents + new text
- Leaves old strings for garbage collection

This is fine for small inputs, but inefficient for large data.

---

## 🚀 Better Alternative: `strings.Join`

```go
fmt.Println(strings.Join(os.Args[1:], " "))
```

Why it’s better:

- Fewer allocations
- Cleaner code
- Idiomatic and performant

---

## 🛠 Debug-Friendly Output

For quick inspection (no formatting concern):

```go
fmt.Println(os.Args[1:])
```

This prints the slice directly and works for **any slice**.

---

## ▶️ How to Run

```bash
go run main.go Hello World from Go
```

Output:

```
Hello World from Go
Hello World from Go
Hello World from Go
```

Each line corresponds to:

1. Classic `for`
2. `range`
3. `strings.Join`

---

## ✅ Key Takeaways

- Programs take input → process → produce output
- `os.Args` exposes command-line input in Go
- `package main` + `func main()` creates an executable
- Prefer `range` for clarity, `strings.Join` for performance
- Go slices use half-open intervals
- Simple examples scale into real CLI tools

---
