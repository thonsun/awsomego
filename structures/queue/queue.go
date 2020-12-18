package queue

type NodeQueue struct {
	tems []string
}

func NewNodeQueue() *NodeQueue {
	q := &NodeQueue{}
	q.tems = []string{}
	return q
}

func (q *NodeQueue) Enqueue(s string) {
	q.tems = append(q.tems,s)
}

func (q *NodeQueue) Dequeue() string {
	item := q.tems[0]
	q.tems = q.tems[1:len(q.tems)]
	return item
}

func (q *NodeQueue) IsEmpty() bool {
	return len(q.tems)==0
}
