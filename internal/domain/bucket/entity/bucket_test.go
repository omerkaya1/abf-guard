package entity

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBucket_Decrement(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	finished := make(chan string)
	count := 5

	b := NewBucket(context.Background(), "test", 5, finished)

	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.GetCount())
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
	defer cancel()
	finished := make(chan string)
	count := 5

	b := NewBucket(context.Background(), "test", 5, finished)

	// These will be allowed
	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.GetCount())
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

func TestBucket_Decrement_Ctx_Error(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	finished := make(chan string)
	count := 5

	b := NewBucket(ctx, "test", 5, finished)

	// These will be allowed
	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.GetCount())
	}

	// These will be disallowed
	for i := 0; i < 3; i++ {
		assert.Equal(t, false, b.Decrement())
	}

	time.AfterFunc(time.Second*2, cancel)

	select {
	case <-ctx.Done():
		assert.Error(t, ctx.Err())
		assert.Equal(t, context.Canceled, ctx.Err())
	case v := <-finished:
		assert.Equal(t, "test", v)
	}
}

func TestBucket_Decrement_Ctx_Deadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	finished := make(chan string)
	count := 5

	b := NewBucket(ctx, "test", 5, finished)

	// These will be allowed
	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.GetCount())
	}

	// These will be disallowed
	for i := 0; i < 3; i++ {
		assert.Equal(t, false, b.Decrement())
	}

	select {
	case <-ctx.Done():
		assert.Error(t, ctx.Err())
		assert.Equal(t, context.DeadlineExceeded, ctx.Err())
	case v := <-finished:
		assert.Equal(t, "test", v)
	}
}

func TestBucket_Stop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	finished := make(chan string)
	count := 5

	b := NewBucket(ctx, "test", 5, finished)

	for i := 1; i < count; i++ {
		assert.Equal(t, true, b.Decrement())
		assert.Equal(t, count-i, b.GetCount())
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
