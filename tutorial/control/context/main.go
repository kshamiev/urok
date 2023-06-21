package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "key", "popcorn")
	go Test1(ctx)
	go Test2(ctx)
	go Test3(ctx)
	time.Sleep(time.Second * 30)
	cancel()
	time.Sleep(time.Hour)
}

func Test1(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 5):
			log.Println("Work Test111")
		}
	}
}

func Test2(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Hour)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 5):
			log.Println("Work Test222")
		}
	}
}

func Test3(ctx context.Context) {
	ctx = ContextDWT(ctx, time.Hour)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 5):
			log.Println("Work Test333 " + ctx.Value("key").(string))
		}
	}
}

type detachedCtx struct{ context.Context }

func (detachedCtx) Deadline() (deadline time.Time, ok bool) { return }
func (detachedCtx) Done() <-chan struct{}                   { return nil }
func (detachedCtx) Err() error                              { return nil }

func ContextDetach(ctx context.Context) context.Context {
	return detachedCtx{ctx}
}

func ContextDWT(ctx context.Context, d time.Duration) context.Context {
	return ContextDetachWithTimeout(ctx, d)
}

func ContextDetachWithTimeout(ctx context.Context, d time.Duration) context.Context {
	ctx, _ = context.WithTimeout(detachedCtx{ctx}, d)
	return ctx

}

func ContextDWTC(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return ContextDetachWithTimeoutCancel(ctx, d)
}

func ContextDetachWithTimeoutCancel(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(detachedCtx{ctx}, d)
}
