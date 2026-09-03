package server

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

func TestWaitForLocalListenerReady(t *testing.T) {
	t.Parallel()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	t.Cleanup(func() {
		if err := listener.Close(); err != nil {
			t.Errorf("close listener: %v", err)
		}
	})

	if err := waitForLocalListener(t.Context(), listener.Addr().String(), time.Second); err != nil {
		t.Fatalf("waitForLocalListener() error = %v", err)
	}
}

func TestWaitForLocalListenerCanceled(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	err := waitForLocalListener(ctx, "127.0.0.1:1", time.Second)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("waitForLocalListener() error = %v, want context.Canceled", err)
	}
}
