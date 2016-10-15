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

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

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
			node := stack.Pop()
			return checkNextSymbolIsCorrect(node, parenthesis1Open)
		case parenthesis2Close:
			node := stack.Pop()
			return checkNextSymbolIsCorrect(node, parenthesis2Open)
		case parenthesis3Close:
			node := stack.Pop()
			return checkNextSymbolIsCorrect(node, parenthesis3Open)
		}
	}
	if stack.count != 0 {
		return false, nil
	}
	return true, nil
}

func checkNextSymbolIsCorrect(node *Node, parenthesisOpen string) (bool, error) {
	if node == nil {
		return false, nil
	}
	if node.Value != parenthesisOpen {
		return false, nil
	}
	return true, nil
}
