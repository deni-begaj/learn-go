# Learning plan

Each folder is one library package. Only the root `main.go` is `package main`.

Status: Day 1 is done. Day 2, part A is the current exercise.

## Day 1 — data

Done.

- Structs, composite literals, and methods (`structs`).
- Slices, `append` results, and a filter into a new slice (`slices`).
- Pointers: `&`, `*`, and a pointer receiver that mutates (`pointers`, `Car.drive`).
- Maps: a `map[string]User` and `range` (`maps`).

## Day 2 — errors and boundaries

A miss is a return value. `panic` is for a bug, not for a missing user.

### Part A — current exercise

In `maps`:

- `var ErrNotFound = errors.New("user not found")`.
- `FindUser(users map[string]User, key string) (User, error)`.
  A hit returns the user and `nil`. A miss returns `User{}` and `ErrNotFound`.
  The lookup is `user, ok := users[key]`. `ok == false` is the only signal; the zero `User` is not.
- `MapsLookup`, called from root `main`. Look up `"sarah"` and `"nobody"`.
  On failure wrap with the key: `fmt.Errorf("find user %s: %w", key, err)`.
  Branch with `errors.Is(err, ErrNotFound)`.

### Part B — next

- A two-method interface declared next to the code that uses it.
- A fake implementation in a `_test.go` file.
- A table-driven test with three rows: success, the sentinel, and a wrapped error checked with `errors.Is`.

## Day 3 — concurrency and context

Blocking a goroutine on I/O is the normal style. The runtime parks it.

- Fan out work with `errgroup`, cancel the rest on the first error.
- `context.Context` is the first parameter of any function that can run long. Do not store it on a struct.
- Exercise: fetch several URLs under one timeout and show that a cancelled context stops the work.

Reach for a mutex when goroutines share memory. Reach for a channel when one goroutine hands a value to another.

## Day 4 — one HTTP service

Standard library only.

- One `GET` on `net/http.ServeMux`.
- `log/slog` JSON logger. Request id on the context.
- Server read, write, and idle timeouts.
- Shutdown on SIGTERM: stop accepting, let in-flight requests finish.
- Pass `r.Context()` into the work.
- Hit the handler with `httptest`, then with a real client, and cancel a slow request.

## Day 5 — service shape

- The handler depends on a small interface. SQL and outbound calls stay behind it.
- `main` builds the dependencies and passes them in. No container.
- The handler test uses a fake and never touches the network.
