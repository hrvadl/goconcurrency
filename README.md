## Go concurrency

This is a small test repository to play with the Go's concurrency. I was practicing
things I've learned from the [Concurrency in Go](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197) book.

## Things implemented

1. To play with `sync.Cond` primitive I have created following programs:
   - [Simple buffered channel implementation](https://github.com/hrvadl/goconcurrency/tree/main/internal/synccond/buffchan)
   - [Dining Philosophers Problem](https://github.com/hrvadl/goconcurrency/tree/main/internal/synccond/philosophers)

## How to run?

Make sure you have [taskfile](https://taskfile.dev/) and [Go](https://go.dev/) installed.

Then, to list all available tasks use following command from the root of the repo:

```sh
task
```

You should get a following output:

```sh
❯ task
task: [default] task --list-all
task: Available tasks for this project:
* default:                Show available tasks
* godoc:                  Host a Godoc web server on the http://localhost:6060/pkg/github.com/hrvadl/converter?m=all
* run-buffchan:
* run-philosophers:
* test:
* install:godoc:          Installed godoc util
* tidy:mod:               Tidy go mod
* update:mod:             Update go mod

```

To run the buffchan experiment, execute the following command:

```sh
task run-buffchan
```

To run the dinin philosophers problem, execute the following command:

```sh
task run-philosophers
```
