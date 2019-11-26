package entity

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBucket_Decrement(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	finished := make(chan string)
	count := 5

	b := NewBucket("test", 5, time.Second*2, finished)

	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.count)
	}

	tick := time.NewTicker(6 * time.Second)

	select {
	case <-ctx.Done():
		cancel()
		assert.Error(t, ctx.Err())
	case v := <-finished:
		assert.Equal(t, "test", v)
	case <-tick.C:
		t.Fail()
	}
}

func TestBucket_Decrement_Prohibited(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	finished := make(chan string)
	count := 5

	b := NewBucket("test", 5, time.Second*2, finished)

	// These will be allowed
	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.count)
	}

	// These will be disallowed
	for i := 0; i < 3; i++ {
		assert.Equal(t, false, b.Decrement())
	}

	tick := time.NewTicker(6 * time.Second)

	select {
	case <-ctx.Done():
		cancel()
		assert.Error(t, ctx.Err())
	case v := <-finished:
		assert.Equal(t, "test", v)
	case <-tick.C:
		t.Fail()
	}
}

func TestBucket_Stop(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	finished := make(chan string)
	count := 5

	b := NewBucket("test", 5, time.Second*2, finished)

	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.count)
	}

	go func() {
		b.Stop()
	}()

	tick := time.NewTicker(6 * time.Second)

	select {
	case <-ctx.Done():
		cancel()
		assert.Error(t, ctx.Err())
	case v := <-finished:
		assert.Equal(t, "test", v)
	case <-tick.C:
		t.Fail()
	}
}
