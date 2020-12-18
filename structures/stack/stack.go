package stack

type NodeStack struct {
	Items []string
}

func NewNodeStack() *NodeStack {
	return &NodeStack{Items: make([]string,0)}
}

func (s *NodeStack) Push(i string)  {
	s.Items = append(s.Items,i)
}

func (s *NodeStack) Pop() string {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0:len(s.Items)-1]
	return item
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.Items)==0
}
