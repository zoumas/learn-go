# learn-go

My solutions and notes for boot.dev's [Learn Go](https://www.boot.dev/courses/learn-golang) course, done lesson by lesson. The course has 148 lessons in 17 chapters and covers idiomatic Go, from variables to concurrency, generics and enums.

This is the first Go course in my run through the boot.dev DevOps path, which focuses on Go, AWS and DevOps.

## Layout

```
chNN/
  lNN/        one lesson: the solution and boot.dev's tests when provided
notes/
  chNN.md     notes for each chapter
```

Run a lesson with `go run ./ch01/l01`. Run all tests with `go test ./...`.

## Learning goals by chapter

The chapter titles come from the course page. The goals under each title are my own expansion, not boot.dev's text.

1. **Variables**: basic types, declaration styles (`var` vs `:=`), zero values, type inference, conversions.
2. **Constants and Formatting**: `const`, untyped constants, the `fmt` verbs and `Sprintf`.
3. **Conditionals**: `if` with an initializer statement, `switch`.
4. **Functions**: multiple returns, named returns, variadics, first-class functions, closures, `defer`.
5. **Structs**: nested and anonymous structs, embedding, methods, memory layout.
6. **Interfaces**: implicit implementation, type assertions, type switches, keeping interfaces small.
7. **Errors**: the `error` interface, wrapping with `%w`, `errors.Is` and `errors.As`, `panic` and `recover`.
8. **Loops**: `for` in its forms, `range`, `break` and `continue`.
9. **Slices**: arrays vs slices, `len` and `cap`, `append` growth, aliasing, `make`, variadic spreading.
10. **Maps**: creation, lookup with `ok`, deletion, nested maps, key comparability.
11. **Pointers**: `&` and `*`, nil, pointer vs value receivers.
12. **Packages and Modules**: package layout, visibility, `go.mod`, remote modules.
13. **Channels**: goroutines, buffered and unbuffered channels, `close`, `range`, `select`, nil and closed channel behavior.
14. **Mutexes**: `sync.Mutex` and `sync.RWMutex`, data races, when a mutex fits better than a channel.
15. **Generics**: type parameters, constraints and type sets, when generics are worth it.
16. **Enums**: `iota`, named types, `String()` methods.
17. **Quiz**: review across the whole course.

## Progress

- [x] 1. Variables
- [x] 2. Constants and Formatting
- [x] 3. Conditionals
- [x] 4. Functions
- [x] 5. Structs
- [x] 6. Interfaces
- [x] 7. Errors
- [x] 8. Loops
- [x] 9. Slices
- [ ] 10. Maps
- [ ] 11. Pointers
- [ ] 12. Packages and Modules
- [ ] 13. Channels
- [ ] 14. Mutexes
- [ ] 15. Generics
- [ ] 16. Enums
- [ ] 17. Quiz
