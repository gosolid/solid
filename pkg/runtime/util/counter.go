//js:package util

package util

import (
	"context"
	"sync"
	"time"

	"github.com/grexie/isolates"
)

type Counter struct {
	mutex    sync.Mutex
	counters []int64
	size     int64
	index    int64
	duration time.Duration
	lastTick time.Time
}

//js:constructor Counter
//js:export Counter
func NewCounter(size int64, duration time.Duration) *Counter {
	counter := &Counter{}
	counter.size = size
	counter.index = 0
	counter.duration = duration
	counter.lastTick = counter.FloorTime(time.UnixMicro(0))
	counter.counters = make([]int64, size)
	return counter
}

//js:method
func (c *Counter) Clone(ctx context.Context) (*Counter, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if counterv, err := isolates.For(ctx).New(NewCounter, c.size, c.duration); err != nil {
		return nil, err
	} else {
		counter := counterv.(*Counter)
		counter.lastTick = c.lastTick
		copy(counter.counters, c.counters)

		return counter, nil
	}
}

//js:method
func (c *Counter) FloorTime(now time.Time) time.Time {
	return now.Add(time.Duration(-now.Sub(time.Unix(0, 0)).Nanoseconds() % (c.duration.Nanoseconds() / c.size)))
}

func (c *Counter) tick(now time.Time) int64 {
	floorTime := c.FloorTime(now)
	isFloorTime := floorTime == now
	now = floorTime

	sinceLastTick := now.Sub(c.lastTick)
	currentIndex := (c.index + (c.size*sinceLastTick.Nanoseconds())/(c.duration.Nanoseconds()))

	for i := int64(0); i <= currentIndex-c.size && i < c.size; i++ {
		if isFloorTime && i == currentIndex-c.size && sinceLastTick < c.duration {
			continue
		}
		c.counters[i] = 0
	}
	// for i := currentIndex - c.size + 1; i > 0 && i < c.size; i++ {
	// 	c.counters[i] = 0
	// }

	for currentIndex%c.size < 0 {
		currentIndex += c.size
	}

	c.index = currentIndex % c.size
	c.lastTick = now
	return currentIndex % c.size
}

//js:method
func (c *Counter) Add(value int64, now time.Time) *Counter {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	index := c.tick(now)
	c.counters[index] += value
	return c
}

//js:method
func (c *Counter) AddNow(value int64) *Counter {
	return c.Add(value, time.Now())
}

//js:method
func (c *Counter) AddDuration(from time.Time, to time.Time) *Counter {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	duration := to.Sub(from).Nanoseconds()

	index := c.tick(to)
	slice := (c.size * duration) / c.duration.Nanoseconds()
	if slice > c.size-1 {
		slice = c.size - 1
	}

	itemDuration := c.duration.Nanoseconds() / c.size

	for i := int64(0); i < slice; i++ {
		id := (c.size + index - i) % c.size

		c.counters[id] += itemDuration
		duration -= itemDuration
	}

	c.counters[index] += duration

	return c
}

//js:method
func (c *Counter) AddDurationNow(from time.Time) *Counter {
	return c.AddDuration(from, time.Now())
}

//js:method
func (c *Counter) Sub(value int64, now time.Time) *Counter {
	return c.Add(-value, now)
}

//js:method
func (c *Counter) SumNow(duration time.Duration) int64 {
	return c.Sum(duration, time.Now())
}

//js:method
func (c *Counter) Sum(duration time.Duration, now time.Time) int64 {
	if duration > c.duration {
		panic("duration greater than recorded duration")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	index := c.tick(now)
	sliceTime := c.duration.Nanoseconds() / c.size
	floorTime := c.FloorTime(now)

	sum := int64(0)
	slice := ((c.size) * duration.Nanoseconds()) / c.duration.Nanoseconds()
	startIndex := index
	endIndex := startIndex - slice

	len := int64(0)

	for i := startIndex - 1; i > endIndex; i-- {
		j := i
		if j < 0 {
			j += c.size
		}
		len++
		sum += c.counters[j]
	}

	upperDuration := now.Sub(floorTime).Nanoseconds()
	lowerTime := now.Add(time.Duration(-duration))
	lowerCeilTime := floorTime.Add(time.Duration(-sliceTime * len))
	lowerDuration := lowerCeilTime.Sub(lowerTime).Nanoseconds()

	if endIndex < 0 {
		endIndex += c.size
	}

	sum += (c.counters[startIndex] * upperDuration) / sliceTime
	sum += (c.counters[endIndex] * lowerDuration) / sliceTime

	return sum
}

//js:method
func (c *Counter) SumAllNow() int64 {
	return c.SumAll(time.Now())
}

//js:method
func (c *Counter) SumAll(now time.Time) int64 {
	return c.Sum(c.duration, now)
}

//js:method
func (c *Counter) SumDurationNow(duration time.Duration) time.Duration {
	return c.SumDuration(duration, time.Now())
}

//js:method
func (c *Counter) SumAllDurationNow() time.Duration {
	return c.SumAllDuration(time.Now())
}

//js:method
func (c *Counter) SumDuration(duration time.Duration, now time.Time) time.Duration {
	return time.Duration(c.Sum(duration, now))
}

//js:method
func (c *Counter) SumAllDuration(now time.Time) time.Duration {
	return c.SumDuration(c.duration, now)
}
