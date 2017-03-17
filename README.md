# timer

timer handle some task

## Install

```
go get github.com/wangh10/timer
```

### Usage

There are three tasks which need handle: t1 handle after 1s, t2 handle after 3s, t3 handle after 5s.
```go
package main

import (
    "time"
    "timer"
)

type PrintTask struct {
    ID uint
}

func (pt *PrintTask) Do() {
    println("print", pt.ID)
}

type WriteTask struct {
    ID uint
}

func (wt *WriteTask) Do() {
    println("write", wt.ID)
}

func main() {
    t := timer.NewTimer()
	t3 := &PrintTask{5}
	t2 := &WriteTask{3}
	t1 := &PrintTask{1}
	t.AddTask(t3, 5)
	t.AddTask(t2, 3)
	t.AddTask(t1, 1)
    time.Sleep(7 * time.Second)
    t.Close()
}
```
