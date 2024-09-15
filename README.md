## Go concurrency

This is a small test repository to play with the Go's concurrency. I was practicing
things I've learned from the [Concurrency in Go](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197) book.

## Things implemented

1. To play with `sync.Cond` primitive I have created following programs:
   - [Simple buffered channel implementation](https://github.com/hrvadl/goconcurrency/tree/main/internal/synccond/buffchan)
   - [Dining Philosophers Problem](https://github.com/hrvadl/goconcurrency/tree/main/internal/synccond/philosophers)
   - [Traffic lights & cars](https://github.com/hrvadl/goconcurrency/tree/main/internal/synccond/trafficgreen)
2. To play with `sync.Pool` primitive I have implemented a following program:
   - [Simple HTTP server, which will cache buffers between requests](https://github.com/hrvadl/goconcurrency/blob/main/internal/syncpool/decoders/handler.go)
3. To play with `concurrency patterns` I have implemented a following programs:
   - [Simple fan-in/fan-out program with or-done-chan pattern and a multiple pipelines](https://github.com/hrvadl/goconcurrency/blob/main/internal/patterns/fan/main.go)
   - [Simple program to demonstrate idiomatic error handling with goroutines](https://github.com/hrvadl/goconcurrency/blob/main/internal/patterns/errorhandle/main.go)
   - [Simple program to practise or-chan pattern](https://github.com/hrvadl/goconcurrency/blob/main/internal/patterns/orchan/main.go)
   - [Simple program to practise tee-chan pattern](https://github.com/hrvadl/goconcurrency/blob/main/internal/patterns/teechan/main.go)
4. To play with `concurrency at scale patterns` I have created following programs:
   - [Simple program to practise using rate limiter & combining limiters to multi-limiter](https://github.com/hrvadl/goconcurrency/blob/main/internal/scalepatterns/ratelimit/main.go)
   - [Simple program to practise using heartbeat pattern](https://github.com/hrvadl/goconcurrency/blob/main/internal/scalepatterns/heartbeat/main.go)

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

To run the dining philosophers problem, execute the following command:

```sh
task run-philosophers
```

To run the traffic lights & cars problem, execute the following command:

```sh
task run-traffic
```

To run the exercise with `sync.Pool`, execute the following command:

```sh
task run-pool
```

To run the exercise with error handle in goroutines, execute the following command:

```sh
task run-errhandle
```

To run the exercise with the or-chan pattern, execute the following command:

```sh
task run-orchan
```

To run the exercise with the fan-in/fan-out patterns, execute the following command:

```sh
task run-fan
```

To run the exercise with the tee-chan pattern, execute the following command:

```sh
task run-teechan
```

To run the exercise with the multi rate limiter, execute the following command:

```sh
task run-multilimiter
```

To run the exercise with the multi rate limiter, execute the following command:

```sh
task run-heartbeat
```

## Main takeaways:

### Channels

1. Owner of the channel is responsible for:
    - instantiate the channel
    - perform writes (remember, that writes can also block), or pass ownership to another goroutine
    - close the channel
2. Consumer of the channel is responsible for:
    - knowing when channel is closed
    - handle blocking (done channel or context with cancellation/timeout/deadline)

### Goroutines

1. Pass errors as values in the result set up
2. Create pipelines (generators) with the goroutines
3. Pipelines are easily composable with the patterns described above
