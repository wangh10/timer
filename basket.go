package timer

type basket struct {
	t []*task
}

func (b basket) Len() int {
	return len(b.t)
}

func (b basket) Less(i, j int) bool {
	return b.t[i].c < b.t[j].c
}

func (b basket) Swap(i, j int) {
	b.t[i], b.t[j] = b.t[j], b.t[i]
}

func (b *basket) Push(x interface{}) {
	b.t = append(b.t, x.(*task))
}

func (b *basket) Pop() interface{} {
	o := b.t
	n := len(o)
	t := o[n-1]
	b.t = o[:n-1]
	return t
}

func (b *basket) next() {
	for i := range b.t {
		b.t[i].c--
	}
}
