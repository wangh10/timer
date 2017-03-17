package timer

import (
	"container/heap"
	"errors"
	"sync/atomic"
	"time"
)

var ErrNilTask = errors.New("task is nil")

type Timer struct {
	baskets [3600]*basket
	index   uint
	ticker  *time.Ticker
	total   int32
}

func NewTimer() *Timer {
	var t Timer
	t.ticker = time.NewTicker(time.Second)
	for i := 0; i < 3600; i++ {
		b := new(basket)
		heap.Init(b)
		t.baskets[i] = b
	}
	go t.start()
	return &t
}

func (t *Timer) AddTask(tk Task, at uint) error {
	if tk == nil {
		return ErrNilTask
	}
	i := at%3600 + t.index
	n := at / 3600
	heap.Push(t.baskets[i], &task{tk, n})
	atomic.AddInt32(&t.total, 1)
	return nil
}

func (t *Timer) Close() {
	t.ticker.Stop()
}

func (t *Timer) Total() int32 {
	return atomic.LoadInt32(&t.total)
}

func (t *Timer) start() {
	for range t.ticker.C {
		t.index = (t.index + 1) % 3600
		go func(index uint) {
			for {
				if t.baskets[index].Len() == 0 {
					return
				}
				x := heap.Pop(t.baskets[index])
				atomic.AddInt32(&t.total, -1)
				tk, ok := x.(*task)
				if !ok {
					continue
				}
				if tk.c > 0 {
					t.baskets[index].next()
					tk.c--
					heap.Push(t.baskets[index], tk)
					atomic.AddInt32(&t.total, 1)
					return
				}
				go tk.t.Do()
			}
		}(t.index)
	}
}
