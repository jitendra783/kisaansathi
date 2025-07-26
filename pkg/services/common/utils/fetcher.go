package utils

import (
	"context"
	"sync"
	"time"
)

type FetchFunc[T any] func(ctx context.Context) (T, error)

func Fetch[T any](
	parentCtx context.Context,
	fetchFunc FetchFunc[T],
	wg *sync.WaitGroup,
	resultCh chan<- T,
	errCh chan<- error,
) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	result, err := fetchFunc(ctx)
	if err != nil {
		errCh <- err
		return
	}
	resultCh <- result
}
