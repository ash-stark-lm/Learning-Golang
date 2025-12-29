# Greetings Module

This folder contains a **Go library module** that provides a simple greeting function.
It is designed to be **imported and used by other Go programs**, not run directly.

---

## 📦 What this module is

- A **Go library** (not an executable)
- Exposes a single exported function: `Hello`
- Demonstrates:

  - Creating a Go module
  - Exported functions
  - Basic string formatting with `fmt.Sprintf`

---

## 📁 Files

```
greetings/
├── go.mod
├── greetings.go
└── README.md
```

---

## 🧠 Important notes

- This module uses `package greetings`
- It **does not** have a `main` package
- Therefore, it **cannot be run directly**

If you try:

```bash
go run greetings.go
```

You will get:

```
package command-line-arguments is not a main package
```

This is expected behavior.

---

## ✨ Provided Function

### `Hello`

```go
func Hello(name string) string
```

Returns a greeting message for the given name.

### Example:

```go
message := greetings.Hello("Ashish")
fmt.Println(message)
```

Output:

```
Hi, Ashish. Welcome!
```

---

## 🔍 How it works

- Uses `fmt.Sprintf` to format the greeting
- `%v` prints the value in its default format
- The function returns a string instead of printing it, making it reusable

---

## 🔗 How to use this module

This module is intended to be **imported by another module**, such as a `hello` application.

Typical usage involves:

- Creating another module
- Adding a `replace` directive in `go.mod`
- Importing `example.com/greetings`

(See the sibling `hello` folder for an example.)

---
