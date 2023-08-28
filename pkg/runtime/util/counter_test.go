package util

import (
	"fmt"
	"testing"
	"time"
)

func testCounter10(t *testing.T, n int64, expected int64) error {
	counter := NewCounter(10, 100*time.Second)

	tm := counter.lastTick

	for i := time.Duration(0); i < 10; i++ {
		counter.Add(1, tm.Add(i*time.Second))
	}

	counter.lastTick = tm

	sum := counter.SumAll(tm.Add(time.Duration(n) * time.Second))
	if sum != expected {
		return fmt.Errorf("counter sum should be %d: %d %p", expected, sum, counter.counters)
	}
	return nil
}

func TestCounter10(t *testing.T) {
	for i := int64(10); i < 1000; i++ {
		if i >= 100 {
			if err := testCounter10(t, i, 0); err != nil {
				t.Error(err)
				return
			}
		} else {
			if err := testCounter10(t, i, 10); err != nil {
				t.Error(err)
				return
			}
		}
	}
}

func testCounterDuration(t *testing.T, n int64) error {
	counter := NewCounter(10, 100*time.Second)

	tm := counter.lastTick

	for i := time.Duration(0); i < 100; i++ {
		counter.Add(1, tm.Add(i*time.Second))
	}

	sum := counter.Sum(time.Duration(n)*time.Second, tm.Add(100*time.Second))
	if sum != n {
		return fmt.Errorf("counter sum should be: %d, got: %d", n, sum)
	}
	return nil
}

func TestCounterDuration(t *testing.T) {
	for i := int64(0); i < 100; i++ {
		if err := testCounterDuration(t, i); err != nil {
			t.Error(err)
			return
		}
	}
}
