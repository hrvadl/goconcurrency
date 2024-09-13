package main

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

const timeout = time.Second * 1

func main() {
	urls := []string{"https://google.com", "https://facebook.com", "https://baddomain"}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resCh := checkURLs(ctx, urls)

	for res := range resCh {
		if res.err != nil {
			slog.Error(
				"Got error while requesting",
				slog.String("from", res.url),
				slog.Any("err", res.err),
			)
			continue
		}

		slog.Info(
			"Got response",
			slog.String("from", res.url),
			slog.Int("status", res.result.StatusCode),
		)
	}

	slog.Info("All requests are made")
}

type result struct {
	err    error
	result *http.Response
	url    string
}

func checkURLs(ctx context.Context, urls []string) <-chan result {
	results := make(chan result, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case results <- getHTTP(ctx, url):
			case <-ctx.Done():
				results <- result{err: ctx.Err(), result: nil, url: url}
			}
		}()
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	return results
}

func getHTTP(ctx context.Context, url string) result {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return result{err: err, result: nil, url: url}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result{err: err, result: nil, url: url}
	}

	if err = res.Body.Close(); err != nil {
		return result{err: err, result: res, url: url}
	}

	return result{err: nil, result: res, url: url}
}
