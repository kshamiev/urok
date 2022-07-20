package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"

	_ "go.opentelemetry.io/otel/exporters/jaeger"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	_ "go.opentelemetry.io/otel/sdk/trace"
	_ "go.opentelemetry.io/otel/semconv/v1.10.0"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// name is the Tracer name used to identify this instrumentation library.
const name = "jaeger"

func main() {
	l := log.New(os.Stdout, "", 0)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}

// App is a Fibonacci computation application.
type App struct {
	r io.Reader
	l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

// Run starts polling users for Fibonacci number requests and writes results.
func (a *App) Run(ctx context.Context) error {
	for {
		// TODO Each execution of the run loop, we should get a new "root" span and context.
		var span trace.Span
		ctx, span = otel.Tracer(name).Start(ctx, "Run")
		defer span.End()

		n, err := a.Poll(ctx)
		if err != nil {
			return err
		}

		a.Write(ctx, n)
	}
}

// Poll asks a user for input and returns the request.
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(name).Start(ctx, "Poll")
	defer span.End()

	a.l.Print("What Fibonacci number would you like to know: ")

	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)

	// Store n as a string to not overflow an int64.
	nStr := strconv.FormatUint(uint64(n), 10)
	span.SetAttributes(attribute.String("request.n", nStr))

	return n, err
}

// Write writes the n-th Fibonacci number back to the user.
func (a *App) Write(ctx context.Context, n uint) {
	var span trace.Span
	ctx, span = otel.Tracer(name).Start(ctx, "Write")
	defer span.End()

	f, err := func(ctx context.Context) (uint64, error) {
		_, span := otel.Tracer(name).Start(ctx, "Fibonacci")
		defer span.End()
		return Fibonacci(n)
	}(ctx)

	if err != nil {
		a.l.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		a.l.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}

// ////

// Fibonacci returns the n-th fibonacci number.
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
