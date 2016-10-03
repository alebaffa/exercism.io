package brackets

const testVersion = 3
const parenthesis1Open = "{"
const parenthesis1Close = "}"
const parenthesis2Open = "["
const parenthesis2Close = "]"
const parenthesis3Open = "("
const parenthesis3Close = ")"

type Node struct {
	Value string
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func Bracket(input string) (bool, error) {
	if input == "" {
		return true, nil
	}
	stack := NewStack()
	for _, char := range input {
		switch string(char) {
		case parenthesis1Open:
			stack.Push(&Node{string(char)})
		case parenthesis2Open:
			stack.Push(&Node{string(char)})
		case parenthesis3Open:
			stack.Push(&Node{string(char)})
		case parenthesis1Close:
			stack.Pop()
		case parenthesis2Close:
			stack.Pop()
		case parenthesis3Close:
			stack.Pop()
		}
	}
	if stack.count == 0 {
		return true, nil
	}
	return false, nil
}
